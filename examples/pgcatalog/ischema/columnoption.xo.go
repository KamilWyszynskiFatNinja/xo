// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/KamilWyszynskiFatNinja/xo/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ColumnOption represents a row from 'information_schema.column_options'.
type ColumnOption struct {
	TableCatalog pgtypes.SQLIdentifier `json:"table_catalog"` // table_catalog
	TableSchema  pgtypes.SQLIdentifier `json:"table_schema"`  // table_schema
	TableName    pgtypes.SQLIdentifier `json:"table_name"`    // table_name
	ColumnName   pgtypes.SQLIdentifier `json:"column_name"`   // column_name
	OptionName   pgtypes.SQLIdentifier `json:"option_name"`   // option_name
	OptionValue  pgtypes.CharacterData `json:"option_value"`  // option_value
}
