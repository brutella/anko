package core

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/brutella/anko/internal/testlib"
)

func TestToX(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testlib.Test{
		{Script: `toBool(-2)`, RunOutput: false},
		{Script: `toBool(-1.5)`, RunOutput: false},
		{Script: `toBool(-1)`, RunOutput: false},
		{Script: `toBool(-0.4)`, RunOutput: false},
		{Script: `toBool(0)`, RunOutput: false},
		{Script: `toBool(0.4)`, RunOutput: true},
		{Script: `toBool(1)`, RunOutput: true},
		{Script: `toBool(1.5)`, RunOutput: true},
		{Script: `toBool(2)`, RunOutput: true},
		{Script: `toBool(true)`, RunOutput: true},
		{Script: `toBool(false)`, RunOutput: false},
		{Script: `toBool("true")`, RunOutput: true},
		{Script: `toBool("false")`, RunOutput: false},
		{Script: `toBool("yes")`, RunOutput: true},
		{Script: `toBool("ye")`, RunOutput: false},
		{Script: `toBool("y")`, RunOutput: true},
		{Script: `toBool("false")`, RunOutput: false},
		{Script: `toBool("f")`, RunOutput: false},
		{Script: `toBool("")`, RunOutput: false},
		{Script: `toBool(nil)`, RunOutput: false},
		{Script: `toBool({})`, RunOutput: false},
		{Script: `toBool([])`, RunOutput: false},
		{Script: `toBool([true])`, RunOutput: false},
		{Script: `toBool({"true": "true"})`, RunOutput: false},
		{Script: `toString(nil)`, RunOutput: "<nil>"},
		{Script: `toString("")`, RunOutput: ""},
		{Script: `toString(1)`, RunOutput: "1"},
		{Script: `toString(1.2)`, RunOutput: "1.2"},
		{Script: `toString(1/3)`, RunOutput: "0.3333333333333333"},
		{Script: `toString(false)`, RunOutput: "false"},
		{Script: `toString(true)`, RunOutput: "true"},
		{Script: `toString({})`, RunOutput: "map[]"},
		{Script: `toString({"foo": "bar"})`, RunOutput: "map[foo:bar]"},
		{Script: `toString([true,nil])`, RunOutput: "[true <nil>]"},
		{Script: `toString(toByteSlice("foo"))`, RunOutput: "foo"},
		{Script: `toInt(nil)`, RunOutput: int64(0)},
		{Script: `toInt(-2)`, RunOutput: int64(-2)},
		{Script: `toInt(-1.4)`, RunOutput: int64(-1)},
		{Script: `toInt(-1)`, RunOutput: int64(-1)},
		{Script: `toInt(0)`, RunOutput: int64(0)},
		{Script: `toInt(1)`, RunOutput: int64(1)},
		{Script: `toInt(1.4)`, RunOutput: int64(1)},
		{Script: `toInt(1.5)`, RunOutput: int64(1)},
		{Script: `toInt(1.9)`, RunOutput: int64(1)},
		{Script: `toInt(2)`, RunOutput: int64(2)},
		{Script: `toInt(2.1)`, RunOutput: int64(2)},
		{Script: `toInt("2")`, RunOutput: int64(2)},
		{Script: `toInt("2.1")`, RunOutput: int64(2)},
		{Script: `toInt(true)`, RunOutput: int64(1)},
		{Script: `toInt(false)`, RunOutput: int64(0)},
		{Script: `toInt({})`, RunOutput: int64(0)},
		{Script: `toInt([])`, RunOutput: int64(0)},
		{Script: `toFloat(nil)`, RunOutput: float64(0.0)},
		{Script: `toFloat(-2)`, RunOutput: float64(-2.0)},
		{Script: `toFloat(-1.4)`, RunOutput: float64(-1.4)},
		{Script: `toFloat(-1)`, RunOutput: float64(-1.0)},
		{Script: `toFloat(0)`, RunOutput: float64(0.0)},
		{Script: `toFloat(1)`, RunOutput: float64(1.0)},
		{Script: `toFloat(1.4)`, RunOutput: float64(1.4)},
		{Script: `toFloat(1.5)`, RunOutput: float64(1.5)},
		{Script: `toFloat(1.9)`, RunOutput: float64(1.9)},
		{Script: `toFloat(2)`, RunOutput: float64(2.0)},
		{Script: `toFloat(2.1)`, RunOutput: float64(2.1)},
		{Script: `toFloat("2")`, RunOutput: float64(2.0)},
		{Script: `toFloat("2.1")`, RunOutput: float64(2.1)},
		{Script: `toFloat(true)`, RunOutput: float64(1.0)},
		{Script: `toFloat(false)`, RunOutput: float64(0.0)},
		{Script: `toFloat({})`, RunOutput: float64(0.0)},
		{Script: `toFloat([])`, RunOutput: float64(0.0)},
		{Script: `toChar(0x1F431)`, RunOutput: "🐱"},
		{Script: `toChar(0)`, RunOutput: "\x00"},
		{Script: `toRune("")`, RunOutput: rune(0)},
		{Script: `toRune("🐱")`, RunOutput: rune(0x1F431)},
		{Script: `toBoolSlice(nil)`, RunOutput: []bool{}},
		{Script: `toBoolSlice(1)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type int64")},
		{Script: `toBoolSlice(1.2)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type float64")},
		{Script: `toBoolSlice(false)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type bool")},
		{Script: `toBoolSlice({})`, RunError: fmt.Errorf("function wants argument type []interface {} but received type map[interface {}]interface {}")},
		{Script: `toBoolSlice([])`, RunOutput: []bool{}},
		{Script: `toBoolSlice([nil])`, RunOutput: []bool{false}},
		{Script: `toBoolSlice([1])`, RunOutput: []bool{false}},
		{Script: `toBoolSlice([1.1])`, RunOutput: []bool{false}},
		{Script: `toBoolSlice([true])`, RunOutput: []bool{true}},
		{Script: `toBoolSlice([[]])`, RunOutput: []bool{false}},
		{Script: `toBoolSlice([{}])`, RunOutput: []bool{false}},
		{Script: `toIntSlice(nil)`, RunOutput: []int64{}},
		{Script: `toIntSlice(1)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type int64")},
		{Script: `toIntSlice(1.2)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type float64")},
		{Script: `toIntSlice(false)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type bool")},
		{Script: `toIntSlice({})`, RunError: fmt.Errorf("function wants argument type []interface {} but received type map[interface {}]interface {}")},
		{Script: `toIntSlice([])`, RunOutput: []int64{}},
		{Script: `toIntSlice([nil])`, RunOutput: []int64{0}},
		{Script: `toIntSlice([1])`, RunOutput: []int64{1}},
		{Script: `toIntSlice([1.1])`, RunOutput: []int64{1}},
		{Script: `toIntSlice([true])`, RunOutput: []int64{0}},
		{Script: `toIntSlice([[]])`, RunOutput: []int64{0}},
		{Script: `toIntSlice([{}])`, RunOutput: []int64{0}},
		{Script: `toFloatSlice(nil)`, RunOutput: []float64{}},
		{Script: `toFloatSlice(1)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type int64")},
		{Script: `toFloatSlice(1.2)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type float64")},
		{Script: `toFloatSlice(false)`, RunError: fmt.Errorf("function wants argument type []interface {} but received type bool")},
		{Script: `toFloatSlice({})`, RunError: fmt.Errorf("function wants argument type []interface {} but received type map[interface {}]interface {}")},
		{Script: `toFloatSlice([])`, RunOutput: []float64{}},
		{Script: `toFloatSlice([nil])`, RunOutput: []float64{0.0}},
		{Script: `toFloatSlice([1])`, RunOutput: []float64{1.0}},
		{Script: `toFloatSlice([1.1])`, RunOutput: []float64{1.1}},
		{Script: `toFloatSlice([true])`, RunOutput: []float64{0.0}},
		{Script: `toFloatSlice([[]])`, RunOutput: []float64{0.0}},
		{Script: `toFloatSlice([{}])`, RunOutput: []float64{0.0}},
		{Script: `toByteSlice(nil)`, RunOutput: []byte{}},
		{Script: `toByteSlice([])`, RunError: fmt.Errorf("function wants argument type string but received type []interface {}")},
		{Script: `toByteSlice(1)`, RunOutput: []byte{0x01}}, // FIXME?
		{Script: `toByteSlice(1.1)`, RunError: fmt.Errorf("function wants argument type string but received type float64")},
		{Script: `toByteSlice(true)`, RunError: fmt.Errorf("function wants argument type string but received type bool")},
		{Script: `toByteSlice("foo")`, RunOutput: []byte{'f', 'o', 'o'}},
		{Script: `toByteSlice("世界")`, RunOutput: []byte{0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}},
		{Script: `toRuneSlice(nil)`, RunOutput: []rune{}},
		{Script: `toRuneSlice([])`, RunError: fmt.Errorf("function wants argument type string but received type []interface {}")},
		{Script: `toRuneSlice(1)`, RunOutput: []rune{0x01}}, // FIXME?
		{Script: `toRuneSlice(1.1)`, RunError: fmt.Errorf("function wants argument type string but received type float64")},
		{Script: `toRuneSlice(true)`, RunError: fmt.Errorf("function wants argument type string but received type bool")},
		{Script: `toRuneSlice("foo")`, RunOutput: []rune{'f', 'o', 'o'}},
		{Script: `toRuneSlice("世界")`, RunOutput: []rune{'世', '界'}},
		{Script: `toStringSlice([true,false,1])`, RunOutput: []string{"", "", "\x01"}}, // FIXME?
		{Script: `toDuration(nil)`, RunOutput: time.Duration(0)},
		{Script: `toDuration(0)`, RunOutput: time.Duration(0)},
		{Script: `toDuration(true)`, RunError: fmt.Errorf("function wants argument type int64 but received type bool")},
		{Script: `toDuration([])`, RunError: fmt.Errorf("function wants argument type int64 but received type []interface {}")},
		{Script: `toDuration({})`, RunError: fmt.Errorf("function wants argument type int64 but received type map[interface {}]interface {}")},
		{Script: `toDuration("")`, RunError: fmt.Errorf("function wants argument type int64 but received type string")},
		{Script: `toDuration("1s")`, RunError: fmt.Errorf("function wants argument type int64 but received type string")}, // TODO
		{Script: `toDuration(a)`, Input: map[string]interface{}{"a": int64(time.Duration(123 * time.Minute))}, RunOutput: time.Duration(123 * time.Minute)},
		{Script: `toDuration(a)`, Input: map[string]interface{}{"a": float64(time.Duration(123 * time.Minute))}, RunOutput: time.Duration(123 * time.Minute)},
		{Script: `toDuration(a)`, Input: map[string]interface{}{"a": time.Duration(123 * time.Minute)}, RunOutput: time.Duration(123 * time.Minute)},
	}
	testlib.Run(t, tests, &testlib.Options{EnvSetupFunc: &testCoreEnvSetupFunc})
}
