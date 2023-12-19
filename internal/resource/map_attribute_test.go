// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resource

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/code"
	"github.com/hashicorp/terraform-plugin-codegen-spec/resource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/model"
	generatorschema "github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

func TestGeneratorMapAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.MapAttribute
		expected      GeneratorMapAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.MapAttribute is nil"),
		},
		"element-type-bool": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
		},
		"element-type-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"element-type-list-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
		},
		"element-type-map-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
		},
		"element-type-list-object-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name:   "str",
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
		},
		"element-type-object-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name:   "str",
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
		},
		"element-type-object-list-string": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										String: &specschema.StringType{},
									},
								},
							},
						},
					},
				},
			},
		},
		"computed": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"computed_optional": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Computed: true,
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"optional": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"required": {
			input: &resource.MapAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"custom_type": {
			input: &resource.MapAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"deprecation_message": {
			input: &resource.MapAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					DeprecationMessage: "deprecation message",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"description": {
			input: &resource.MapAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"sensitive": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
		},
		"validators": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.MapValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.MapValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/.../myvalidator",
								},
							},
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
		"plan-modifiers": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.MapPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.MapPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							Imports: []code.Import{
								{
									Path: "github.com/.../my_planmodifier",
								},
							},
							SchemaDefinition: "my_planmodifier.Modify()",
						},
					},
				},
			},
		},
		"default": {
			input: &resource.MapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.MapDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
			expected: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.MapDefault{
					Custom: &specschema.CustomDefault{
						Imports: []code.Import{
							{
								Path: "github.com/.../my_default",
							},
						},
						SchemaDefinition: "my_default.Default()",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := NewGeneratorMapAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			// TODO: This prevents misleading failure when ElementType for both got and expected are nil.
			// TODO: Could overwrite comparison using an option to cmp.Diff()?
			if got.MapAttribute.ElementType == nil && testCase.expected.MapAttribute.ElementType == nil {
				return
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapAttribute
		expected      string
		expectedError error
	}{
		"element-type-bool": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.BoolType,
},`,
		},

		"element-type-list": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-list-list": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							List: &specschema.ListType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.ListType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-list-object": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ListType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"element-type-map": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.BoolType,
},
},`,
		},

		"element-type-map-map": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Map: &specschema.MapType{
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.MapType{
ElemType: types.BoolType,
},
},
},`,
		},

		"element-type-map-object": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							Object: &specschema.ObjectType{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "bool",
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.MapType{
ElemType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},`,
		},

		"element-type-object": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},`,
		},

		"element-type-object-object": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "obj",
								Object: &specschema.ObjectType{
									AttributeTypes: specschema.ObjectAttributeTypes{
										{
											Name: "bool",
											Bool: &specschema.BoolType{},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"obj": types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": types.BoolType,
},
},
},
},
},`,
		},

		"element-type-object-list": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "list",
								List: &specschema.ListType{
									ElementType: specschema.ElementType{
										Bool: &specschema.BoolType{},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"list": types.ListType{
ElemType: types.BoolType,
},
},
},
},`,
		},

		"element-type-string": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
},`,
		},

		"custom-type": {
			input: GeneratorMapAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
CustomType: my_custom_type,
},`,
		},

		"associated-external-type": {
			input: GeneratorMapAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.MapAttribute",
					},
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
CustomType: MapAttributeType{
types.MapType{
ElemType: types.StringType,
},
},
},`,
		},

		"custom-type-overriding-associated-external-type": {
			input: GeneratorMapAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.MapAttribute",
					},
				},
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Required: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Required: true,
},`,
		},

		"optional": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Optional: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Computed: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Sensitive: true,
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Sensitive: true,
},`,
		},

		// TODO: Do we need separate description and markdown description?
		"description": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					Description: "description",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorMapAttribute{
				MapAttribute: schema.MapAttribute{
					DeprecationMessage: "deprecated",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: specschema.MapValidators{
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_validator.Validate()",
						},
					},
					{
						Custom: &specschema.CustomValidator{
							SchemaDefinition: "my_other_validator.Validate()",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Validators: []validator.Map{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				PlanModifiers: specschema.MapPlanModifiers{
					{
						Custom: &specschema.CustomPlanModifier{
							SchemaDefinition: "my_plan_modifier.Modify()",
						},
					},
					{
						Custom: &specschema.CustomPlanModifier{
							SchemaDefinition: "my_other_plan_modifier.Modify()",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
PlanModifiers: []planmodifier.Map{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Default: &specschema.MapDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_map_default.Default()",
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.StringType,
Default: my_map_default.Default(),
},`,
		},

		"element-type-bool-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{
						CustomType: &specschema.CustomType{
							Type: "boolCustomType",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: boolCustomType,
},`,
		},

		"element-type-float64-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Float64: &specschema.Float64Type{
						CustomType: &specschema.CustomType{
							Type: "float64CustomType",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: float64CustomType,
},`,
		},

		"element-type-int64-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Int64: &specschema.Int64Type{
						CustomType: &specschema.CustomType{
							Type: "int64CustomType",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: int64CustomType,
},`,
		},

		"element-type-list-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						CustomType: &specschema.CustomType{
							Type: "customListType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: customListType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-map-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						CustomType: &specschema.CustomType{
							Type: "customMapType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: customMapType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-number-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Number: &specschema.NumberType{
						CustomType: &specschema.CustomType{
							Type: "numberCustomType",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: numberCustomType,
},`,
		},

		"element-type-object-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Object: &specschema.ObjectType{
						AttributeTypes: specschema.ObjectAttributeTypes{
							{
								Name: "bool",
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Type: "customBoolType",
									},
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: types.ObjectType{
AttrTypes: map[string]attr.Type{
"bool": customBoolType,
},
},
},`,
		},

		"element-type-set-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					Set: &specschema.SetType{
						CustomType: &specschema.CustomType{
							Type: "customSetType",
						},
						ElementType: specschema.ElementType{
							Bool: &specschema.BoolType{
								CustomType: &specschema.CustomType{
									Type: "customBoolType",
								},
							},
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: customSetType{
ElemType: customBoolType,
},
},`,
		},

		"element-type-string-custom": {
			input: GeneratorMapAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{
						CustomType: &specschema.CustomType{
							Type: "stringCustomType",
						},
					},
				},
			},
			expected: `
"map_attribute": schema.MapAttribute{
ElementType: stringCustomType,
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("map_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorMapAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorMapAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "types.Map",
				TfsdkName: "map_attribute",
			},
		},
		"custom-type": {
			input: GeneratorMapAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "map_attribute",
			},
		},
		"associated-external-type": {
			input: GeneratorMapAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
					},
				},
			},
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "MapAttributeValue",
				TfsdkName: "map_attribute",
			},
		},
		"custom-type-overriding-associated-external-type": {
			input: GeneratorMapAttribute{
				AssociatedExternalType: &generatorschema.AssocExtType{
					AssociatedExternalType: &specschema.AssociatedExternalType{
						Type: "*api.BoolAttribute",
					},
				},
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "MapAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "map_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("map_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}