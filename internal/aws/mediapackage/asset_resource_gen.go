// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediapackage_asset", assetResource)
}

// assetResource returns the Terraform awscc_mediapackage_asset resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaPackage::Asset resource.
func assetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Asset.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Asset.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time the Asset was initially submitted for Ingest.",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time the Asset was initially submitted for Ingest.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EgressEndpoints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of egress endpoints available for the Asset.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "The endpoint URL used to access an Asset using one PackagingConfiguration.",
		//	    "properties": {
		//	      "PackagingConfigurationId": {
		//	        "description": "The ID of the PackagingConfiguration being applied to the Asset.",
		//	        "type": "string"
		//	      },
		//	      "Url": {
		//	        "description": "The URL of the parent manifest for the repackaged Asset.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "PackagingConfigurationId",
		//	      "Url"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"egress_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: PackagingConfigurationId
					"packaging_configuration_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the PackagingConfiguration being applied to the Asset.",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Url
					"url": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The URL of the parent manifest for the repackaged Asset.",
						Required:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The list of egress endpoints available for the Asset.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier for the Asset.",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier for the Asset.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: PackagingGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the PackagingGroup for the Asset.",
		//	  "type": "string"
		//	}
		"packaging_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the PackagingGroup for the Asset.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource ID to include in SPEKE key requests.",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource ID to include in SPEKE key requests.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN of the source object in S3.",
		//	  "type": "string"
		//	}
		"source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN of the source object in S3.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM role_arn used to access the source S3 bucket.",
		//	  "type": "string"
		//	}
		"source_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM role_arn used to access the source S3 bucket.",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of tags associated with a resource",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A collection of tags associated with a resource",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaPackage::Asset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::Asset").WithTerraformTypeName("awscc_mediapackage_asset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                        "Arn",
		"created_at":                 "CreatedAt",
		"egress_endpoints":           "EgressEndpoints",
		"id":                         "Id",
		"key":                        "Key",
		"packaging_configuration_id": "PackagingConfigurationId",
		"packaging_group_id":         "PackagingGroupId",
		"resource_id":                "ResourceId",
		"source_arn":                 "SourceArn",
		"source_role_arn":            "SourceRoleArn",
		"tags":                       "Tags",
		"url":                        "Url",
		"value":                      "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
