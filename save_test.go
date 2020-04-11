package jsonformatting

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestSaveToFile(t *testing.T) {

	type easyStruct struct {
		FieldString  string  `json:"field_string"`
		FieldNumeric float64 `json:"field_numeric"`
		FieldStruct  struct {
			FieldMap     map[uint64]interface{} `json:"field_map"`
			FieldSlice   []string               `json:"field_slice"`
			FieldNumeric uint8                  `json:"field_numeric"`
		} `json:"field_struct"`
		FieldStruct2 struct {
			FieldString string `json:"field_string"`
		} `json:"field_struct_2"`
	}

	easyData := easyStruct{
		FieldString:  "test",
		FieldNumeric: 0.2334,
		FieldStruct: struct {
			FieldMap     map[uint64]interface{} `json:"field_map"`
			FieldSlice   []string               `json:"field_slice"`
			FieldNumeric uint8                  `json:"field_numeric"`
		}{
			FieldMap:     map[uint64]interface{}{1: uint64(23), 2: "test", 3: true, 4: []byte("hello world")},
			FieldSlice:   []string{"1", "2", "3"},
			FieldNumeric: 0,
		},
		FieldStruct2: struct {
			FieldString string `json:"field_string"`
		}{FieldString: "as}{[]},\""},
	}

	b, err := json.Marshal(easyData)
	if err != nil {
		t.Fatal(err)
	}

	// write the whole body at once
	err = ioutil.WriteFile("test.json", ConvertToFormatJSON(b), 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConvertToFormatJSONWithError(t *testing.T) {

	errorJSON := []byte(`
					{
						"test": "a"
					}}
				`)

	if _, err := ConvertToFormatJSONWithError(errorJSON); err == nil {
		t.Fatalf("incorrect json format: %s, but ConvertToFormatJSONWithError return error == nil ", string(errorJSON))
	}

	correctJSON := []byte(`
					{
						"test": "a"
					}
				`)

	if _, err := ConvertToFormatJSONWithError(correctJSON); err != nil {
		t.Fatalf("correct json format: %s, but ConvertToFormatJSONWithError return error == %s ", string(errorJSON), err)
	}
}
