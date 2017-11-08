package mock

import "github.com/eriktate/simple/todo"

type MockRepo struct {
	CreateFn     func(todo todo.Todo) (string, error)
	CreateCalled int

	Get       func(id string) (todo.Todo, error)
	GetCalled int

	GetAll       func() ([]todo.Todo, error)
	GetAllCalled int

	DeleteFn     func(id string) (todo.Todo, error)
	DeleteCalled int

	SetTitleFn     func(id, title string) error
	SetTitleCalled int

	SetDescriptionFn func(id, description string) error
	SetTitleCalled   int

	CompleteFn     func(id string) error
	CompleteCalled int

	IncompleteFn     func(id string) error
	IncompleteCalled int

	PassThru bool
}

func (m MockRepo) Create(todo todo.Todo) (string, error) {
	m.CreateCalled++

	if PassThru {
		return m.CreateFn(todo)
	}

	return "", nil
}

func (m MockRepo) Get(id string) (todo.Todo, error) {
	m.GetCalled++

	if PassThru {
		return m.GetFn(todo)
	}

	return todo.Todo{}, nil
}

func (m MockRepo) GetAll() ([]todo.Todo, error) {
	m.GetAllCalled++

	if PassThru {
		return m.GetAllFn(todo)
	}

	return []todo.Todo{}, nil
}

func (m MockRepo) Delete(id string) (todo.Todo, error) {
	m.DeleteCalled++

	if PassThru {
		return m.DeleteFn(todo)
	}

	return todo.Todo{}, nil
}

func (m MockRepo) Delete(id string) (todo.Todo, error) {
	m.DeleteCalled++

	if PassThru {
		return m.DeleteFn(todo)
	}

	return todo.Todo{}, nil
}

func (m MockRepo) SetTitle(id, title string) error {
	m.SetTitleCalled++

	if PassThru {
		return m.SetTitleFn(todo)
	}

	return nil
}

func (m MockRepo) SetDescription(id, description string) error {
	m.SetDescriptionCalled++

	if PassThru {
		return m.SetDescriptionFn(todo)
	}

	return nil
}

func (m MockRepo) Complete(id string) error {
	m.CompleteCalled++

	if PassThru {
		return m.CompleteFn(todo)
	}

	return nil
}

func (m MockRepo) Incomplete(id string) error {
	m.IncompleteCalled++

	if PassThru {
		return m.IncompleteFn(todo)
	}

	return nil
}
