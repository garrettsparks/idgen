// Copyright Garrett Sparks.
// All Rights Reserved

package idgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("it should properly calculate maxCharIndex", func(t *testing.T) {
		builder := New()
		assert.Equal(t, 61, builder.maxCharIndex)
	})
}

func TestWithCharset(t *testing.T) {
	t.Run("it should build an ID when only one character is supplied", func(t *testing.T) {
		idLength := 5
		builder := New().WithCharset("a").WithLength(idLength)
		id := builder.BuildID()
		assert.Equal(t, "aaaaa", id)
	})

	t.Run("it should return an empty ID when no characters are supplied", func(t *testing.T) {
		idLength := 5
		builder := New().WithCharset("").WithLength(idLength)
		id := builder.BuildID()
		assert.Equal(t, "", id)
	})
}

func TestWithLength(t *testing.T) {
	t.Run("it should build an ID of size 10", func(t *testing.T) {
		idLength := 10
		builder := New().WithLength(idLength)
		id := builder.BuildID()
		assert.Len(t, id, idLength)
	})

	t.Run("it should build an ID of size 0", func(t *testing.T) {
		idLength := 0
		builder := New().WithLength(idLength)
		id := builder.BuildID()
		assert.Len(t, id, idLength)
	})

	t.Run("it should build an ID of size 1", func(t *testing.T) {
		idLength := 1
		builder := New().WithLength(idLength)
		id := builder.BuildID()
		assert.Len(t, id, idLength)
	})
}

func TestChooseChar(t *testing.T) {
	builder := New()

	t.Run("it should choose a 0 byte character", func(t *testing.T) {
		char := string([]byte{builder.chooseChar(0)})
		assert.Equal(t, "a", char)
	})

	t.Run("it should choose a max byte character", func(t *testing.T) {
		char := string([]byte{builder.chooseChar(255)})
		assert.Equal(t, "9", char)
	})

	t.Run("it should choose a middle byte character", func(t *testing.T) {
		char := string([]byte{builder.chooseChar(128)})
		assert.Equal(t, "E", char)
	})
}

func TestBucketByte(t *testing.T) {
	builder := New()

	t.Run("it should bucket a 0 byte", func(t *testing.T) {
		bucket := builder.bucketByte(0)
		assert.Zero(t, bucket)
	})

	t.Run("it should bucket a max byte", func(t *testing.T) {
		bucket := builder.bucketByte(255)
		assert.Equal(t, 61, bucket)
	})

	t.Run("it should bucket a middle byte", func(t *testing.T) {
		bucket := builder.bucketByte(128)
		assert.Equal(t, 30, bucket)
	})
}
