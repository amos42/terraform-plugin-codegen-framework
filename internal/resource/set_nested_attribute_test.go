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

func TestGeneratorSetNestedAttribute_New(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         *resource.SetNestedAttribute
		expected      GeneratorSetNestedAttribute
		expectedError error
	}{
		"nil": {
			expectedError: fmt.Errorf("*resource.SetNestedAttribute is nil"),
		},
		"attribute-nil": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "empty",
						},
					},
				},
			},
			expectedError: fmt.Errorf("attribute type not defined: %+v", resource.Attribute{
				Name: "empty",
			}),
		},
		"attributes-bool": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "bool_attribute",
							Bool: &resource.BoolAttribute{
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool_attribute": GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-list-bool": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "list_attribute",
							List: &resource.ListAttribute{
								ComputedOptionalRequired: "optional",
								ElementType: specschema.ElementType{
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
			expected: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list_attribute": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
		},
		"attributes-list-nested-bool": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							SetNested: &resource.SetNestedAttribute{
								NestedObject: resource.NestedAttributeObject{
									Attributes: []resource.Attribute{
										{
											Name: "nested_bool",
											Bool: &resource.BoolAttribute{
												ComputedOptionalRequired: "computed",
											},
										},
									},
								},
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSetNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"nested_bool": GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Computed: true,
										},
									},
								},
							},
							SetNestedAttribute: schema.SetNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"attributes-object-bool": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "object_attribute",
							Object: &resource.ObjectAttribute{
								AttributeTypes: specschema.ObjectAttributeTypes{
									{
										Name: "obj_bool",
										Bool: &specschema.BoolType{},
									},
								},
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"object_attribute": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "obj_bool",
									Bool: &specschema.BoolType{},
								},
							},
						},
					},
				},
			},
		},
		"attributes-single-nested-bool": {
			input: &resource.SetNestedAttribute{
				NestedObject: resource.NestedAttributeObject{
					Attributes: []resource.Attribute{
						{
							Name: "nested_attribute",
							SingleNested: &resource.SingleNestedAttribute{
								Attributes: []resource.Attribute{
									{
										Name: "nested_bool",
										Bool: &resource.BoolAttribute{
											ComputedOptionalRequired: "computed",
										},
									},
								},
								ComputedOptionalRequired: "optional",
							},
						},
					},
				},
			},
			expected: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_attribute": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"nested_bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Computed: true,
									},
								},
							},
							SingleNestedAttribute: schema.SingleNestedAttribute{
								Optional: true,
							},
						},
					},
				},
			},
		},
		"computed": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "computed",
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Computed: true,
				},
			},
		},
		"computed_optional": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "computed_optional",
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Computed: true,
					Optional: true,
				},
			},
		},
		"optional": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "optional",
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
		},
		"required": {
			input: &resource.SetNestedAttribute{
				ComputedOptionalRequired: "required",
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Required: true,
				},
			},
		},
		"custom_type": {
			input: &resource.SetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
			expected: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/",
					},
					Type:      "my_type",
					ValueType: "myvalue_type",
				},
			},
		},
		"deprecation_message": {
			input: &resource.SetNestedAttribute{
				DeprecationMessage: pointer("deprecation message"),
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					DeprecationMessage: "deprecation message",
				},
			},
		},
		"description": {
			input: &resource.SetNestedAttribute{
				Description: pointer("description"),
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Description:         "description",
					MarkdownDescription: "description",
				},
			},
		},
		"sensitive": {
			input: &resource.SetNestedAttribute{
				Sensitive: pointer(true),
			},
			expected: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Sensitive: true,
				},
			},
		},
		"validators": {
			input: &resource.SetNestedAttribute{
				Validators: specschema.SetValidators{
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
			expected: GeneratorSetNestedAttribute{
				Validators: specschema.SetValidators{
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
			input: &resource.SetNestedAttribute{
				PlanModifiers: specschema.SetPlanModifiers{
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
			expected: GeneratorSetNestedAttribute{
				PlanModifiers: specschema.SetPlanModifiers{
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
			input: &resource.SetNestedAttribute{
				Default: &specschema.SetDefault{
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
			expected: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
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

			got, err := NewGeneratorSetNestedAttribute(testCase.input)

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetNestedAttribute_Imports(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    GeneratorSetNestedAttribute
		expected []code.Import
	}{
		"default": {
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-without-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-without-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-and-nested-object-custom-type-without-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-and-nested-object-custom-type-with-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "",
					},
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-custom-type-with-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/attribute",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"custom-type-with-import-with-nested-object-custom-type-with-import": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Import: &code.Import{
						Path: "github.com/my_account/my_project/attribute",
					},
				},
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Import: &code.Import{
							Path: "github.com/my_account/my_project/nested_object",
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: "github.com/my_account/my_project/attribute",
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-set": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"set": GeneratorSetAttribute{
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				}, {
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-set-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"set": GeneratorSetAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_list",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-set-with-custom-type-with-element-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"set": GeneratorSetAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_list",
								},
							},
							ElementType: specschema.ElementType{
								Bool: &specschema.BoolType{
									CustomType: &specschema.CustomType{
										Import: &code.Import{
											Path: "github.com/my_account/my_project/bool",
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_list",
				},
				{
					Path: "github.com/my_account/my_project/bool",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"obj": GeneratorObjectAttribute{
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
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-with-custom-type-with-attribute-with-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"obj": GeneratorObjectAttribute{
							CustomType: &specschema.CustomType{
								Import: &code.Import{
									Path: "github.com/my_account/my_project/nested_object",
								},
							},
							AttributeTypes: specschema.ObjectAttributeTypes{
								{
									Name: "bool",
									Bool: &specschema.BoolType{
										CustomType: &specschema.CustomType{
											Import: &code.Import{
												Path: "github.com/my_account/my_project/bool",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: "github.com/my_account/my_project/nested_object",
				},
				{
					Path: "github.com/my_account/my_project/bool",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-nil": {
			input: GeneratorSetNestedAttribute{
				Validators: specschema.SetValidators{
					{
						Custom: nil,
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				Validators: specschema.SetValidators{
					{
						Custom: &specschema.CustomValidator{},
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				Validators: specschema.SetValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "",
								},
							},
						},
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"validator-custom-import": {
			input: GeneratorSetNestedAttribute{
				Validators: specschema.SetValidators{
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myotherproject/myvalidators/validator",
								},
							},
						},
					},
					{
						Custom: &specschema.CustomValidator{
							Imports: []code.Import{
								{
									Path: "github.com/myproject/myvalidators/validator",
								},
							},
						},
					},
				}},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.ValidatorImport,
				},
				{
					Path: "github.com/myotherproject/myvalidators/validator",
				},
				{
					Path: "github.com/myproject/myvalidators/validator",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-validator-custom-nil": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: nil,
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-validator-custom-import-nil": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-validator-custom-import-empty-string": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "",
									},
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
		"nested-object-validator-custom-import": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "github.com/myotherproject/myvalidators/validator",
									},
								},
							},
						},
						{
							Custom: &specschema.CustomValidator{
								Imports: []code.Import{
									{
										Path: "github.com/myproject/myvalidators/validator",
									},
								},
							},
						},
					},
				},
			},
			expected: []code.Import{
				{
					Path: generatorschema.TypesImport,
				},
				{
					Path: generatorschema.ValidatorImport,
				},
				{
					Path: "github.com/myotherproject/myvalidators/validator",
				},
				{
					Path: "github.com/myproject/myvalidators/validator",
				},
				{
					Path: generatorschema.AttrImport,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.input.Imports().All()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetNestedAttribute_Schema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetNestedAttribute
		expected      string
		expectedError error
	}{
		"attribute-bool": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"bool": GeneratorBoolAttribute{
							BoolAttribute: schema.BoolAttribute{
								Optional: true,
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"list": GeneratorListAttribute{
							ListAttribute: schema.ListAttribute{
								Optional: true,
							},
							ElementType: specschema.ElementType{
								String: &specschema.StringType{},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"list": schema.ListAttribute{
ElementType: types.StringType,
Optional: true,
},
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-list-nested": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_list_nested": GeneratorSetNestedAttribute{
							NestedObject: GeneratorNestedAttributeObject{
								Attributes: generatorschema.GeneratorAttributes{
									"bool": GeneratorBoolAttribute{
										BoolAttribute: schema.BoolAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_list_nested": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: NestedListNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedListNestedValue{}.AttributeTypes(ctx),
},
},
},
},
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-object": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"object": GeneratorObjectAttribute{
							ObjectAttribute: schema.ObjectAttribute{
								Optional: true,
							},
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
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"object": schema.ObjectAttribute{
AttributeTypes: map[string]attr.Type{
"str": types.StringType,
},
Optional: true,
},
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"attribute-single-nested-bool": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Attributes: generatorschema.GeneratorAttributes{
						"nested_single_nested": GeneratorSingleNestedAttribute{
							Attributes: generatorschema.GeneratorAttributes{
								"bool": GeneratorBoolAttribute{
									BoolAttribute: schema.BoolAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
"nested_single_nested": schema.SingleNestedAttribute{
Attributes: map[string]schema.Attribute{
"bool": schema.BoolAttribute{
Optional: true,
},
},
CustomType: NestedSingleNestedType{
ObjectType: types.ObjectType{
AttrTypes: NestedSingleNestedValue{}.AttributeTypes(ctx),
},
},
},
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
},`,
		},

		"custom-type": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					Type: "my_custom_type",
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
CustomType: my_custom_type,
},`,
		},

		"required": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Required: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Required: true,
},`,
		},

		"optional": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Optional: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Optional: true,
},`,
		},

		"computed": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Computed: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Computed: true,
},`,
		},

		"sensitive": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Sensitive: true,
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Sensitive: true,
},`,
		},

		"description": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					Description: "description",
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Description: "description",
MarkdownDescription: "description",
},`,
		},

		"deprecation-message": {
			input: GeneratorSetNestedAttribute{
				SetNestedAttribute: schema.SetNestedAttribute{
					DeprecationMessage: "deprecated",
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
DeprecationMessage: "deprecated",
},`,
		},

		"validators": {
			input: GeneratorSetNestedAttribute{
				Validators: specschema.SetValidators{
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
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Validators: []validator.Set{
my_validator.Validate(),
my_other_validator.Validate(),
},
},`,
		},

		"nested-object-custom-type": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					CustomType: &specschema.CustomType{
						Type: "my_custom_type",
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: my_custom_type,
},
},`,
		},

		"nested-object-validators": {
			input: GeneratorSetNestedAttribute{
				NestedObject: GeneratorNestedAttributeObject{
					Validators: specschema.ObjectValidators{
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
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
Validators: []validator.Object{
my_validator.Validate(),
my_other_validator.Validate(),
},
},
},`,
		},

		"plan-modifiers": {
			input: GeneratorSetNestedAttribute{
				PlanModifiers: specschema.SetPlanModifiers{
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
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
PlanModifiers: []planmodifier.Set{
my_plan_modifier.Modify(),
my_other_plan_modifier.Modify(),
},
},`,
		},

		"default-custom": {
			input: GeneratorSetNestedAttribute{
				Default: &specschema.SetDefault{
					Custom: &specschema.CustomDefault{
						SchemaDefinition: "my_set_default.Default()",
					},
				},
			},
			expected: `
"set_nested_attribute": schema.SetNestedAttribute{
NestedObject: schema.NestedAttributeObject{
Attributes: map[string]schema.Attribute{
},
CustomType: SetNestedAttributeType{
ObjectType: types.ObjectType{
AttrTypes: SetNestedAttributeValue{}.AttributeTypes(ctx),
},
},
},
Default: my_set_default.Default(),
},`,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.Schema("set_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGeneratorSetNestedAttribute_ModelField(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input         GeneratorSetNestedAttribute
		expected      model.Field
		expectedError error
	}{
		"default": {
			expected: model.Field{
				Name:      "SetNestedAttribute",
				ValueType: "types.Set",
				TfsdkName: "set_nested_attribute",
			},
		},
		"custom-type": {
			input: GeneratorSetNestedAttribute{
				CustomType: &specschema.CustomType{
					ValueType: "my_custom_value_type",
				},
			},
			expected: model.Field{
				Name:      "SetNestedAttribute",
				ValueType: "my_custom_value_type",
				TfsdkName: "set_nested_attribute",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.input.ModelField("set_nested_attribute")

			if diff := cmp.Diff(err, testCase.expectedError, equateErrorMessage); diff != "" {
				t.Errorf("unexpected error: %s", diff)
			}

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}