package benchmarks_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	gogenModel "github.com/boguslaw-wojcik/encoding-benchmarks/avro/model/gogen"
	linkedinModel "github.com/boguslaw-wojcik/encoding-benchmarks/avro/model/linkedin"
	"github.com/linkedin/goavro"
)

var avroPayload []byte

var avroResult interface{}

var avroCodec *goavro.Codec

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

func BenchmarkAVRODecodeLinkedinMap(b *testing.B) {

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		entityMap, _, err := avroCodec.NativeFromBinary(avroPayload)
		if err != nil {
			b.Fatal(err)
		}

		avroResult = entityMap
	}
}

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
