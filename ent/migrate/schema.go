// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "hash", Type: field.TypeString, Unique: true},
		{Name: "title", Type: field.TypeString},
		{Name: "dt", Type: field.TypeString},
		{Name: "cat", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt},
		{Name: "ext_id", Type: field.TypeString, Nullable: true},
		{Name: "imdb", Type: field.TypeString, Nullable: true},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "item_cat",
				Unique:  true,
				Columns: []*schema.Column{ItemsColumns[4]},
			},
			{
				Name:    "item_ext_id",
				Unique:  true,
				Columns: []*schema.Column{ItemsColumns[6]},
			},
			{
				Name:    "item_imdb",
				Unique:  true,
				Columns: []*schema.Column{ItemsColumns[7]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ItemsTable,
	}
)

func init() {
}
