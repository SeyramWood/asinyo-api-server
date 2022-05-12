// Code generated by entc, DO NOT EDIT.

package merchant

import (
	"time"
)

const (
	// Label holds the string label denoting the merchant type in the database.
	Label = "merchant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeSupplier holds the string denoting the supplier edge name in mutations.
	EdgeSupplier = "supplier"
	// EdgeRetailer holds the string denoting the retailer edge name in mutations.
	EdgeRetailer = "retailer"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the merchant in the database.
	Table = "merchants"
	// SupplierTable is the table that holds the supplier relation/edge.
	SupplierTable = "supplier_merchants"
	// SupplierInverseTable is the table name for the SupplierMerchant entity.
	// It exists in this package in order to avoid circular dependency with the "suppliermerchant" package.
	SupplierInverseTable = "supplier_merchants"
	// SupplierColumn is the table column denoting the supplier relation/edge.
	SupplierColumn = "merchant_supplier"
	// RetailerTable is the table that holds the retailer relation/edge.
	RetailerTable = "retail_merchants"
	// RetailerInverseTable is the table name for the RetailMerchant entity.
	// It exists in this package in order to avoid circular dependency with the "retailmerchant" package.
	RetailerInverseTable = "retail_merchants"
	// RetailerColumn is the table column denoting the retailer relation/edge.
	RetailerColumn = "merchant_retailer"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "merchants"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "merchant_products"
)

// Columns holds all SQL columns for merchant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPassword,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "merchants"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"merchant_products",
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
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func([]byte) error
)
