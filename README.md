# Encoding Benchmarks
A simple benchmark comparing Protobuf, AVRO and JSON encoding implementations in Go.

## Why yet another benchmark?
- You cannot directly compare performance of cross-platform technologies. You can only compare performance of their available implementations in a specific language. What is true for one language might not be the case for another. Most benchmarks I have initially found focus mostly on C and Java implementations and these not necessarily carry over to Go which is still relatively young.
- Experience shows that very often unofficial implementations provide a game-changing performance boost being sometimes even ten times faster than a standard solution. Most benchmarks I have found compare official or standard implementations. On top of that, new improvements and solutions are being constantly developed. In some cases new version of the same library brought improvements eliminating previous performance gap that was not reflected in the old benchmarks.
- Many of the published benchmarks tend to share two characteristics. First, they often focus on the message size, benchmarking large and small but mostly flat payloads, ignoring message complexity arising from nested structures. Secondly, some benchmarks entirely skip the model-mapping part of the encoding/decoding process, which results in hiding the CPU and memory cost we will have to eventually pay in all our applications. Wherever required I have provided custom code to complete mapping process in order to measure cost of the entire process.

## Decoding Performance
| Format | Library | Sample | Performance | Memory | Allocations | Custom Work |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| Protobuf | gogo/protobuf | 2000000 | 579 ns/op | 439 B/op | 7 allocs/op | no |
| Protobuf | golang/protobuf (official library) | 2000000 | 827 ns/op | 360 B/op | 10 allocs/op | no |
| AVRO | actgardner/gogen-avro | 1000000 | 1335 ns/op | 472 B/op | 35 allocs/op | no |
| JSON | json-iterator/go | 1000000 | 1341 ns/op | 64 B/op | 4 allocs/op | no |
| AVRO | linkedin/goavro (map) | 1000000 | 2101 ns/op | 1776 B/op | 39 allocs/op | no |
| AVRO | linkedin/goavro (struct) | 500000 | 2637 ns/op | 2016 B/op | 44 allocs/op | yes |
| JSON | buger/jsonparser | 500000 | 3647 ns/op | 669 B/op | 15 allocs/op | yes |
| JSON | encoding/json (standard library) | 200000 | 7914 ns/op | 320 B/op | 19 allocs/op | no |

## Encoding Performance
| Format | Library | Sample | Performance | Memory | Allocations | Custom Work |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| Protobuf | gogo/protobuf | 10000000 | 225 ns/op | 160 B/op | 2 allocs/op | no |
| Protobuf | golang/protobuf (official) | 3000000 | 560 ns/op | 160 B/op | 2 allocs/op | no |
| AVRO | linkedin/goavro (map) | 2000000 | 722 ns/op | 320 B/op | 11 allocs/op | no |
| AVRO | actgardner/gogen-avro | 1000000 | 1526 ns/op | 501 B/op | 7 allocs/op | no |
| JSON | json-iterator/go | 1000000 | 1659 ns/op | 360 B/op | 3 allocs/op | no |
| AVRO | linkedin/goavro (struct) | 1000000 | 2294 ns/op | 1856 B/op | 37 allocs/op | yes |
| JSON | encoding/json (standard library) | 500000 | 2566 ns/op | 352 B/op | 2 allocs/op | no |

## Payload Size
| Format | Size | 
| :---: | :---: |
| AVRO | 100 B |
| AVRO + wire | 105 B |
| Protobuf | 116 B |
| JSON | 302 B |
_Note: In order to benefit from AVRO forward compatibility you need to always transmit messages with unique ID of schema held in a schema registry, which adds additional 5 bytes to the beginning of each payload as per its universally adopted [wire format](https://docs.confluent.io/current/schema-registry/docs/serializer-formatter.html#wire-format)._

## Highlights
* Winning _gogo/protobuf_ library is in general over **12 times** faster than standard JSON library but really only a bit over **2 times** faster than _json-iterator/go_ in decoding. While these are still incredible results, they slightly contrast with commonly exaggerated opinions.
* Both AVRO implementations are generally **on par** with best available JSON solution in terms of performance while utilizing more memory for each operation. It is important to note that despite slightly better overall performance, _actgardner/gogen-avro_ cannot be easily used in system relying on forward compatibility of transmitted payloads, which in turn makes LinkedIn solution preferable.
* At any time any strongly typed binary format such as Protobuf or AVRO should be chosen over JSON for internal communication - if not for performance boost then at least because of **3 times** smaller payloads.
* Sometimes considered the fastest JSON decoding library _buger/jsonparser_ lives up to this expectation only if a set of specific conditions are met. In typical use cases, while still being huge improvement over standard library, it falls short of _json-iterator/go_. That being said it is still a great library which is a life-saver when you have to parse just one or two fields out of the entire JSON payload or you are dealing with an unsafe and unstable third party API.

## Tested libraries
* AVRO
  * [linkedin/goavro](https://github.com/linkedin/goavro)
  * [actgardner/gogen-avro](https://github.com/actgardner/gogen-avro)
* JSON
  * [encoding/json](https://golang.org/pkg/encoding/json/) (standard library)
  * [json-iterator/go](https://github.com/json-iterator/go)
  * [buger/jsonparser](https://github.com/buger/jsonparser)
* Protobuf
  * [golang/protobuf](https://github.com/golang/protobuf) (official library)
  * [gogo/protobuf](https://github.com/gogo/protobuf)

## How to run the benchmark locally?
```bash
$ git clone git@github.com:boguslaw-wojcik/encoding-benchmarks.git
$ cd encoding-benchmarks
$ go get
$ go test -bench=.
```

## Reference Message
All benchmarks have been performed with following reference message containing mix of textual, numeric and logical attributes combined with nested object in a collection.
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

## Contributions welcomed
Feel free to report issue or create a pull request.