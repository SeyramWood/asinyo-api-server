// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/SeyramWood/app/domain/models"
	"github.com/SeyramWood/ent/agent"
)

// Agent is the model entity for the Agent schema.
type Agent struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Password holds the value of the "password" field.
	Password []byte `json:"-"`
	// GhanaCard holds the value of the "ghana_card" field.
	GhanaCard string `json:"ghana_card,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// OtherName holds the value of the "other_name" field.
	OtherName string `json:"other_name,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// OtherPhone holds the value of the "other_phone" field.
	OtherPhone string `json:"other_phone,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// DigitalAddress holds the value of the "digital_address" field.
	DigitalAddress string `json:"digital_address,omitempty"`
	// Region holds the value of the "region" field.
	Region string `json:"region,omitempty"`
	// District holds the value of the "district" field.
	District string `json:"district,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// DefaultAccount holds the value of the "default_account" field.
	DefaultAccount agent.DefaultAccount `json:"default_account,omitempty"`
	// BankAccount holds the value of the "bank_account" field.
	BankAccount *models.AgentBankAccount `json:"bank_account,omitempty"`
	// MomoAccount holds the value of the "momo_account" field.
	MomoAccount *models.AgentMomoAccount `json:"momo_account,omitempty"`
	// Verified holds the value of the "verified" field.
	Verified bool `json:"verified,omitempty"`
	// Compliance holds the value of the "compliance" field.
	Compliance *models.AgentComplianceModel `json:"compliance,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AgentQuery when eager-loading is set.
	Edges        AgentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AgentEdges holds the relations/edges for other nodes in the graph.
type AgentEdges struct {
	// Addresses holds the value of the addresses edge.
	Addresses []*Address `json:"addresses,omitempty"`
	// Orders holds the value of the orders edge.
	Orders []*Order `json:"orders,omitempty"`
	// Favourites holds the value of the favourites edge.
	Favourites []*Favourite `json:"favourites,omitempty"`
	// Store holds the value of the store edge.
	Store []*MerchantStore `json:"store,omitempty"`
	// Requests holds the value of the requests edge.
	Requests []*AgentRequest `json:"requests,omitempty"`
	// Notifications holds the value of the notifications edge.
	Notifications []*Notification `json:"notifications,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AddressesOrErr returns the Addresses value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) AddressesOrErr() ([]*Address, error) {
	if e.loadedTypes[0] {
		return e.Addresses, nil
	}
	return nil, &NotLoadedError{edge: "addresses"}
}

// OrdersOrErr returns the Orders value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) OrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[1] {
		return e.Orders, nil
	}
	return nil, &NotLoadedError{edge: "orders"}
}

// FavouritesOrErr returns the Favourites value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) FavouritesOrErr() ([]*Favourite, error) {
	if e.loadedTypes[2] {
		return e.Favourites, nil
	}
	return nil, &NotLoadedError{edge: "favourites"}
}

// StoreOrErr returns the Store value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) StoreOrErr() ([]*MerchantStore, error) {
	if e.loadedTypes[3] {
		return e.Store, nil
	}
	return nil, &NotLoadedError{edge: "store"}
}

// RequestsOrErr returns the Requests value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) RequestsOrErr() ([]*AgentRequest, error) {
	if e.loadedTypes[4] {
		return e.Requests, nil
	}
	return nil, &NotLoadedError{edge: "requests"}
}

