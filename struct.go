package typeutils

import "reflect"

func StructTags(key string, v interface{}) []string {
	var tags []string
	structType := reflect.TypeOf(v)
	for i := 0; i < structType.NumField(); i++ {
		tags = append(tags, structType.Field(i).Tag.Get(key))
	}
	return tags
}
