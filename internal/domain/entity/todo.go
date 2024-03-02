package entity

import "time"

type Todo struct {
	ID        int
	Title     string
	DoneFlag  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ListTodos []Todo

func NewRegisterTodo(
	title string,
) *Todo {
	return &Todo{
		Title: title,
	}
}

func NewChangeTodoDoneFlag(
	id int,
	doneFlag bool,
) *Todo {
	return &Todo{
		ID:    id,
		DoneFlag: doneFlag,
	}
}
