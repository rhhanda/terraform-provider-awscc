// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ec2_subnet_cidr_block", subnetCidrBlockResource)
}

// subnetCidrBlockResource returns the Terraform awscc_ec2_subnet_cidr_block resource.
// This Terraform resource corresponds to the CloudFormation AWS::EC2::SubnetCidrBlock resource.
func subnetCidrBlockResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Information about the IPv6 association.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Information about the IPv6 association.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Ipv6CidrBlock
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length",
		//	  "maxLength": 42,
		//	  "type": "string"
		//	}
		"ipv_6_cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(42),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "The AWS::EC2::SubnetCidrBlock resource creates association between subnet and IPv6 CIDR",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::SubnetCidrBlock").WithTerraformTypeName("awscc_ec2_subnet_cidr_block")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"id":               "Id",
		"ipv_6_cidr_block": "Ipv6CidrBlock",
		"subnet_id":        "SubnetId",
	})

	opts = opts.IsImmutableType(true)

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
