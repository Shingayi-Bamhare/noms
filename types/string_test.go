// Copyright 2016 The Noms Authors. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

package types

import (
	"testing"

	"github.com/attic-labs/testify/assert"
)

func TestStringEquals(t *testing.T) {
	assert := assert.New(t)
	s1 := NewString("foo")
	s2 := NewString("foo")
	s3 := s2
	s4 := NewString("bar")
	assert.True(s1.Equals(s2))
	assert.True(s2.Equals(s1))
	assert.True(s1.Equals(s3))
	assert.True(s3.Equals(s1))
	assert.False(s1.Equals(s4))
	assert.False(s4.Equals(s1))
}

func TestStringString(t *testing.T) {
	assert := assert.New(t)
	s1 := NewString("")
	s2 := NewString("foo")
	assert.Equal("", s1.String())
	assert.Equal("foo", s2.String())
}

func TestStringType(t *testing.T) {
	assert.True(t, NewString("hi").Type().Equals(StringType))
}
