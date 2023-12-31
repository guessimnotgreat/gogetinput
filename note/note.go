package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func (n *Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n\n%v\n", n.Title, n.Content)
}

func (n *Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func New(Title, Content string) (Note, error) {
	if Title == "" || Content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     Title,
		Content:   Content,
		CreatedAt: time.Now(),
	}, nil
}
