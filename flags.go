// mauflag - An extendable command-line argument parser for Golang
// Copyright (C) 2016 Maunium
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package flag

var flags []*Flag

// Flag represents a single flag
type Flag struct {
	longKeys  []string
	shortKeys []string
	Value     Value
	defaul    string
	usage     string
}

// Make creates and registers a flag
func Make() *Flag {
	flag := &Flag{}
	val := stringValue("")
	flag.Value = &val
	flag.activateDefaultValue()
	flags = append(flags, flag)
	return flag
}

func (flag *Flag) setValue(val string) error {
	return flag.Value.Set(val)
}

// Usage sets the usage of this Flag
func (flag *Flag) Usage(usage string) *Flag {
	flag.usage = usage
	return flag
}

// Default sets the default value of this Flag
func (flag *Flag) Default(defaul string) *Flag {
	flag.defaul = defaul
	return flag
}

func (flag *Flag) activateDefaultValue() {
	if len(flag.defaul) > 0 {
		flag.Value.Set(flag.defaul)
	}
}

// LongKey adds a long key to this Flag
func (flag *Flag) LongKey(key string) *Flag {
	flag.longKeys = append(flag.longKeys, key)
	return flag
}

// ShortKey adds a short key to this Flag
func (flag *Flag) ShortKey(key string) *Flag {
	flag.shortKeys = append(flag.shortKeys, key)
	return flag
}

// Custom sets a custom object that implemets Value as the value of this flag
func (flag *Flag) Custom(val Value) {
	flag.Value = val
	flag.activateDefaultValue()
}

// Bool sets the type of this flag to a boolean and returns a pointer to the value
func (flag *Flag) Bool() *bool {
	val := boolValue(false)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*bool)(&val)
}

// String sets the type of this flag to a string and returns a pointer to the value
func (flag *Flag) String() *string {
	val := stringValue("")
	flag.Value = &val
	flag.activateDefaultValue()
	return (*string)(&val)
}

// StringArray sets the type of this flag to a string array and returns a pointer to the value
func (flag *Flag) StringArray() *[]string {
	val := stringArrayValue([]string{})
	flag.Value = &val
	flag.activateDefaultValue()
	return (*[]string)(&val)
}

// Int sets the type of this flag to a signed default-length integer and returns a pointer to the value
func (flag *Flag) Int() *int {
	val := intValue(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*int)(&val)
}

// Uint sets the type of this flag to an unsigned default-length integer and returns a pointer to the value
func (flag *Flag) Uint() *uint {
	val := uintValue(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*uint)(&val)
}

// Int8 sets the type of this flag to a signed 8-bit integer and returns a pointer to the value
func (flag *Flag) Int8() *int8 {
	val := int8Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*int8)(&val)
}

// Uint8 sets the type of this flag to an unsigned 8-bit integer and returns a pointer to the value
func (flag *Flag) Uint8() *uint8 {
	val := uint8Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*uint8)(&val)
}

// Byte sets the type of this flag to a byte (unsigned 8-bit integer) and returns a pointer to the value
func (flag *Flag) Byte() *byte {
	val := byteValue(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*byte)(&val)
}

// Int16 sets the type of this flag to a signed 16-bit integer and returns a pointer to the value
func (flag *Flag) Int16() *int16 {
	val := int16Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*int16)(&val)
}

// Uint16 sets the type of this flag to an unsigned 16-bit integer and returns a pointer to the value
func (flag *Flag) Uint16() *uint16 {
	val := uint16Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*uint16)(&val)
}

// Int32 sets the type of this flag to a signed 32-bit integer and returns a pointer to the value
func (flag *Flag) Int32() *int32 {
	val := int32Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*int32)(&val)
}

// Rune sets the type of this flag to a rune (signed 32-bit integer) and returns a pointer to the value
func (flag *Flag) Rune() *rune {
	val := runeValue(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*rune)(&val)
}

// Uint32 sets the type of this flag to an unsigned 32-bit integer and returns a pointer to the value
func (flag *Flag) Uint32() *uint32 {
	val := uint32Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*uint32)(&val)
}

// Int64 sets the type of this flag to a signed 64-bit integer and returns a pointer to the value
func (flag *Flag) Int64() *int64 {
	val := int64Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*int64)(&val)
}

// Uint64 sets the type of this flag to an unsigned 64-bit integer and returns a pointer to the value
func (flag *Flag) Uint64() *uint64 {
	val := uint64Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*uint64)(&val)
}

// Float32 sets the type of this flag to an 32-bit float and returns a pointer to the value
func (flag *Flag) Float32() *float32 {
	val := float32Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*float32)(&val)
}

// Float64 sets the type of this flag to an 64-bit float and returns a pointer to the value
func (flag *Flag) Float64() *float64 {
	val := float64Value(0)
	flag.Value = &val
	flag.activateDefaultValue()
	return (*float64)(&val)
}
