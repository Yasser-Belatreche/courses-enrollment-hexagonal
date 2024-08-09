package domain

import (
	"errors"
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/lib"
)

type Course struct {
	id               string
	name             string
	minSize          int
	maxSize          int
	totalEnrollments int
	status           CourseStatus
	events           []CourseEvent[interface{}]
}

type CourseState struct {
	Id               string
	Name             string
	MinSize          int
	MaxSize          int
	TotalEnrollments int
	Status           string
}

func ScheduleCourse(name string, minSize int, maxSize int) (*Course, error) {
	if minSize > maxSize {
		return nil, errors.New("minSize cannot be greater than maxSize")
	}

	if minSize <= 0 {
		return nil, errors.New("minSize must be greater than 0")
	}

	if maxSize <= 0 {
		return nil, errors.New("maxSize must be greater than 0")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	course := &Course{
		id:               lib.GenerateUlid(),
		name:             name,
		minSize:          minSize,
		maxSize:          maxSize,
		totalEnrollments: 0,
		status:           CourseNotViable,
		events:           make([]CourseEvent[interface{}], 0),
	}

	course.addEvent(NewCourseScheduledEvent(course))

	return course, nil
}

func CourseFromState(state *CourseState) *Course {
	return &Course{
		id:               state.Id,
		name:             state.Name,
		minSize:          state.MinSize,
		maxSize:          state.MaxSize,
		totalEnrollments: state.TotalEnrollments,
		status:           CourseStatus(state.Status),
		events:           make([]CourseEvent[interface{}], 0),
	}
}

func (c *Course) State() *CourseState {
	return &CourseState{
		Id:               c.id,
		Name:             c.name,
		MinSize:          c.minSize,
		MaxSize:          c.maxSize,
		TotalEnrollments: c.totalEnrollments,
		Status:           string(c.status),
	}
}

func (c *Course) addEvent(event *CourseEvent[interface{}]) {
	c.events = append(c.events, *event)
}

func (c *Course) PullEvents() []CourseEvent[interface{}] {
	allEvents := c.events

	c.events = make([]CourseEvent[interface{}], 0)

	return allEvents
}
