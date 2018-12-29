package benchmarks_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	gogenModel "github.com/boguslaw-wojcik/encoding-benchmarks/avro/model/gogen"
	linkedinModel "github.com/boguslaw-wojcik/encoding-benchmarks/avro/model/linkedin"
	"github.com/linkedin/goavro"
	"github.com/stretchr/testify/assert"
)

// avroPayload is a variable holding encoded AVRO reference payload used in all benchmarks.
var avroPayload []byte

// avroResult is a dummy output variable for each benchmark. In benchmarks all results must be copied over to an exported variable to prevent Go compiler from skipping parts of code which results are never used.
var avroResult interface{}

// avroCodec is a runtime encoding/decoding codec required to benchmark linkedin/goavro library.
var avroCodec *goavro.Codec

// init reads AVRO reference payload and prepares AVRO codec.
func init() {
	var err error

	avroPayload, err = ioutil.ReadFile("./avro/payload/superhero.avro")
	if err != nil {
		log.Fatal(err)
	}

	schema, err := ioutil.ReadFile("./avro/schema/superhero.avsc")
	if err != nil {
		log.Fatal(err)
	}

	avroCodec, err = goavro.NewCodec(string(schema))
	if err != nil {
		log.Fatal(err)
	}
}

// TestAVROLinkedinDecodingEncoding tests decoding and encoding back and forth by linkedin/goavro utilizing model hydration and extraction.
// Test is required to make sure that custom code for data hydration and extraction is working properly.
func TestAVROLinkedinDecodingEncoding(t *testing.T) {
	entityMap, _, err := avroCodec.NativeFromBinary(avroPayload)
	if err != nil {
		t.Fatal(err)
	}

	entity := &linkedinModel.Superhero{}
	entity.FromMap(entityMap.(map[string]interface{}))

	p, err := avroCodec.BinaryFromNative(nil, entity.ToMap())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, p, avroPayload)
}

// BenchmarkAVRODecodeLinkedinMap performs benchmark of AVRO decoding by linkedin/goavro library.
func BenchmarkAVRODecodeLinkedinMap(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entityMap, _, err := avroCodec.NativeFromBinary(avroPayload)
		if err != nil {
			b.Fatal(err)
		}

		avroResult = entityMap
	}
}

// BenchmarkAVRODecodeLinkedinStruct performs benchmark of AVRO decoding by linkedin/goavro library including additional hydration of data into a designated struct model.
// By default payloads are decoded into Go native interface{} type, which in practice means that we will be dealing mostly with map[string]interface{} type.
// In this benchmark the map resulting from decoding hydrates a struct model in order to mimic the cost of achieving results similar to other libraries.
func BenchmarkAVRODecodeLinkedinStruct(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		entityMap, _, err := avroCodec.NativeFromBinary(avroPayload)
		if err != nil {
			b.Fatal(err)
		}

		entity := &linkedinModel.Superhero{}
		entity.FromMap(entityMap.(map[string]interface{}))

		avroResult = entity
	}
}

// BenchmarkAVRODecodeGogen performs benchmark of AVRO decoding by actgardner/gogen-avro library.
func BenchmarkAVRODecodeGogen(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := bytes.NewReader(avroPayload)

		entity, err := gogenModel.DeserializeSuperhero(r)
		if err != nil {
			b.Fatal(err)
		}

		avroResult = entity
	}
}

// BenchmarkAVROEncodeLinkedinMap performs benchmark of AVRO encoding by linkedin/goavro library.
func BenchmarkAVROEncodeLinkedinMap(b *testing.B) {
	entityMap, _, err := avroCodec.NativeFromBinary(avroPayload)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p, err := avroCodec.BinaryFromNative(nil, entityMap)
		if err != nil {
			b.Fatal(err)
		}

		avroResult = p
	}
}

// BenchmarkAVROEncodeLinkedinStruct performs benchmark of AVRO encoding by linkedin/goavro library including additional extraction of data from a designated struct model.
// By default payloads are encoded from a Go native interface{} type, which in practice means that we will be dealing mostly with map[string]interface{} type.
// In this benchmark the map is extracted from a struct model in order to mimic the cost of achieving results similar to other libraries.
func BenchmarkAVROEncodeLinkedinStruct(b *testing.B) {
	entityMap, _, err := avroCodec.NativeFromBinary(avroPayload)
	if err != nil {
		b.Fatal(err)
	}

	entity := &linkedinModel.Superhero{}
	entity.FromMap(entityMap.(map[string]interface{}))

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p, err := avroCodec.BinaryFromNative(nil, entity.ToMap())
		if err != nil {
			b.Fatal(err)
		}

		avroResult = p
	}
}

// BenchmarkAVROEncodeGogen performs benchmark of AVRO encoding by actgardner/gogen-avro library.
func BenchmarkAVROEncodeGogen(b *testing.B) {
	r := bytes.NewReader(avroPayload)
	entity, err := gogenModel.DeserializeSuperhero(r)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := bytes.NewBuffer([]byte{})

		err := entity.Serialize(w)
		if err != nil {
			b.Fatal(err)
		}

		avroResult = w.Bytes()
	}
}
