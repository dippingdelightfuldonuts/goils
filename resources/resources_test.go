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
				Name: "Setting",
				Type: "show",
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
				Name: "ListSetting",
				Type: "index",
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
