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
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeSupplier holds the string denoting the supplier edge name in mutations.
	EdgeSupplier = "supplier"
	// EdgeRetailer holds the string denoting the retailer edge name in mutations.
	EdgeRetailer = "retailer"
	// EdgeStore holds the string denoting the store edge name in mutations.
	EdgeStore = "store"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeAddresses holds the string denoting the addresses edge name in mutations.
	EdgeAddresses = "addresses"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// EdgeBaskets holds the string denoting the baskets edge name in mutations.
	EdgeBaskets = "baskets"
	// EdgeFavourites holds the string denoting the favourites edge name in mutations.
	EdgeFavourites = "favourites"
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
	// StoreTable is the table that holds the store relation/edge.
	StoreTable = "merchant_stores"
	// StoreInverseTable is the table name for the MerchantStore entity.
	// It exists in this package in order to avoid circular dependency with the "merchantstore" package.
	StoreInverseTable = "merchant_stores"
	// StoreColumn is the table column denoting the store relation/edge.
	StoreColumn = "merchant_store"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "merchant_products"
	// AddressesTable is the table that holds the addresses relation/edge.
	AddressesTable = "addresses"
	// AddressesInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressesInverseTable = "addresses"
	// AddressesColumn is the table column denoting the addresses relation/edge.
	AddressesColumn = "merchant_addresses"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "merchant_orders"
	// BasketsTable is the table that holds the baskets relation/edge.
	BasketsTable = "baskets"
	// BasketsInverseTable is the table name for the Basket entity.
	// It exists in this package in order to avoid circular dependency with the "basket" package.
	BasketsInverseTable = "baskets"
	// BasketsColumn is the table column denoting the baskets relation/edge.
	BasketsColumn = "merchant_baskets"
	// FavouritesTable is the table that holds the favourites relation/edge.
	FavouritesTable = "favourites"
	// FavouritesInverseTable is the table name for the Favourite entity.
	// It exists in this package in order to avoid circular dependency with the "favourite" package.
	FavouritesInverseTable = "favourites"
	// FavouritesColumn is the table column denoting the favourites relation/edge.
	FavouritesColumn = "merchant_favourites"
)

// Columns holds all SQL columns for merchant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPassword,
	FieldType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)
