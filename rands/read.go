package rands

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	ErrInsufficientRandom = errors.New("insufficient random bytes")
)

func RandomUint64Array(length int) ([]uint64, error) {
	if length < 1 {
		return []uint64{}, nil
	}

	buf := make([]byte, length*8)

	if n, err := rand.Read(buf); err != nil {
		return nil, fmt.Errorf("reading random for RandomInt64Array: %+v", err)
	} else if n != len(buf) {
		return nil, ErrInsufficientRandom
	}

	// convert byte array to uint64 array
	ret := make([]uint64, length)
	for i := 0; i < length; i++ {
		ret[i] = uint64(binary.BigEndian.Uint64(buf[i*8 : (i+1)*8]))
	}

	return ret, nil
}

func RandomUint32Array(length int) ([]uint32, error) {
	if length < 1 {
		return []uint32{}, nil
	}

	buf := make([]byte, length*4)

	if n, err := rand.Read(buf); err != nil {
		return nil, fmt.Errorf("reading random for RandomInt32Array: %+v", err)
	} else if n != len(buf) {
		return nil, ErrInsufficientRandom
	}

	// convert byte array to uint32 array
	ret := make([]uint32, length)
	for i := 0; i < length; i++ {
		ret[i] = binary.BigEndian.Uint32(buf[i*4 : (i+1)*4])
	}

	return ret, nil
}
