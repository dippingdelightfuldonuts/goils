package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
	"weavelab.xyz/goils/resources"
)

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

	table := resources.CreateTable{
		TableName: "Bob",
		Attributes: []resources.Attribute{
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
		Indexes: []resources.Index{
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

	resource := resources.Resource{
		CreateTable: table,
		CrudOptions: []resources.CrudOption{"show"},
	}

	protoTemplate := template.Must(
		template.New("proto").Funcs(AllFunctions()).Parse(string(protoTemp)),
	)
	err = protoTemplate.Execute(os.Stdout, resource)
	if err != nil {
		fmt.Println("err:", err)
	}
}
