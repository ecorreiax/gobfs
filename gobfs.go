// Package gobfs offers utilities to generate hash values based on provided inputs.
//
// This package utilizes the standard library's hash interface and custom parser utilities
// from the "github.com/ecorreiax/gobfs/internal/parser" package.
package gobfs

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash"
)

// bitset is a boolean array used to store hash indexes.
var bitset []bool

type CreateHashStruct struct {
	Idx   int
	Error error
}

// CreateHash computes a hash index based on a given string and hash algorithm.
// It returns a CreateHashStruct containing the hash index and an error, if any.
// The function uses the given hash.Hash compliant hashing algorithm to hash the input string.
// If the string is empty or longer than 1 character, an error is returned.
//
// Parameters:
//
// h: A hash.Hash compliant hashing algorithm.
// s: The input string to hash.
//
// Returns:
//
// CreateHashStruct: A struct containing the calculated hash index and an error, if any.
//
// Example Usage:
//
// h := sha256.New()
// input := "a"
// result := CreateHash(h, input)
// if result.Error != nil {
//     log.Fatalf("Error creating hash: %v", result.Error)
// }
// fmt.Printf("Hash index: %d\n", result.Idx)

func CreateHash(h hash.Hash, s string) CreateHashStruct {
	if len(s) > 1 {
		return CreateHashStruct{
			Idx:   0,
			Error: errors.New("string can't be empty"),
		}
	}

	h.Write([]byte(s))
	bits := h.Sum(nil)
	buf := bytes.NewBuffer(bits)
	bufId, err := binary.ReadVarint(buf)
	if err != nil {
		return CreateHashStruct{
			Idx:   0,
			Error: err,
		}
	}

	return CreateHashStruct{
		Idx:   parseIdx(int(bufId)),
		Error: nil,
	}
}

// AddToHash adds a value into bitset based on the given index.
//
// This function first ensures that the index is non-negative. It then checks
// if the index falls outside the bounds of the existing bitset. If it does,
// the bitset is resized to accommodate the new index.
//
// Parameters:
//
//	idx int: The index at which the bit should be set. Negative values are converted to positive.
//
// Example:
//
//	AddHash(42)
func AddToHash(idx int) {
	if idx < 0 {
		idx = -idx
	}

	if parseIdx(idx) >= len(bitset) {
		newBitset := make([]bool, idx+1)
		copy(newBitset, bitset)
		bitset = newBitset
	}

	bitset[idx] = true
}

// VerifyHash checks if a bit is set at a given index in the global bitset.
//
// The function returns true if the bit at the specified index is set,
// otherwise false. It is the caller's responsibility to ensure that the
// index is within the bounds of the bitset.
//
// Parameters:
//
//	idx int: The index of the bit to be verified.
//
// Returns:
//
//	bool: True if the bit at the given index is set, otherwise false.
//
// Example:
//
//	result := VerifyHash(42)
func GetFromHash(idx int) bool {
	return bitset[idx]
}

func parseIdx(idx int) int {
	if idx < 0 {
		idx = -idx
		return idx
	}
	return idx
}
