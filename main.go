package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
	"packer-plugin-digitalocean-image-lifecycle/lifecycle"
)

type LifecyclePostProcessor struct{}

func main() {
    pps := plugin.NewSet()
    pps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(lifecycle.PostProcessor))
    err := pps.Run()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        os.Exit(1)
    }
}
