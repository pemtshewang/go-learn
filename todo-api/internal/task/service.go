package task

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateTask(title string) Task {
	return s.repo.Create(title)
}

func (s *Service) DeleteTask(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetTasks() []Task {
	return s.repo.GetAll()
}
