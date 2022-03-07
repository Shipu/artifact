package artifact

import (
	"github.com/goldeneggg/structil"
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"strings"
	"unsafe"
	"github.com/shipu/artifact/env"
)

var Config *Configuration

type Configuration struct {
	RegisteredConfigStruct map[string]interface{}
	LoadedConfig           map[string]interface{}
}

type Getter struct {
	*structil.Getter
}

func NewConfig() *Configuration {
	return &Configuration{
		RegisteredConfigStruct: make(map[string]interface{}),
	}
}

func (configuration *Configuration) Load() map[string]interface{} {
	newConfig := make(map[string]interface{})
	v := env.New(viper.New())
	v.AutomaticEnv()
	for name, value := range configuration.RegisteredConfigStruct {
		err := v.Unmarshal(&value)
		if err != nil {
			log.Fatal("environment cant be loaded: ", err)
		}

		defaults.SetDefaults(value)

		newConfig[name] = value
	}

	configuration.LoadedConfig = newConfig

	return newConfig
}

func (configuration *Configuration) AddConfig(name string, userConfig interface{}) *Configuration {
	configuration.RegisteredConfigStruct[name] = userConfig

	return configuration
}

func (configuration Configuration) keySplit(key string) (string, string) {
	var rootKey string

	splitKey := strings.Split(key, ".")

	rootKey, splitKey = splitKey[0], splitKey[1:]

	key = strings.Join(splitKey, ".")

	return rootKey, key
}

func (configuration Configuration) prepareConfigKey(key string) (string, *structil.Getter, error) {
	rootKey, key := configuration.keySplit(key)

	newGetter, err := structil.NewGetter(configuration.LoadedConfig[rootKey])

	return key, newGetter, err
}

func (configuration Configuration) GetString(key string) string {
	value, _ := configuration.String(key)

	return value
}

// String returns the string of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not string.
func (configuration *Configuration) String(key string) (string, bool) {

	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.String(key)
}

// Int returns the int of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not int.
func (configuration Configuration) Int(key string) (int, bool) {

	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Int(key)
}

// Get returns the interface of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
func (configuration Configuration) Get(key string) (interface{}, bool) {

	newKey, newGetter, _ := configuration.prepareConfigKey(key)

	if newKey == "" {
		return newGetter.ToMap(), true
	}

	return newGetter.Get(newKey)
}

func (configuration Configuration) Getter(key string) *Getter {
	_, newGetter, _ := configuration.prepareConfigKey(key)

	return &Getter{newGetter}
}

// NumField returns num of struct field.
func (configuration Configuration) NumField(key string) int {
	newGetter := configuration.Getter(key)

	return newGetter.NumField()
}

// Names returns names of struct field.
func (configuration Configuration) Names(key string) []string {
	newGetter := configuration.Getter(key)

	return newGetter.Names()
}

// Has tests whether the original struct has a field named "name".
func (configuration Configuration) Has(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Has(key)
}

// GetType returns the reflect.Type object of the original struct field named "name".
// 2nd return value will be false if the original struct does not have a "name" field.
func (configuration Configuration) GetType(key string) (reflect.Type, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.GetType(key)
}

// GetValue returns the reflect.Value object of the original struct field named "name".
// 2nd return value will be false if the original struct does not have a "name" field.
func (configuration Configuration) GetValue(key string) (reflect.Value, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.GetValue(key)
}

// ToMap returns a map converted from this Getter.
func (configuration Configuration) ToMap(key string) map[string]interface{} {
	newGetter := configuration.Getter(key)

	return newGetter.ToMap()
}

// IsSlice reports whether type of the original struct field named name is slice.
func (configuration Configuration) IsSlice(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsSlice(key)
}

// Slice returns the slice of interface of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not slice of interface.
func (configuration Configuration) Slice(key string) ([]interface{}, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Slice(key)
}

// IsBool reports whether type of the original struct field named name is bool.
func (configuration Configuration) IsBool(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsBool(key)
}

// Bool returns the byte of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not bool.
func (configuration Configuration) Bool(key string) (bool, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Bool(key)
}

// IsByte reports whether type of the original struct field named name is byte.
func (configuration Configuration) IsByte(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsByte(key)
}

// Byte returns the byte of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not byte.
func (configuration Configuration) Byte(key string) (byte, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Byte(key)
}

// IsBytes reports whether type of the original struct field named name is []byte.
func (configuration Configuration) IsBytes(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsBytes(key)
}

// Bytes returns the []byte of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not []byte.
func (configuration Configuration) Bytes(key string) ([]byte, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Bytes(key)
}

// IsString reports whether type of the original struct field named name is string.
func (configuration Configuration) IsString(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsString(key)
}

// IsInt reports whether type of the original struct field named name is int.
func (configuration Configuration) IsInt(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsInt(key)
}

// IsInt8 reports whether type of the original struct field named name is int8.
func (configuration Configuration) IsInt8(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsInt8(key)
}

