package roach

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"github.com/eriktate/simple/todo"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	// Because we need it...
	_ "github.com/lib/pq"
)

// A Repo is a cockroach implementation of a TodoRepo.
type Repo struct {
	db *sql.DB
}

// NewRepo returns a Repo ready to talk to CockroachDB using the connection info provided.
func NewRepo(host, user, password, dbName string, port uint) (Repo, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?application_name=cockroach&sslmode=disable", user, password, host, port, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return Repo{}, errors.Wrap(err, "Failed to establish connection to CockroachDB")
	}

	return Repo{
		db: db,
	}, nil
}

// NewRepoWithDB returns a new Repo using an existing database connection (primarily for unit tests)
func NewRepoWithDB(db *sql.DB) Repo {
	return Repo{
		db: db,
	}
}

// Create will insert a new Todo into cockroach.
func (r Repo) Create(todo todo.Todo) (string, error) {
	if todo.ID == "" {
		todo.ID = xid.New().String()
	}

	if _, err := r.db.Exec("insert into todo (id, title, description) values($1, $2, $3)", todo.ID, todo.Title, todo.Description); err != nil {
		return "", errors.Wrap(err, "Failed to Create Todo")
	}

	return todo.ID, nil
}

// Get will retrieve an existing Todo from cockroach by ID.
func (r Repo) Get(id string) (todo.Todo, error) {
	todo, err := scanRow(r.db.QueryRow("select * from todo where id = $1", id))
	if err != nil {
		return todo, errors.Wrap(err, "Failed to Get todo")
	}

	return todo, nil
}

// GetAll will retrieve all existing Todos from cockroach.
func (r Repo) GetAll() ([]todo.Todo, error) {
	var todos []todo.Todo

	res, err := r.db.Query("select * from todo")
	if err != nil {
		return todos, errors.Wrap(err, "Failed to select todos from database")
	}

	todos, err = scanRows(res)
	if err != nil {
		return todos, err
	}

	return todos, nil
}

// Delete will delete an existing Todo from cockroach by ID.
// It will also return the deleted Todo.
func (r Repo) Delete(id string) (todo.Todo, error) {
	todo, err := r.Get(id)
	if err != nil {
		return todo, errors.Wrap(err, "Failed to retrieve todo before deleting")
	}

	if _, err := r.db.Exec("delete from todo where id = $1", id); err != nil {
		return todo, errors.Wrap(err, "Failed to delete todo")
	}

	return todo, nil
}

// SetTitle will update the title on the Todo with the given ID.
func (r Repo) SetTitle(id, title string) error {
	if _, err := r.db.Exec("update todo set title = $1 where id = $2", title, id); err != nil {
		return errors.Wrap(err, "Failed to SetTitle")
	}

	return nil
}

// SetDescription will update the description on the Todo with the given ID.
func (r Repo) SetDescription(id, description string) error {
	if _, err := r.db.Exec("update todo set description = $1 where id = $2", description, id); err != nil {
		return errors.Wrap(err, "Failed to SetDescription")
	}

	return nil
}

// Complete marks a given Todo as completed.
func (r Repo) Complete(id string) error {
	if _, err := r.db.Exec("update todo set completed = true where id = $1", id); err != nil {
		return errors.Wrap(err, "Failed to Complete Todo")
	}

	return nil

}

// Incomplete marks a given Todo as incomplete.
func (r Repo) Incomplete(id string) error {
	if _, err := r.db.Exec("update todo set completed = false where id = $1", id); err != nil {
		return errors.Wrap(err, "Failed to Incomplete Todo")
	}

	return nil
}

// RunQueryFile reads from the file specified in path and executes it against the database
// pointer configured in Repo. This should be used for bootstrapping the database or
// running migrations in an automated way.
func (r Repo) RunQueryFile(path string) error {
	query, err := ioutil.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "Failed to read query file")
	}

	_, err = r.db.Exec(string(query))
	if err != nil {
		return errors.Wrap(err, "Failed to execute query file")
	}

	return nil
}

func scanRow(r *sql.Row) (todo.Todo, error) {
	var todo todo.Todo
	if err := r.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Complete, &todo.CompletedAt, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return todo, err
		}

		return todo, errors.Wrap(err, "Failed to scan row into struct")
	}

	return todo, nil
}

func scanRows(r *sql.Rows) ([]todo.Todo, error) {
	var t todo.Todo
	todos := make([]todo.Todo, 0)

	for r.Next() {
		if err := r.Scan(&t.ID, &t.Title, &t.Description, &t.Complete, &t.CompletedAt, &t.CreatedAt, &t.UpdatedAt); err != nil {
			if err == sql.ErrNoRows {
				return todos, err
			}

			return todos, errors.Wrap(err, "Failed to scan rows into slice")
		}

		todos = append(todos, t)
	}

	return todos, nil
}
