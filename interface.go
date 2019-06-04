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

package key_value

// KeyValuer is thread-safe Key-Value object store
type KeyValuer interface {
	// Set allows you to set a value that is keyed to a string
	// @param key the string key that identifies the record
	// @param value is the new value to set, if the setIfTrue returns true
	Set(key string, value interface{})

	// Get allows you to retrieve a previously Set value
	// @param key the string key that identifies the record
	// @return v the value that was previously set, or nil, if not set
	// @return ok false if no value was found, true if the value is valid
	Get(key string) (v interface{}, ok bool)

	// CheckAndSet will set the value if and only if the setIfTrue returns true.
	// @param key the string key that identifies the record
	// @param value is the new value to set, if the setIfTrue returns true
	// @param setIfTrue is the test method that returns true if the set should occur, or
	// false if it should not. This function takes in 3 params:
	//   0 @param currentValue is the value that is currently in the map
	//   1 @param ok is true if original is found in the map, false if not
	CheckAndSet(key string, value interface{}, setIfTrue func(currentValue interface{}, ok bool) bool)

	// Del removes a value from the KeyValuer. Using Get after calling this on the same key
	// will cause the KeyValuer to behave as though the key was never Set.
	Del(key string)

	// MustGet is just like Get, but instead of returning the sentinel value if the value is
	// not found, this method will throw a panic
	MustGet(key string) (v interface{})

	// Copy creates a shallow clone of this KeyValuer so that it may be used in nested transactions in isolation
	// The shallow close does not copy the values in the map, but it will copy all of the keys and the references
	// to their values. If you stored the actual values and not references, those values will get their own copies,
	// but if you stored Interface references or struct references, the references will be copied and not the data
	// associated with that reference.
	Copy() KeyValuer
}