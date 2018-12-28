# encoding-benchmarks
A simple benchmark comparing Protobuf, AVRO and JSON encoding implementations in Go.


## Decoding Benchmarks
| Format | Library | Sample | Performance | Memory | Allocations | Custom Work |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| Protobuf | gogo/protobuf | 2000000 | 579 ns/op | 439 B/op | 7 allocs/op | no |
| Protobuf | golang/protobuf (official) | 2000000 | 827 ns/op | 360 B/op | 10 allocs/op | no |
| AVRO | actgardner/gogen-avro | 1000000 | 1335 ns/op | 472 B/op | 35 allocs/op | no |
| JSON | json-iterator/go | 1000000 | 1341 ns/op | 64 B/op | 4 allocs/op | no |
| AVRO | linkedin/goavro (map) | 1000000 | 2101 ns/op | 1776 B/op | 39 allocs/op | no |
| AVRO | linkedin/goavro (struct) | 500000 | 2637 ns/op | 2016 B/op | 44 allocs/op | yes |
| JSON | buger/jsonparser | 500000 | 3647 ns/op | 669 B/op | 15 allocs/op | yes |
| JSON | standard library | 200000 | 7914 ns/op | 320 B/op | 19 allocs/op | no |

## Encoding Benchmarks
| Format | Library | Sample | Performance | Memory | Allocations | Custom Work |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| Protobuf | gogo/protobuf | 10000000 | 225 ns/op | 160 B/op | 2 allocs/op | no |
| Protobuf | golang/protobuf (official) | 3000000 | 560 ns/op | 160 B/op | 2 allocs/op | no |
| AVRO | linkedin/goavro (map) | 2000000 | 722 ns/op | 320 B/op | 11 allocs/op | no |
| AVRO | actgardner/gogen-avro | 1000000 | 1526 ns/op | 501 B/op | 7 allocs/op | no |
| JSON | json-iterator/go | 1000000 | 1659 ns/op | 360 B/op | 3 allocs/op | no |
| AVRO | linkedin/goavro (struct) | 1000000 | 2294 ns/op | 1856 B/op | 37 allocs/op | yes |
| JSON | standard library | 500000 | 2566 ns/op | 352 B/op | 2 allocs/op | no |

## Payload Size
| Format | Size | 
| :---: | :---: |
| AVRO | 100 B |
| AVRO + wire | 105 B|
| Protobuf | 116 B |
| JSON | 302 B |

## Reference Message

All benchmarks have been performed with following reference message.

```json
{
  "id": 234765,
  "affiliation_id": 9867,
  "name": "Wolverine",
  "life": 85.25,
  "energy": 32.75,
  "powers": [
    {
      "id": 2345,
      "name": "Bone Claws",
      "damage": 5,
      "energy": 1.15,
      "passive": false,
    },
    {
      "id": 2346,
      "name": "Regeneration",
      "damage": -2,
      "energy": 0.55,
      "passive": true
    },
    {
      "id": 2347,
      "name": "Adamant skeleton",
      "damage": -10,
      "energy": 0,
      "passive": true
    }
  ]
}
```
