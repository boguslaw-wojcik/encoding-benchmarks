package model

import (
	"github.com/buger/jsonparser"
)

// superheroPaths is a list of paths to be parsed by buger/jsonparser.
var superheroPaths = [][]string{
	{"id"},
	{"affiliation_id"},
	{"name"},
	{"life"},
	{"energy"},
	{"powers"},
}

// Superhero is a model to which a custom unmarshal method utilizing buger/jsonparser API is attached.
type Superhero struct {
	ID            int
	AffiliationID int
	Name          string
	Life          float64
	Energy        float64
	Powers        []*Superpower
}

// UnmarshalJSON is a method implementing unmarshaler interface and utilizing buger/jsonparser low level API.
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
				s.ID = int(id)
			}
		case 1: // affiliation_id
			if affiliationID, err := jsonparser.ParseInt(v); err == nil {
				s.AffiliationID = int(affiliationID)
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

// superpowerPaths is a list of paths to be parsed by buger/jsonparser.
var superpowerPaths = [][]string{
	{"id"},
	{"name"},
	{"damage"},
	{"energy"},
	{"passive"},
}

// Superpower is a model to which a custom unmarshal method utilizing buger/jsonparser API is attached.
type Superpower struct {
	ID      int
	Name    string
	Damage  float64
	Energy  float64
	Passive bool
}

// UnmarshalJSON is a method implementing unmarshaler interface and utilizing buger/jsonparser low level API.
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
				s.ID = int(id)
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
