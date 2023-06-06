// Code generated by ent, DO NOT EDIT.

package individualcustomer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the individualcustomer type in the database.
	Label = "individual_customer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldOtherName holds the string denoting the other_name field in the database.
	FieldOtherName = "other_name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldOtherPhone holds the string denoting the other_phone field in the database.
	FieldOtherPhone = "other_phone"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// Table holds the table name of the individualcustomer in the database.
	Table = "individual_customers"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "individual_customers"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_individual"
)

// Columns holds all SQL columns for individualcustomer fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastName,
	FieldOtherName,
	FieldPhone,
	FieldOtherPhone,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "individual_customers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"customer_individual",
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
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	OtherNameValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
)

// OrderOption defines the ordering options for the IndividualCustomer queries.
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

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByOtherName orders the results by the other_name field.
func ByOtherName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtherName, opts...).ToFunc()
}

// ByPhone orders the results by the phone field.
func ByPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhone, opts...).ToFunc()
}

// ByOtherPhone orders the results by the other_phone field.
func ByOtherPhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtherPhone, opts...).ToFunc()
}

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, CustomerTable, CustomerColumn),
	)
}
