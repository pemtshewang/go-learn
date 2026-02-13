package task

import "errors"

type Repository struct {
	tasks  []Task
	nextID int
}

// a database
// return by pointer so it persists
func NewRepository() *Repository {
	return &Repository{
		tasks:  []Task{},
		nextID: 1,
	}
}

func (r *Repository) Create(title string) Task {
	task := Task{
		ID:    r.nextID,
		Title: title,
		Done:  false,
	}

	r.tasks = append(r.tasks, task)
	// increment the id
	r.nextID++
	return task
}

func (r *Repository) Delete(id int) error {
	for i := range r.tasks {
		if r.tasks[i].ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("ID not found")
}

func (r *Repository) GetAll() []Task {
	return r.tasks
}

func (r *Repository) GetByID(id int) (Task, error) {
	for _, t := range r.tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return Task{}, errors.New("task not found")
}
