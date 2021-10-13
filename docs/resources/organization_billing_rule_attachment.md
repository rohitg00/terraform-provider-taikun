---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "taikun_organization_billing_rule_attachment Resource - terraform-provider-taikun"
subcategory: ""
description: |-
  Taikun Organization - Billing Rule Attachment
---

# taikun_organization_billing_rule_attachment (Resource)

Taikun Organization - Billing Rule Attachment

## Example Usage

```terraform
resource "taikun_billing_credential" "foo" {
  name                = "foo"
  prometheus_password = "password"
  prometheus_url      = "url"
  prometheus_username = "username"
}

resource "taikun_billing_rule" "foo" {
  name        = "foo"
  metric_name = "coredns_forward_request_duration_seconds"
  price       = 1
  type        = "Sum"

  billing_credential_id = resource.taikun_billing_credential.foo.id

  label {
    key   = "key"
    value = "value"
  }
}

resource "taikun_organization" "foo" {
  name          = "foo"
  full_name     = "foo"
  discount_rate = 100
}

resource "taikun_organization_billing_rule_attachment" "foo" {
  # Required
  billing_rule_id = resource.taikun_billing_rule.foo.id
  organization_id = resource.taikun_organization.foo.id

  # Optional
  discount_rate = 100
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **billing_rule_id** (String) Id of the billing rule.
- **organization_id** (String) Id of the organisation.

### Optional

- **discount_rate** (Number) Discount rate in percents (0-1000 %), 100 equals one. Defaults to `100`.
- **id** (String) The ID of this resource.

### Read-Only

- **organization_name** (String) Name of the organisation.

