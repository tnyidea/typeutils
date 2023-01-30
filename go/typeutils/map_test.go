package typeutils

import (
	"log"
	"testing"
)

func TestMapNoEmptyValuesEmptyMap(t *testing.T) {
	testMap := make(map[string]string)

	err := MapNoEmptyValues(testMap)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

func TestMapNoEmptyValuesCompleteMap(t *testing.T) {
	testMap := make(map[string]string)
	testMap["One"] = "One"
	testMap["Two"] = "Two"

	err := MapNoEmptyValues(testMap)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

func TestMapNoEmptyValuesIncompleteMap(t *testing.T) {
	testMap := make(map[string]string)
	testMap["One"] = "One"
	testMap["Two"] = ""

	err := MapNoEmptyValues(testMap)
	if err == nil {
		log.Println("test: TestMapNoEmptyValuesIncompleteMap(expected error for incomplete map)")
		t.FailNow()
	}

	log.Println("test: TestMapNoEmptyValuesIncompleteMap(received expected error)")
	log.Println(err)
}