// NotificationsOrErr returns the Notifications value or an error if the edge
// was not loaded in eager-loading.
func (e AgentEdges) NotificationsOrErr() ([]*Notification, error) {
	if e.loadedTypes[5] {
		return e.Notifications, nil
	}
	return nil, &NotLoadedError{edge: "notifications"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Agent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case agent.FieldPassword, agent.FieldBankAccount, agent.FieldMomoAccount, agent.FieldCompliance:
			values[i] = new([]byte)
		case agent.FieldVerified:
			values[i] = new(sql.NullBool)
		case agent.FieldID:
			values[i] = new(sql.NullInt64)
		case agent.FieldUsername, agent.FieldGhanaCard, agent.FieldLastName, agent.FieldOtherName, agent.FieldPhone, agent.FieldOtherPhone, agent.FieldAddress, agent.FieldDigitalAddress, agent.FieldRegion, agent.FieldDistrict, agent.FieldCity, agent.FieldDefaultAccount:
			values[i] = new(sql.NullString)
		case agent.FieldCreatedAt, agent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Agent fields.
func (a *Agent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case agent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case agent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		case agent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				a.UpdatedAt = value.Time
			}
		case agent.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				a.Username = value.String
			}
		case agent.FieldPassword:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value != nil {
				a.Password = *value
			}
		case agent.FieldGhanaCard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ghana_card", values[i])
			} else if value.Valid {
				a.GhanaCard = value.String
			}
		case agent.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				a.LastName = value.String
			}
		case agent.FieldOtherName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_name", values[i])
			} else if value.Valid {
				a.OtherName = value.String
			}
		case agent.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				a.Phone = value.String
			}
		case agent.FieldOtherPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_phone", values[i])
			} else if value.Valid {
				a.OtherPhone = value.String
			}
		case agent.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				a.Address = value.String
			}
		case agent.FieldDigitalAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field digital_address", values[i])
			} else if value.Valid {
				a.DigitalAddress = value.String
			}
		case agent.FieldRegion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field region", values[i])
			} else if value.Valid {
				a.Region = value.String
			}
		case agent.FieldDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field district", values[i])
			} else if value.Valid {
				a.District = value.String
			}
		case agent.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				a.City = value.String
			}
		case agent.FieldDefaultAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_account", values[i])
			} else if value.Valid {
				a.DefaultAccount = agent.DefaultAccount(value.String)
			}
		case agent.FieldBankAccount:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field bank_account", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.BankAccount); err != nil {
					return fmt.Errorf("unmarshal field bank_account: %w", err)
				}
			}
		case agent.FieldMomoAccount:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field momo_account", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.MomoAccount); err != nil {
					return fmt.Errorf("unmarshal field momo_account: %w", err)
				}
			}
		case agent.FieldVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field verified", values[i])
			} else if value.Valid {
				a.Verified = value.Bool
			}
		case agent.FieldCompliance:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field compliance", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &a.Compliance); err != nil {
					return fmt.Errorf("unmarshal field compliance: %w", err)
				}
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Agent.
// This includes values selected through modifiers, order, etc.
func (a *Agent) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryAddresses queries the "addresses" edge of the Agent entity.
func (a *Agent) QueryAddresses() *AddressQuery {
	return NewAgentClient(a.config).QueryAddresses(a)
}

// QueryOrders queries the "orders" edge of the Agent entity.
func (a *Agent) QueryOrders() *OrderQuery {
	return NewAgentClient(a.config).QueryOrders(a)
}

// QueryFavourites queries the "favourites" edge of the Agent entity.
func (a *Agent) QueryFavourites() *FavouriteQuery {
	return NewAgentClient(a.config).QueryFavourites(a)
}

// QueryStore queries the "store" edge of the Agent entity.
func (a *Agent) QueryStore() *MerchantStoreQuery {
	return NewAgentClient(a.config).QueryStore(a)
}

// QueryRequests queries the "requests" edge of the Agent entity.
func (a *Agent) QueryRequests() *AgentRequestQuery {
	return NewAgentClient(a.config).QueryRequests(a)
}

// QueryNotifications queries the "notifications" edge of the Agent entity.
func (a *Agent) QueryNotifications() *NotificationQuery {
	return NewAgentClient(a.config).QueryNotifications(a)
}

// Update returns a builder for updating this Agent.
// Note that you need to call Agent.Unwrap() before calling this method if this Agent
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Agent) Update() *AgentUpdateOne {
	return NewAgentClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Agent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Agent) Unwrap() *Agent {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Agent is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Agent) String() string {
	var builder strings.Builder
	builder.WriteString("Agent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(a.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(a.Username)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("ghana_card=")
	builder.WriteString(a.GhanaCard)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(a.LastName)
	builder.WriteString(", ")
	builder.WriteString("other_name=")
	builder.WriteString(a.OtherName)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(a.Phone)
	builder.WriteString(", ")
	builder.WriteString("other_phone=")
	builder.WriteString(a.OtherPhone)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(a.Address)
	builder.WriteString(", ")
	builder.WriteString("digital_address=")
	builder.WriteString(a.DigitalAddress)
	builder.WriteString(", ")
	builder.WriteString("region=")
	builder.WriteString(a.Region)
	builder.WriteString(", ")
	builder.WriteString("district=")
	builder.WriteString(a.District)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(a.City)
	builder.WriteString(", ")
	builder.WriteString("default_account=")
	builder.WriteString(fmt.Sprintf("%v", a.DefaultAccount))
	builder.WriteString(", ")
	builder.WriteString("bank_account=")
	builder.WriteString(fmt.Sprintf("%v", a.BankAccount))
	builder.WriteString(", ")
	builder.WriteString("momo_account=")
	builder.WriteString(fmt.Sprintf("%v", a.MomoAccount))
	builder.WriteString(", ")
	builder.WriteString("verified=")
	builder.WriteString(fmt.Sprintf("%v", a.Verified))
	builder.WriteString(", ")
	builder.WriteString("compliance=")
	builder.WriteString(fmt.Sprintf("%v", a.Compliance))
	builder.WriteByte(')')
	return builder.String()
}

// Agents is a parsable slice of Agent.
type Agents []*Agent
