package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const SPC = "  "

func main() {
	dec := json.NewDecoder(os.Stdin)
	doc := make(map[string]interface{})
	if err := dec.Decode(&doc); err != nil {
		log.Panic(err)
	}

	var lines []string
	for k, v := range doc {
		if subschema := createSchema(v, SPC+SPC); subschema != "" {
			lines = append(lines, SPC+k+" "+subschema)
		}
	}
	fmt.Println("CREATE EXTERNAL TABLE test (")
	fmt.Println(strings.Join(lines, ",\n"))
	fmt.Println(")")
}

func createSchema(doc interface{}, indent string) (schema string) {
	const epsilon = 0.000001
	switch doc := doc.(type) {
	case nil:
	case bool:
		return "BOOLEAN"
	case string:
		return "STRING"
	case float64:
		if math.Abs((doc-math.Floor(doc))) < epsilon || math.Abs((doc-math.Ceil(doc))) < epsilon {
			value := int64(doc)
			if value >= -2147483648 && value <= 2147483647 {
				return "INT"
			} else {
				return "BIGINT"
			}
		} else {
			return "FLOAT"
		}
	case map[string]interface{}:
		var struct_type string
		if len(doc) > 0 {
			// count types
			types := make(map[string]bool)
			for _, value := range doc {
				if subschema := createSchema(value, indent+SPC); subschema != "" {
					types[subschema] = true
				}
			}

			if len(types) == 1 && len(doc) > 1 { // treat as map
				for typ := range types {
					struct_type = "\n" + indent + "MAP<STRING, " + typ + ">"
					break
				}
			} else { // treat as struct
				struct_type = "\n" + indent + "STRUCT<"
				var fields []string
				for name, value := range doc {
					if subschema := createSchema(value, indent+SPC); subschema != "" {
						field_schema := "\n" + indent + name + ":" + subschema
						fields = append(fields, field_schema)
					}
				}
				struct_type += strings.Join(fields, ",")
				struct_type += ">"
			}
		}
		return struct_type
	case []interface{}:
		if len(doc) > 0 {
			array_type := "\n" + indent + "ARRAY<"
			array_type += createSchema(doc[0], indent+SPC)
			array_type += ">"
			return array_type
		}
	default:
		log.Panic("unknown:", fmt.Sprintf("%#v", doc))
	}
	return
}
