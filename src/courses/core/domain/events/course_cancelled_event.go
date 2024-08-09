package events

const CourseCancelledEventType = "Courses.Cancelled"

type CourseCancelledEventPayload struct {
	Id string
}
