---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_logs_account_policy Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::Logs::AccountPolicy
---

# awscc_logs_account_policy (Data Source)

Data Source schema for AWS::Logs::AccountPolicy



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `account_id` (String) User account id
- `policy_document` (String) The body of the policy document you want to use for this topic.

You can only add one policy per PolicyType.

The policy must be in JSON string format.

Length Constraints: Maximum length of 30720
- `policy_name` (String) The name of the account policy
- `policy_type` (String) Type of the policy.
- `scope` (String) Scope for policy application
