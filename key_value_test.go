//Copyright 2019 Chris Wojno
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
// documentation files (the "Software"), to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
// Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
// WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS
// OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package go_keyvaluer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValue_MustGetExists(t *testing.T) {
	key := "test"
	value := "puppy"
	kvo := New()
	kvo.Set(key, value)
	assert.Equal(t, value, kvo.MustGet(key))
}

func TestKeyValue_MustGetNotExists(t *testing.T) {
	key := "test"
	kvo := New()
	assert.Panics(t, func() {
		kvo.MustGet(key)
	})
}

func TestKeyValue_GetExists(t *testing.T) {
	key := "test"
	value := "puppy"
	kvo := New()
	kvo.Set(key, value)
	actual, ok := kvo.Get(key)
	assert.Equal(t, value, actual)
	assert.True(t, ok)
}

func TestKeyValue_GetNotExists(t *testing.T) {
	key := "test"
	kvo := New()
	actual, ok := kvo.Get(key)
	assert.Empty(t, actual)
	assert.False(t, ok)
}

func TestKeyValue_CASYes(t *testing.T) {
	key := "test"
	value := "puppy"
	kvo := New()
	kvo.CheckAndSet(key, value, func(currentValue interface{}, ok bool) bool {
		return !ok
	})
	_, ok := kvo.Get(key)
	assert.True(t, ok)
}

func TestKeyValue_CASNo(t *testing.T) {
	key := "test"
	value := "puppy"
	value2 := "puppy2"
	kvo := New()
	kvo.Set(key, value)
	kvo.CheckAndSet(key, value2, func(currentValue interface{}, ok bool) bool {
		return !ok
	})
	actual, _ := kvo.Get(key)
	assert.Equal(t, value, actual)
}

func TestKeyValue_Del(t *testing.T) {
	key := "test"
	value := "puppy"
	kvo := New()
	kvo.Set(key, value)
	kvo.Del(key)
	_, ok := kvo.Get(key)
	assert.False(t, ok)
}

func TestKeyValue_Copy(t *testing.T) {
	key := "test"
	value := "puppy"
	kvo := New()
	kvo.Set(key, value)

	kvoCopy := kvo.Copy()
	_, ok := kvoCopy.Get(key)
	assert.True(t, ok)

	kvo.Del(key)
	_, ok = kvoCopy.Get(key)
	assert.True(t, ok)
}
