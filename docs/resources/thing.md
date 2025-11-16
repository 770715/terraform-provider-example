---
page_title: "example_thing Resource - terraform-provider-example"
subcategory: ""
description: |-
  Manages a example thing.
---

# example_thing (Resource)

Manages a example thing demonstrating resource lifecycle.

## Example Usage

```terraform
resource "example_thing" "example" {
  name        = "my-thing"
  description = "Example resource"
}
```

## Schema

### Required

- `name` (String) Name of the thing.

### Optional

- `description` (String) Description of the thing.

### Read-Only

- `id` (String) Unique identifier.
- `last_updated` (String) Timestamp of last update.
