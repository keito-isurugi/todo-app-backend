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

func NewChangeTodo(
	id int,
	title string,
) *Todo {
	return &Todo{
		ID:    id,
		Title: title,
	}
}
