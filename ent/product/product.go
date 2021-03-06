// Code generated by entc, DO NOT EDIT.

package product

import (
	"time"
)

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldPromoPrice holds the string denoting the promo_price field in the database.
	FieldPromoPrice = "promo_price"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// EdgeFavourites holds the string denoting the favourites edge name in mutations.
	EdgeFavourites = "favourites"
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// EdgeMajor holds the string denoting the major edge name in mutations.
	EdgeMajor = "major"
	// EdgeMinor holds the string denoting the minor edge name in mutations.
	EdgeMinor = "minor"
	// Table holds the table name of the product in the database.
	Table = "products"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "order_details"
	// OrdersInverseTable is the table name for the OrderDetail entity.
	// It exists in this package in order to avoid circular dependency with the "orderdetail" package.
	OrdersInverseTable = "order_details"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "product_orders"
	// FavouritesTable is the table that holds the favourites relation/edge.
	FavouritesTable = "favourites"
	// FavouritesInverseTable is the table name for the Favourite entity.
	// It exists in this package in order to avoid circular dependency with the "favourite" package.
	FavouritesInverseTable = "favourites"
	// FavouritesColumn is the table column denoting the favourites relation/edge.
	FavouritesColumn = "product_favourites"
	// MerchantTable is the table that holds the merchant relation/edge.
	MerchantTable = "products"
	// MerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantInverseTable = "merchants"
	// MerchantColumn is the table column denoting the merchant relation/edge.
	MerchantColumn = "merchant_products"
	// MajorTable is the table that holds the major relation/edge.
	MajorTable = "products"
	// MajorInverseTable is the table name for the ProductCategoryMajor entity.
	// It exists in this package in order to avoid circular dependency with the "productcategorymajor" package.
	MajorInverseTable = "product_category_majors"
	// MajorColumn is the table column denoting the major relation/edge.
	MajorColumn = "product_category_major_products"
	// MinorTable is the table that holds the minor relation/edge.
	MinorTable = "products"
	// MinorInverseTable is the table name for the ProductCategoryMinor entity.
	// It exists in this package in order to avoid circular dependency with the "productcategoryminor" package.
	MinorInverseTable = "product_category_minors"
	// MinorColumn is the table column denoting the minor relation/edge.
	MinorColumn = "product_category_minor_products"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldPrice,
	FieldPromoPrice,
	FieldQuantity,
	FieldUnit,
	FieldDescription,
	FieldImage,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"merchant_products",
	"product_category_major_products",
	"product_category_minor_products",
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice float64
	// DefaultQuantity holds the default value on creation for the "quantity" field.
	DefaultQuantity uint32
	// UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	UnitValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
)
