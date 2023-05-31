package datasource_convert

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-codegen-spec/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"

	"github/hashicorp/terraform-provider-code-generator/internal/datasource_generate"
)

func convertBoolAttribute(a *datasource.BoolAttribute) (datasource_generate.GeneratorBoolAttribute, error) {
	if a == nil {
		return datasource_generate.GeneratorBoolAttribute{}, fmt.Errorf("*datasource.BoolAttribute is nil")
	}

	return datasource_generate.GeneratorBoolAttribute{
		BoolAttribute: schema.BoolAttribute{
			Required:            isRequired(a.ComputedOptionalRequired),
			Optional:            isOptional(a.ComputedOptionalRequired),
			Computed:            isComputed(a.ComputedOptionalRequired),
			Sensitive:           isSensitive(a.Sensitive),
			Description:         description(a.Description),
			MarkdownDescription: description(a.Description),
			DeprecationMessage:  deprecationMessage(a.DeprecationMessage),
		},

		CustomType: a.CustomType,
		Validators: a.Validators,
	}, nil
}