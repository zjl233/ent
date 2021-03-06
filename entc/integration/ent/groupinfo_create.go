// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupInfoCreate is the builder for creating a GroupInfo entity.
type GroupInfoCreate struct {
	config
	desc      *string
	max_users *int
	groups    map[string]struct{}
}

// SetDesc sets the desc field.
func (gic *GroupInfoCreate) SetDesc(s string) *GroupInfoCreate {
	gic.desc = &s
	return gic
}

// SetMaxUsers sets the max_users field.
func (gic *GroupInfoCreate) SetMaxUsers(i int) *GroupInfoCreate {
	gic.max_users = &i
	return gic
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (gic *GroupInfoCreate) SetNillableMaxUsers(i *int) *GroupInfoCreate {
	if i != nil {
		gic.SetMaxUsers(*i)
	}
	return gic
}

// AddGroupIDs adds the groups edge to Group by ids.
func (gic *GroupInfoCreate) AddGroupIDs(ids ...string) *GroupInfoCreate {
	if gic.groups == nil {
		gic.groups = make(map[string]struct{})
	}
	for i := range ids {
		gic.groups[ids[i]] = struct{}{}
	}
	return gic
}

// AddGroups adds the groups edges to Group.
func (gic *GroupInfoCreate) AddGroups(g ...*Group) *GroupInfoCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gic.AddGroupIDs(ids...)
}

// Save creates the GroupInfo in the database.
func (gic *GroupInfoCreate) Save(ctx context.Context) (*GroupInfo, error) {
	if gic.desc == nil {
		return nil, errors.New("ent: missing required field \"desc\"")
	}
	if gic.max_users == nil {
		v := groupinfo.DefaultMaxUsers
		gic.max_users = &v
	}
	return gic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GroupInfoCreate) SaveX(ctx context.Context) *GroupInfo {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gic *GroupInfoCreate) sqlSave(ctx context.Context) (*GroupInfo, error) {
	var (
		gi    = &GroupInfo{config: gic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupinfo.FieldID,
			},
		}
	)
	if value := gic.desc; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: groupinfo.FieldDesc,
		})
		gi.Desc = *value
	}
	if value := gic.max_users; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: groupinfo.FieldMaxUsers,
		})
		gi.MaxUsers = *value
	}
	if nodes := gic.groups; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, gic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	gi.ID = strconv.FormatInt(id, 10)
	return gi, nil
}
