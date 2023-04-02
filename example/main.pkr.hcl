// In development comment out the `required_plugins` block 
// and move the built go binary to the example folder to use
packer {
  required_plugins {
    digitalocean-image-lifecycle = {
      version = ">=v1.0.0"
      source  = "github.com/androidwiltron/digitalocean-image-lifecycle"
    }
  }
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["sources.null.basic-example"]

  post-processor "digitalocean-image-lifecycle" {
    name_prefix     = "example-image"
    days_older_than = 9
    dry_run         = false
  }
}