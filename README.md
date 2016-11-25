# json2hive
[![Build Status][1]][2]
[1]: https://travis-ci.org/xtaci/json2hive.svg?branch=master
[2]: https://travis-ci.org/xtaci/json2hive

generate hive schema from a json document

# usage
```sql
$ go get github.com/xtaci/json2hive
$ json2hive < test.json
CREATE EXTERNAL TABLE test (
  wobble
    ARRAY<
      STRUCT<
      entry:INT,
      EntryDetails:
        STRUCT<
        details2:INT,
        details1:STRING>>>,
  items
    ARRAY<
      STRUCT<
      property:INT,
      name:STRING,
      id:INT,
      count:INT>>,
  items2
    ARRAY<
      MAP<STRING, INT>>,
  points
    ARRAY<INT>,
  pointsfloat
    ARRAY<FLOAT>,
  description STRING,
  foo
    STRUCT<
    bar:STRING,
    quux:STRING,
    level1:
      STRUCT<
      l2string:STRING,
      l2struct:
        STRUCT<
        level3:STRING>>>,
  wibble STRING
)
```
