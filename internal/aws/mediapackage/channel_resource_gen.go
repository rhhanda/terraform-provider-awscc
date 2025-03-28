// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediapackage_channel", channelResource)
}

// channelResource returns the Terraform awscc_mediapackage_channel resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaPackage::Channel resource.
func channelResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) assigned to the Channel.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) assigned to the Channel.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A short text description of the Channel.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A short text description of the Channel.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EgressAccessLogs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration parameters for egress access logging.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"egress_access_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 256),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration parameters for egress access logging.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HlsIngest
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An HTTP Live Streaming (HLS) ingest resource configuration.",
		//	  "properties": {
		//	    "ingestEndpoints": {
		//	      "description": "A list of endpoints to which the source stream should be sent.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An endpoint for ingesting source content for a Channel.",
		//	        "properties": {
		//	          "Id": {
		//	            "description": "The system generated unique identifier for the IngestEndpoint",
		//	            "type": "string"
		//	          },
		//	          "Password": {
		//	            "description": "The system generated password for ingest authentication.",
		//	            "type": "string"
		//	          },
		//	          "Url": {
		//	            "description": "The ingest URL to which the source stream should be sent.",
		//	            "type": "string"
		//	          },
		//	          "Username": {
		//	            "description": "The system generated username for ingest authentication.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Id",
		//	          "Username",
		//	          "Password",
		//	          "Url"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"hls_ingest": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ingestEndpoints
				"ingest_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Id
							"id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The system generated unique identifier for the IngestEndpoint",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Password
							"password": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The system generated password for ingest authentication.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Url
							"url": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The ingest URL to which the source stream should be sent.",
								Required:    true,
							}, /*END ATTRIBUTE*/
							// Property: Username
							"username": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The system generated username for ingest authentication.",
								Required:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "A list of endpoints to which the source stream should be sent.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An HTTP Live Streaming (HLS) ingest resource configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Channel.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Channel.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 256),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngressAccessLogs
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration parameters for egress access logging.",
		//	  "properties": {
		//	    "LogGroupName": {
		//	      "description": "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
		//	      "maxLength": 256,
		//	      "minLength": 1,
		//	      "pattern": "",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"ingress_access_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: LogGroupName
				"log_group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Sets a custom AWS CloudWatch log group name for access logs. If a log group name isn't specified, the defaults are used: /aws/MediaPackage/EgressAccessLogs for egress access logs and /aws/MediaPackage/IngressAccessLogs for ingress access logs.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 256),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration parameters for egress access logging.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
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
				listplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	schema := schema.Schema{
		Description: "Resource schema for AWS::MediaPackage::Channel",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaPackage::Channel").WithTerraformTypeName("awscc_mediapackage_channel")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(false)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                 "Arn",
		"description":         "Description",
		"egress_access_logs":  "EgressAccessLogs",
		"hls_ingest":          "HlsIngest",
		"id":                  "Id",
		"ingest_endpoints":    "ingestEndpoints",
		"ingress_access_logs": "IngressAccessLogs",
		"key":                 "Key",
		"log_group_name":      "LogGroupName",
		"password":            "Password",
		"tags":                "Tags",
		"url":                 "Url",
		"username":            "Username",
		"value":               "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
