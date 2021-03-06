package resources

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GenerateMigration(t *testing.T) {
	MockedTime = time.Date(2020, 6, 15, 12, 0, 0, 0, time.UTC)

	type args struct {
		resource Resource
	}
	tests := []struct {
		name string
		args args
		want GeneratedGroup
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
			want: GeneratedGroup{
				GeneratedResult{
					Output:  goldenFile("generatemigration"),
					FileOut: "20200615120000_create_sms_table.sql",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateMigration", t, func() {
				got := GenerateMigration(tt.args.resource)
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
		name string
		args args
		want GeneratedGroup
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
			want: GeneratedGroup{
				GeneratedResult{
					Output:  goldenFile("generateproto"),
					FileOut: "proto.proto",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateProto", t, func() {
				got := GenerateProto(tt.args.resource)
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
		name string
		args args
		want GeneratedGroup
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
			want: GeneratedGroup{
				GeneratedResult{
					Output:  goldenFile("generatesql"),
					FileOut: "queries.sql",
				},
				GeneratedResult{
					Output:  goldenFile("generatesqlyaml"),
					FileOut: "sqlc.yaml",
				},
				GeneratedResult{
					Output:  goldenFile("generatesqlschema"),
					FileOut: "schema.sql",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateSQL", t, func() {
				got := GenerateSQL(tt.args.resource)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}

func Test_GenerateTests(t *testing.T) {
	type args struct {
		resource Resource
	}
	tests := []struct {
		name string
		args args
		want GeneratedGroup
	}{
		{
			name: "should generate sqlc templates given resource",
			args: args{
				resource: Resource{
					Package: "main",
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
			want: GeneratedGroup{
				GeneratedResult{
					Output:  goldenFile("generatetests"),
					FileOut: "queries_test.go",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("generateTests", t, func() {
				got := GenerateTests(tt.args.resource)
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}
