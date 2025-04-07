package services

import (
	"fmt"
	"mirzaadr/todo-cli/models"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
)

type Todos []models.Todo

func (todos *Todos) Add(desc string) {
	t := *todos

	var id int
	lastIndex := len(t) - 1
	if len(t) > 0 {
		id = t[lastIndex].ID + 1
	} else {
		id = 1
	}
	todo := models.Todo{
		ID:          id,
		Description: desc,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) findIndexByID(ID int) (int, error) {
	var todoIndex int = -1
	for index, todo := range *todos {
		if todo.ID == ID {
			todoIndex = index
		}
	}
	if todoIndex == -1 {
		return todoIndex, fmt.Errorf("ID not found")
	}
	return todoIndex, nil
}

func (todos *Todos) Delete(ID int) error {
	t := *todos

	index, err := t.findIndexByID(ID)
	if err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) Complete(ID int) error {
	t := *todos

	index, err := t.findIndexByID(ID)
	if err != nil {
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

func (todos *Todos) Print(isAll bool) error {
	t := *todos
	// Create a new tabwriter.Writer instance.
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 4, ' ', 0)

	// Write some data to the Writer.
	if isAll {
		fmt.Fprintln(w, "ID\tTask\tCreated\tDone")
		for _, todo := range t {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", todo.ID, todo.Description, timediff.TimeDiff(todo.CreatedAt), todo.IsComplete)
		}
	} else {
		fmt.Fprintln(w, "ID\tTask\tCreatedAt")
		for _, todo := range t {
			fmt.Fprintf(w, "%v\t%v\t%v\n", todo.ID, todo.Description, timediff.TimeDiff(todo.CreatedAt))
		}
	}

	// Flush the Writer to ensure all data is written to the output.
	return w.Flush()
}
