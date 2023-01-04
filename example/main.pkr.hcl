packer {}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["sources.null.basic-example"]

  post-processor "digitalocean-image-lifecycle" {
    name_prefix = "example-image"
    days_older_than = 9
  }
}