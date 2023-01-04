//go:generate packer-sdc mapstructure-to-hcl2 -type Config

package lifecycle

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/digitalocean/godo"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/common"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	NamePrefix			string `mapstructure:"name_prefix"`
	Days				int`mapstructure:"days_older_than"` 
	ctx                 interpolate.Context
}

type PostProcessor struct {
	config Config
}

func (p *PostProcessor) ConfigSpec() hcldec.ObjectSpec { return p.config.FlatMapstructure().HCL2Spec() }

func (p *PostProcessor) Configure(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		PluginType:         "packer.post-processor.lifecycle",
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
		InterpolateFilter: &interpolate.RenderFilter{
			Exclude: []string{},
		},
	}, raws...)
	if err != nil {
		return err
	}

	if p.config.Days == 0 {
		return fmt.Errorf("days_older_than must be defined and at least 1 day")
	}

	if len(p.config.NamePrefix) < 3 {
		return fmt.Errorf("name_prefix must filter images with at least 3 characters this avoids targeting all images in the account")
	}

	return nil
}

func (p *PostProcessor) PostProcess(ctx context.Context, ui packersdk.Ui, source packersdk.Artifact) (packersdk.Artifact, bool, bool, error) {
	ui.Say(fmt.Sprintf("Deleting images with the name prefix %v older than %d days. Before %v", p.config.NamePrefix, p.config.Days, time.Now().AddDate(0,0, -p.config.Days).UTC()))

	var do_token string
	if len(os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")) > 0 {
		do_token = os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	} else if len(os.Getenv("DO_API_TOKEN")) > 0 {
		do_token = os.Getenv("DO_API_TOKEN")
	} else {
		ui.Error("set the access token environment variable: DIGITALOCEAN_ACCESS_TOKEN or DO_API_TOKEN (depreciated)")
	}
	
    client := godo.NewFromToken(do_token)

	images, _, err := client.Images.ListUser(ctx, &godo.ListOptions{})

	var filteredImages []godo.Image

	for _, image := range images {

        if strings.HasPrefix(image.Name, p.config.NamePrefix) {
            filteredImages = append(filteredImages, image)
        }
    }

	if len(filteredImages) < 1 {
		ui.Message("no images found")
	}

	for _, image := range filteredImages {
		createdTime, _ := time.Parse("2006-01-02T15:04:05Z0700", image.Created)
        if isOlderThanDays(createdTime, p.config.Days) {
			jsonOut, err := json.Marshal(image)
			if err != nil {
				return source, true, true, err
			}
			
			ui.Say(fmt.Sprintf("Deleting image: %v", string(jsonOut)))
        }
    }

	return source, true, true, err
}

func isOlderThanDays(dateTime time.Time, numDays int) bool {
	return time.Since(dateTime.UTC()) > time.Since(time.Now().AddDate(0,0, -numDays).UTC())
}