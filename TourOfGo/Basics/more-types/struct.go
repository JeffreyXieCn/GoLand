package main

import "fmt"

type Vertex struct {
	X, Y int
}

type Vertex3D struct {
	X, Y int
	Z    *int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Person struct {
	Name   string
	Age    int
	Scores map[string]int
}

func (p Person) Clone() Person {
	c := Person{}
	c.Name = p.Name
	c.Age = p.Age
	c.Scores = map[string]int{}
	for k, v := range p.Scores {
		c.Scores[k] = v
	}

	return c
}

func main() {
	fmt.Println(v1, v2, v3, p)
	changeVertex(v1)
	fmt.Printf("\nAfter calling changeVertex: %+v", v1)

	v1c := v1
	v1c.X++
	v1c.Y++
	fmt.Printf("\nAfter changing v1c: %+v", v1c)
	fmt.Printf("\nv1: %+v", v1)

	changeVertexPtr(p)
	fmt.Printf("\nAfter calling changeVertexPtr: %+v", p)

	z := 0
	v3d1 := Vertex3D{X: 0, Y: 0, Z: &z}
	fmt.Printf("\nv3d1: %+v", v3d1)

	v3d1c := v3d1
	v3d1c.X++
	v3d1c.Y++
	*v3d1c.Z++
	fmt.Printf("\nv3d1c: %+v, v3d1c.Z: %d\n", v3d1c, *v3d1c.Z)
	fmt.Printf("\nv3d1: %+v, v3d1.Z: %d\n", v3d1, *v3d1.Z)

	// deep copy (clone) struct
	alice1 := Person{"Alice", 30, map[string]int{"math": 90, "english": 100}}
	alice2 := alice1
	alice3 := alice1.Clone()
	//fmt.Println(alice1 == alice2)   // => struct containing map[string]int cannot be compared
	fmt.Println(&alice1 == &alice2) // => false, they have different addresses
	fmt.Println(&alice1 == &alice3) // => false, they have different addresses

	alice1.Age += 10
	delete(alice1.Scores, "math")
	fmt.Printf("alice1: %v\n", alice1)
	fmt.Printf("alice2: %v\n", alice2)
	fmt.Printf("alice3: %v\n", alice3)
}

func changeVertex(v Vertex) {
	v.X++
	v.Y++
}

func changeVertexPtr(v *Vertex) {
	v.X++
	v.Y++
}
