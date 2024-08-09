package infra

import (
	"errors"
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"
)

func NewFakeCourseRepository() domain.CourseRepository {
	return &inMemoryCourseRepository{
		courses: make(map[string]*domain.CourseState),
	}
}

type inMemoryCourseRepository struct {
	courses map[string]*domain.CourseState
}

func (i *inMemoryCourseRepository) Create(course *domain.Course) error {
	i.courses[course.State().Id] = course.State()

	return nil
}

func (i *inMemoryCourseRepository) Update(course *domain.Course) error {
	i.courses[course.State().Id] = course.State()

	return nil
}

func (i *inMemoryCourseRepository) FindById(id string) (*domain.Course, error) {
	course := i.courses[id]

	if course == nil {
		return nil, errors.New("course not found")
	}

	return domain.CourseFromState(course), nil
}
