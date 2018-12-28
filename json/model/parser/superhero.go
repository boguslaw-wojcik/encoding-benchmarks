package parser

import (
	"github.com/buger/jsonparser"
)

type Superhero struct {
	Id            int
	AffiliationId int
	Name          string
	Life          float64
	Energy        float64
	Powers        []*Superpower
}

type Superpower struct {
	Id      int
	Name    string
	Damage  float64
	Energy  float64
	Passive bool
}

var superheroPaths = [][]string{
	{"id"},
	{"affiliation_id"},
	{"name"},
	{"life"},
	{"energy"},
	{"powers"},
}

var superpowerPaths = [][]string{
	{"id"},
	{"name"},
	{"damage"},
	{"energy"},
	{"passive"},
}

func (s *Superhero) UnmarshalJSON(b []byte) error {
	var lastErr error

	jsonparser.EachKey(b, func(i int, v []byte, vt jsonparser.ValueType, err error) {
		if err != nil {
			lastErr = err
			return
		}

		if vt == jsonparser.Null || vt == jsonparser.NotExist {
			return
		}

		switch i {
		case 0: // id
			if id, err := jsonparser.ParseInt(v); err == nil {
				s.Id = int(id)
			}
		case 1: // affiliation_id
			if affiliationID, err := jsonparser.ParseInt(v); err == nil {
				s.AffiliationId = int(affiliationID)
			}
		case 2: // name
			s.Name, err = jsonparser.ParseString(v)
		case 3: // life
			s.Life, err = jsonparser.ParseFloat(v)
		case 4: // energy
			s.Energy, err = jsonparser.ParseFloat(v)
		case 5: // powers
			jsonparser.ArrayEach(v, func(v []byte, vt jsonparser.ValueType, offset int, err error) {
				p := &Superpower{}
				p.UnmarshalJSON(v)
				s.Powers = append(s.Powers, p)
			})
		}
	}, superheroPaths...)

	return lastErr
}

func (s *Superpower) UnmarshalJSON(b []byte) error {
	var lastErr error

	jsonparser.EachKey(b, func(i int, v []byte, vt jsonparser.ValueType, err error) {
		if err != nil {
			lastErr = err
			return
		}

		if vt == jsonparser.Null || vt == jsonparser.NotExist {
			return
		}

		switch i {
		case 0: // id
			if id, err := jsonparser.ParseInt(v); err == nil {
				s.Id = int(id)
			}
		case 1: // name
			s.Name, err = jsonparser.ParseString(v)
		case 2: // damage
			s.Damage, err = jsonparser.ParseFloat(v)
		case 3: // energy
			s.Energy, err = jsonparser.ParseFloat(v)
		case 4: // passive
			s.Passive, err = jsonparser.ParseBoolean(v)
		}
	}, superpowerPaths...)

	return lastErr
}
