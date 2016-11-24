# json2hive
generate hive schema from a json document

# usage
```
$ go get github.com/xtaci/json2hive
$ json2hive < test.json
CREATE EXTERNAL TABLE test (
  points ARRAY<INT>,
  pointsfloat ARRAY<FLOAT>,
  description STRING,
  foo STRUCT<
	quux:STRING,
	level1:STRUCT<
	  l2string:STRING,
	  l2struct:STRUCT<
	    level3:STRING>>,
	bar:STRING>,
  wibble STRING,
  wobble ARRAY<STRUCT<
	  EntryDetails:STRUCT<
	    details1:STRING,
	    details2:INT>,
	  entry:INT>>,
  items ARRAY<STRUCT<
	  id:INT,
	  count:INT,
	  property:INT,
	  name:STRING>>,
  items2 ARRAY<MAP<STRING, INT>>
)
```
