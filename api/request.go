package api

import (
	"bytes"
	"encoding/binary"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var dtFormats = []string{
	"2006-01-02T15:04:05Z",
	"2006-01-02T15:04:05.999Z",
	"20060102T150405Z",
	"20060102T150405.999Z",
}

func ParamAsString(r *http.Request, name string, def string) string {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		return v
	}
	return def
}

func ParamAsInt(r *http.Request, name string, def int) int {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}

func ParamAsInt8(r *http.Request, name string, def int8) int8 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseInt(v, 10, 8); err == nil {
			return int8(i)
		}
	}
	return def
}

func ParamAsInt16(r *http.Request, name string, def int16) int16 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseInt(v, 10, 16); err == nil {
			return int16(i)
		}
	}
	return def
}

func ParamAsInt32(r *http.Request, name string, def int32) int32 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseInt(v, 10, 32); err == nil {
			return int32(i)
		}
	}
	return def
}

func ParamAsInt64(r *http.Request, name string, def int64) int64 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
	}
	return def
}

func ParamAsUint(r *http.Request, name string, def uint) uint {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if strconv.IntSize == 32 {
			if i, err := strconv.ParseUint(v, 10, 32); err == nil {
				return uint(i)
			}
		} else {
			if i, err := strconv.ParseUint(v, 10, 64); err == nil {
				return uint(i)
			}
		}
	}
	return def
}

func ParamAsUInt8(r *http.Request, name string, def uint8) uint8 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseUint(v, 10, 8); err == nil {
			return uint8(i)
		}
	}
	return def
}

func ParamAsUint16(r *http.Request, name string, def uint16) uint16 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseUint(v, 10, 16); err == nil {
			return uint16(i)
		}
	}
	return def

}

func ParamAsUint32(r *http.Request, name string, def uint32) uint32 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseUint(v, 10, 32); err == nil {
			return uint32(i)
		}
	}
	return def

}

func ParamAsUint64(r *http.Request, name string, def uint64) uint64 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseUint(v, 10, 64); err == nil {
			return i
		}
	}
	return def
}

func ParamAsFloat32(r *http.Request, name string, def float32) float32 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseFloat(v, 32); err == nil {
			return float32(i)
		}
	}
	return def
}

func ParamAsFloat64(r *http.Request, name string, def float64) float64 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		if i, err := strconv.ParseFloat(v, 64); err == nil {
			return i
		}
	}
	return def
}

func writeByteArray(v interface{}, order binary.ByteOrder, def []byte) []byte {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, order, v); err != nil {
		return def
	}
	return buf.Bytes()
}

func toByteArray(v string, kind reflect.Kind, def []byte) []byte {

	switch kind {
	case reflect.String:
		return []byte(v)
	case reflect.Int:
		i, err := strconv.Atoi(v)

		if err != nil {
			return def
		}

		return writeByteArray(int64(i), binary.LittleEndian, def)
	case reflect.Int8:
		i, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return def
		}

		return writeByteArray(int8(i), binary.LittleEndian, def)
	case reflect.Int16:
		i, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return def
		}

		return writeByteArray(int16(i), binary.LittleEndian, def)
	case reflect.Int32:
		i, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return def
		}

		return writeByteArray(int32(i), binary.LittleEndian, def)
	case reflect.Int64:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return def
		}

		return writeByteArray(i, binary.LittleEndian, def)
	case reflect.Uint:
		i, err := strconv.ParseUint(v, 10, strconv.IntSize)

		if err != nil {
			return def
		}

		return writeByteArray(i, binary.LittleEndian, def)
	case reflect.Uint8:
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return def
		}

		return writeByteArray(uint8(i), binary.LittleEndian, def)
	case reflect.Uint16:
		i, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return def
		}

		return writeByteArray(uint16(i), binary.LittleEndian, def)
	case reflect.Uint32:
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return def
		}

		return writeByteArray(uint32(i), binary.LittleEndian, def)
	case reflect.Uint64:
		i, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return def
		}

		return writeByteArray(i, binary.LittleEndian, def)
	case reflect.Float32:
		f, err := strconv.ParseFloat(v, 32)

		if err != nil {
			return def
		}

		return writeByteArray(float32(f), binary.LittleEndian, def)
	case reflect.Float64:
		f, err := strconv.ParseFloat(v, 64)

		if err != nil {
			return def
		}

		return writeByteArray(f, binary.LittleEndian, def)
	case reflect.Bool:
		b := toBoolean(v, false)

		return writeByteArray(b, binary.LittleEndian, def)
	default:
		return def
	}
}

