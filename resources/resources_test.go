package resources

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_newProtoMessage(t *testing.T) {
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

	type args struct {
		resource Resource
		name     string
		typ      string
	}
	tests := []struct {
		name string
		args args
		want ProtoMessage
	}{
		{
			name: "",
			args: args{
				resource: Resource{
					CreateTable: table,
					CrudOptions: []CrudOption{"show"},
				},
				name: "schedule_requests",
				typ:  "show",
			},
			want: ProtoMessage{
				Name: "Bob",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("expected tests", t, func() {
				got := newProtoMessage(tt.args.resource, tt.args.name, tt.args.typ)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}
