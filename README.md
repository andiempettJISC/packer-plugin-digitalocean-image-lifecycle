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
      source  = "github.com/androidwiltron/digitalocean-image-lifecycle"
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