terraform {
  required_providers {
    example = {
      source = "yourorg/example"
    }
  }
}

provider "example" {
  endpoint = "https://api.example.com"
}
