package resources

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	createTableTemplate = "templates/database/create_table.sql.tmpl"
	grpcMessageTemplate = "templates/grpc/message.proto.tmpl"
	sqlTemplate         = "templates/database/sqlc.tmpl"
	sqlYamlTemplate     = "templates/database/sqlc.yaml.tmpl"
	sqlSchemeTemplate   = "templates/database/sqlc.schema.tmpl"
	sqlTestTemplate     = "templates/testing/sql.test.tmpl"

	directory         = "output"
	templateTimestamp = "{timestamp}"
)

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"camelcase": func(a string) string {
			return strcase.ToCamel(a)
		},
		"pluralize": func(a string) string {
			if strings.HasSuffix(a, "s") {
				return a + "es"
			}

			return a + "s"
		},
		"join": func(a []string) string {
			return strings.Join(a, ", ")
		},
		"empty": func(a []string) bool {
			if len(a) == 0 {
				return true
			}

			return false
		},
		"present": func(a []string) bool {
			if len(a) > 0 {
				return true
			}

			return false
		},
	}
}

type Template struct {
	Label        string
	TemplateFile string
	FileOut      string
}
type Templates []Template

func NewTemplate(label string, templateFile string, fileOut string) Template {
	return Template{
		Label:        label,
		TemplateFile: templateFile,
		FileOut:      fileOut,
	}
}

func GenerateTemplates(resource Resource, templates ...Template) GeneratedGroup {
	return Templates(templates).Run(resource)
}

type GeneratedResult struct {
	Output  string
	FileOut string
	Error   error
}

func (g GeneratedResult) HasError() bool {
	return g.Error != nil
}

func (g GeneratedResult) PrintError() {
	fmt.Println("err:", g.Error)
}

func (g GeneratedResult) CreateFile() error {
	return ioutil.WriteFile(filepath.Join(directory, g.FileOut), []byte(g.Output), 0644)
}

// parseFileOut takes file name string and replaces templates with template values (i.e. date)
func parseFileOut(fileOut string) string {
	if strings.Contains(fileOut, templateTimestamp) {
		return strings.ReplaceAll(fileOut, templateTimestamp, CurrentTime().Format("20060102150405"))
	}

	return fileOut
}

type GeneratedGroup []GeneratedResult

func (g GeneratedGroup) CreateFiles() {
	for i, group := range g {
		err := group.CreateFile()
		if err != nil {
			fmt.Println("error:", i, ",", err)
		}
	}
}

func (g GeneratedGroup) AnyErrors() bool {
	for _, group := range g {
		if group.HasError() {
			return true
		}
	}
	return false
}

func (g GeneratedGroup) PrintErrors() {
	for i, group := range g {
		if group.HasError() {
			fmt.Println("group:", i, " error:", group.Error)
		}
	}
}

type GeneratedGroups []GeneratedGroup

func (gs GeneratedGroups) Each(f func(r GeneratedGroup)) {
	for _, g := range gs {
		f(g)
	}
}

func (t Templates) Run(resource Resource) GeneratedGroup {
	generated := make(GeneratedGroup, len(t))

	for i, templ := range t {
		gen, err := generateStandardTemplate(resource, templ.Label, templ.TemplateFile)

		generated[i] = GeneratedResult{
			Output:  gen,
			FileOut: templ.FileOut,
			Error:   err,
		}
	}

	return generated
}

func GenerateMigration(resource Resource) GeneratedGroup {
	// TODO: use template system with filename too
	output := parseFileOut("{timestamp}_create_" + resource.TableName + "_table.sql")

	return GenerateTemplates(
		resource,
		NewTemplate("createTableTemplate", createTableTemplate, output),
	)
}

func GenerateProto(resource Resource) GeneratedGroup {
	return GenerateTemplates(
		resource,
		NewTemplate("grpcMessageTemplate", grpcMessageTemplate, "proto.proto"),
	)
}

// GenerateSQL these generate templates for sqlc
func GenerateSQL(resource Resource) GeneratedGroup {
	return GenerateTemplates(
		resource,
		NewTemplate("sqlTemplate", sqlTemplate, "queries.sql"),
		NewTemplate("sqlYamlTemplate", sqlYamlTemplate, "sqlc.yaml"),
		NewTemplate("sqlSchemeTemplate", sqlSchemeTemplate, "schema.sql"),
	)
}

func GenerateTests(resource Resource) GeneratedGroup {
	return GenerateTemplates(
		resource,
		NewTemplate("sqlTestTemplate", sqlTestTemplate, "queries_test.go"),
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
