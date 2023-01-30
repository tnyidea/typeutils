package typeutils

import (
	"errors"
	"reflect"
)

func MapNoEmptyValues(v any) error {
	t := reflect.TypeOf(v)
	if t == nil {
		return errors.New("typeutils: MapNoEmptyValues(nil)")
	}
	if t.Kind() != reflect.Map {
		return errors.New("typeutils: MapNoEmptyValues(non-map " + t.String() + ")")
	}

	mapValue := reflect.ValueOf(v)
	for _, key := range mapValue.MapKeys() {
		keyValue := mapValue.MapIndex(key)
		if keyValue.IsZero() {
			return errors.New("typeutils: MapNoEmptyValues(map contains at least one empty key value)")
		}
	}

	return nil
}
