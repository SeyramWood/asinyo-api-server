// Code generated by entc, DO NOT EDIT.

package agent

import (
	"time"
)

const (
	// Label holds the string label denoting the agent type in the database.
	Label = "agent"
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
	// FieldGhanaCard holds the string denoting the ghana_card field in the database.
	FieldGhanaCard = "ghana_card"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldOtherName holds the string denoting the other_name field in the database.
	FieldOtherName = "other_name"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldOtherPhone holds the string denoting the other_phone field in the database.
	FieldOtherPhone = "other_phone"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldDigitalAddress holds the string denoting the digital_address field in the database.
	FieldDigitalAddress = "digital_address"
	// EdgeAddresses holds the string denoting the addresses edge name in mutations.
	EdgeAddresses = "addresses"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// EdgeFavourites holds the string denoting the favourites edge name in mutations.
	EdgeFavourites = "favourites"
	// Table holds the table name of the agent in the database.
	Table = "agents"
	// AddressesTable is the table that holds the addresses relation/edge.
	AddressesTable = "addresses"
	// AddressesInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressesInverseTable = "addresses"
	// AddressesColumn is the table column denoting the addresses relation/edge.
	AddressesColumn = "agent_addresses"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "agent_orders"
	// FavouritesTable is the table that holds the favourites relation/edge.
	FavouritesTable = "favourites"
	// FavouritesInverseTable is the table name for the Favourite entity.
	// It exists in this package in order to avoid circular dependency with the "favourite" package.
	FavouritesInverseTable = "favourites"
	// FavouritesColumn is the table column denoting the favourites relation/edge.
	FavouritesColumn = "agent_favourites"
)

// Columns holds all SQL columns for agent fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPassword,
	FieldGhanaCard,
	FieldLastName,
	FieldOtherName,
	FieldPhone,
	FieldOtherPhone,
	FieldAddress,
	FieldDigitalAddress,
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
	// GhanaCardValidator is a validator for the "ghana_card" field. It is called by the builders before save.
	GhanaCardValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// OtherNameValidator is a validator for the "other_name" field. It is called by the builders before save.
	OtherNameValidator func(string) error
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// AddressValidator is a validator for the "address" field. It is called by the builders before save.
	AddressValidator func(string) error
	// DigitalAddressValidator is a validator for the "digital_address" field. It is called by the builders before save.
	DigitalAddressValidator func(string) error
)
