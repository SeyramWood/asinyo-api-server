// Code generated by entc, DO NOT EDIT.

package order

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldOrderNumber holds the string denoting the order_number field in the database.
	FieldOrderNumber = "order_number"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldDeliveryFee holds the string denoting the delivery_fee field in the database.
	FieldDeliveryFee = "delivery_fee"
	// FieldReference holds the string denoting the reference field in the database.
	FieldReference = "reference"
	// FieldChannel holds the string denoting the channel field in the database.
	FieldChannel = "channel"
	// FieldPaidAt holds the string denoting the paid_at field in the database.
	FieldPaidAt = "paid_at"
	// FieldDeliveryMethod holds the string denoting the delivery_method field in the database.
	FieldDeliveryMethod = "delivery_method"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDeliveredAt holds the string denoting the delivered_at field in the database.
	FieldDeliveredAt = "delivered_at"
	// EdgeDetails holds the string denoting the details edge name in mutations.
	EdgeDetails = "details"
	// EdgeMerchant holds the string denoting the merchant edge name in mutations.
	EdgeMerchant = "merchant"
	// EdgeAgent holds the string denoting the agent edge name in mutations.
	EdgeAgent = "agent"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeAddress holds the string denoting the address edge name in mutations.
	EdgeAddress = "address"
	// EdgePickup holds the string denoting the pickup edge name in mutations.
	EdgePickup = "pickup"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// DetailsTable is the table that holds the details relation/edge.
	DetailsTable = "order_details"
	// DetailsInverseTable is the table name for the OrderDetail entity.
	// It exists in this package in order to avoid circular dependency with the "orderdetail" package.
	DetailsInverseTable = "order_details"
	// DetailsColumn is the table column denoting the details relation/edge.
	DetailsColumn = "order_details"
	// MerchantTable is the table that holds the merchant relation/edge.
	MerchantTable = "orders"
	// MerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	MerchantInverseTable = "merchants"
	// MerchantColumn is the table column denoting the merchant relation/edge.
	MerchantColumn = "merchant_orders"
	// AgentTable is the table that holds the agent relation/edge.
	AgentTable = "orders"
	// AgentInverseTable is the table name for the Agent entity.
	// It exists in this package in order to avoid circular dependency with the "agent" package.
	AgentInverseTable = "agents"
	// AgentColumn is the table column denoting the agent relation/edge.
	AgentColumn = "agent_orders"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "orders"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_orders"
	// AddressTable is the table that holds the address relation/edge.
	AddressTable = "orders"
	// AddressInverseTable is the table name for the Address entity.
	// It exists in this package in order to avoid circular dependency with the "address" package.
	AddressInverseTable = "addresses"
	// AddressColumn is the table column denoting the address relation/edge.
	AddressColumn = "address_orders"
	// PickupTable is the table that holds the pickup relation/edge.
	PickupTable = "orders"
	// PickupInverseTable is the table name for the PickupStation entity.
	// It exists in this package in order to avoid circular dependency with the "pickupstation" package.
	PickupInverseTable = "pickup_stations"
	// PickupColumn is the table column denoting the pickup relation/edge.
	PickupColumn = "pickup_station_orders"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldOrderNumber,
	FieldCurrency,
	FieldAmount,
	FieldDeliveryFee,
	FieldReference,
	FieldChannel,
	FieldPaidAt,
	FieldDeliveryMethod,
	FieldStatus,
	FieldDeliveredAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"address_orders",
	"agent_orders",
	"customer_orders",
	"merchant_orders",
	"pickup_station_orders",
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
	// OrderNumberValidator is a validator for the "order_number" field. It is called by the builders before save.
	OrderNumberValidator func(string) error
	// CurrencyValidator is a validator for the "currency" field. It is called by the builders before save.
	CurrencyValidator func(string) error
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount float64
	// DefaultDeliveryFee holds the default value on creation for the "delivery_fee" field.
	DefaultDeliveryFee float64
	// ReferenceValidator is a validator for the "reference" field. It is called by the builders before save.
	ReferenceValidator func(string) error
	// ChannelValidator is a validator for the "channel" field. It is called by the builders before save.
	ChannelValidator func(string) error
	// PaidAtValidator is a validator for the "paid_at" field. It is called by the builders before save.
	PaidAtValidator func(string) error
)

// DeliveryMethod defines the type for the "delivery_method" enum field.
type DeliveryMethod string

// DeliveryMethod values.
const (
	DeliveryMethodHOD DeliveryMethod = "HOD"
	DeliveryMethodPSD DeliveryMethod = "PSD"
)

func (dm DeliveryMethod) String() string {
	return string(dm)
}

// DeliveryMethodValidator is a validator for the "delivery_method" field enum values. It is called by the builders before save.
func DeliveryMethodValidator(dm DeliveryMethod) error {
	switch dm {
	case DeliveryMethodHOD, DeliveryMethodPSD:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for delivery_method field: %q", dm)
	}
}

// Status defines the type for the "status" enum field.
type Status string

// StatusPending is the default value of the Status enum.
const DefaultStatus = StatusPending

// Status values.
const (
	StatusPending   Status = "pending"
	StatusShipping  Status = "shipping"
	StatusDelivered Status = "delivered"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPending, StatusShipping, StatusDelivered:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for status field: %q", s)
	}
}
