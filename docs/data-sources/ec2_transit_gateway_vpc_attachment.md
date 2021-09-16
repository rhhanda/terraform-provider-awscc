---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_transit_gateway_vpc_attachment Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::EC2::TransitGatewayVpcAttachment
---

# awscc_ec2_transit_gateway_vpc_attachment (Data Source)

Data Source schema for AWS::EC2::TransitGatewayVpcAttachment



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) Uniquely identifies the resource.

### Read-Only

- **add_subnet_ids** (List of String)
- **options** (Attributes) The options for the transit gateway vpc attachment. (see [below for nested schema](#nestedatt--options))
- **remove_subnet_ids** (List of String)
- **subnet_ids** (List of String)
- **tags** (Attributes List) (see [below for nested schema](#nestedatt--tags))
- **transit_gateway_id** (String)
- **vpc_id** (String)

<a id="nestedatt--options"></a>
### Nested Schema for `options`

Read-Only:

- **appliance_mode_support** (String) Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable
- **dns_support** (String) Indicates whether to enable DNS Support for Vpc Attachment. Valid Values: enable | disable
- **ipv_6_support** (String) Indicates whether to enable Ipv6 Support for Vpc Attachment. Valid Values: enable | disable


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Read-Only:

- **key** (String)
- **value** (String)

