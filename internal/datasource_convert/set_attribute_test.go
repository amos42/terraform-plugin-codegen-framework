package datasource_convert

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	specschema "github.com/hashicorp/terraform-plugin-codegen-spec/schema"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func TestConvertSetAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *datasource.SetAttribute
		expected      datasource_generate.GeneratorSetAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*datasource.SetAttribute is nil"),
		},
		"element-type-nil": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{},
			},
			expectedError: fmt.Errorf("element type is not defined: %+v", specschema.ElementType{}),
		},
		"element-type-bool": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					Bool: &specschema.BoolType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.BoolType,
				},
			},
		},
		"element-type-string": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.StringType,
				},
			},
		},
		"element-type-list-string": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.ListType{
						ElemType: types.StringType,
					},
				},
			},
		},
		"element-type-map-string": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					Map: &specschema.MapType{
						ElementType: specschema.ElementType{
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.MapType{
						ElemType: types.StringType,
					},
				},
			},
		},
		"element-type-list-object-string": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					List: &specschema.ListType{
						ElementType: specschema.ElementType{
							Object: []specschema.ObjectAttributeType{
								{
									Name:   "str",
									String: &specschema.StringType{},
								},
							},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.ListType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"str": types.StringType,
							},
						},
					},
				},
			},
		},
		"element-type-object-string": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
						{
							Name:   "str",
							String: &specschema.StringType{},
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"str": types.StringType,
						},
					},
				},
			},
		},
		"element-type-object-list-string": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					Object: []specschema.ObjectAttributeType{
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
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.ObjectType{
						AttrTypes: map[string]attr.Type{
							"list": types.ListType{
								ElemType: types.StringType,
							},
						},
					},
				},
			},
		},
		"computed": {
			input: &datasource.SetAttribute{
				ComputedOptionalRequired: "computed",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Computed:    true,
					ElementType: types.StringType,
				},
			},
		},
		"computed_optional": {
			input: &datasource.SetAttribute{
				ComputedOptionalRequired: "computed_optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Computed:    true,
					Optional:    true,
					ElementType: types.StringType,
				},
			},
		},
		"optional": {
			input: &datasource.SetAttribute{
				ComputedOptionalRequired: "optional",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Optional:    true,
					ElementType: types.StringType,
				},
			},
		},
		"required": {
			input: &datasource.SetAttribute{
				ComputedOptionalRequired: "required",
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Required:    true,
					ElementType: types.StringType,
				},
			},
		},
		"custom_type": {
			input: &datasource.SetAttribute{
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.StringType,
				},
				CustomType: &specschema.CustomType{
					Import:    pointer("github.com/"),
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &datasource.SetAttribute{
				DeprecationMessage: pointer("deprecation message"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					DeprecationMessage: "deprecation message",
					ElementType:        types.StringType,
				},
			},
		},
		"description": {
			input: &datasource.SetAttribute{
				Description: pointer("description"),
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Description:         "description",
					MarkdownDescription: "description",
					ElementType:         types.StringType,
				},
			},
		},
		"sensitive": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Sensitive: pointer(true),
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					Sensitive:   true,
					ElementType: types.StringType,
				},
			},
		},
		"validators": {
			input: &datasource.SetAttribute{
				ElementType: specschema.ElementType{
					String: &specschema.StringType{},
				},
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
			expected: datasource_generate.GeneratorSetAttribute{
				SetAttribute: schema.SetAttribute{
					ElementType: types.StringType,
				},
				Validators: []specschema.SetValidator{
					{
						Custom: &specschema.CustomValidator{
							Import:           pointer("github.com/.../myvalidator"),
							SchemaDefinition: "myvalidator.Validate()",
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := convertSetAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			// TODO: This prevents misleading failure when ElementType for both got and expected are nil.
			// TODO: Could overwrite comparison using an option to cmp.Diff()?
			if got.SetAttribute.ElementType == nil && testCase.expected.SetAttribute.ElementType == nil {
				return
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}