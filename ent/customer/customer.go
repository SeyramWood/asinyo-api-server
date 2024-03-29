// Code generated by ent, DO NOT EDIT.

package customer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the customer type in the database.
	Label = "customer"
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
	// EdgeBusiness holds the string denoting the business edge name in mutations.
	EdgeBusiness = "business"
	// EdgeIndividual holds the string denoting the individual edge name in mutations.
	EdgeIndividual = "individual"
	// EdgeAddresses holds the string denoting the addresses edge name in mutations.
	EdgeAddresses = "addresses"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// EdgeFavourites holds the string denoting the favourites edge name in mutations.
	EdgeFavourites = "favourites"
	// EdgeNotifications holds the string denoting the notifications edge name in mutations.
	EdgeNotifications = "notifications"
	// EdgePurchaseRequest holds the string denoting the purchase_request edge name in mutations.
	EdgePurchaseRequest = "purchase_request"
	// EdgeAdmin holds the string denoting the admin edge name in mutations.
	EdgeAdmin = "admin"
	// Table holds the table name of the customer in the database.
	Table = "customers"
	// BusinessTable is the table that holds the business relation/edge.
	BusinessTable = "business_customers"
	// BusinessInverseTable is the table name for the BusinessCustomer entity.
	// It exists in this package in order to avoid circular dependency with the "businesscustomer" package.
	BusinessInverseTable = "business_customers"
	// BusinessColumn is the table column denoting the business relation/edge.
	BusinessColumn = "customer_business"
	// IndividualTable is the table that holds the individual relation/edge.
	IndividualTable = "individual_customers"
	// IndividualInverseTable is the table name for the IndividualCustomer entity.
	// It exists in this package in order to avoid circular dependency with the "individualcustomer" package.
	IndividualInverseTable = "individual_customers"
	// IndividualColumn is the table column denoting the individual relation/edge.
	IndividualColumn = "customer_individual"
	// AddressesTable is the table that holds the addresses relation/edge.
	AddressesTable = "addresses"
	// AddressesInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressesInverseTable = "addresses"
	// AddressesColumn is the table column denoting the addresses relation/edge.
	AddressesColumn = "customer_addresses"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "customer_orders"
	// FavouritesTable is the table that holds the favourites relation/edge.
	FavouritesTable = "favourites"
	// FavouritesInverseTable is the table name for the Favourite entity.
	// It exists in this package in order to avoid circular dependency with the "favourite" package.
	FavouritesInverseTable = "favourites"
	// FavouritesColumn is the table column denoting the favourites relation/edge.
	FavouritesColumn = "customer_favourites"
	// NotificationsTable is the table that holds the notifications relation/edge. The primary key declared below.
	NotificationsTable = "customer_notifications"
	// NotificationsInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	NotificationsInverseTable = "notifications"
	// PurchaseRequestTable is the table that holds the purchase_request relation/edge.
	PurchaseRequestTable = "purchase_requests"
	// PurchaseRequestInverseTable is the table name for the PurchaseRequest entity.
	// It exists in this package in order to avoid circular dependency with the "purchaserequest" package.
	PurchaseRequestInverseTable = "purchase_requests"
	// PurchaseRequestColumn is the table column denoting the purchase_request relation/edge.
	PurchaseRequestColumn = "customer_purchase_request"
	// AdminTable is the table that holds the admin relation/edge.
	AdminTable = "customers"
	// AdminInverseTable is the table name for the Admin entity.
	// It exists in this package in order to avoid circular dependency with the "admin" package.
	AdminInverseTable = "admins"
	// AdminColumn is the table column denoting the admin relation/edge.
	AdminColumn = "admin_customers"
)

// Columns holds all SQL columns for customer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPassword,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "customers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"admin_customers",
}

var (
	// NotificationsPrimaryKey and NotificationsColumn2 are the table columns denoting the
	// primary key for the notifications relation (M2M).
	NotificationsPrimaryKey = []string{"customer_id", "notification_id"}
)

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
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)

// OrderOption defines the ordering options for the Customer queries.
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

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByBusinessField orders the results by business field.
func ByBusinessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessStep(), sql.OrderByField(field, opts...))
	}
}

// ByIndividualField orders the results by individual field.
func ByIndividualField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newIndividualStep(), sql.OrderByField(field, opts...))
	}
}

// ByAddressesCount orders the results by addresses count.
func ByAddressesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAddressesStep(), opts...)
	}
}

// ByAddresses orders the results by addresses terms.
func ByAddresses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAddressesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrdersCount orders the results by orders count.
func ByOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdersStep(), opts...)
	}
}

// ByOrders orders the results by orders terms.
func ByOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByNotificationsCount orders the results by notifications count.
func ByNotificationsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNotificationsStep(), opts...)
	}
}

// ByNotifications orders the results by notifications terms.
func ByNotifications(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNotificationsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPurchaseRequestCount orders the results by purchase_request count.
func ByPurchaseRequestCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPurchaseRequestStep(), opts...)
	}
}

// ByPurchaseRequest orders the results by purchase_request terms.
func ByPurchaseRequest(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPurchaseRequestStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAdminField orders the results by admin field.
func ByAdminField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminStep(), sql.OrderByField(field, opts...))
	}
}
func newBusinessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, BusinessTable, BusinessColumn),
	)
}
func newIndividualStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(IndividualInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, IndividualTable, IndividualColumn),
	)
}
func newAddressesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AddressesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AddressesTable, AddressesColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
	)
}
func newFavouritesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FavouritesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, FavouritesTable, FavouritesColumn),
	)
}
func newNotificationsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NotificationsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, NotificationsTable, NotificationsPrimaryKey...),
	)
}
func newPurchaseRequestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PurchaseRequestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PurchaseRequestTable, PurchaseRequestColumn),
	)
}
func newAdminStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdminInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AdminTable, AdminColumn),
	)
}
