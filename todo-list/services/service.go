package services

import (
	"errors"
	"fmt"
	"mirzaadr/todo-cli/models"
	"time"
)

type Todos []models.Todo

func (todos *Todos) Add(desc string) {
	todo := models.Todo{
		Description: desc,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index > len(*todos) {
		err := errors.New("index out of reach")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) Complete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	// to add completion time
	// isComplete := t[index].IsComplete
	// if !isComplete {
	// 	completionTime := time.Now()
	// 	t[index].CompletedAt = &completionTime
	// }

	t[index].IsComplete = true

	return nil
}

func (todos *Todos) Print() error {
	return nil
}
