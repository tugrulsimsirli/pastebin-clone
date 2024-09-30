package mapper

import (
	"fmt"
	"reflect"
)

func Map(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
	}

	if srcVal.Kind() == reflect.Slice && dstVal.Kind() == reflect.Slice {
		return mapSlice(srcVal, dstVal)
	}

	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i)
		dstField := dstVal.FieldByName(field.Name)

		if dstField.IsValid() && dstField.CanSet() {
			srcField := srcVal.Field(i)

			// Veri tipi uyumu kontrolü
			if srcField.Type() == dstField.Type() {
				dstField.Set(srcField)
			} else {
				return fmt.Errorf("field type mismatch for field: %s", field.Name)
			}
		}
	}

	return nil
}

func mapSlice(srcVal reflect.Value, dstVal reflect.Value) error {
	newSlice := reflect.MakeSlice(dstVal.Type(), srcVal.Len(), srcVal.Cap())

	for i := 0; i < srcVal.Len(); i++ {
		srcElem := srcVal.Index(i)
		dstElem := newSlice.Index(i)
		if err := Map(srcElem.Addr().Interface(), dstElem.Addr().Interface()); err != nil {
			return err
		}
	}

	dstVal.Set(newSlice)
	return nil
}
