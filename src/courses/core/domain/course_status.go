package domain

type CourseStatus string

const (
	CourseViable    CourseStatus = "viable"
	CourseNotViable CourseStatus = "not-viable"
	CourseCancelled CourseStatus = "cancelled"
)
