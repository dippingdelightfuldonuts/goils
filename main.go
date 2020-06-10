package main

import (
	"fmt"

	"weavelab.xyz/goils/resources"
)

func main() {
	fmt.Println("Hello, and welcome to Goils")

	table := resources.CreateTable{
		TableName: "setting",
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

	resource := resources.Resource{
		CreateTable: table,
		CrudOptions: []resources.CrudOption{"show", "index", "create"},
		Package:     "main",
	}

	groups := resources.GeneratedGroups{
		resources.GenerateMigration(resource),
		resources.GenerateProto(resource),
		resources.GenerateSQL(resource),
		resources.GenerateTests(resource),
	}

	groups.Each(func(group resources.GeneratedGroup) {
		if group.AnyErrors() {
			group.PrintErrors()
			return
		}

		group.CreateFiles()
	})
}
