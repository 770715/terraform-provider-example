# Terraform Provider Example

This is a Terraform provider for managing example resources.

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.24

## Building The Provider

```bash
go build -o terraform-provider-example
```

## Using the Provider

```terraform
terraform {
  required_providers {
    example = {
      source  = "yourorg/example"
      version = "~> 1.0"
    }
  }
}

provider "example" {
  endpoint = "https://api.example.com"
}
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed.

To compile the provider, run:

```bash
make build
```

To install the provider locally for testing:

```bash
make install
```

To run tests:

```bash
make test
```

To run acceptance tests:

```bash
make testacc
```

## License

MPL-2.0
