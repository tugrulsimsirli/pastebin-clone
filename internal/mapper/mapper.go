package mapper

import (
	"fmt"
	"reflect"
)

func Map(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// if src and dst are pointer, dereference it
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
	}

	// if src and dst are slice, call slice mapping function
	if srcVal.Kind() == reflect.Slice && dstVal.Kind() == reflect.Slice {
		return mapSlice(srcVal, dstVal)
	}

	// src and dst must be struct, return error if they are not
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("src and dst must be structs")
	}

	// mappingo of struct fields
	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i)
		dstField := dstVal.FieldByName(field.Name)

		// if destination field is not valid and mutable, skip
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		srcField := srcVal.Field(i)

		// data type compatibility check
		if srcField.Type() == dstField.Type() {
			dstField.Set(srcField)
		} else {
			continue
		}
	}

	return nil
}

func mapSlice(srcVal reflect.Value, dstVal reflect.Value) error {
	// if source slice is empty, destination slice will be empty
	if srcVal.IsNil() {
		dstVal.Set(reflect.MakeSlice(dstVal.Type(), 0, 0))
		return nil
	}

	// creating a new slice
	newSlice := reflect.MakeSlice(dstVal.Type(), srcVal.Len(), srcVal.Cap())

	// mapping of each slice element
	for i := 0; i < srcVal.Len(); i++ {
		srcElem := srcVal.Index(i)
		dstElem := newSlice.Index(i)

		// if element is not a pointer or mutable, process it different
		if srcElem.Kind() == reflect.Ptr && !srcElem.IsNil() {
			newDstElem := reflect.New(dstElem.Type().Elem())
			if err := Map(srcElem.Interface(), newDstElem.Interface()); err != nil {
				return err
			}
			dstElem.Set(newDstElem.Elem())
		} else {
			newElem := reflect.New(dstElem.Type()).Elem()
			if err := Map(srcElem.Interface(), newElem.Addr().Interface()); err != nil {
				return err
			}
			dstElem.Set(newElem)
		}
	}

	// set response to destination
	dstVal.Set(newSlice)
	return nil
}
