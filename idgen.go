// Copyright Garrett Sparks.
// All Rights Reserved

// Package idgen uses crypt/rand to generate a random string of a given length.
package idgen

import "crypto/rand"

const (
	// DefaultIDChars defines the default set of characters as alphanumeric plus uppercase.
	DefaultIDChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// DefaultIDLength defines the default ID length as 8.
	DefaultIDLength = 8
)

// IDBuilder is a builder for an ID.
type IDBuilder struct {
	idLength     int
	idChars      string
	maxCharIndex int
}

// New creates a new IDBuilder with a specified length.
func New() *IDBuilder {
	return &IDBuilder{
		idLength:     DefaultIDLength,
		idChars:      DefaultIDChars,
		maxCharIndex: len(DefaultIDChars) - 1,
	}
}

// WithLength sets the length of the randomly generated ID.
// The default length is 8.
func (builder *IDBuilder) WithLength(length int) *IDBuilder {
	builder.idLength = length
	return builder
}

// WithCharset sets the string of available characters for the builder to use when generating an ID.
// The default is alphanumeric plus uppercase.
func (builder *IDBuilder) WithCharset(chars string) *IDBuilder {
	builder.idChars = chars
	builder.maxCharIndex = len(chars) - 1
	return builder
}

// BuildID builds an ID.
func (builder *IDBuilder) BuildID() string {
	if len(builder.idChars) < 1 {
		return ""
	}

	randBytes, randID := make([]byte, builder.idLength), make([]byte, builder.idLength)
	rand.Read(randBytes)

	for byteIndex, byteVal := range randBytes {
		newChar := builder.chooseChar(byteVal)
		randID[byteIndex] = newChar
	}

	return string(randID)
}

func (builder *IDBuilder) chooseChar(byteVal byte) byte {
	return builder.idChars[builder.bucketByte(byteVal)]
}

func (builder *IDBuilder) bucketByte(byteVal byte) int {
	return int(byteVal) * builder.maxCharIndex / 255
}
