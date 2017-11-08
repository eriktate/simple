package todo

import "time"

// A Todo is something that needs to be done.
type Todo struct {
	ID          string
	Title       string
	Description string
	Complete    bool

	CompletedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// A TodoRepo is something that handles storing and retrieving Todos.
type TodoRepo interface {
	Create(todo Todo) (string, error)
	Get(id string) (Todo, error)
	GetAll() ([]Todo, error)
	Delete(id string) (Todo, error)

	SetTitle(id, title string) error
	SetDescription(id, description string) error
	Complete(id string) error
	Incomplete(id string) error
}
