package resources

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	createTableTemplate = "templates/database/create_table.sql.tmpl"
	grpcMessageTemplate = "templates/grpc/message.proto.tmpl"
	sqlTemplate         = "templates/database/sqlc.tmpl"
	sqlYamlTemplate     = "templates/database/sqlc.yaml.tmpl"
)

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"camelcase": func(a string) string {
			return strcase.ToCamel(a)
		},
	}
}

func GenerateMigration(resource Resource) (string, error) {
	return generateStandardTemplate(resource, "createTableTemplate", createTableTemplate)
}

// TODO: make all generate funcs follow same pattern (i.e. array of funcs to call)
func GenerateProto(resource Resource) (string, error) {
	return generateStandardTemplate(resource, "grpcMessageTemplate", grpcMessageTemplate)
}

func GenerateSQL(resource Resource) ([]string, error) {
	str, err := generateSQLTemplate(resource)
	if err != nil {
		return nil, err
	}

	str2, err := generateSQLYamlTemplate(resource)
	if err != nil {
		return nil, err
	}

	return []string{str, str2}, nil
}

func generateSQLTemplate(resource Resource) (string, error) {
	return generateStandardTemplate(resource, "sqlTemplate", sqlTemplate)
}

func generateSQLYamlTemplate(resource Resource) (string, error) {
	return generateStandardTemplate(resource, "sqlYamlTemplate", sqlYamlTemplate)
}

func generateStandardTemplate(resource Resource, label string, templateFile string) (string, error) {
	s := ""
	buffer := bytes.NewBufferString(s)

	temp, err := ioutil.ReadFile(filepath.Join(templateFile))
	if err != nil {
		temp, err = ioutil.ReadFile(filepath.Join("..", templateFile))
		if err != nil {
			return "", err
		}
	}

	t := template.Must(
		template.New(label).Funcs(templateFunctions()).Parse(string(temp)),
	)
	err = t.Execute(buffer, resource)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
