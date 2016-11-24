package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	doc := make(map[string]interface{})
	if err := dec.Decode(&doc); err != nil {
		log.Panic(err)
	}

	var fields []string
	for k, v := range doc {
		subschema := createSchema(v)
		if subschema != "" {
			fields = append(fields, k+" "+createSchema(v))
		}
	}
	fmt.Println(strings.Join(fields, ","))
}

func createSchema(doc interface{}) (schema string) {
	switch doc := doc.(type) {
	case bool:
		return "BOOLEAN"
	case string:
		return "STRING"
	case float64:
		return "FLOAT"
	case map[string]interface{}:
		if len(doc) > 0 {
			struct_type := "STRUCT<"
			var fields []string
			for name, value := range doc {
				subschema := createSchema(value)
				if subschema != "" {
					field_schema := name + ":" + createSchema(value)
					fields = append(fields, field_schema)
				}
			}
			struct_type += strings.Join(fields, ",")
			struct_type += ">"
			return struct_type
		}
	case []interface{}:
		if len(doc) > 0 {
			array_type := "ARRAY<"
			array_type += createSchema(doc[0])
			array_type += ">"
			return array_type
		}
	default:
		log.Println("unknown:", fmt.Sprintf("%#v %T", doc, doc))
	}
	return
}
