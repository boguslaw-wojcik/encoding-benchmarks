// Code generated by gopkg.in/actgardner/gogen-avro.v5. DO NOT EDIT.
/*
 * SOURCE:
 *     superhero.avsc
 */

package model

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type ByteReader interface {
	ReadByte() (byte, error)
}

type ByteWriter interface {
	Grow(int)
	WriteByte(byte) error
}

type StringWriter interface {
	WriteString(string) (int, error)
}

func encodeFloat(w io.Writer, byteCount int, bits uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}
	for i := 0; i < byteCount; i++ {
		if bw != nil {
			err = bw.WriteByte(byte(bits & 255))
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(bits&255))
		}
		bits = bits >> 8
	}
	if bw == nil {
		_, err = w.Write(bb)
		return err
	}
	return nil
}

func encodeInt(w io.Writer, byteCount int, encoded uint64) error {
	var err error
	var bb []byte
	bw, ok := w.(ByteWriter)
	// To avoid reallocations, grow capacity to the largest possible size
	// for this integer
	if ok {
		bw.Grow(byteCount)
	} else {
		bb = make([]byte, 0, byteCount)
	}

	if encoded == 0 {
		if bw != nil {
			err = bw.WriteByte(0)
			if err != nil {
				return err
			}
		} else {
			bb = append(bb, byte(0))
		}
	} else {
		for encoded > 0 {
			b := byte(encoded & 127)
			encoded = encoded >> 7
			if !(encoded == 0) {
				b |= 128
			}
			if bw != nil {
				err = bw.WriteByte(b)
				if err != nil {
					return err
				}
			} else {
				bb = append(bb, b)
			}
		}
	}
	if bw == nil {
		_, err := w.Write(bb)
		return err
	}
	return nil

}

func readArraySuperpower(r io.Reader) ([]*Superpower, error) {
	var err error
	var blkSize int64
	var arr = make([]*Superpower, 0)
	for {
		blkSize, err = readLong(r)
		if err != nil {
			return nil, err
		}
		if blkSize == 0 {
			break
		}
		if blkSize < 0 {
			blkSize = -blkSize
			_, err = readLong(r)
			if err != nil {
				return nil, err
			}
		}
		for i := int64(0); i < blkSize; i++ {
			elem, err := readSuperpower(r)
			if err != nil {
				return nil, err
			}
			arr = append(arr, elem)
		}
	}
	return arr, nil
}

func readBool(r io.Reader) (bool, error) {
	var b byte
	var err error
	if br, ok := r.(ByteReader); ok {
		b, err = br.ReadByte()
	} else {
		bs := make([]byte, 1)
		_, err = io.ReadFull(r, bs)
		if err != nil {
			return false, err
		}
		b = bs[0]
	}
	return b == 1, nil
}

func readFloat(r io.Reader) (float32, error) {
	buf := make([]byte, 4)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint32(buf)
	val := math.Float32frombits(bits)
	return val, nil

}

func readInt(r io.Reader) (int32, error) {
	var v int
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= int(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int32(v>>1) ^ -int32(v&1))
	return datum, nil
}

func readLong(r io.Reader) (int64, error) {
	var v uint64
	buf := make([]byte, 1)
	for shift := uint(0); ; shift += 7 {
		if _, err := io.ReadFull(r, buf); err != nil {
			return 0, err
		}
		b := buf[0]
		v |= uint64(b&127) << shift
		if b&128 == 0 {
			break
		}
	}
	datum := (int64(v>>1) ^ -int64(v&1))
	return datum, nil
}

func readString(r io.Reader) (string, error) {
	len, err := readLong(r)
	if err != nil {
		return "", err
	}

	// makeslice can fail depending on available memory.
	// We arbitrarily limit string size to sane default (~2.2GB).
	if len < 0 || len > math.MaxInt32 {
		return "", fmt.Errorf("string length out of range: %d", len)
	}

	bb := make([]byte, len)
	_, err = io.ReadFull(r, bb)
	if err != nil {
		return "", err
	}
	return string(bb), nil
}

func readSuperhero(r io.Reader) (*Superhero, error) {
	var str = &Superhero{}
	var err error
	str.Id, err = readInt(r)
	if err != nil {
		return nil, err
	}
	str.Affiliation_id, err = readInt(r)
	if err != nil {
		return nil, err
	}
	str.Name, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Life, err = readFloat(r)
	if err != nil {
		return nil, err
	}
	str.Energy, err = readFloat(r)
	if err != nil {
		return nil, err
	}
	str.Powers, err = readArraySuperpower(r)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func readSuperpower(r io.Reader) (*Superpower, error) {
	var str = &Superpower{}
	var err error
	str.Id, err = readInt(r)
	if err != nil {
		return nil, err
	}
	str.Name, err = readString(r)
	if err != nil {
		return nil, err
	}
	str.Damage, err = readFloat(r)
	if err != nil {
		return nil, err
	}
	str.Energy, err = readFloat(r)
	if err != nil {
		return nil, err
	}
	str.Passive, err = readBool(r)
	if err != nil {
		return nil, err
	}

	return str, nil
}

func writeArraySuperpower(r []*Superpower, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = writeSuperpower(e, w)
		if err != nil {
			return err
		}
	}
	return writeLong(0, w)
}

func writeBool(r bool, w io.Writer) error {
	var b byte
	if r {
		b = byte(1)
	}

	var err error
	if bw, ok := w.(ByteWriter); ok {
		err = bw.WriteByte(b)
	} else {
		bb := make([]byte, 1)
		bb[0] = b
		_, err = w.Write(bb)
	}
	if err != nil {
		return err
	}
	return nil
}

func writeFloat(r float32, w io.Writer) error {
	bits := uint64(math.Float32bits(r))
	const byteCount = 4
	return encodeFloat(w, byteCount, bits)
}

func writeInt(r int32, w io.Writer) error {
	downShift := uint32(31)
	encoded := uint64((uint32(r) << 1) ^ uint32(r>>downShift))
	const maxByteSize = 5
	return encodeInt(w, maxByteSize, encoded)
}

func writeLong(r int64, w io.Writer) error {
	downShift := uint64(63)
	encoded := uint64((r << 1) ^ (r >> downShift))
	const maxByteSize = 10
	return encodeInt(w, maxByteSize, encoded)
}

func writeString(r string, w io.Writer) error {
	err := writeLong(int64(len(r)), w)
	if err != nil {
		return err
	}
	if sw, ok := w.(StringWriter); ok {
		_, err = sw.WriteString(r)
	} else {
		_, err = w.Write([]byte(r))
	}
	return err
}

func writeSuperhero(r *Superhero, w io.Writer) error {
	var err error
	err = writeInt(r.Id, w)
	if err != nil {
		return err
	}
	err = writeInt(r.Affiliation_id, w)
	if err != nil {
		return err
	}
	err = writeString(r.Name, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Life, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Energy, w)
	if err != nil {
		return err
	}
	err = writeArraySuperpower(r.Powers, w)
	if err != nil {
		return err
	}

	return nil
}
func writeSuperpower(r *Superpower, w io.Writer) error {
	var err error
	err = writeInt(r.Id, w)
	if err != nil {
		return err
	}
	err = writeString(r.Name, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Damage, w)
	if err != nil {
		return err
	}
	err = writeFloat(r.Energy, w)
	if err != nil {
		return err
	}
	err = writeBool(r.Passive, w)
	if err != nil {
		return err
	}

	return nil
}
