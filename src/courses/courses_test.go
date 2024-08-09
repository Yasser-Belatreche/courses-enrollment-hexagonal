package courses

import (
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/usecases"
	"github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/infra"
	"testing"
)

func TestCoursesFacade(t *testing.T) {
	facade := &Facade{
		CourseRepository:     infra.NewFakeCourseRepository(),
		EnrollmentRepository: infra.NewFakeEnrollmentRepository(),
		EventsPublisher:      infra.NewFakeEventsPublisher(),
	}

	runTestsOn(facade, t)
}

func runTestsOn(facade *Facade, t *testing.T) {
	t.Run("Course Scheduling", func(t *testing.T) {
		t.Run("Scheduled courses should not have an empty name", func(t *testing.T) {
			_, err := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "",
				MaxSize: 3,
				MinSize: 1,
			})

			if err == nil {
				t.Error("Expected an error, got nil")
			}
		})

		t.Run("Scheduled courses should have a max and min size > 0", func(t *testing.T) {
			_, err := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 0,
				MinSize: 3,
			})

			if err == nil {
				t.Error("Expected an error, got nil")
			}

			_, err = facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 0,
			})

			if err == nil {
				t.Error("Expected an error, got nil")
			}
		})

		t.Run("Scheduled courses max size should be bigger than min size", func(t *testing.T) {
			_, err := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 5,
			})

			if err == nil {
				t.Error("Expected an error, got nil")
			}
		})

		t.Run("Scheduled courses should have unique ids for each one", func(t *testing.T) {
			response, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 5,
				MinSize: 3,
			})

			response2, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 5,
				MinSize: 3,
			})

			if response.Id == response2.Id {
				t.Error("Expected different ids, got the same")
			}
		})

		t.Run("Scheduled course at first should have no enrollments and have status not-viable", func(t *testing.T) {
			res, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 5,
				MinSize: 3,
			})

			course, _ := facade.GetCourse(&usecases.GetCourseQuery{Id: res.Id})

			if len(course.Enrollments) != 0 {
				t.Errorf("Expected 0 enrollments, got %v", len(course.Enrollments))
			}

			if course.Status != "not-viable" {
				t.Errorf("Expected status not-viable, got %v", course.Status)
			}
		})
	})

	t.Run("Course Enrollment", func(t *testing.T) {

	})

	t.Run("Course Cancellation", func(t *testing.T) {})
}
