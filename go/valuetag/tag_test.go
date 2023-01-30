package valuetag

import (
	typeutils "github.com/tnyidea/typeutils/go"
	"log"
	"testing"
)

type ValueTagTest struct {
	Field1 *string  `json:"field1"`
	Field2 *int     `json:"field2" value:"required"`
	Field3 *bool    `json:"field3" value:"type1:required"`
	Field4 []string `json:"field4" value:"type2:invalid"`
	Type   *int     `json:"type"`
}

var testGoodValueTag = ValueTagTest{
	Field1: typeutils.StringPtr("field1Value"),
	Field2: typeutils.IntPtr(5),
	Field3: typeutils.BoolPtr(true),
	Field4: []string{
		"value1", "value2",
	},
	Type: typeutils.IntPtr(5),
}

func TestValidateValueTags(t *testing.T) {
	err := valuetag.Validate(testGoodValueTag)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}
