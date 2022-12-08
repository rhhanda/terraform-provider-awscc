// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package oam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_oam_link", linkDataSource)
}

// linkDataSource returns the Terraform awscc_oam_link data source.
// This Terraform data source corresponds to the CloudFormation AWS::Oam::Link resource.
func linkDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"label": {
			// Property: Label
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"label_template": {
			// Property: LabelTemplate
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 64,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"resource_types": {
			// Property: ResourceTypes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "enum": [
			//	      "AWS::CloudWatch::Metric",
			//	      "AWS::Logs::LogGroup",
			//	      "AWS::XRay::Trace"
			//	    ],
			//	    "type": "string"
			//	  },
			//	  "maxItems": 50,
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Type:     types.SetType{ElemType: types.StringType},
			Computed: true,
		},
		"sink_identifier": {
			// Property: SinkIdentifier
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "Tags to apply to the link",
			//	  "patternProperties": {
			//	    "": {
			//	      "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	      "maxLength": 256,
			//	      "minLength": 0,
			//	      "pattern": "",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "Tags to apply to the link",
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Oam::Link",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Oam::Link").WithTerraformTypeName("awscc_oam_link")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":             "Arn",
		"label":           "Label",
		"label_template":  "LabelTemplate",
		"resource_types":  "ResourceTypes",
		"sink_identifier": "SinkIdentifier",
		"tags":            "Tags",
	})

	v, err := NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}