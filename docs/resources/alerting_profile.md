---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "taikun_alerting_profile Resource - terraform-provider-taikun"
subcategory: ""
description: |-
  Taikun Alerting Profile
---

# taikun_alerting_profile (Resource)

Taikun Alerting Profile

## Example Usage

```terraform
resource "taikun_alerting_profile" "foo" {
  # Required
  name     = "foo"
  reminder = "None"

  # Optional
  emails = ["test@example.com", "test@example.org", "test@example.net"]

  is_locked = false

  organization_id = resource.taikun_organization.foo.id

  webhook {
    url = "https://www.example.com"
  }

  webhook {
    header {
      key   = "key"
      value = "value"
    }
    url = "https://www.example.com"
  }

  webhook {
    header {
      key   = "key"
      value = "value"
    }
    header {
      key   = "key2"
      value = "value"
    }
    url = "https://www.example.org"
  }

  integration {
    type  = "Opsgenie"
    url   = "https://www.opsgenie.example"
    token = "secret_token"
  }
  integration {
    type = "MicrosoftTeams"
    url  = "https://www.teams.example"
  }
  integration {
    type  = "Splunk"
    url   = "https://www.splunk.example"
    token = "secret_token"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The alerting profile's name.
- **reminder** (String) The frequency of notifications (HalfHour, Hourly, Daily or None).

### Optional

- **emails** (List of String) The list of e-mails to notify.
- **integration** (Block List) list of alerting integrations (see [below for nested schema](#nestedblock--integration))
- **is_locked** (Boolean) Whether the profile is locked or not. Defaults to `false`.
- **organization_id** (String) The ID of the organization which owns the profile.
- **slack_configuration_id** (String) The ID of the Slack configuration to notify. Defaults to `0`.
- **webhook** (Block Set) The list of webhooks to notify. (see [below for nested schema](#nestedblock--webhook))

### Read-Only

- **created_by** (String) The profile creator.
- **id** (String) The alerting profile's ID.
- **last_modified** (String) The time and date of last modification.
- **last_modified_by** (String) The last user to have modified the profile.
- **organization_name** (String) The name of the organization which owns the profile.
- **slack_configuration_name** (String) The name of the Slack configuration to notify.

<a id="nestedblock--integration"></a>
### Nested Schema for `integration`

Required:

- **type** (String) type of integration (Opsgenie, Pagerduty, Splunk or MicrosoftTeams)
- **url** (String) URL

Optional:

- **token** (String) token (required from Opsgenie, Pagerduty and Splunk) Defaults to ` `.


<a id="nestedblock--webhook"></a>
### Nested Schema for `webhook`

Required:

- **url** (String) The webhook URL.

Optional:

- **header** (Block Set) The list of headers. (see [below for nested schema](#nestedblock--webhook--header))

<a id="nestedblock--webhook--header"></a>
### Nested Schema for `webhook.header`

Required:

- **key** (String) The header key.
- **value** (String) The header value.

## Import

Import is supported using the following syntax:

```shell
# import with Taikun ID
terraform import taikun_alerting_profile.myalertingprofile 42
```