// Code generated by entc, DO NOT EDIT.

package favourite

import (
	"time"
)

const (
	// Label holds the string label denoting the favourite type in the database.
	Label = "favourite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// EdgeAgent holds the string denoting the agent edge name in mutations.
	EdgeAgent = "agent"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// Table holds the table name of the favourite in the database.
	Table = "favourites"
	// MerchantTable is the table that holds the merchant relation/edge.
	MerchantTable = "favourites"
	// MerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantInverseTable = "merchants"
	// MerchantColumn is the table column denoting the merchant relation/edge.
	MerchantColumn = "merchant_favourites"
	// AgentTable is the table that holds the agent relation/edge.
	AgentTable = "favourites"
	// AgentInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentInverseTable = "agents"
	// AgentColumn is the table column denoting the agent relation/edge.
	AgentColumn = "agent_favourites"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "favourites"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_favourites"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "favourites"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_favourites"
)

// Columns holds all SQL columns for favourite fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "favourites"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"agent_favourites",
	"customer_favourites",
	"merchant_favourites",
	"product_favourites",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
