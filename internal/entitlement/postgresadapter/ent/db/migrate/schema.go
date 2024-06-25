// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EntitlementsColumns holds the columns for the "entitlements" table.
	EntitlementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "metadata", Type: field.TypeJSON, Nullable: true, SchemaType: map[string]string{"postgres": "jsonb"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "entitlement_type", Type: field.TypeEnum, Enums: []string{"metered", "static", "boolean"}},
		{Name: "feature_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "subject_key", Type: field.TypeString},
		{Name: "measure_usage_from", Type: field.TypeTime, Nullable: true},
		{Name: "issue_after_reset", Type: field.TypeFloat64, Nullable: true},
		{Name: "is_soft_limit", Type: field.TypeBool, Nullable: true},
		{Name: "config", Type: field.TypeJSON, Nullable: true},
		{Name: "usage_period_interval", Type: field.TypeEnum, Nullable: true, Enums: []string{"DAY", "WEEK", "MONTH", "YEAR"}},
		{Name: "usage_period_anchor", Type: field.TypeTime, Nullable: true},
	}
	// EntitlementsTable holds the schema information for the "entitlements" table.
	EntitlementsTable = &schema.Table{
		Name:       "entitlements",
		Columns:    EntitlementsColumns,
		PrimaryKey: []*schema.Column{EntitlementsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "entitlement_id",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[0]},
			},
			{
				Name:    "entitlement_namespace_id",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[0]},
			},
			{
				Name:    "entitlement_namespace_subject_key",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[8]},
			},
			{
				Name:    "entitlement_namespace_id_subject_key",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[0], EntitlementsColumns[8]},
			},
			{
				Name:    "entitlement_namespace_feature_id_id",
				Unique:  false,
				Columns: []*schema.Column{EntitlementsColumns[1], EntitlementsColumns[7], EntitlementsColumns[0]},
			},
		},
	}
	// UsageResetsColumns holds the columns for the "usage_resets" table.
	UsageResetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(26)"}},
		{Name: "namespace", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "reset_time", Type: field.TypeTime},
		{Name: "entitlement_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(26)"}},
	}
	// UsageResetsTable holds the schema information for the "usage_resets" table.
	UsageResetsTable = &schema.Table{
		Name:       "usage_resets",
		Columns:    UsageResetsColumns,
		PrimaryKey: []*schema.Column{UsageResetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "usage_resets_entitlements_usage_reset",
				Columns:    []*schema.Column{UsageResetsColumns[6]},
				RefColumns: []*schema.Column{EntitlementsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usagereset_id",
				Unique:  false,
				Columns: []*schema.Column{UsageResetsColumns[0]},
			},
			{
				Name:    "usagereset_namespace_entitlement_id",
				Unique:  false,
				Columns: []*schema.Column{UsageResetsColumns[1], UsageResetsColumns[6]},
			},
			{
				Name:    "usagereset_namespace_entitlement_id_reset_time",
				Unique:  false,
				Columns: []*schema.Column{UsageResetsColumns[1], UsageResetsColumns[6], UsageResetsColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EntitlementsTable,
		UsageResetsTable,
	}
)

func init() {
	UsageResetsTable.ForeignKeys[0].RefTable = EntitlementsTable
}
