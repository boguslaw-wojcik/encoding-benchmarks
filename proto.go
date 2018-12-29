package benchmarks

//go:generate protoc -I ./proto/schema/ ./proto/schema/superhero.proto --go_out=proto/model/standard
//go:generate protoc -I ./proto/schema/ ./proto/schema/superhero.proto --gofast_out=proto/model/gogo
