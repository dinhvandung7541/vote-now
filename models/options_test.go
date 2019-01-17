// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOptions(t *testing.T) {
	t.Parallel()

	query := Options()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOptionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOptionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Options().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOptionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OptionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOptionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OptionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Option exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OptionExists to return true, but got false.")
	}
}

func testOptionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	optionFound, err := FindOption(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if optionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOptionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Options().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOptionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Options().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOptionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	optionOne := &Option{}
	optionTwo := &Option{}
	if err = randomize.Struct(seed, optionOne, optionDBTypes, false, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}
	if err = randomize.Struct(seed, optionTwo, optionDBTypes, false, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = optionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = optionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Options().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOptionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	optionOne := &Option{}
	optionTwo := &Option{}
	if err = randomize.Struct(seed, optionOne, optionDBTypes, false, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}
	if err = randomize.Struct(seed, optionTwo, optionDBTypes, false, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = optionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = optionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func optionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func optionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Option) error {
	*o = Option{}
	return nil
}

func testOptionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Option{}
	o := &Option{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, optionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Option object: %s", err)
	}

	AddOptionHook(boil.BeforeInsertHook, optionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	optionBeforeInsertHooks = []OptionHook{}

	AddOptionHook(boil.AfterInsertHook, optionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	optionAfterInsertHooks = []OptionHook{}

	AddOptionHook(boil.AfterSelectHook, optionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	optionAfterSelectHooks = []OptionHook{}

	AddOptionHook(boil.BeforeUpdateHook, optionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	optionBeforeUpdateHooks = []OptionHook{}

	AddOptionHook(boil.AfterUpdateHook, optionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	optionAfterUpdateHooks = []OptionHook{}

	AddOptionHook(boil.BeforeDeleteHook, optionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	optionBeforeDeleteHooks = []OptionHook{}

	AddOptionHook(boil.AfterDeleteHook, optionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	optionAfterDeleteHooks = []OptionHook{}

	AddOptionHook(boil.BeforeUpsertHook, optionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	optionBeforeUpsertHooks = []OptionHook{}

	AddOptionHook(boil.AfterUpsertHook, optionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	optionAfterUpsertHooks = []OptionHook{}
}

func testOptionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOptionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(optionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOptionToManyVotes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b, c Vote

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, voteDBTypes, false, voteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, voteDBTypes, false, voteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.OptionID, a.ID)
	queries.Assign(&c.OptionID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	vote, err := a.Votes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range vote {
		if queries.Equal(v.OptionID, b.OptionID) {
			bFound = true
		}
		if queries.Equal(v.OptionID, c.OptionID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OptionSlice{&a}
	if err = a.L.LoadVotes(ctx, tx, false, (*[]*Option)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Votes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Votes = nil
	if err = a.L.LoadVotes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Votes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", vote)
	}
}

func testOptionToManyAddOpVotes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b, c, d, e Vote

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Vote{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, voteDBTypes, false, strmangle.SetComplement(votePrimaryKeyColumns, voteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Vote{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddVotes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.OptionID) {
			t.Error("foreign key was wrong value", a.ID, first.OptionID)
		}
		if !queries.Equal(a.ID, second.OptionID) {
			t.Error("foreign key was wrong value", a.ID, second.OptionID)
		}

		if first.R.Option != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Option != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Votes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Votes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Votes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOptionToManySetOpVotes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b, c, d, e Vote

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Vote{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, voteDBTypes, false, strmangle.SetComplement(votePrimaryKeyColumns, voteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetVotes(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Votes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetVotes(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Votes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.OptionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.OptionID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.OptionID) {
		t.Error("foreign key was wrong value", a.ID, d.OptionID)
	}
	if !queries.Equal(a.ID, e.OptionID) {
		t.Error("foreign key was wrong value", a.ID, e.OptionID)
	}

	if b.R.Option != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Option != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Option != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Option != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Votes[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Votes[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testOptionToManyRemoveOpVotes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b, c, d, e Vote

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Vote{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, voteDBTypes, false, strmangle.SetComplement(votePrimaryKeyColumns, voteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddVotes(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Votes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveVotes(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Votes().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.OptionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.OptionID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Option != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Option != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Option != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Option != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Votes) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Votes[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Votes[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testOptionToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Option
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.UserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OptionSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Option)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOptionToOneQuestionUsingQuestion(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Option
	var foreign Question

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, questionDBTypes, false, questionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Question struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.QuestionID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Question().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OptionSlice{&local}
	if err = local.L.LoadQuestion(ctx, tx, false, (*[]*Option)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Question == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Question = nil
	if err = local.L.LoadQuestion(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Question == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOptionToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Options[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.UserID, x.ID) {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testOptionToOneRemoveOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetUser(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveUser(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.User().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.User != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.UserID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Options) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOptionToOneSetOpQuestionUsingQuestion(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b, c Question

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Question{&b, &c} {
		err = a.SetQuestion(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Question != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Options[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.QuestionID, x.ID) {
			t.Error("foreign key was wrong value", a.QuestionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.QuestionID))
		reflect.Indirect(reflect.ValueOf(&a.QuestionID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.QuestionID, x.ID) {
			t.Error("foreign key was wrong value", a.QuestionID, x.ID)
		}
	}
}

func testOptionToOneRemoveOpQuestionUsingQuestion(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Option
	var b Question

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, optionDBTypes, false, strmangle.SetComplement(optionPrimaryKeyColumns, optionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, questionDBTypes, false, strmangle.SetComplement(questionPrimaryKeyColumns, questionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetQuestion(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveQuestion(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.Question().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.Question != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.QuestionID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Options) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOptionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOptionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OptionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOptionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Options().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	optionDBTypes = map[string]string{`Content`: `text`, `CreatedAt`: `timestamp with time zone`, `ID`: `integer`, `QuestionID`: `integer`, `UpdatedAt`: `timestamp with time zone`, `UserID`: `integer`}
	_             = bytes.MinRead
)

func testOptionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(optionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(optionColumns) == len(optionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, optionDBTypes, true, optionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOptionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(optionColumns) == len(optionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Option{}
	if err = randomize.Struct(seed, o, optionDBTypes, true, optionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, optionDBTypes, true, optionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(optionColumns, optionPrimaryKeyColumns) {
		fields = optionColumns
	} else {
		fields = strmangle.SetComplement(
			optionColumns,
			optionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OptionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOptionsUpsert(t *testing.T) {
	t.Parallel()

	if len(optionColumns) == len(optionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Option{}
	if err = randomize.Struct(seed, &o, optionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Option: %s", err)
	}

	count, err := Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, optionDBTypes, false, optionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Option struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Option: %s", err)
	}

	count, err = Options().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}