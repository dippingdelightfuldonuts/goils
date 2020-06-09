package resources

import (
	"bytes"
	"fmt"
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

type Template struct {
	label        string
	templateFile string
}
type Templates []Template

func NewTemplate(label string, templateFile string) Template {
	return Template{
		label:        label,
		templateFile: templateFile,
	}
}

func GenerateTemplates(resource Resource, templates ...Template) GeneratedResult {
	return Templates(templates).Run(resource)
}

type GeneratedResult struct {
	Output []string
	Error  error
}
type GeneratedResults []GeneratedResult

func (g GeneratedResult) HasError() bool {
	return g.Error != nil
}

func (g GeneratedResult) PrintError() {
	fmt.Println("err:", g.Error)
}

func (t Templates) Run(resource Resource) GeneratedResult {
	var runtimeError error
	generated := make([]string, len(t))

	for i, templ := range t {
		gen, err := generateStandardTemplate(resource, templ.label, templ.templateFile)
		if err != nil {
			runtimeError = err
			break
		}
		generated[i] = gen
	}

	return GeneratedResult{
		Output: generated,
		Error:  runtimeError,
	}
}

func GenerateMigration(resource Resource) GeneratedResult {
	return GenerateTemplates(
		resource,
		NewTemplate("createTableTemplate", createTableTemplate),
	)
}

func GenerateProto(resource Resource) GeneratedResult {
	return GenerateTemplates(
		resource,
		NewTemplate("grpcMessageTemplate", grpcMessageTemplate),
	)
}

func GenerateSQL(resource Resource) GeneratedResult {
	return GenerateTemplates(
		resource,
		NewTemplate("sqlTemplate", sqlTemplate),
		NewTemplate("sqlYamlTemplate", sqlYamlTemplate),
	)
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
