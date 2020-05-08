package resources

import (
	"fmt"
	"strings"

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

func newProtoMessage(resource Resource, typ string) ProtoMessage {
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
		pm.Name = strcase.ToCamel(resource.TableName) // todo: maybe expand string so titleize, downcase, etc accessible
		pm.Attributes = indexed
	case "index":
		var suitable []Attribute
		for _, attr := range resource.Attributes {
			if attr.Name == "id" {
				continue
			}
			suitable = append(suitable, attr)
		}
		pm.Name = "List" + strcase.ToCamel(resource.TableName)
		pm.Attributes = suitable
	}
	return pm
}

func (r Resource) CrudMessages() []ProtoMessage {
	messages := make([]ProtoMessage, len(r.CrudOptions))
	for i, msg := range r.CrudOptions {
		messages[i] = newProtoMessage(r, msg.MessageName())
	}

	return messages
}
