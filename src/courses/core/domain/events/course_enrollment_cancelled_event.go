package events

const CourseEnrollmentCancelledEventType = "Courses.EnrollmentCancelled"

type CourseEnrollmentCancelledEventPayload struct {
	CourseId     string
	EnrollmentId string
	Student      string
}
