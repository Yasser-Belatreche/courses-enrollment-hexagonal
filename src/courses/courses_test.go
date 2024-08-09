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
		t.Run("Should not be able to enroll in a non existing course", func(t *testing.T) {
			_, err := facade.Enroll(&usecases.EnrollCommand{
				CourseId: "random",
				Student:  "Alice",
			})

			if err == nil {
				t.Error("Expected an error, got nil")
			}
		})

		t.Run("student name should not be empty", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 5,
				MinSize: 3,
			})

			_, err := facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "",
			})

			if err == nil {
				t.Error("Expected an error, got nil")
			}
		})

		t.Run("Should be able to enroll to existing courses", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 5,
				MinSize: 3,
			})
			facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Alice",
			})

			courseAfterEnrollment, _ := facade.GetCourse(&usecases.GetCourseQuery{Id: course.Id})

			if len(courseAfterEnrollment.Enrollments) != 1 {
				t.Error("Expected to have an enrollment in the course by got 0")
			}
		})

		t.Run("when the course reach the minimum enrollments, it becomes viable", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 5,
				MinSize: 2,
			})
			facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Alice",
			})
			facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Bob",
			})

			courseAfterEnrollment, _ := facade.GetCourse(&usecases.GetCourseQuery{Id: course.Id})

			if courseAfterEnrollment.Status != "viable" {
				t.Error("course status expected to viable, got ", courseAfterEnrollment.Status)
			}
		})

		t.Run("should not be able to enroll in the course after reaching the maximum size", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 1,
			})
			students := []string{
				"Alice", "Bob", "Charlie",
			}
			for _, student := range students {
				_, err := facade.Enroll(&usecases.EnrollCommand{
					CourseId: course.Id,
					Student:  student,
				})
				if err != nil {
					t.Error("Expect no error, but got", err)
				}
			}

			_, err := facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Derek",
			})

			if err == nil {
				t.Error("expected to have an error, got nil")
			}
		})

		t.Run("should be able to cancel existing enrollments", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 2,
			})

			res, _ := facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Alice",
			})

			facade.CancelEnrollment(&usecases.CancelEnrollmentCommand{
				CourseId:     course.Id,
				EnrollmentId: res.EnrollmentId,
			})

			courseAfterEnrollment, _ := facade.GetCourse(&usecases.GetCourseQuery{Id: course.Id})

			if len(courseAfterEnrollment.Enrollments) != 0 {
				t.Error("expected to not have any enrollments, but got", len(courseAfterEnrollment.Enrollments))
			}
		})

		t.Run("if the course reach a size smaller than the minimum size due to enrollment cancellation, then it should become not-viable again", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 1,
			})

			res, _ := facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Alice",
			})

			courseAfter, _ := facade.GetCourse(&usecases.GetCourseQuery{Id: course.Id})
			if courseAfter.Status != "viable" {
				t.Error("expect the course status to be viable, but got", courseAfter.Status)
			}

			facade.CancelEnrollment(&usecases.CancelEnrollmentCommand{
				CourseId:     course.Id,
				EnrollmentId: res.EnrollmentId,
			})

			courseAfter, _ = facade.GetCourse(&usecases.GetCourseQuery{Id: course.Id})
			if courseAfter.Status != "not-viable" {
				t.Error("expect the course status to be not-viable, but got", courseAfter.Status)
			}
		})
	})

	t.Run("Course Cancellation", func(t *testing.T) {
		t.Run("Should not be able to cancel a non existing courses", func(t *testing.T) {
			_, err := facade.CancelCourse(&usecases.CancelCourseCommand{Id: "random"})
			if err == nil {
				t.Error("expected error, got nil")
			}
		})

		t.Run("Should be able to cancel existing courses", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 1,
			})

			facade.CancelCourse(&usecases.CancelCourseCommand{Id: course.Id})

			courseAfter, _ := facade.GetCourse(&usecases.GetCourseQuery{Id: course.Id})

			if courseAfter.Status != "cancelled" {
				t.Error("expected course status to be cancelled, got", courseAfter.Status)
			}
		})

		t.Run("should not be able to cancel a course twice", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 1,
			})

			facade.CancelCourse(&usecases.CancelCourseCommand{Id: course.Id})

			_, err := facade.CancelCourse(&usecases.CancelCourseCommand{Id: course.Id})
			if err == nil {
				t.Error("expected error, got nil")
			}
		})

		t.Run("cannot enroll into cancelled course", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 1,
			})

			facade.CancelCourse(&usecases.CancelCourseCommand{Id: course.Id})

			_, err := facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Alice",
			})

			if err == nil {
				t.Error("expected error, got nil")
			}
		})

		t.Run("cannot cancel enrollment of a cancelled course", func(t *testing.T) {
			course, _ := facade.ScheduleCourse(&usecases.ScheduleCourseCommand{
				Name:    "My Course",
				MaxSize: 3,
				MinSize: 1,
			})

			res, _ := facade.Enroll(&usecases.EnrollCommand{
				CourseId: course.Id,
				Student:  "Alice",
			})

			facade.CancelCourse(&usecases.CancelCourseCommand{Id: course.Id})

			_, err := facade.CancelEnrollment(&usecases.CancelEnrollmentCommand{
				CourseId:     course.Id,
				EnrollmentId: res.EnrollmentId,
			})
			if err == nil {
				t.Error("expected error, got nil")
			}
		})
	})
}
