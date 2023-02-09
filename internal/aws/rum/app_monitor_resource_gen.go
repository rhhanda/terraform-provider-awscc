// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rum

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"regexp"
)

func init() {
	registry.AddResourceFactory("awscc_rum_app_monitor", appMonitorResource)
}

// appMonitorResource returns the Terraform awscc_rum_app_monitor resource.
// This Terraform resource corresponds to the CloudFormation AWS::RUM::AppMonitor resource.
func appMonitorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AppMonitorConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "AppMonitor configuration",
		//	  "properties": {
		//	    "AllowCookies": {
		//	      "description": "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
		//	      "type": "boolean"
		//	    },
		//	    "EnableXRay": {
		//	      "description": "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
		//	      "type": "boolean"
		//	    },
		//	    "ExcludedPages": {
		//	      "description": "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Page Url",
		//	        "maxLength": 1260,
		//	        "minLength": 1,
		//	        "pattern": "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?\u0026//=]*)",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "FavoritePages": {
		//	      "description": "A list of pages in the RUM console that are to be displayed with a favorite icon.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "GuestRoleArn": {
		//	      "description": "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
		//	      "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
		//	      "type": "string"
		//	    },
		//	    "IdentityPoolId": {
		//	      "description": "The ID of the identity pool that is used to authorize the sending of data to RUM.",
		//	      "maxLength": 55,
		//	      "minLength": 1,
		//	      "pattern": "[\\w-]+:[0-9a-f-]+",
		//	      "type": "string"
		//	    },
		//	    "IncludedPages": {
		//	      "description": "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "Page Url",
		//	        "maxLength": 1260,
		//	        "minLength": 1,
		//	        "pattern": "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?\u0026//=]*)",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 50,
		//	      "minItems": 0,
		//	      "type": "array"
		//	    },
		//	    "MetricDestinations": {
		//	      "description": "An array of structures which define the destinations and the metrics that you want to send.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "An structure which defines the destination and the metrics that you want to send.",
		//	        "properties": {
		//	          "Destination": {
		//	            "description": "Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.",
		//	            "enum": [
		//	              "CloudWatch",
		//	              "Evidently"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "DestinationArn": {
		//	            "description": "Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.",
		//	            "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
		//	            "type": "string"
		//	          },
		//	          "IamRoleArn": {
		//	            "description": "This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.\n\nThis parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.",
		//	            "pattern": "arn:[^:]*:[^:]*:[^:]*:[^:]*:.*",
		//	            "type": "string"
		//	          },
		//	          "MetricDefinitions": {
		//	            "description": "An array of structures which define the metrics that you want to send.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "additionalProperties": false,
		//	              "description": "A single metric definition",
		//	              "properties": {
		//	                "DimensionKeys": {
		//	                  "additionalProperties": false,
		//	                  "description": "Use this field only if you are sending the metric to CloudWatch.\n\nThis field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. Valid values for the entries in this field are the following:\n\n\"metadata.pageId\": \"PageId\"\n\n\"metadata.browserName\": \"BrowserName\"\n\n\"metadata.deviceType\": \"DeviceType\"\n\n\"metadata.osName\": \"OSName\"\n\n\"metadata.countryCode\": \"CountryCode\"\n\n\"event_details.fileType\": \"FileType\"\n\nAll dimensions listed in this field must also be included in EventPattern.",
		//	                  "patternProperties": {
		//	                    "": {
		//	                      "maxLength": 255,
		//	                      "minLength": 1,
		//	                      "pattern": ".*[^\\s].*",
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "EventPattern": {
		//	                  "description": "The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.\n\nWhen you define extended metrics, the metric definition is not valid if EventPattern is omitted.\n\nExample event patterns:\n\n'{ \"event_type\": [\"com.amazon.rum.js_error_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Firefox\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"\u003c\", 2000 ] }] } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], \"countryCode\": [ \"US\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"\u003e=\", 2000, \"\u003c\", 8000 ] }] } }'\n\nIf the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.",
		//	                  "maxLength": 4000,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "Name": {
		//	                  "description": "The name for the metric that is defined in this structure. Valid values are the following:\n\nPerformanceNavigationDuration\n\nPerformanceResourceDuration\n\nNavigationSatisfiedTransaction\n\nNavigationToleratedTransaction\n\nNavigationFrustratedTransaction\n\nWebVitalsCumulativeLayoutShift\n\nWebVitalsFirstInputDelay\n\nWebVitalsLargestContentfulPaint\n\nJsErrorCount\n\nHttpErrorCount\n\nSessionCount",
		//	                  "maxLength": 255,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "UnitLabel": {
		//	                  "description": "The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.",
		//	                  "maxLength": 256,
		//	                  "minLength": 1,
		//	                  "type": "string"
		//	                },
		//	                "ValueKey": {
		//	                  "description": "The field within the event object that the metric value is sourced from.\n\nIf you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.\n\nIf this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.",
		//	                  "maxLength": 256,
		//	                  "minLength": 1,
		//	                  "pattern": ".*",
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Name"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "maxItems": 2000,
		//	            "minItems": 0,
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "required": [
		//	          "Destination"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "maxItems": 20,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "SessionSampleRate": {
		//	      "description": "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
		//	      "maximum": 1,
		//	      "minimum": 0,
		//	      "type": "number"
		//	    },
		//	    "Telemetries": {
		//	      "description": "An array that lists the types of telemetry data that this app monitor is to collect.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "errors",
		//	          "performance",
		//	          "http"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"app_monitor_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AllowCookies
				"allow_cookies": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie. The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EnableXRay
				"enable_x_ray": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If you set this to true, RUM enables xray tracing for the user sessions that RUM samples. RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ExcludedPages
				"excluded_pages": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of URLs in your website or application to exclude from RUM data collection. You can't include both ExcludedPages and IncludedPages in the same operation.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 50),
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 1260),
							stringvalidator.RegexMatches(regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: FavoritePages
				"favorite_pages": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of pages in the RUM console that are to be displayed with a favorite icon.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 50),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: GuestRoleArn
				"guest_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("arn:[^:]*:[^:]*:[^:]*:[^:]*:.*"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IdentityPoolId
				"identity_pool_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the identity pool that is used to authorize the sending of data to RUM.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 55),
						stringvalidator.RegexMatches(regexp.MustCompile("[\\w-]+:[0-9a-f-]+"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IncludedPages
				"included_pages": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "If this app monitor is to collect data from only certain pages in your application, this structure lists those pages. You can't include both ExcludedPages and IncludedPages in the same operation.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(0, 50),
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(1, 1260),
							stringvalidator.RegexMatches(regexp.MustCompile("https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: MetricDestinations
				"metric_destinations": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Destination
							"destination": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Defines the destination to send the metrics to. Valid values are CloudWatch and Evidently. If you specify Evidently, you must also specify the ARN of the Evidently experiment that is to be the destination and an IAM role that has permission to write to the experiment.",
								Required:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"CloudWatch",
										"Evidently",
									),
								}, /*END VALIDATORS*/
							}, /*END ATTRIBUTE*/
							// Property: DestinationArn
							"destination_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Use this parameter only if Destination is Evidently. This parameter specifies the ARN of the Evidently experiment that will receive the extended metrics.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("arn:[^:]*:[^:]*:[^:]*:[^:]*:.*"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: IamRoleArn
							"iam_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "This parameter is required if Destination is Evidently. If Destination is CloudWatch, do not use this parameter.\n\nThis parameter specifies the ARN of an IAM role that RUM will assume to write to the Evidently experiment that you are sending metrics to. This role must have permission to write to that experiment.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.RegexMatches(regexp.MustCompile("arn:[^:]*:[^:]*:[^:]*:[^:]*:.*"), ""),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: MetricDefinitions
							"metric_definitions": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
								NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: DimensionKeys
										"dimension_keys":    // Pattern: ""
										schema.MapAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Description: "Use this field only if you are sending the metric to CloudWatch.\n\nThis field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. Valid values for the entries in this field are the following:\n\n\"metadata.pageId\": \"PageId\"\n\n\"metadata.browserName\": \"BrowserName\"\n\n\"metadata.deviceType\": \"DeviceType\"\n\n\"metadata.osName\": \"OSName\"\n\n\"metadata.countryCode\": \"CountryCode\"\n\n\"event_details.fileType\": \"FileType\"\n\nAll dimensions listed in this field must also be included in EventPattern.",
											Optional:    true,
											Computed:    true,
											PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
												mapplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: EventPattern
										"event_pattern": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The pattern that defines the metric, specified as a JSON object. RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.\n\nWhen you define extended metrics, the metric definition is not valid if EventPattern is omitted.\n\nExample event patterns:\n\n'{ \"event_type\": [\"com.amazon.rum.js_error_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Firefox\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \"<\", 2000 ] }] } }'\n\n'{ \"event_type\": [\"com.amazon.rum.performance_navigation_event\"], \"metadata\": { \"browserName\": [ \"Chrome\", \"Safari\" ], \"countryCode\": [ \"US\" ] }, \"event_details\": { \"duration\": [{ \"numeric\": [ \">=\", 2000, \"<\", 8000 ] }] } }'\n\nIf the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.",
											Optional:    true,
											Computed:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 4000),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Name
										"name": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The name for the metric that is defined in this structure. Valid values are the following:\n\nPerformanceNavigationDuration\n\nPerformanceResourceDuration\n\nNavigationSatisfiedTransaction\n\nNavigationToleratedTransaction\n\nNavigationFrustratedTransaction\n\nWebVitalsCumulativeLayoutShift\n\nWebVitalsFirstInputDelay\n\nWebVitalsLargestContentfulPaint\n\nJsErrorCount\n\nHttpErrorCount\n\nSessionCount",
											Required:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 255),
											}, /*END VALIDATORS*/
										}, /*END ATTRIBUTE*/
										// Property: UnitLabel
										"unit_label": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The CloudWatch metric unit to use for this metric. If you omit this field, the metric is recorded with no unit.",
											Optional:    true,
											Computed:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 256),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: ValueKey
										"value_key": schema.StringAttribute{ /*START ATTRIBUTE*/
											Description: "The field within the event object that the metric value is sourced from.\n\nIf you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.\n\nIf this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event.",
											Optional:    true,
											Computed:    true,
											Validators: []validator.String{ /*START VALIDATORS*/
												stringvalidator.LengthBetween(1, 256),
												stringvalidator.RegexMatches(regexp.MustCompile(".*"), ""),
											}, /*END VALIDATORS*/
											PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
												stringplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
								}, /*END NESTED OBJECT*/
								Description: "An array of structures which define the metrics that you want to send.",
								Optional:    true,
								Computed:    true,
								Validators: []validator.Set{ /*START VALIDATORS*/
									setvalidator.SizeBetween(0, 2000),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
									setplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "An array of structures which define the destinations and the metrics that you want to send.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(0, 20),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SessionSampleRate
				"session_sample_rate": schema.Float64Attribute{ /*START ATTRIBUTE*/
					Description: "Specifies the percentage of user sessions to use for RUM data collection. Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Float64{ /*START VALIDATORS*/
						float64validator.Between(0.000000, 1.000000),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Float64{ /*START PLAN MODIFIERS*/
						float64planmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Telemetries
				"telemetries": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "An array that lists the types of telemetry data that this app monitor is to collect.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"errors",
								"performance",
								"http",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "AppMonitor configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CustomEvents
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "AppMonitor custom events configuration",
		//	  "properties": {
		//	    "Status": {
		//	      "description": "Indicates whether AppMonitor accepts custom events.",
		//	      "enum": [
		//	        "ENABLED",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"custom_events": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether AppMonitor accepts custom events.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ENABLED",
							"DISABLED",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "AppMonitor custom events configuration",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CwLogEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
		//	  "type": "boolean"
		//	}
		"cw_log_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Data collected by RUM is kept by RUM for 30 days and then deleted. This parameter specifies whether RUM sends a copy of this telemetry data to CWLlong in your account. This enables you to keep the telemetry data for more than 30 days, but it does incur CWLlong charges. If you omit this parameter, the default is false",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Domain
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The top-level internet domain name for which your application has administrative authority.",
		//	  "maxLength": 253,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"domain": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The top-level internet domain name for which your application has administrative authority.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 253),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A name for the app monitor",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[\\.\\-_/#A-Za-z0-9]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A name for the app monitor",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 255),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\.\\-_/#A-Za-z0-9]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "minLength": 0,
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
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Assigns one or more tags (key-value pairs) to the app monitor. Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values. Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.You can associate as many as 50 tags with an app monitor.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource Type definition for AWS::RUM::AppMonitor",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RUM::AppMonitor").WithTerraformTypeName("awscc_rum_app_monitor")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"allow_cookies":             "AllowCookies",
		"app_monitor_configuration": "AppMonitorConfiguration",
		"custom_events":             "CustomEvents",
		"cw_log_enabled":            "CwLogEnabled",
		"destination":               "Destination",
		"destination_arn":           "DestinationArn",
		"dimension_keys":            "DimensionKeys",
		"domain":                    "Domain",
		"enable_x_ray":              "EnableXRay",
		"event_pattern":             "EventPattern",
		"excluded_pages":            "ExcludedPages",
		"favorite_pages":            "FavoritePages",
		"guest_role_arn":            "GuestRoleArn",
		"iam_role_arn":              "IamRoleArn",
		"identity_pool_id":          "IdentityPoolId",
		"included_pages":            "IncludedPages",
		"key":                       "Key",
		"metric_definitions":        "MetricDefinitions",
		"metric_destinations":       "MetricDestinations",
		"name":                      "Name",
		"session_sample_rate":       "SessionSampleRate",
		"status":                    "Status",
		"tags":                      "Tags",
		"telemetries":               "Telemetries",
		"unit_label":                "UnitLabel",
		"value":                     "Value",
		"value_key":                 "ValueKey",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
