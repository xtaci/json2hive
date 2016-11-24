# json2hive
generate hive schema from a json document

# usage
```
$ go get github.com/xtaci/json2hive
$ json2hive < test.json
(
points ARRAY<INT>,
pointsfloat ARRAY<FLOAT>,
description STRING,
foo STRUCT<bar:STRING,quux:STRING,level1:STRUCT<l2string:STRING,l2struct:STRUCT<level3:STRING>>>,
wibble STRING,
wobble ARRAY<STRUCT<entry:INT,EntryDetails:STRUCT<details1:STRING,details2:INT>>>,
items ARRAY<STRUCT<id:INT,count:INT,property:INT,name:STRING>>,
items2 ARRAY<MAP<STRING, INT>>
)
```
