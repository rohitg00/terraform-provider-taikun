---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "taikun_showback_rule Data Source - terraform-provider-taikun"
subcategory: ""
description: |-
  Get a showback rule by its id.
---

# taikun_showback_rule (Data Source)

Get a showback rule by its id.

## Example Usage

```terraform
data "taikun_showback_rule" "foo" {
  id = "42"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The id of the showback rule.

### Read-Only

- **created_by** (String) The creator of the showback rule.
- **global_alert_limit** (Number) Set limit of alerts for all projects.
- **kind** (String) Type of the showback rule. `General` (data source is taikun) or `External` (data source is external see `showback_credential_id`)
- **label** (List of Object) Labels linked to this showback rule. (see [below for nested schema](#nestedatt--label))
- **last_modified** (String) Time of last modification.
- **last_modified_by** (String) The last user who modified the showback rule.
- **metric_name** (String) The metric name.
- **name** (String) The name of the showback rule.
- **organization_id** (String) The id of the organization which owns the showback rule.
- **organization_name** (String) The name of the organization which owns the showback rule.
- **price** (Number) Billing in CZK per selected unit.
- **project_alert_limit** (Number) Set limit of alerts for one project.
- **showback_credential_id** (String) Id of the showback rule.
- **showback_credential_name** (String) Name of the showback rule.
- **type** (String) Type of the showback rule. `Count` (calculate package as unit) or `Sum` (calculate per quantity)

<a id="nestedatt--label"></a>
### Nested Schema for `label`

Read-Only:

- **key** (String)
- **value** (String)

