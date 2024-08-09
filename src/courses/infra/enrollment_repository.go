package infra

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

func NewFakeEnrollmentRepository() domain.EnrollmentRepository {
	return &inMemoryEnrollmentRepository{
		enrollments: make(map[string]*domain.Enrollment),
	}
}

type inMemoryEnrollmentRepository struct {
	enrollments map[string]*domain.Enrollment
}

func (i *inMemoryEnrollmentRepository) Create(enrollment *domain.Enrollment) error {
	i.enrollments[enrollment.Id] = enrollment

	return nil
}

func (i *inMemoryEnrollmentRepository) FindById(id string) (*domain.Enrollment, error) {
	return i.enrollments[id], nil
}

func (i *inMemoryEnrollmentRepository) Delete(id string) error {
	delete(i.enrollments, id)

	return nil
}

func (i *inMemoryEnrollmentRepository) FindByCourseId(id string) ([]*domain.Enrollment, error) {
	list := make([]*domain.Enrollment, 0)

	for _, enrollment := range i.enrollments {
		if enrollment.CourseId == id {
			list = append(list, enrollment)
		}
	}

	return list, nil
}
