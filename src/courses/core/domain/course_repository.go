package domain

type CourseRepository interface {
	Create(course *Course) error

	Update(course *Course) error

	FindById(id string) (*Course, error)
}
