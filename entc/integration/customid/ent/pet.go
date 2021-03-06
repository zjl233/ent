// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/pet"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pet) scanValues() []interface{} {
	return []interface{}{
		&sql.NullString{}, // id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pet fields.
func (pe *Pet) assignValues(values ...interface{}) error {
	if m, n := len(values), len(pet.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value.Valid {
		pe.ID = value.String
	}
	values = values[1:]
	return nil
}

// Update returns a builder for updating this Pet.
// Note that, you need to call Pet.Unwrap() before calling this method, if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (pe *Pet) id() int {
	id, _ := strconv.Atoi(pe.ID)
	return id
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
