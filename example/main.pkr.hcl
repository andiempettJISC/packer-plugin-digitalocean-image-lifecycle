packer {}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = ["sources.null.basic-example"]

  post-processor "digitalocean-image-lifecycle" {
    name_prefix = "mpz-ubuntu-18-04-x64-main-mysql"
    days_older_than = 9
  }
}