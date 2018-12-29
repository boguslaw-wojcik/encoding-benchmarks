package benchmarks_test

import (
	"io/ioutil"
	"log"
	"testing"

	gogoModel "github.com/boguslaw-wojcik/encoding-benchmarks/proto/model/gogo"
	standardModel "github.com/boguslaw-wojcik/encoding-benchmarks/proto/model/standard"
	"github.com/golang/protobuf/proto"
)

// protoPayload is a variable holding encoded Protobuf reference payload used in all benchmarks.
var protoPayload []byte

// protoResult is a dummy output variable for each benchmark. In benchmarks all results must be copied over to an exported variable to prevent Go compiler from skipping parts of code which results are never used.
var protoResult interface{}

// init reads Protobuf reference payload.
func init() {
	var err error

	protoPayload, err = ioutil.ReadFile("./proto/payload/superhero.pb")
	if err != nil {
		log.Fatal(err)
	}
}

// BenchmarkProtoDecodeStandard performs benchmark of Protobuf decoding by golang/protobuf Google official library.
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

// BenchmarkProtoDecodeGogo performs benchmark of Protobuf decoding by gogo/protobuf library.
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

// BenchmarkProtoEncodeStandard performs benchmark of Protobuf encoding by golang/protobuf Google official library.
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

// BenchmarkProtoEncodeGogo performs benchmark of Protobuf encoding by gogo/protobuf library.
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
