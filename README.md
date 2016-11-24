# json2hive
generate hive schema from a json document

# usage
```
$ go get github.com/xtaci/json2hive
$ json2hive < test.json
items ARRAY<STRUCT<id:INT,count:INT,property:INT,name:STRING>>
description STRING
foo STRUCT<bar:STRING,quux:STRING,level1:STRUCT<l2struct:STRUCT<level3:STRING>,l2string:STRING>>
wibble STRING
wobble ARRAY<STRUCT<entry:INT,EntryDetails:STRUCT<details1:STRING,details2:INT>>>
```
