---
page_title: "Example Provider"
subcategory: ""
description: |-
  Terraform provider for managing example resources.
---

# Example Provider

The Example provider demonstrates Terraform Plugin Framework v6 capabilities.

## Example Usage

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

## Schema

### Optional

- `endpoint` (String) API endpoint for the example service.
