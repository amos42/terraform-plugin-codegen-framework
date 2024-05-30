// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package scaffold

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/hashicorp/terraform-plugin-codegen-framework/internal/schema"
)

// DataSourceBytes will create scaffolding Go code bytes for a Terraform Plugin Framework data source
func DataSourceBytes(dataSourceIdentifier schema.FrameworkIdentifier, packageName string, templateDir string) ([]byte, error) {
	var t *template.Template
	var err error
	if len(templateDir) > 0 {
		var pattern = filepath.Join(templateDir, "*.gotmpl")
		t = template.Must(template.ParseGlob(pattern))
	} else {
		t, err = template.New("data_source_scaffold").Parse(dataSourceScaffoldGoTemplate)
	}
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	templateData := struct {
		PackageName string
		NameSnake   string
		NameCamel   string
		NamePascal  string
	}{
		PackageName: packageName,
		NameSnake:   string(dataSourceIdentifier),
		NameCamel:   dataSourceIdentifier.ToCamelCase(),
		NamePascal:  dataSourceIdentifier.ToPascalCase(),
	}

	err = t.ExecuteTemplate(&buf, "data_source_scaffold", templateData)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
