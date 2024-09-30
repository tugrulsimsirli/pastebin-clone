package mapper

import (
	"fmt"
	"reflect"
)

func Map(src interface{}, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// src ve dst pointer ise dereference et
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
	}

	// Eğer src ve dst slice ise, slice mapleme fonksiyonunu çağır
	if srcVal.Kind() == reflect.Slice && dstVal.Kind() == reflect.Slice {
		return mapSlice(srcVal, dstVal)
	}

	// src ve dst struct olmalı, eğer değilse hata döndür
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("src and dst must be structs")
	}

	// Struct field'larını mapleyelim
	for i := 0; i < srcVal.NumField(); i++ {
		field := srcVal.Type().Field(i)
		dstField := dstVal.FieldByName(field.Name)

		// Eğer destination field geçerli ve set edilebilir değilse, atla
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		srcField := srcVal.Field(i)

		// Veri tipi uyumu kontrolü
		if srcField.Type() == dstField.Type() {
			dstField.Set(srcField)
		} else {
			continue
		}
	}

	return nil
}

func mapSlice(srcVal reflect.Value, dstVal reflect.Value) error {
	// Eğer source slice boş ise, destination slice de boş olacak
	if srcVal.IsNil() {
		dstVal.Set(reflect.MakeSlice(dstVal.Type(), 0, 0))
		return nil
	}

	// Yeni bir slice oluştur
	newSlice := reflect.MakeSlice(dstVal.Type(), srcVal.Len(), srcVal.Cap())

	// Her bir slice elemanını mapleyelim
	for i := 0; i < srcVal.Len(); i++ {
		srcElem := srcVal.Index(i)
		dstElem := newSlice.Index(i)

		// Eleman pointer ise veya adreslenebilir değilse, farklı şekilde işleyelim
		if srcElem.Kind() == reflect.Ptr && !srcElem.IsNil() {
			fmt.Println("Processing pointer type")
			newDstElem := reflect.New(dstElem.Type().Elem())
			if err := Map(srcElem.Interface(), newDstElem.Interface()); err != nil {
				return err
			}
			dstElem.Set(newDstElem.Elem())
		} else {
			fmt.Println("Processing non-pointer type")
			newElem := reflect.New(dstElem.Type()).Elem()
			if err := Map(srcElem.Interface(), newElem.Addr().Interface()); err != nil {
				return err
			}
			dstElem.Set(newElem)
		}
	}

	// Sonucu destination'a set et
	dstVal.Set(newSlice)
	return nil
}
