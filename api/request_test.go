package api

import (
	"bytes"
	"encoding/binary"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

func TestParamAsBoolean(t *testing.T) {

	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"true":  "true",
		"TRUE":  "TRUE",
		"false": "false",
		"FALSE": "FALSE",
		"t":     "t",
		"T":     "T",
		"f":     "f",
		"F":     "F",
		"yes":   "yes",
		"YES":   "YES",
		"no":    "no",
		"NO":    "NO",
		"y":     "y",
		"Y":     "Y",
		"n":     "n",
		"N":     "N",
	})

	type args struct {
		r    *http.Request
		name string
		def  bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test true",
			args{
				r:    req,
				name: "true",
				def:  false,
			},
			true,
		},
		{
			"Test TRUE",
			args{
				r:    req,
				name: "TRUE",
				def:  false,
			},
			true,
		},
		{
			"Test false",
			args{
				r:    req,
				name: "false",
				def:  true,
			},
			false,
		},
		{
			"Test FALSE",
			args{
				r:    req,
				name: "FALSE",
				def:  true,
			},
			false,
		},
		{
			"Test t",
			args{
				r:    req,
				name: "t",
				def:  false,
			},
			true,
		},
		{
			"Test T",
			args{
				r:    req,
				name: "T",
				def:  false,
			},
			true,
		},
		{
			"Test f",
			args{
				r:    req,
				name: "f",
				def:  true,
			},
			false,
		},
		{
			"Test F",
			args{
				r:    req,
				name: "F",
				def:  true,
			},
			false,
		},
		{
			"Test yes",
			args{
				r:    req,
				name: "yes",
				def:  false,
			},
			true,
		},
		{
			"Test YES",
			args{
				r:    req,
				name: "YES",
				def:  false,
			},
			true,
		},
		{
			"Test no",
			args{
				r:    req,
				name: "no",
				def:  true,
			},
			false,
		},
		{
			"Test NO",
			args{
				r:    req,
				name: "NO",
				def:  true,
			},
			false,
		},
		{
			"Test y",
			args{
				r:    req,
				name: "y",
				def:  false,
			},
			true,
		},
		{
			"Test Y",
			args{
				r:    req,
				name: "Y",
				def:  false,
			},
			true,
		},
		{
			"Test n",
			args{
				r:    req,
				name: "n",
				def:  true,
			},
			false,
		},
		{
			"Test N",
			args{
				r:    req,
				name: "N",
				def:  true,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsBoolean(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getBytes(v interface{}) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, v)
	if err != nil {
		return nil
	}

	return buf.Bytes()
}

func TestParamAsByteArray(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"string":  "String",
		"int":     "1234",
		"int8":    "123",
		"int16":   "1234",
		"int32":   "1234",
		"int64":   "1234",
		"uint":    "1234",
		"uint8":   "123",
		"uint16":  "1234",
		"uint32":  "1234",
		"uint64":  "1234",
		"float32": "123.999",
		"float64": "123.999",
		"true":    "true",
		"false":   "false",
		"TRUE":    "TRUE",
		"FALSE":   "FALSE",
		"t":       "t",
		"f":       "f",
		"T":       "T",
		"F":       "F",
		"yes":     "yes",
		"no":      "no",
		"YES":     "YES",
		"NO":      "NO",
		"y":       "y",
		"n":       "n",
		"Y":       "Y",
		"N":       "N",
	})

	type args struct {
		r    *http.Request
		name string
		kind reflect.Kind
		def  []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"String",
			args{
				req,
				"string",
				reflect.String,
				nil,
			},
			[]byte("String"),
		},
		{
			"int",
			args{
				req,
				"int",
				reflect.Int,
				nil,
			},
			getBytes(int64(1234)),
		},
		{
			"int8",
			args{
				req,
				"int8",
				reflect.Int8,
				nil,
			},
			getBytes(int8(123)),
		},
		{
			"int16",
			args{
				req,
				"int16",
				reflect.Int16,
				nil,
			},
			getBytes(int16(1234)),
		},
		{
			"int32",
			args{
				req,
				"int32",
				reflect.Int32,
				nil,
			},
			getBytes(int32(1234)),
		},
		{
			"int64",
			args{
				req,
				"int64",
				reflect.Int64,
				nil,
			},
			getBytes(int64(1234)),
		},
		{
			"uint",
			args{
				req,
				"uint",
				reflect.Uint,
				nil,
			},
			getBytes(uint64(1234)),
		},
		{
			"uint8",
			args{
				req,
				"uint8",
				reflect.Uint8,
				nil,
			},
			getBytes(uint8(123)),
		},
		{
			"uint16",
			args{
				req,
				"uint16",
				reflect.Uint16,
				nil,
			},
			getBytes(uint16(1234)),
		},
		{
			"uint32",
			args{
				req,
				"uint32",
				reflect.Uint32,
				nil,
			},
			getBytes(uint32(1234)),
		},
		{
			"uint64",
			args{
				req,
				"uint64",
				reflect.Uint64,
				nil,
			},
			getBytes(uint64(1234)),
		},
		{
			"float32",
			args{
				req,
				"float32",
				reflect.Float32,
				nil,
			},
			getBytes(float32(123.999)),
		},
		{
			"float64",
			args{
				req,
				"float64",
				reflect.Float64,
				nil,
			},
			getBytes(123.999),
		},
		{
			"true",
			args{
				req,
				"true",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"false",
			args{
				req,
				"false",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"TRUE",
			args{
				req,
				"TRUE",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"FALSE",
			args{
				req,
				"FALSE",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"t",
			args{
				req,
				"t",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"f",
			args{
				req,
				"f",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"T",
			args{
				req,
				"T",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"F",
			args{
				req,
				"F",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"yes",
			args{
				req,
				"yes",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"no",
			args{
				req,
				"no",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"YES",
			args{
				req,
				"YES",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"NO",
			args{
				req,
				"NO",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"y",
			args{
				req,
				"y",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"n",
			args{
				req,
				"n",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
		{
			"Y",
			args{
				req,
				"Y",
				reflect.Bool,
				nil,
			},
			getBytes(true),
		},
		{
			"N",
			args{
				req,
				"N",
				reflect.Bool,
				nil,
			},
			getBytes(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsByteArray(tt.args.r, tt.args.name, tt.args.kind, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamAsByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsFloat32(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"value": "123.999",
		"int":   "123",
		"text":  "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			"Parse valid float",
			args{
				req,
				"value",
				0.0,
			},
			123.999,
		},
		{
			"Parse invalid float",
			args{
				req,
				"text",
				111.11,
			},
			111.11,
		},
		{
			"Parse int",
			args{
				req,
				"int",
				111.11,
			},
			123,
		},
		{
			"nil request",
			args{
				nil,
				"value",
				111.11,
			},
			111.11,
		},
		{
			"empty name",
			args{
				req,
				"",
				111.11,
			},
			111.11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsFloat32(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsFloat64(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"value": "123.999",
		"int":   "123",
		"text":  "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"Parse valid float",
			args{
				req,
				"value",
				0.0,
			},
			123.999,
		},
		{
			"Parse invalid float",
			args{
				req,
				"text",
				111.11,
			},
			111.11,
		},
		{
			"Parse int",
			args{
				req,
				"int",
				111.11,
			},
			123,
		},
		{
			"nil request",
			args{
				nil,
				"value",
				111.11,
			},
			111.11,
		},
		{
			"empty name",
			args{
				req,
				"",
				111.11,
			},
			111.11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsFloat64(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsInt(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "123456",
		"negative": "-123456",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			123456,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			-123456,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsInt(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsInt32(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "123456",
		"negative": "-123456",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			123456,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			-123456,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsInt32(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsInt64(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "123456",
		"negative": "-123456",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			123456,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			-123456,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsInt64(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsString(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"string": "string",
		"int":    "123456",
		"float":  "1234.56",
		"bool":   "true",
	})

	type args struct {
		r    *http.Request
		name string
		def  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"string",
			args{
				req,
				"string",
				"fail",
			},
			"string",
		},
		{
			"int",
			args{
				req,
				"int",
				"fail",
			},
			"123456",
		},
		{
			"float",
			args{
				req,
				"float",
				"fail",
			},
			"1234.56",
		},
		{
			"bool",
			args{
				req,
				"bool",
				"fail",
			},
			"true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsString(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamsAsTime(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"2006-01-02T15:04:05Z":     "2020-07-15T09:43:26Z",
		"2006-01-02T15:04:05.999Z": "2020-07-15T09:43:26.578Z",
		"20060102T150405Z":         "20200715T094326Z",
		"20060102T150405.999Z":     "20200715T094326.578Z",
	})

	type args struct {
		r    *http.Request
		name string
		def  time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"2006-01-02T15:04:05Z",
			args{
				req,
				"2006-01-02T15:04:05Z",
				time.Unix(0, 0),
			},
			time.Date(2020, 7, 15, 9, 43, 26, 0, time.UTC),
		},
		{"2006-01-02T15:04:05.999Z",
			args{
				req,
				"2006-01-02T15:04:05.999Z",
				time.Unix(0, 0),
			},
			time.Date(2020, 7, 15, 9, 43, 26, 578000000, time.UTC),
		},
		{"20060102T150405Z",
			args{
				req,
				"20060102T150405Z",
				time.Unix(0, 0),
			},
			time.Date(2020, 7, 15, 9, 43, 26, 0, time.UTC),
		},
		{"20060102T150405.999Z",
			args{
				req,
				"20060102T150405.999Z",
				time.Unix(0, 0),
			},
			time.Date(2020, 7, 15, 9, 43, 26, 578000000, time.UTC),
		},
		{"Test nil request",
			args{
				nil,
				"",
				time.Unix(0, 0),
			},
			time.Unix(0, 0),
		},
		{"Test empty name",
			args{
				req,
				"",
				time.Unix(0, 0),
			},
			time.Unix(0, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamsAsTime(tt.args.r, tt.args.name, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamsAsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsBool(t *testing.T) {
	req := httptest.NewRequest("GET",
		"/test?true=true&TRUE=TRUE&false=false&FALSE=FALSE&t=t&f=f&T=T&F=F&yes=yes&no=no&YES=YES&NO=NO&y=y&n=n&Y=Y&N=N",
		nil,
	)

	type args struct {
		r    *http.Request
		name string
		def  bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"true", args{req, "true", false}, true},
		{"TRUE", args{req, "TRUE", false}, true},
		{"FALSE", args{req, "FALSE", true}, false},
		{"false", args{req, "false", true}, false},
		{"t", args{req, "t", false}, true},
		{"T", args{req, "T", false}, true},
		{"f", args{req, "f", true}, false},
		{"F", args{req, "F", true}, false},
		{"yes", args{req, "yes", false}, true},
		{"YES", args{req, "YES", false}, true},
		{"NO", args{req, "NO", true}, false},
		{"no", args{req, "no", true}, false},
		{"y", args{req, "y", false}, true},
		{"Y", args{req, "Y", false}, true},
		{"n", args{req, "n", true}, false},
		{"N", args{req, "N", true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsBoolean(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsBoolean() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsByteArray(t *testing.T) {
	type args struct {
		r    *http.Request
		name string
		kind reflect.Kind
		def  []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"string", args{httptest.NewRequest("GET", "/test?string=String", nil), "string", reflect.String, nil}, []byte("String")},
		{"int", args{httptest.NewRequest("GET", "/test?int=1234", nil), "int", reflect.Int, nil}, getBytes(int64(1234))},
		{"int8", args{httptest.NewRequest("GET", "/test?int8=123", nil), "int8", reflect.Int8, nil}, getBytes(int8(123))},
		{"int16", args{httptest.NewRequest("GET", "/test?int16=1234", nil), "int16", reflect.Int16, nil}, getBytes(int16(1234))},
		{"int32", args{httptest.NewRequest("GET", "/test?int32=1234", nil), "int32", reflect.Int32, nil}, getBytes(int32(1234))},
		{"int64", args{httptest.NewRequest("GET", "/test?int64=1234", nil), "int64", reflect.Int64, nil}, getBytes(int64(1234))},
		{"uint", args{httptest.NewRequest("GET", "/test?uint=1234", nil), "uint", reflect.Uint, nil}, getBytes(uint64(1234))},
		{"uint8", args{httptest.NewRequest("GET", "/test?uint8=123", nil), "uint8", reflect.Uint8, nil}, getBytes(uint8(123))},
		{"uint16", args{httptest.NewRequest("GET", "/test?uint16=1234", nil), "uint16", reflect.Uint16, nil}, getBytes(uint16(1234))},
		{"uint32", args{httptest.NewRequest("GET", "/test?uint32=1234", nil), "uint32", reflect.Uint32, nil}, getBytes(uint32(1234))},
		{"uint64", args{httptest.NewRequest("GET", "/test?uint64=1234", nil), "uint64", reflect.Uint64, nil}, getBytes(uint64(1234))},
		{"float32", args{httptest.NewRequest("GET", "/test?float32=123.99", nil), "float32", reflect.Float32, nil}, getBytes(float32(123.99))},
		{"float64", args{httptest.NewRequest("GET", "/test?float64=123.99", nil), "float64", reflect.Float64, nil}, getBytes(123.99)},
		{"true", args{httptest.NewRequest("GET", "/test?true=true", nil), "true", reflect.Bool, nil}, getBytes(true)},
		{"false", args{httptest.NewRequest("GET", "/test?false=false", nil), "false", reflect.Bool, nil}, getBytes(false)},
		{"TRUE", args{httptest.NewRequest("GET", "/test?TRUE=TRUE", nil), "TRUE", reflect.Bool, nil}, getBytes(true)},
		{"FALSE", args{httptest.NewRequest("GET", "/test?FALSE=FALSE", nil), "FALSE", reflect.Bool, nil}, getBytes(false)},
		{"t", args{httptest.NewRequest("GET", "/test?t=t", nil), "t", reflect.Bool, nil}, getBytes(true)},
		{"f", args{httptest.NewRequest("GET", "/test?f=f", nil), "f", reflect.Bool, nil}, getBytes(false)},
		{"T", args{httptest.NewRequest("GET", "/test?T=T", nil), "T", reflect.Bool, nil}, getBytes(true)},
		{"F", args{httptest.NewRequest("GET", "/test?F=F", nil), "F", reflect.Bool, nil}, getBytes(false)},
		{"yes", args{httptest.NewRequest("GET", "/test?yes=yes", nil), "yes", reflect.Bool, nil}, getBytes(true)},
		{"no", args{httptest.NewRequest("GET", "/test?no=no", nil), "no", reflect.Bool, nil}, getBytes(false)},
		{"YES", args{httptest.NewRequest("GET", "/test?YES=YES", nil), "YES", reflect.Bool, nil}, getBytes(true)},
		{"NO", args{httptest.NewRequest("GET", "/test?NO=NO", nil), "NO", reflect.Bool, nil}, getBytes(false)},
		{"y", args{httptest.NewRequest("GET", "/test?y=y", nil), "y", reflect.Bool, nil}, getBytes(true)},
		{"n", args{httptest.NewRequest("GET", "/test?n=n", nil), "n", reflect.Bool, nil}, getBytes(false)},
		{"Y", args{httptest.NewRequest("GET", "/test?Y=Y", nil), "Y", reflect.Bool, nil}, getBytes(true)},
		{"N", args{httptest.NewRequest("GET", "/test?N=N", nil), "N", reflect.Bool, nil}, getBytes(false)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsByteArray(tt.args.r, tt.args.name, tt.args.kind, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryParamAsByteArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsFloat32(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&i=100&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{"float", args{req, "f", 100.0}, 123.99},
		{"int", args{req, "i", 10.0}, 100.0},
		{"text", args{req, "txt", 100.0}, 100.0},
		{"nil request", args{nil, "f", 123.99}, 123.99},
		{"empty name", args{req, "", 100.0}, 100.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsFloat32(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsFloat64(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&i=100&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"float", args{req, "f", 100.0}, 123.99},
		{"int", args{req, "i", 10.0}, 100.0},
		{"text", args{req, "txt", 100.0}, 100.0},
		{"nil request", args{nil, "f", 123.99}, 123.99},
		{"empty name", args{req, "", 100.0}, 100.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsFloat64(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsInt32(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", -10}, -111},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsInt32(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsInt64(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", -10}, -111},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsInt64(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsString(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?s=String&f=123.99&i=111&b=true", nil)

	type args struct {
		r    *http.Request
		name string
		def  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"String", args{req, "s", "fail"}, "String"},
		{"Float", args{req, "f", "fail"}, "123.99"},
		{"Int", args{req, "i", "fail"}, "111"},
		{"Bool", args{req, "b", "fail"}, "true"},
		{"nil request", args{nil, "test", "fail"}, "fail"},
		{"empty name", args{req, "", "fail"}, "fail"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsString(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsTime(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?a=2020-07-15T09:43:26Z&b=2020-07-15T09:43:26.578Z&c=20200715T094326Z&d=20200715T094326.578Z", nil)
	type args struct {
		r    *http.Request
		name string
		def  time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
		{"2006-01-02T15:04:05Z", args{req, "a", time.Unix(0, 0)}, time.Date(2020, 7, 15, 9, 43, 26, 0, time.UTC)},
		{"2006-01-02T15:04:05.999Z", args{req, "b", time.Unix(0, 0)}, time.Date(2020, 7, 15, 9, 43, 26, 578000000, time.UTC)},
		{"20060102T150405Z", args{req, "c", time.Unix(0, 0)}, time.Date(2020, 7, 15, 9, 43, 26, 0, time.UTC)},
		{"20060102T150405.999Z", args{req, "d", time.Unix(0, 0)}, time.Date(2020, 7, 15, 9, 43, 26, 578000000, time.UTC)},
		{"nil request", args{nil, "a", time.Unix(1000, 0)}, time.Unix(1000, 0)},
		{"empty name", args{req, "", time.Unix(1000, 0)}, time.Unix(1000, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsTime(tt.args.r, tt.args.name, tt.args.def); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryParamAsTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamsAsInt(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", -10}, -111},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamsAsInt(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamsAsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsInt8(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "123",
		"negative": "-123",
		"float":    "123.45",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			123,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			-123,
		},
		{
			"Test float",
			args{
				req,
				"float",
				123,
			},
			123,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsInt8(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsInt16(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "1234",
		"negative": "-1234",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			1234,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			-1234,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsInt16(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsUint(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "1234",
		"negative": "-1234",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			1234,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			0,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsUint(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsUInt8(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "123",
		"negative": "-123",
		"float":    "123.45",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			123,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			0,
		},
		{
			"Test float",
			args{
				req,
				"float",
				123,
			},
			123,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsUInt8(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsUint16(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "1234",
		"negative": "-1234",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			1234,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			0,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsUint16(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsUint32(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "1234",
		"negative": "-1234",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			1234,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			0,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsUint32(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamAsUint64(t *testing.T) {
	req := httptest.NewRequest("GET", "/test", nil)

	req = mux.SetURLVars(req, map[string]string{
		"positive": "1234",
		"negative": "-1234",
		"float":    "1234.56",
		"text":     "Test",
	})

	type args struct {
		r    *http.Request
		name string
		def  uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			"Test positive",
			args{
				req,
				"positive",
				0,
			},
			1234,
		},
		{
			"Test negative",
			args{
				req,
				"negative",
				0,
			},
			0,
		},
		{
			"Test float",
			args{
				req,
				"float",
				1234,
			},
			1234,
		},
		{
			"Test nil request",
			args{
				nil,
				"positive",
				100,
			},
			100,
		},
		{
			"Test empty name",
			args{
				req,
				"",
				100,
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamAsUint64(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("ParamAsUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsInt8(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", -10}, -111},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsInt8(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsInt16(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", -10}, -111},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsInt16(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamsAsUint(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", 10}, 10},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamsAsUint(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamsAsUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsUint8(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", 10}, 10},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsUint8(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsUint16(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", 10}, 10},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsUint16(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsUint32(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", 10}, 10},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsUint32(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryParamAsUint64(t *testing.T) {
	req := httptest.NewRequest("GET", "/test?f=123.99&p=100&n=-111&txt=Test", nil)

	type args struct {
		r    *http.Request
		name string
		def  uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{"positive", args{req, "p", 10}, 100},
		{"negative", args{req, "n", 10}, 10},
		{"float", args{req, "f", 100}, 100},
		{"text", args{req, "txt", 100}, 100},
		{"nil request", args{nil, "", 0}, 0},
		{"empty name", args{req, "", 100}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParamAsUint64(tt.args.r, tt.args.name, tt.args.def); got != tt.want {
				t.Errorf("QueryParamAsUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
