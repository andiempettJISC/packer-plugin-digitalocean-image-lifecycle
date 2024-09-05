# Digitalocean image lifecycle

A multi-component packer plugin to manage the lifescycle of existing digitialocean images.

## Post-Processors: digitalocean-image-lifecycle

### Usage

Initialise the plugin in a `required_plugins` block in configuration. See the [ packer installing plugins docs](https://developer.hashicorp.com/packer/docs/plugins/install-plugins):

```
packer {
  required_plugins {
    digitalocean-image-lifecycle = {
      version = ">=v1.0.0"
      source  = "github.com/andiempettJISC/digitalocean-image-lifecycle"
    }
  }
}
```

Define your sources, typically you would be using [the digitalocean builder](https://developer.hashicorp.com/packer/plugins/builders/digitalocean) then in the build block define the `digitalocean-image-lifecycle` post-processor:

```
build {
  sources = ["sources.null.basic-example"]

  post-processor "digitalocean-image-lifecycle" {
    name_prefix = "example-image"
    days_older_than = 9
  }
}
```

### Inputs

| Name            | Description                                      | Type   |
|-----------------|--------------------------------------------------|--------|
| name_prefix     | The start of the name of the image to filter by. | string |
| days_older_than | The Age of existing images to delete in days.    | number |

## Development

This plugin follows the packer multi-component [plugin development workflow](https://developer.hashicorp.com/packer/docs/plugins/creation).

TLDR:

Its worth noting we need to generate the HCL2 of the plugin with a helper. See https://developer.hashicorp.com/packer/guides/hcl/component-object-spec

```
go install
go generate ./lifecycle/lifecycle.go 
go build
cp packer-plugin-digitalocean-image-lifecycle example
```

comment out the `required_plugins` block in [`example/main.pkr.hcl`](example/main.pkr.hcl)

```
cd example
packer init .
packer build .
```
<<<<<<< HEAD


### If you're up to this for the first time, ensure you've got `Go` set up

```
brew install go
brew install goreleaser
```

For a missing `packer-sdc` command, run the below command.
```
go install github.com/hashicorp/packer-plugin-sdk/cmd/packer-sdc@latest
# if not done before, add the your go/bin to your PATH to be able to run Go commands:
export PATH=$PATH:$HOME/go/bin
```

For missing dependencies:
```
go mod tidy
```

When trying `goreleaser` or `go build` locally, a few variables have to be set locally, too. Variables can be checked by running `go env <variable_name>` (e.g. `go env GOOS`.)
```
export API_VERSION=$(go run . describe | jq -r '.api_version')
# the below will be needed for a local go build
export VERSION=<your_version, in 0.0.0 format>
export MODULE_PATH=$(go list -m)
export GOOS=${GOOS:-$(go env GOOS)}
export GOARCH=${GOARCH:-$(go env GOARCH)}
```

Running `goreleaser` requires a Git tag to be populated, as that value is used to fill the plugin's version variable.

To test a plugin, run `<path_to_the_plugin> describe` - that should return a JSON block with its _description_, e.g. `version`.

The Packer plugin can be then validated with Packer by running `packer plugins install -path <path_to_the_generate_plugin> "github.com/zestia/digitalocean-image-lifecycle"` - it also does a `describe` but also confirms all is valid for a Packer plugin.
=======
>>>>>>> 52cb53695a733c333c89dd49e15b4bfb3de77147
