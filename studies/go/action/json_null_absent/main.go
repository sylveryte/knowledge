package main
// full copy pasta from https://github.com/golang/go/issues/64515#issuecomment-1842973794

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type jsonValueStruct[T any] struct {
	Value T
}

type JSONValue[T any] map[bool]jsonValueStruct[T] // true is for value, false is for null, empty is for undefined.

func (t JSONValue[T]) MarshalJSON() ([]byte, error) {
	if t.IsNull() {
		return []byte("null"), nil
	} else if !t.IsDefined() {
		return []byte{}, fmt.Errorf("value is not set, forgot to set omitempty flag?")
	}
	return json.Marshal(t[true].Value)
}

func (t *JSONValue[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		t.SetNull()
		return nil
	}
	var v T
	err := json.Unmarshal(data, &v)
	t.SetValue(v)
	return err
}

func (t *JSONValue[T]) SetValue(v T) {
	*t = map[bool]jsonValueStruct[T]{true: {Value: v}}
}

func (t JSONValue[T]) GetValue() T {
	return t[true].Value
}

func (t *JSONValue[T]) SetNull() {
	*t = map[bool]jsonValueStruct[T]{false: {}}
}

func (t JSONValue[T]) IsNull() bool {
	_, found := t[false]
	return found
}

func (t *JSONValue[T]) Reset() {
	*t = map[bool]jsonValueStruct[T]{}
}

func (t JSONValue[T]) IsDefined() bool {
	return len(t) != 0
}

func main() {
	type struct5 struct {
		I0 JSONValue[int32]          `json:"int32_0,omitempty"`
		I1 JSONValue[int32]          `json:"int32_1,omitempty"`
		I2 JSONValue[int32]          `json:"int32_2,omitempty"`
		I3 JSONValue[int32]          `json:"int32_3,omitempty"`
		J0 JSONValue[string]         `json:"string_0,omitempty"`
		J1 JSONValue[string]         `json:"string_1,omitempty"`
		J2 JSONValue[string]         `json:"string_2,omitempty"`
		J3 JSONValue[string]         `json:"string_3,omitempty"`
		K0 JSONValue[map[string]any] `json:"map_0,omitempty"`
		K1 JSONValue[map[string]any] `json:"map_1,omitempty"`
		K2 JSONValue[map[string]any] `json:"map_2,omitempty"`
		K3 JSONValue[map[string]any] `json:"map_3,omitempty"`
		L0 JSONValue[[]int]          `json:"intArray_0,omitempty"`
		L1 JSONValue[[]int]          `json:"intAttay_1,omitempty"`
		L2 JSONValue[[]int]          `json:"intAttay_2,omitempty"`
		L3 JSONValue[[]int]          `json:"intAttay_3,omitempty"`
	}
	var v [2]struct5
	v[0].I0.SetValue(100)
	v[0].I1.SetNull()
	v[0].I2.Reset()
	v[0].J0.SetValue(`abc`)
	v[0].J1.SetNull()
	v[0].J2.Reset()
	v[0].K0.SetValue(map[string]any{`abc`: 100})
	v[0].K1.SetNull()
	v[0].K2.Reset()
	v[0].L0.SetValue([]int{100, 200})
	v[0].L1.SetNull()
	v[0].L2.Reset()
	if out, err := json.Marshal(v[0]); err == nil && json.Unmarshal(out, &v[1]) == nil {
		// {"int32_0":100,"int32_1":null,"string_0":"abc","string_1":null,"map_0":{"abc":100},"map_1":null,"intArray_0":[100,200],"intAttay_1":null}
		fmt.Println(string(out))
		for i := 0; i < 2; i++ {
			// int32: true false 100 true true 0 false false 0 false false 0
			fmt.Println(`int32:`, v[i].I0.IsDefined(), v[i].I0.IsNull(), v[i].I0.GetValue(), v[i].I1.IsDefined(), v[i].I1.IsNull(), v[i].I1.GetValue(), v[i].I2.IsDefined(), v[i].I2.IsNull(), v[i].I2.GetValue(), v[i].I3.IsDefined(), v[i].I3.IsNull(), v[i].I3.GetValue())
		}
		for i := 0; i < 2; i++ {
			// string: true false abc true true  false false  false false
			fmt.Println(`string:`, v[i].J0.IsDefined(), v[i].J0.IsNull(), v[i].J0.GetValue(), v[i].J1.IsDefined(), v[i].J1.IsNull(), v[i].J1.GetValue(), v[i].J2.IsDefined(), v[i].J2.IsNull(), v[i].J2.GetValue(), v[i].J3.IsDefined(), v[i].J3.IsNull(), v[i].J3.GetValue())
		}
		for i := 0; i < 2; i++ {
			// map[string]any: true false map[abc:100] true true map[] false false map[] false false map[]
			fmt.Println(`map[string]any:`, v[i].K0.IsDefined(), v[i].K0.IsNull(), v[i].K0.GetValue(), v[i].K1.IsDefined(), v[i].K1.IsNull(), v[i].K1.GetValue(), v[i].K2.IsDefined(), v[i].K2.IsNull(), v[i].K2.GetValue(), v[i].K3.IsDefined(), v[i].K3.IsNull(), v[i].K3.GetValue())
		}
		for i := 0; i < 2; i++ {
			// []int: true false [100 200] true true [] false false [] false false []
			fmt.Println(`[]int:`, v[i].L0.IsDefined(), v[i].L0.IsNull(), v[i].L0.GetValue(), v[i].L1.IsDefined(), v[i].L1.IsNull(), v[i].L1.GetValue(), v[i].L2.IsDefined(), v[i].L2.IsNull(), v[i].L2.GetValue(), v[i].L3.IsDefined(), v[i].L3.IsNull(), v[i].L3.GetValue())
		}
	}
}
