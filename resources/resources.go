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

type Indexes []Index

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

func (i Indexes) Any(f func(i Index) bool) bool {
	for _, index := range i {
		if f(index) {
			return true
		}
	}

	return false
}

type Attributes []Attribute

type CreateTable struct {
	TableName  string
	Attributes Attributes
	Indexes    Indexes
	Owner      string
}

func (a Attributes) Any(f func(a Attribute) bool) bool {
	for _, attribute := range a {
		if f(attribute) {
			return true
		}
	}

	return false
}

func (a Attributes) Select(f func(a Attribute) bool) []Attribute {
	attrs := make([]Attribute, 0)
	for _, attribute := range a {
		if f(attribute) {
			attrs = append(attrs, attribute)
		}
	}

	return attrs
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
		return newShowProtoMessage(resource)
	case "index":
		return newIndexProtoMessage(resource)
	}
	return pm
}

func newIndexProtoMessage(resource Resource) ProtoMessage {
	suitable := resource.Attributes.Select(func(attr Attribute) bool {
		return attr.Name != "id"
	})

	return ProtoMessage{
		Type:       "index",
		Name:       "List" + strcase.ToCamel(resource.TableName),
		Attributes: suitable,
	}
}

func newShowProtoMessage(resource Resource) ProtoMessage {
	indexed := resource.Attributes.Select(func(attr Attribute) bool {
		return resource.Indexes.Any(func(index Index) bool {
			return index.HasAttribute(attr.Name)
		})
	})

	return ProtoMessage{
		Type:       "show",
		Name:       strcase.ToCamel(resource.TableName),
		Attributes: indexed,
	}
}

func (r Resource) CrudMessages() []ProtoMessage {
	messages := make([]ProtoMessage, len(r.CrudOptions))
	for i, msg := range r.CrudOptions {
		messages[i] = newProtoMessage(r, msg.MessageName())
	}

	return messages
}
