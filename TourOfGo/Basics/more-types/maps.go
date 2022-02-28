package main

import "fmt"

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func main() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println("=============Map Literals=============")
	var m2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	var m3 = map[string]Vertex2{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m3)

	bl := m3["Bell Labs"]
	fmt.Println(bl)

	bl.Lat += 1
	fmt.Println("Bell Labs's Lat changed", bl)
	fmt.Println("But the map is intact: ", m3)

	fmt.Println("=============Iterate Map Key Values=============")
	for key, val := range m3 {
		fmt.Printf("key=%v, value=%v\n", key, val)
	}

	fmt.Println("=============Iterate Map Keys=============")
	for key := range m3 {
		fmt.Printf("key=%v\n", key)
	}

	fmt.Println("=============Iterate Map Values=============")
	for _, val := range m3 {
		fmt.Printf("val=%v\n", val)
	}

	fmt.Println("=============Mutating Maps=============")
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	//If key is in m, ok is true. If not, ok is false.
	//
	//If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
