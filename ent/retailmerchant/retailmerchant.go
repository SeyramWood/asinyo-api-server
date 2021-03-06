// Code generated by entc, DO NOT EDIT.

package retailmerchant

import (
	"time"
)

const (
	// Label holds the string label denoting the retailmerchant type in the database.
	Label = "retail_merchant"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
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
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// Table holds the table name of the retailmerchant in the database.
	Table = "retail_merchants"
	// MerchantTable is the table that holds the merchant relation/edge.
	MerchantTable = "retail_merchants"
	// MerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantInverseTable = "merchants"
	// MerchantColumn is the table column denoting the merchant relation/edge.
	MerchantColumn = "merchant_retailer"
)

// Columns holds all SQL columns for retailmerchant fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldGhanaCard,
	FieldLastName,
	FieldOtherName,
	FieldPhone,
	FieldOtherPhone,
	FieldAddress,
	FieldDigitalAddress,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "retail_merchants"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"merchant_retailer",
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