func ParamAsByteArray(r *http.Request, name string, kind reflect.Kind, def []byte) []byte {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		return toByteArray(v, kind, def)
	}

	return def
}

func toBoolean(v string, def bool) bool {
	if len(v) == 0 {
		return def
	}

	v = strings.ToUpper(v)

	switch v[0] {
	case 'Y':
		v = "true"
	case 'N':
		v = "false"
	}

	if b, err := strconv.ParseBool(v); err == nil {
		return b
	}

	return def
}

func ParamAsBoolean(r *http.Request, name string, def bool) bool {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		return toBoolean(v, def)
	}
	return def
}

func ParamsAsTime(r *http.Request, name string, def time.Time) time.Time {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vars := mux.Vars(r)

	if v, ok := vars[name]; ok {
		for _, dtFormat := range dtFormats {
			if t, err := time.Parse(dtFormat, v); err == nil {
				return t
			}
		}
	}
	return def
}

func QueryParamAsString(r *http.Request, name string, def string) string {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		return v
	}

	return def
}

func QueryParamsAsInt(r *http.Request, name string, def int) int {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}

	return def
}

func QueryParamAsInt8(r *http.Request, name string, def int8) int8 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseInt(v, 10, 8); err == nil {
			return int8(i)
		}
	}

	return def
}

func QueryParamAsInt16(r *http.Request, name string, def int16) int16 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseInt(v, 10, 16); err == nil {
			return int16(i)
		}
	}

	return def
}

func QueryParamAsInt32(r *http.Request, name string, def int32) int32 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseInt(v, 10, 32); err == nil {
			return int32(i)
		}
	}

	return def
}

func QueryParamAsInt64(r *http.Request, name string, def int64) int64 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return i
		}
	}

	return def
}

func QueryParamsAsUint(r *http.Request, name string, def uint) uint {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if strconv.IntSize == 32 {
			if i, err := strconv.ParseUint(v, 10, 32); err == nil {
				return uint(i)
			}
		} else {
			if i, err := strconv.ParseUint(v, 10, 64); err == nil {
				return uint(i)
			}
		}
	}

	return def
}

func QueryParamAsUint8(r *http.Request, name string, def uint8) uint8 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseUint(v, 10, 8); err == nil {
			return uint8(i)
		}
	}

	return def
}

func QueryParamAsUint16(r *http.Request, name string, def uint16) uint16 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseUint(v, 10, 16); err == nil {
			return uint16(i)
		}
	}

	return def
}

func QueryParamAsUint32(r *http.Request, name string, def uint32) uint32 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseUint(v, 10, 32); err == nil {
			return uint32(i)
		}
	}

	return def
}

func QueryParamAsUint64(r *http.Request, name string, def uint64) uint64 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseUint(v, 10, 64); err == nil {
			return i
		}
	}

	return def
}

func QueryParamAsFloat32(r *http.Request, name string, def float32) float32 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseFloat(v, 32); err == nil {
			return float32(i)
		}
	}

	return def
}

func QueryParamAsFloat64(r *http.Request, name string, def float64) float64 {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		if i, err := strconv.ParseFloat(v, 64); err == nil {
			return i
		}
	}

	return def
}

func QueryParamAsByteArray(r *http.Request, name string, kind reflect.Kind, def []byte) []byte {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		return toByteArray(v, kind, def)
	}

	return def
}

func QueryParamAsBoolean(r *http.Request, name string, def bool) bool {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		return toBoolean(v, def)
	}

	return def
}

func QueryParamAsTime(r *http.Request, name string, def time.Time) time.Time {
	if r == nil {
		return def
	}

	if name == "" {
		return def
	}

	vs := r.URL.Query()

	if v := vs.Get(name); v != "" {
		for _, dtFormat := range dtFormats {
			if dt, err := time.Parse(dtFormat, v); err == nil {
				return dt
			}
		}
	}

	return def
}
