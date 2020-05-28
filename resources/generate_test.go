package resources

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GenerateMigration(t *testing.T) {
	type args struct {
		resource Resource
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should generate template given resource",
			args: args{
				resource: Resource{
					CreateTable: CreateTable{
						TableName: "sms",
						Attributes: Attributes{
							{
								Name:     "id",
								Type:     "UUID",
								Nullable: false,
							},
							{
								Name:     "text",
								Type:     "string",
								Nullable: false,
							},
							{
								Name:     "created_at",
								Type:     "date",
								Nullable: false,
							},
							{
								Name:     "auto",
								Type:     "boolean",
								Nullable: false,
							},
						},
					},
				},
			},
			want: goldenFile("generatemigration"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateMigration", t, func() {
				got, err := GenerateMigration(tt.args.resource)
				So(err != nil, ShouldEqual, tt.wantErr)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}

func Test_GenerateProto(t *testing.T) {
	type args struct {
		resource Resource
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should generate template given resource",
			args: args{
				resource: Resource{
					CreateTable: CreateTable{
						TableName: "sms",
						Attributes: Attributes{
							{
								Name:     "id",
								Type:     "UUID",
								Nullable: false,
							},
							{
								Name:     "text",
								Type:     "string",
								Nullable: false,
							},
							{
								Name:     "created_at",
								Type:     "date",
								Nullable: false,
							},
							{
								Name:     "auto",
								Type:     "boolean",
								Nullable: false,
							},
						},
					},
					CrudOptions: []CrudOption{
						"show",
					},
				},
			},
			want: goldenFile("generateproto"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateProto", t, func() {
				got, err := GenerateProto(tt.args.resource)
				So(err != nil, ShouldEqual, tt.wantErr)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}

func Test_GenerateSQL(t *testing.T) {
	type args struct {
		resource Resource
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "should generate sqlc templates given resource",
			args: args{
				resource: Resource{
					CreateTable: CreateTable{
						TableName: "sms",
						Attributes: Attributes{
							{
								Name:     "id",
								Type:     "UUID",
								Nullable: false,
							},
							{
								Name:     "text",
								Type:     "string",
								Nullable: false,
							},
							{
								Name:     "created_at",
								Type:     "date",
								Nullable: false,
							},
							{
								Name:     "auto",
								Type:     "boolean",
								Nullable: false,
							},
						},
					},
					CrudOptions: []CrudOption{
						"show",
						"index",
						"create",
					},
				},
			},
			want: goldenFile("generatesql"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateSQL", t, func() {
				got, err := GenerateSQL(tt.args.resource)
				So(err != nil, ShouldEqual, tt.wantErr)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}
