package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// TaskState represents the state of task, moving through Created, Running then Finished or Errorred
type TaskState int

const (
	// Created represents the task has been created but not started yet
	Created TaskState = iota
	//Running represents the task has started
	Running
	// Finished represents the task is complete
	Finished
	// Errorred represents the task has encountered a problem and is no longer running
	Errorred
)

func (s TaskState) String() string {
	return toString[s]
}

var toString = map[TaskState]string{
	Created:  "Created",
	Running:  "Running",
	Finished: "Finished",
	Errorred: "Errorred",
}

var toID = map[string]TaskState{
	"Created":  Created,
	"Running":  Running,
	"Finished": Finished,
	"Errorred": Errorred,
}

// MarshalJSON marshals the enum as a quoted json string
func (s TaskState) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(toString[s])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *TaskState) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = toID[j]
	return nil
}

type Task struct {
	Name  string    `json:"name"`
	State TaskState `json:"state"`
	// "Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON object."
	// byte is the same as uint8
	// 	UrlIndices []uint8     `json:"urlIndices"` // https://stackoverflow.com/questions/14177862/how-to-marshal-a-byte-uint8-array-as-json-array-in-go
	UrlIndices []int `json:"urlIndices"`
}

func main() {
	task1 := Task{Name: "go to church", State: Created, UrlIndices: []int{0, 1}}
	task2 := Task{Name: "go to beach", State: Running, UrlIndices: []int{1, 2, 3}}
	tasks := []Task{task1, task2}

	data, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)
	err = os.WriteFile("tasks.json", data, 0600)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.ReadFile("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	loadedTasks := make([]Task, 0)
	err = json.Unmarshal(file, &loadedTasks)
	if err != nil {
		log.Fatal(err)
	}

	for _, task := range loadedTasks {
		fmt.Printf("%v\n", task)
		fmt.Printf("%+v\n", task)
	}

	if fmt.Sprintf("%+v", loadedTasks[0]) == fmt.Sprintf("%+v", task1) {
		fmt.Println("First task loaded successfully")
	}

	if fmt.Sprintf("%+v", loadedTasks[1]) == fmt.Sprintf("%+v", task2) {
		fmt.Println("Second task loaded successfully")
	}
}
