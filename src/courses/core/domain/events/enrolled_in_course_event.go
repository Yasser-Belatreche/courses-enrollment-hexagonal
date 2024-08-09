package events

const EnrolledInCourseEventType = "Courses.Enrolled"

type EnrolledInCourseEventPayload struct {
	CourseId     string
	EnrollmentId string
	Student      string
}
