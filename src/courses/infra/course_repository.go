package infra

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

func NewFakeCourseRepository() domain.CourseRepository {
	return &inMemoryCourseRepository{
		courses: make(map[string]*domain.Course),
	}
}

type inMemoryCourseRepository struct {
	courses map[string]*domain.Course
}

func (i *inMemoryCourseRepository) Create(course *domain.Course) error {
	i.courses[course.State().Id] = course

	return nil
}

func (i *inMemoryCourseRepository) Update(course *domain.Course) error {
	i.courses[course.State().Id] = course

	return nil
}

func (i *inMemoryCourseRepository) FindById(id string) (*domain.Course, error) {
	return i.courses[id], nil
}
