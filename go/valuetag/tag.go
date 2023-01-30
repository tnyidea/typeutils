package valuetag

import (
	"errors"
	"reflect"
	"strings"
)

func Validate(v interface{}) error {
	var errMessage []string
	packageType := reflect.TypeOf(v)
	fieldCount := packageType.NumField()

	// Find the Type field
	valueTypeField := reflect.ValueOf(v).FieldByName("Type")

	// Find and process the value tag (value:"...") for each struct field
	for i := 0; i < fieldCount; i++ {
		field := packageType.Field(i)
		valueTag := field.Tag.Get("value")

		// If the field name is 'Type' then ignore and continue
		// TODO: should this be something else... like ValueType or ValueTagType?
		if field.Name == "Type" {
			continue
		}

		// Process the value tag if not empty
		if valueTag != "" {
			// If the field is not a pointer type, ignore and continue
			fieldKind := reflect.ValueOf(v).FieldByName(field.Name).Kind()
			if fieldKind != reflect.Ptr && fieldKind != reflect.Slice && fieldKind != reflect.Map {
				errMessage = append(errMessage, "field '"+packageType.Name()+"."+field.Name+
					"' with value tag is not an addressable type")
			}

			fieldIsNil := reflect.ValueOf(v).FieldByName(field.Name).IsNil()

			// If the tag only contains 'required' and the field is nil return an error
			// Otherwise the tag is assumed to be of correct format and has a value type constraint
			if valueTag == "required" {
				if fieldIsNil {
					errMessage = append(errMessage, "required value is nil for field: "+
						packageType.Name()+"."+field.Name)
				}
				continue
			} else {
				// If the struct does not have a 'Type' field, then ignore the tag, be silent and continue
				if valueTypeField.IsZero() {
					continue
				}

				// If the get the value of the type field.  If it is nil or empty, be silent and continue
				var valueTypeFieldKey string
				switch valueTypeField.Kind() {
				case reflect.Ptr:
					if valueTypeField.IsNil() {
						continue
					}
					// Todo: this is weak: if it is not a *string type then we get a string that doesn't match a tag, but we ignore anyway if no match
					valueTypeFieldKey = valueTypeField.Elem().String()
				case reflect.String:
					valueTypeFieldKey = valueTypeField.String()
					if valueTypeFieldKey == "" {
						continue
					}
				default:
					errMessage = append(errMessage, "cannot determine value type for value tag: '"+
						packageType.Name()+".Type' must be of type string or pointer to a string")
					continue
				}

				// Process each value tag item (of the format <type>:<constraint>)
				valueTagItems := strings.Split(valueTag, ",")
				for j := range valueTagItems {
					tokens := strings.Split(valueTagItems[j], ":")

					// If tokens is not len 2, then the tag is not formatted properly
					if len(tokens) != 2 {
						errMessage = append(errMessage, "invalid value tag format for field: "+
							packageType.Name()+"."+field.Name+": "+valueTagItems[j])
						continue
					}

					// Otherwise check that the value type key matches the field type key
					// If not, be silent and continue
					valueTagKey := tokens[0]
					if valueTagKey != strings.ToLower(valueTypeFieldKey) {
						continue
					}

					// Check the constraints of the field.  Cannot be:
					// 1) neither required or invalid
					// 2) both required and invalid
					fieldRequired := tokens[1] == "required"
					fieldInvalid := tokens[1] == "invalid"
					if (fieldRequired && fieldInvalid) || (!fieldRequired && !fieldInvalid) {
						errMessage = append(errMessage, "invalid value tag format for field: "+
							packageType.Name()+"."+field.Name+": "+valueTagItems[j])
						continue
					}

					// Process the constraint of the field
					if fieldRequired && fieldIsNil {
						errMessage = append(errMessage, "value is nil for field '"+
							packageType.Name()+"."+field.Name+"' with '"+packageType.Name()+".Type' of '"+
							valueTypeFieldKey+"' and value tag '"+valueTagItems[j]+"'")
						continue
					}
					if fieldInvalid && !fieldIsNil {
						errMessage = append(errMessage, "value is not nil for field '"+
							packageType.Name()+"."+field.Name+"' with '"+packageType.Name()+".Type' of '"+
							valueTypeFieldKey+"' and value tag '"+valueTagItems[j]+"'")
						continue
					}
					// either field is required and not nil or invalid and nil, so fall through to end of loop
				}
			}
		}
	}

	if errMessage != nil {
		return errors.New(strings.Join(errMessage, "\n"))
	}
	return nil
}
