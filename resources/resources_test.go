package resources

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_newProtoMessage(t *testing.T) {
	table := CreateTable{
		TableName: "setting",
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
			{
				Name:     "name",
				Type:     "string",
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

	type args struct {
		resource Resource
		typ      string
	}
	tests := []struct {
		name string
		args args
		want ProtoMessage
	}{
		{
			name: "given 'show' creates suitable ProtoMessage",
			args: args{
				resource: Resource{
					CreateTable: table,
					CrudOptions: []CrudOption{"show"},
				},
				typ: "show",
			},
			want: ProtoMessage{
				Name:      "Setting",
				ModelName: "Setting",
				Type:      "show",
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
			},
		},
		{
			name: "given 'index' creates suitable ProtoMessage",
			args: args{
				resource: Resource{
					CreateTable: table,
					CrudOptions: []CrudOption{"show"},
				},
				typ: "index",
			},
			want: ProtoMessage{
				Name:      "ListSetting",
				ModelName: "Setting",
				Type:      "index",
				Attributes: []Attribute{
					{
						Name:     "locationid",
						Type:     "UUID",
						Nullable: false,
					},
					{
						Name:     "name",
						Type:     "string",
						Nullable: false,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("newProtoMessage", t, func() {
				got := newProtoMessage(tt.args.resource, tt.args.typ)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}

func TestProtoMessage_CrudFuncName(t *testing.T) {
	type fields struct {
		Name       string
		ModelName  string
		Type       string // TODO: maybe Type should be renamed to CrudType or CrudAction
		Attributes []Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "given Type is 'show', it returns Get+ModelName",
			fields: fields{
				ModelName: "setting",
				Type:      "show",
			},
			want: "GetSetting",
		},
		{
			name: "given Type is 'index', it returns List+ModelName",
			fields: fields{
				ModelName: "setting",
				Type:      "index",
			},
			want: "ListSetting",
		},
		{
			name: "given Type is 'update', it returns Update+ModelName",
			fields: fields{
				ModelName: "setting",
				Type:      "update",
			},
			want: "UpdateSetting",
		},
		{
			name: "given Type is 'create', it returns Create+ModelName",
			fields: fields{
				ModelName: "setting",
				Type:      "create",
			},
			want: "CreateSetting",
		},
		{
			name: "given Type is 'delete', it returns Delete+ModelName",
			fields: fields{
				ModelName: "setting",
				Type:      "delete",
			},
			want: "DeleteSetting",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("CrudFuncName", t, func() {
				pm := ProtoMessage{
					Name:       tt.fields.Name,
					ModelName:  tt.fields.ModelName,
					Type:       tt.fields.Type,
					Attributes: tt.fields.Attributes,
				}
				got := pm.CrudFuncName()
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}
