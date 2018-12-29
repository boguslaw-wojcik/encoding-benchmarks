package benchmarks_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"

	parserModel "github.com/boguslaw-wojcik/encoding-benchmarks/json/model/parser"
	standardModel "github.com/boguslaw-wojcik/encoding-benchmarks/json/model/standard"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
)

// jsonPayload is a variable holding encoded JSON reference payload used in all benchmarks.
var jsonPayload []byte

// jsonResult is a dummy output variable for each benchmark. In benchmarks all results must be copied over to an exported variable to prevent Go compiler from skipping parts of code which results are never used.
var jsonResult interface{}

// init reads JSON reference payload.
func init() {
	var err error

	jsonPayload, err = ioutil.ReadFile("./json/payload/superhero.json")
	if err != nil {
		log.Fatal(err)
	}
}

// TestJSONParserDecoding tests decoding through low level API provided by buger/jsonparser library,
// Test is required to make sure that custom unmarshal method is working properly.
func TestJSONParserDecoding(t *testing.T) {
	entityParser := &parserModel.Superhero{}
	err := entityParser.UnmarshalJSON(jsonPayload)
	if err != nil {
		t.Fatal(err)
	}

	entity := &standardModel.Superhero{}
	err = jsoniter.ConfigFastest.Unmarshal(jsonPayload, entity)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, entity.Id, entityParser.ID)
	assert.Equal(t, entity.AffiliationId, entityParser.AffiliationID)
	assert.Equal(t, entity.Name, entityParser.Name)
	assert.Equal(t, entity.Life, entityParser.Life)
	assert.Equal(t, entity.Energy, entityParser.Energy)

	for i, power := range entity.Powers {
		parserPower := entityParser.Powers[i]
		assert.Equal(t, power.Id, parserPower.ID)
		assert.Equal(t, power.Name, parserPower.Name)
		assert.Equal(t, power.Energy, parserPower.Energy)
		assert.Equal(t, power.Damage, parserPower.Damage)
		assert.Equal(t, power.Passive, parserPower.Passive)
	}
}

// BenchmarkJSONDecodeStandard performs benchmark of JSON decoding by encoding/json GO standard library.
func BenchmarkJSONDecodeStandard(b *testing.B) {
	entity := &standardModel.Superhero{}

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

// BenchmarkJSONDecodeIterator performs benchmark of JSON decoding by json-iterator/go library.
func BenchmarkJSONDecodeIterator(b *testing.B) {
	e := &standardModel.Superhero{}

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

// BenchmarkJSONDecodeParser performs benchmark of JSON decoding by using low level API provided by buger/jsonparser library.
func BenchmarkJSONDecodeParser(b *testing.B) {
	e := &parserModel.Superhero{}

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

// BenchmarkJSONEncodeStandard performs benchmark of JSON encoding by encoding/json GO standard library.
func BenchmarkJSONEncodeStandard(b *testing.B) {
	entity := &standardModel.Superhero{}
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

// BenchmarkJSONEncodeIterator performs benchmark of JSON encoding by json-iterator/go library.
func BenchmarkJSONEncodeIterator(b *testing.B) {
	e := &standardModel.Superhero{}
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
