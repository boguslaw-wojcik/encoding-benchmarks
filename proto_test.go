package benchmarks_test

import (
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	"log"
	"testing"

	gogoModel "github.com/boguslaw-wojcik/encoding-benchmarks/proto/model/gogo"
	standardModel "github.com/boguslaw-wojcik/encoding-benchmarks/proto/model/standard"
)

var protoPayload []byte

var protoResult interface{}

func init() {
	p, err := ioutil.ReadFile("./proto/payload/superhero.pb")
	if err != nil {
		log.Fatal(err)
	}

	protoPayload = p
}


func BenchmarkProtoDecodeStandard(b *testing.B) {
	entity := &standardModel.Superhero{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := proto.Unmarshal(protoPayload, entity)
		if err != nil {
			b.Fatal(err)
		}

		protoResult = entity
	}
}

func BenchmarkProtoDecodeGogo(b *testing.B) {
	entity := &gogoModel.Superhero{}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		err := entity.Unmarshal(protoPayload)
		if err != nil {
			b.Fatal(err)
		}

		protoResult = entity
	}
}

func BenchmarkProtoEncodeStandard(b *testing.B) {
	entity := &standardModel.Superhero{}
	err := proto.Unmarshal(protoPayload, entity)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p, err := proto.Marshal(entity)
		if err != nil {
			b.Fatal(err)
		}

		protoResult = p
	}
}

func BenchmarkProtoEncodeGogo(b *testing.B) {
	entity := &gogoModel.Superhero{}
	err := entity.Unmarshal(protoPayload)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		p, err := entity.Marshal()
		if err != nil {
			b.Fatal(err)
		}

		protoResult = p
	}
}