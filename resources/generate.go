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
	s := ""
	buffer := bytes.NewBufferString(s)

	temp, err := ioutil.ReadFile(filepath.Join(createTableTemplate))
	if err != nil {
		temp, err = ioutil.ReadFile(filepath.Join("..", createTableTemplate))
		if err != nil {
			return "", err
		}
	}

	t := template.Must(
		template.New("createTableTemplate").Funcs(templateFunctions()).Parse(string(temp)),
	)
	err = t.Execute(buffer, resource.CreateTable)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
