package domain

type EnrollmentRepository interface {
	Create(enrollment *Enrollment) error

	FindById(id string) (*Enrollment, error)

	Delete(id string) error

	FindByCourseId(id string) ([]*Enrollment, error)
}
