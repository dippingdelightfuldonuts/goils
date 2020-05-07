package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

type AttributeType string

func (a AttributeType) ToProto() string {
	switch a {
	case "UUID":
		return "shared.UUID"
	}
	return ""
}

type Attribute struct {
	Name     string
	Type     AttributeType
	Nullable bool
}

func (a Attribute) ToTemplate() string {
	nullableStr := ""
	if !a.Nullable {
		nullableStr = " NOT NULL"
	}

	return fmt.Sprintf("%s %s%s", a.Name, a.Type, nullableStr)
}

func (a Attribute) ToProto() string {
	return fmt.Sprintf("%s %s", a.Type.ToProto(), strings.ReplaceAll(strcase.ToCamel(a.Name), "id", "ID"))
}

type Index struct {
	Name    string
	Type    string // (i.e. BTREE)
	Columns []string
}

func (i Index) SqlIndex() string {
	return strings.Join(i.Columns, ", ")
}

func (i Index) HasAttribute(attribute string) bool {
	for _, column := range i.Columns {
		if column == attribute {
			return true
		}
	}

	return false
}

type CreateTable struct {
	TableName  string
	Attributes []Attribute
	Indexes    []Index
	Owner      string
}

type CrudOption string

func (c CrudOption) MessageName() string {
	switch c {
	case "show":
		return string(c)
	}

	return ""
}

type Resource struct {
	CreateTable
	CrudOptions []CrudOption
}

type ProtoMessage struct {
	Name       string
	Type       string
	Attributes []Attribute
}

func newProtoMessage(resource Resource, name string, typ string) ProtoMessage {
	pm := ProtoMessage{
		Type: typ,
	}

	switch typ {
	case "show":
		var indexed []Attribute
		for _, attr := range resource.Attributes {
			for _, index := range resource.Indexes {
				fmt.Println("attr:", attr, "index:", index)
				if index.HasAttribute(attr.Name) {
					indexed = append(indexed, attr)
				}
			}
		}
		pm.Name = resource.TableName // todo: maybe expand string so titleize, downcase, etc accessible
		pm.Attributes = indexed
	}
	return pm
}

func (r Resource) CrudMessages() []ProtoMessage {
	messages := make([]ProtoMessage, len(r.CrudOptions))
	for i, msg := range r.CrudOptions {
		messages[i] = newProtoMessage(r, string(msg), msg.MessageName())
	}

	return messages
}

func AllFunctions() template.FuncMap {
	return template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"camelcase": func(a string) string {
			return strcase.ToCamel(a)
		},
	}
}

func main() {
	fmt.Println("Hello, and welcome to Goils")

	data, err := ioutil.ReadFile("templates/database/create_table.sql.tmpl")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	protoTemp, err := ioutil.ReadFile("templates/grpc/message.proto.tmpl")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	table := CreateTable{
		TableName: "Bob",
		Attributes: []Attribute{
			{
				Name:     "id",
				Type:     "UUID",
				Nullable: false,
			},
			{
				Name:     "locationid",
				Type:     "UUID",
				Nullable: false,
			},
		},
		Indexes: []Index{
			{
				Name:    "idx_appt_type_location_id",
				Type:    "BTREE",
				Columns: []string{"locationid", "id"},
			},
		},
		Owner: "schedule",
	}
	t := template.Must(
		template.New("letter").Funcs(AllFunctions()).Parse(string(data)),
	)
	err = t.Execute(os.Stdout, table)
	if err != nil {
		fmt.Println("err:", err)
	}

	resource := Resource{
		CreateTable: table,
		CrudOptions: []CrudOption{"show"},
	}

	protoTemplate := template.Must(
		template.New("proto").Funcs(AllFunctions()).Parse(string(protoTemp)),
	)
	err = protoTemplate.Execute(os.Stdout, resource)
	if err != nil {
		fmt.Println("err:", err)
	}
}
