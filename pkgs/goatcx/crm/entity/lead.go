// Code generated by ent, DO NOT EDIT.

package entity

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/pulkitbhardwaj/voyager/pkgs/goatcx/crm/entity/lead"
)

// Lead is the model entity for the Lead schema.
type Lead struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Lead) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case lead.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Lead fields.
func (l *Lead) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case lead.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			l.ID = int(value.Int64)
		default:
			l.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Lead.
// This includes values selected through modifiers, order, etc.
func (l *Lead) Value(name string) (ent.Value, error) {
	return l.selectValues.Get(name)
}

// Update returns a builder for updating this Lead.
// Note that you need to call Lead.Unwrap() before calling this method if this Lead
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Lead) Update() *LeadUpdateOne {
	return NewLeadClient(l.config).UpdateOne(l)
}

// Unwrap unwraps the Lead entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Lead) Unwrap() *Lead {
	_tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("entity: Lead is not a transactional entity")
	}
	l.config.driver = _tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Lead) String() string {
	var builder strings.Builder
	builder.WriteString("Lead(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Leads is a parsable slice of Lead.
type Leads []*Lead
