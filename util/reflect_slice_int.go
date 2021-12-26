package util

import (
	"errors"
	"reflect"
)

func ReflectSlice(thisSlice interface{}, ofThis interface{}) error {
	ptrSlice := reflect.ValueOf(thisSlice)
	reflected := reflect.ValueOf(ofThis)

	if ptrSlice.Kind() != reflect.Ptr {
		return errors.New("o primeiro argumento de ReflectSlice deve ser um ponteiro")
	}

	if reflected.Kind() != reflect.Slice {
		return errors.New("o segundo argumento de ReflectSlice deve ser um slice")
	}

	if reflected.Len() <= 0 || reflected.IsZero() {
		ptrSlice.Elem().Set(reflect.Zero(
			ptrSlice.Elem().Type(),
		))
		return nil
	}

	reflectedSlice := reflect.MakeSlice(
		reflect.SliceOf(reflected.Index(0).Type()),
		0,
		reflected.Len(),
	)

	for i := 0; i < reflected.Len(); i++ {
		el := reflected.Index(i)

		reflectedSlice = reflect.Append(reflectedSlice, el)
	}

	ptrSlice.Elem().Set(reflectedSlice)

	return nil
}
