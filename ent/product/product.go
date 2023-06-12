// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
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
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldUnit holds the string denoting the unit field in the database.
	FieldUnit = "unit"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldBestDeal holds the string denoting the best_deal field in the database.
	FieldBestDeal = "best_deal"
	// EdgeOrderDetails holds the string denoting the order_details edge name in mutations.
	EdgeOrderDetails = "order_details"
	// EdgeFavourites holds the string denoting the favourites edge name in mutations.
	EdgeFavourites = "favourites"
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// EdgeMajor holds the string denoting the major edge name in mutations.
	EdgeMajor = "major"
	// EdgeMinor holds the string denoting the minor edge name in mutations.
	EdgeMinor = "minor"
	// EdgePriceModel holds the string denoting the price_model edge name in mutations.
	EdgePriceModel = "price_model"
	// Table holds the table name of the product in the database.
	Table = "products"
	// OrderDetailsTable is the table that holds the order_details relation/edge.
	OrderDetailsTable = "order_details"
	// OrderDetailsInverseTable is the table name for the OrderDetail entity.
	// It exists in this package in order to avoid circular dependency with the "orderdetail" package.
	OrderDetailsInverseTable = "order_details"
	// OrderDetailsColumn is the table column denoting the order_details relation/edge.
	OrderDetailsColumn = "product_order_details"
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
	// PriceModelTable is the table that holds the price_model relation/edge.
	PriceModelTable = "products"
	// PriceModelInverseTable is the table name for the PriceModel entity.
	// It exists in this package in order to avoid circular dependency with the "pricemodel" package.
	PriceModelInverseTable = "price_models"
	// PriceModelColumn is the table column denoting the price_model relation/edge.
	PriceModelColumn = "price_model_model"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldPrice,
	FieldPromoPrice,
	FieldWeight,
	FieldQuantity,
	FieldUnit,
	FieldDescription,
	FieldImage,
	FieldBestDeal,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "products"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"merchant_products",
	"price_model_model",
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
	// DefaultPromoPrice holds the default value on creation for the "promo_price" field.
	DefaultPromoPrice float64
	// DefaultWeight holds the default value on creation for the "weight" field.
	DefaultWeight uint32
	// DefaultQuantity holds the default value on creation for the "quantity" field.
	DefaultQuantity uint32
	// UnitValidator is a validator for the "unit" field. It is called by the builders before save.
	UnitValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// ImageValidator is a validator for the "image" field. It is called by the builders before save.
	ImageValidator func(string) error
	// DefaultBestDeal holds the default value on creation for the "best_deal" field.
	DefaultBestDeal uint64
)

// OrderOption defines the ordering options for the Product queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByPromoPrice orders the results by the promo_price field.
func ByPromoPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPromoPrice, opts...).ToFunc()
}

// ByWeight orders the results by the weight field.
func ByWeight(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWeight, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByUnit orders the results by the unit field.
func ByUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnit, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByImage orders the results by the image field.
func ByImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImage, opts...).ToFunc()
}

// ByBestDeal orders the results by the best_deal field.
func ByBestDeal(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBestDeal, opts...).ToFunc()
}

// ByOrderDetailsCount orders the results by order_details count.
func ByOrderDetailsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderDetailsStep(), opts...)
	}
}

// ByOrderDetails orders the results by order_details terms.
func ByOrderDetails(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderDetailsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByFavouritesCount orders the results by favourites count.
func ByFavouritesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFavouritesStep(), opts...)
	}
}

// ByFavourites orders the results by favourites terms.
func ByFavourites(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFavouritesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMerchantField orders the results by merchant field.
func ByMerchantField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMerchantStep(), sql.OrderByField(field, opts...))
	}
}

// ByMajorField orders the results by major field.
func ByMajorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMajorStep(), sql.OrderByField(field, opts...))
	}
}

// ByMinorField orders the results by minor field.
func ByMinorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMinorStep(), sql.OrderByField(field, opts...))
	}
}

// ByPriceModelField orders the results by price_model field.
func ByPriceModelField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPriceModelStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderDetailsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderDetailsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrderDetailsTable, OrderDetailsColumn),
	)
}
func newFavouritesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FavouritesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FavouritesTable, FavouritesColumn),
	)
}
func newMerchantStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MerchantInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MerchantTable, MerchantColumn),
	)
}
func newMajorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MajorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MajorTable, MajorColumn),
	)
}
func newMinorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MinorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MinorTable, MinorColumn),
	)
}
func newPriceModelStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PriceModelInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PriceModelTable, PriceModelColumn),
	)
}
