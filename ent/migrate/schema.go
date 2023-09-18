// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TeasColumns holds the columns for the "teas" table.
	TeasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(100)"}},
		{Name: "name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(100)"}},
		{Name: "color", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(100)"}},
	}
	// TeasTable holds the schema information for the "teas" table.
	TeasTable = &schema.Table{
		Name:       "teas",
		Columns:    TeasColumns,
		PrimaryKey: []*schema.Column{TeasColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TeasTable,
	}
)

func init() {
}
