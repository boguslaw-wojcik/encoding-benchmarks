package benchmarks_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	"github.com/boguslaw-wojcik/encoding-benchmarks/json/model/parser"
	"github.com/boguslaw-wojcik/encoding-benchmarks/json/model/standard"
	"github.com/json-iterator/go"
)

var jsonPayload []byte

var jsonResult interface{}

func init() {
	var err error

	jsonPayload, err = ioutil.ReadFile("./json/payload/superhero.json")
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkJSONDecodeStandard(b *testing.B) {
	entity := &model.Superhero{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := json.Unmarshal(jsonPayload, entity)
		if err != nil {
			b.Fatal(err)
		}

		jsonResult = entity
	}
}

func BenchmarkJSONDecodeIterator(b *testing.B) {
	e := &model.Superhero{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := jsoniter.ConfigFastest.Unmarshal(jsonPayload, e)
		if err != nil {
			b.Fatal(err)
		}

		jsonResult = e
	}
}

func BenchmarkJSONDecodeParser(b *testing.B) {
	e := &parser.Superhero{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON(jsonPayload)
		if err != nil {
			b.Fatal(err)
		}

		jsonResult = e
	}
}

func BenchmarkJSONEncodeStandard(b *testing.B) {
	entity := &model.Superhero{}
	err := json.Unmarshal(jsonPayload, entity)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p, err := json.Marshal(entity)
		if err != nil {
			b.Fatal(err)
		}

		jsonResult = p
	}
}

func BenchmarkJSONEncodeIterator(b *testing.B) {
	e := &model.Superhero{}
	err := jsoniter.ConfigFastest.Unmarshal(jsonPayload, e)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p, err := jsoniter.ConfigFastest.Marshal(e)
		if err != nil {
			b.Fatal(err)
		}
		jsonResult = p
	}
}
