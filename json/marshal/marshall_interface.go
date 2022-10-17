// https://play.golang.org/p/l_GmjX31XbM

package main

import (
	"encoding/json"
	"fmt"
)

// Main has a Field2 whose type we don't know at the time of unmarshaling. We
// will demonstrate a very simple case of custom unmarshaling when the interface
// type is a base type (int or string.) If you have a more complex type, then
// you may have to use more maps and additional logic to discover what concrete
// type is held by the interface
type Main struct {
	Field1 string    `json:"field1"`
	Field2 Interface `json:"field2"`
}

type Interface interface {
	Print()
}

type Impl1 int

func (i Impl1) Print() {
	fmt.Println(i)
}

type Impl2 string

func (i Impl2) Print() {
	fmt.Println(i)
}

func (m *Main) UnmarshalJSON(data []byte) (err error) {
	temp := make(map[string]interface{}, 2)
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return
	}

	m.Field1 = temp["field1"].(string)
	switch val := temp["field2"].(type) {
	case string:
		m.Field2 = Impl2(val)
	case float64:
		m.Field2 = Impl1(int(val))
	default:
		err = fmt.Errorf("type %T not supported", val)
	}

	return
}

func main() {
	src1 := Main{
		Field1: "field",
		Field2: Impl1(1),
	}
	test(src1)

	src2 := Main{
		Field1: "field",
		Field2: Impl2("My2ndField"),
	}
	test(src2)
}

func test(src Main) {
	fmt.Println("\n\nsrc:", src)
	//fmt.Println(src)

	ser, err := json.Marshal(&src)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(ser))

	var dst Main
	err = json.Unmarshal(ser, &dst)
	if err != nil {
		panic(err)
	}

	fmt.Println("dst:", dst)
}
