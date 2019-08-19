// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"github.com/facebookincubator/ent/entc/integration/ent/comment"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	unique_int   *int
	unique_float *float64
}

// SetUniqueInt sets the unique_int field.
func (cc *CommentCreate) SetUniqueInt(i int) *CommentCreate {
	cc.unique_int = &i
	return cc
}

// SetUniqueFloat sets the unique_float field.
func (cc *CommentCreate) SetUniqueFloat(f float64) *CommentCreate {
	cc.unique_float = &f
	return cc
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	if cc.unique_int == nil {
		return nil, errors.New("ent: missing required field \"unique_int\"")
	}
	if cc.unique_float == nil {
		return nil, errors.New("ent: missing required field \"unique_float\"")
	}
	switch cc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return cc.sqlSave(ctx)
	case dialect.Neptune:
		return cc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	var (
		res sql.Result
		c   = &Comment{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(comment.Table).Default(cc.driver.Dialect())
	if cc.unique_int != nil {
		builder.Set(comment.FieldUniqueInt, *cc.unique_int)
		c.UniqueInt = *cc.unique_int
	}
	if cc.unique_float != nil {
		builder.Set(comment.FieldUniqueFloat, *cc.unique_float)
		c.UniqueFloat = *cc.unique_float
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CommentCreate) gremlinSave(ctx context.Context) (*Comment, error) {
	res := &gremlin.Response{}
	query, bindings := cc.gremlin().Query()
	if err := cc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	c := &Comment{config: cc.config}
	if err := c.FromResponse(res); err != nil {
		return nil, err
	}
	return c, nil
}

func (cc *CommentCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(comment.Label)
	if cc.unique_int != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueInt, *cc.unique_int).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueInt, *cc.unique_int)),
		})
		v.Property(dsl.Single, comment.FieldUniqueInt, *cc.unique_int)
	}
	if cc.unique_float != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(comment.Label, comment.FieldUniqueFloat, *cc.unique_float).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(comment.Label, comment.FieldUniqueFloat, *cc.unique_float)),
		})
		v.Property(dsl.Single, comment.FieldUniqueFloat, *cc.unique_float)
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}
