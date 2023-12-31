package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (t *Todo) Display() {
	fmt.Printf(t.Text)
}

func (t *Todo) Save() error {
	filename := "todo.json"

	json, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
