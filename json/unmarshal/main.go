package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type TableSchema struct {
	Schema *Schema `json:"schema"`
}

type Schema struct {
	Fields []*Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
}

func main() {
	file, err := os.ReadFile("BQtableSchema.json")
	if err != nil {
		log.Fatal(err)
	}

	tableSchema := &TableSchema{}
	err = json.Unmarshal(file, tableSchema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n\n", *tableSchema)
	fmt.Printf("%+v\n\n", *tableSchema)
	fmt.Printf("%#v\n\n", *tableSchema)

	fmt.Printf("%v\n\n", tableSchema)
	fmt.Printf("%+v\n\n", tableSchema)
	fmt.Printf("%#v\n\n", tableSchema)

	fmt.Printf("%v\n\n", tableSchema.Schema.Fields)
	fmt.Printf("%+v\n\n", tableSchema.Schema.Fields)
	fmt.Printf("%#v\n\n", tableSchema.Schema.Fields)

	for i, f := range tableSchema.Schema.Fields {
		fmt.Printf("%d, %v\n", i, f.Name)
	}
}