// Int8 returns the int8 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not int8.
func (configuration Configuration) Int8(key string) (int8, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Int8(key)
}

// IsInt16 reports whether type of the original struct field named name is int16.
func (configuration Configuration) IsInt16(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsInt16(key)
}

// Int16 returns the int16 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not int16.
func (configuration Configuration) Int16(key string) (int16, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Int16(key)
}

// IsInt32 reports whether type of the original struct field named name is int32.
func (configuration Configuration) IsInt32(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsInt32(key)
}

// Int32 returns the int32 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not int32.
func (configuration Configuration) Int32(key string) (int32, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Int32(key)
}

// IsInt64 reports whether type of the original struct field named name is int64.
func (configuration Configuration) IsInt64(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsInt64(key)
}

// Int64 returns the int64 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not int64.
func (configuration Configuration) Int64(key string) (int64, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Int64(key)
}

// IsUint reports whether type of the original struct field named name is uint.
func (configuration Configuration) IsUint(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUint(key)
}

// Uint returns the uint of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not uint.
func (configuration Configuration) Uint(key string) (uint, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Uint(key)
}

// IsUint8 reports whether type of the original struct field named name is uint8.
func (configuration Configuration) IsUint8(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUint8(key)
}

// Uint8 returns the uint8 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not uint8.
func (configuration Configuration) Uint8(key string) (uint8, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Uint8(key)
}

// IsUint16 reports whether type of the original struct field named name is uint16.
func (configuration Configuration) IsUint16(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUint16(key)
}

// Uint16 returns the uint16 of the original struct field named name.Getter
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not uint16.
func (configuration Configuration) Uint16(key string) (uint16, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Uint16(key)
}

// IsUint32 reports whether type of the original struct field named name is uint32.
func (configuration Configuration) IsUint32(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUint32(key)
}

// Uint32 returns the uint32 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not uint32.
func (configuration Configuration) Uint32(key string) (uint32, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Uint32(key)
}

// IsUint64 reports whether type of the original struct field named name is uint64.
func (configuration Configuration) IsUint64(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUint64(key)
}

// Uint64 returns the uint64 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not uint64.
func (configuration Configuration) Uint64(key string) (uint64, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Uint64(key)
}

// IsUintptr reports whether type of the original struct field named name is uintptr.
func (configuration Configuration) IsUintptr(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUintptr(key)
}

// Uintptr returns the uintptr of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not uintptr.
func (configuration Configuration) Uintptr(key string) (uintptr, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Uintptr(key)
}

// IsFloat32 reports whether type of the original struct field named name is float32.
func (configuration Configuration) IsFloat32(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsFloat32(key)
}

// Float32 returns the float32 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not float32.
func (configuration Configuration) Float32(key string) (float32, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Float32(key)
}

// IsFloat64 reports whether type of the original struct field named name is float64.
func (configuration Configuration) IsFloat64(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsFloat64(key)
}

// Float64 returns the float64 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not float64.
func (configuration Configuration) Float64(key string) (float64, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Float64(key)
}

// IsComplex64 reports whether type of the original struct field named name is []byte.
func (configuration Configuration) IsComplex64(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsComplex64(key)
}

// Complex64 returns the complex64 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not complex64.
func (configuration Configuration) Complex64(key string) (complex64, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Complex64(key)
}

// IsComplex128 reports whether type of the original struct field named name is []byte.
func (configuration Configuration) IsComplex128(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsComplex128(key)
}

// Complex128 returns the complex128 of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not complex128.
func (configuration Configuration) Complex128(key string) (complex128, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.Complex128(key)
}

// IsUnsafePointer reports whether type of the original struct field named name is []byte.
func (configuration Configuration) IsUnsafePointer(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsUnsafePointer(key)
}

// UnsafePointer returns the unsafe.Pointer of the original struct field named name.
// 2nd return value will be false if the original struct does not have a "name" field.
// 2nd return value will be false if type of the original struct "name" field is not unsafe.Pointer.
func (configuration Configuration) UnsafePointer(key string) (unsafe.Pointer, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.UnsafePointer(key)
}

// IsMap reports whether type of the original struct field named name is map.
func (configuration Configuration) IsMap(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsMap(key)
}

// IsFunc reports whether type of the original struct field named name is func.
func (configuration Configuration) IsFunc(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsFunc(key)
}

// IsChan reports whether type of the original struct field named name is chan.
func (configuration Configuration) IsChan(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsChan(key)
}

// IsStruct reports whether type of the original struct field named name is struct.
func (configuration Configuration) IsStruct(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsStruct(key)
}

// IsArray reports whether type of the original struct field named name is slice.
func (configuration Configuration) IsArray(key string) bool {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.IsArray(key)
}

func (configuration Configuration) GetGetter(key string) (*structil.Getter, bool) {
	key, newGetter, _ := configuration.prepareConfigKey(key)

	return newGetter.GetGetter(key)
}
