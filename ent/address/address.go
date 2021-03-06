// Code generated by entc, DO NOT EDIT.

package address

import (
	"time"
)

const (
	// Label holds the string label denoting the address type in the database.
	Label = "address"
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
	// FieldDigitalAddress holds the string denoting the digital_address field in the database.
	FieldDigitalAddress = "digital_address"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldRegion holds the string denoting the region field in the database.
	FieldRegion = "region"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldOtherInformation holds the string denoting the other_information field in the database.
	FieldOtherInformation = "other_information"
	// FieldDefault holds the string denoting the default field in the database.
	FieldDefault = "default"
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// EdgeAgent holds the string denoting the agent edge name in mutations.
	EdgeAgent = "agent"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the address in the database.
	Table = "addresses"
	// MerchantTable is the table that holds the merchant relation/edge.
	MerchantTable = "addresses"
	// MerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantInverseTable = "merchants"
	// MerchantColumn is the table column denoting the merchant relation/edge.
	MerchantColumn = "merchant_addresses"
	// AgentTable is the table that holds the agent relation/edge.
	AgentTable = "addresses"
	// AgentInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentInverseTable = "agents"
	// AgentColumn is the table column denoting the agent relation/edge.
	AgentColumn = "agent_addresses"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "addresses"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_addresses"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "address_orders"
)

// Columns holds all SQL columns for address fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldLastName,
	FieldOtherName,
	FieldPhone,
	FieldOtherPhone,
	FieldDigitalAddress,
	FieldCity,
	FieldRegion,
	FieldAddress,
	FieldOtherInformation,
	FieldDefault,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "addresses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"agent_addresses",
	"customer_addresses",
	"merchant_addresses",
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
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// RegionValidator is a validator for the "Region" field. It is called by the builders before save.
	RegionValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// DefaultDefault holds the default value on creation for the "default" field.
	DefaultDefault bool
)
