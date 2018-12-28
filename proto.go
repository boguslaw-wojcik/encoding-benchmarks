package benchmarks

import (
	model "github.com/boguslaw-wojcik/encoding-benchmarks/proto/model/standard"
)

//go:generate protoc -I ./proto/schema/ ./proto/schema/superhero.proto --go_out=proto/model/standard
//go:generate protoc -I ./proto/schema/ ./proto/schema/superhero.proto --gofast_out=proto/model/gogo

var ProtoEntity = &model.Superhero{
	Id:            234765,
	AffiliationId: 9867,
	Name:          "Wolverine",
	Life:          85.25,
	Energy:        32.75,
	Powers: []*model.Superpower{
		{
			Id:      2345,
			Name:    "Bone Claws",
			Damage:  5,
			Energy:  1.15,
			Passive: false,
		},
		{
			Id:      2346,
			Name:    "Regeneration",
			Damage:  -2,
			Energy:  0.55,
			Passive: true,
		},
		{
			Id:      2347,
			Name:    "Adamant skeleton",
			Damage:  -10,
			Energy:  0,
			Passive: true,
		},
	},
}