package resources

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	. "weavelab.xyz/goils/extstring"
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
				Name:      "ListSettingsRequest",
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

func TestProtoMessage_CrudAttributes(t *testing.T) {
	type fields struct {
		Name       string
		ModelName  string
		Type       string
		Attributes []Attribute
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "given Type is 'show' returns id",
			fields: fields{
				Type: "show",
			},
			want: []string{
				"id",
			},
		},
		{
			name: "given Type is 'index' returns id",
			fields: fields{
				Type: "index",
			},
			want: []string{},
		},
		{
			name: "given Type is 'create' returns id",
			fields: fields{
				Type: "create",
			},
			want: []string{},
		},
		{
			name: "given Type is 'update' returns id",
			fields: fields{
				Type: "update",
			},
			want: []string{
				"id",
			},
		},
		{
			name: "given Type is 'delete' returns id",
			fields: fields{
				Type: "delete",
			},
			want: []string{
				"id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("CrudAttributes", t, func() {
				pm := ProtoMessage{
					Name:       tt.fields.Name,
					ModelName:  tt.fields.ModelName,
					Type:       tt.fields.Type,
					Attributes: tt.fields.Attributes,
				}
				got := pm.CrudAttributes()
				So(got, ShouldResemble, tt.want)
			})
		})
	}
}

func TestProtoMessage_MethodName(t *testing.T) {
	type fields struct {
		Name       string
		ModelName  string
		Type       string
		Attributes []Attribute
		Verb       ExtString
		Noun       ExtString
	}
	tests := []struct {
		name   string
		fields fields
		want   ExtString
	}{
		{
			name: "Returns MethodName representing 'List' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "list",
				Noun:       "book",
			},
			want: "ListBooks",
		},
		{
			name: "Returns MethodName representing 'Get' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "get",
				Noun:       "book",
			},
			want: "GetBook",
		},
		{
			name: "Returns MethodName representing 'Create' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "create",
				Noun:       "book",
			},
			want: "CreateBook",
		},
		{
			name: "Returns MethodName representing 'Update' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "update",
				Noun:       "book",
			},
			want: "UpdateBook",
		},
		{
			name: "Returns MethodName representing 'Rename' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "rename",
				Noun:       "book",
			},
			want: "RenameBook",
		},
		{
			name: "Returns MethodName representing 'Delete' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "delete",
				Noun:       "book",
			},
			want: "DeleteBook",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("MethodName", t, func() {

				pm := ProtoMessage{
					Name:       tt.fields.Name,
					ModelName:  tt.fields.ModelName,
					Type:       tt.fields.Type,
					Attributes: tt.fields.Attributes,
					Verb:       tt.fields.Verb,
					Noun:       tt.fields.Noun,
				}

				So(pm.MethodName(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestProtoMessage_RequestMessageName(t *testing.T) {
	type fields struct {
		Name       string
		ModelName  string
		Type       string
		Attributes []Attribute
		Verb       ExtString
		Noun       ExtString
	}
	tests := []struct {
		name   string
		fields fields
		want   ExtString
	}{
		{
			name: "Returns RequestMessageName representing 'List' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "list",
				Noun:       "book",
			},
			want: "ListBooksRequest",
		},
		{
			name: "Returns RequestMessageName representing 'Get' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "get",
				Noun:       "book",
			},
			want: "GetBookRequest",
		},
		{
			name: "Returns RequestMessageName representing 'Create' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "create",
				Noun:       "book",
			},
			want: "CreateBookRequest",
		},
		{
			name: "Returns RequestMessageName representing 'Update' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "update",
				Noun:       "book",
			},
			want: "UpdateBookRequest",
		},
		{
			name: "Returns RequestMessageName representing 'Rename' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "rename",
				Noun:       "book",
			},
			want: "RenameBookRequest",
		},
		{
			name: "Returns RequestMessageName representing 'Delete' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "delete",
				Noun:       "book",
			},
			want: "DeleteBookRequest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("RequestMessageName", t, func() {
				pm := ProtoMessage{
					Name:       tt.fields.Name,
					ModelName:  tt.fields.ModelName,
					Type:       tt.fields.Type,
					Attributes: tt.fields.Attributes,
					Verb:       tt.fields.Verb,
					Noun:       tt.fields.Noun,
				}
				So(pm.RequestMessageName(), ShouldResemble, tt.want)
			})
		})
	}
}

func TestProtoMessage_ResponseMessageName(t *testing.T) {
	type fields struct {
		Name       string
		ModelName  string
		Type       string
		Attributes []Attribute
		Verb       ExtString
		Noun       ExtString
	}
	tests := []struct {
		name   string
		fields fields
		want   ExtString
	}{
		{
			name: "Returns ResponseMessageName representing 'List' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "list",
				Noun:       "book",
			},
			want: "ListBooksResponse",
		},
		{
			name: "Returns ResponseMessageName representing 'Get' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "get",
				Noun:       "book",
			},
			want: "Book",
		},
		{
			name: "Returns ResponseMessageName representing 'Create' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "create",
				Noun:       "book",
			},
			want: "Book",
		},
		{
			name: "Returns ResponseMessageName representing 'Update' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "update",
				Noun:       "book",
			},
			want: "Book",
		},
		{
			name: "Returns ResponseMessageName representing 'Rename' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "rename",
				Noun:       "book",
			},
			want: "RenameBookResponse",
		},
		{
			name: "Returns ResponseMessageName representing 'Delete' ProtoMessage",
			fields: fields{
				Name:       "book", // TODO: revisit Name and ModelName and Noun
				ModelName:  "book",
				Type:       "",            // TODO: revisit Type
				Attributes: []Attribute{}, // ignored
				Verb:       "delete",
				Noun:       "book",
			},
			want: "google.protobuf.Empty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convey("RequestMessageName", t, func() {
				pm := ProtoMessage{
					Name:       tt.fields.Name,
					ModelName:  tt.fields.ModelName,
					Type:       tt.fields.Type,
					Attributes: tt.fields.Attributes,
					Verb:       tt.fields.Verb,
					Noun:       tt.fields.Noun,
				}
				So(pm.ResponseMessageName(), ShouldResemble, tt.want)
			})
		})
	}
}
