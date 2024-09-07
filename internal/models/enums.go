package models

import (
	"database/sql/driver"
	"fmt"
	"reflect"
)

type namer interface {
	Name() string
}

type baseIntEnum int

func (enum baseIntEnum) Name() string {
	return reflect.TypeOf(enum).Name()
}

func (enum baseIntEnum) Value() (driver.Value, error) {
	return int64(enum), nil
}

func (enum *baseIntEnum) Scan(src any) error {
	v, ok := src.(int64)
	if !ok {
		panic(fmt.Errorf("cannot scan type %s into int64", enum.Name()))
	}

	*enum = baseIntEnum(v)
	return nil
}
