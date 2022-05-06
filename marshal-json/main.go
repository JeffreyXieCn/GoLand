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
}

func main() {
	task1 := Task{Name: "go to church", State: Created}
	task2 := Task{Name: "go to beach", State: Running}
	tasks := []Task{task1, task2}

	data, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("JSON masshashing failed: %s", err)
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
		fmt.Printf("%+v\n", task)
		fmt.Printf("%v\n", task)
	}

	if loadedTasks[0] == task1 {
		fmt.Println("First task loaded successfully")
	}

	if loadedTasks[1] == task2 {
		fmt.Println("Second task loaded successfully")
	}
}
