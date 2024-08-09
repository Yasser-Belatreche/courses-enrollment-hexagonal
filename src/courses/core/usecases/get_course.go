package usecases

import "github.com/Yasser-Belatreche/courses-enrollment-hexagonal/src/courses/core/domain"

type GetCourseQuery struct {
	Id string
}

type GetCourseQueryResponse struct {
	Id          string
	Name        string
	MaxSize     int
	MinSize     int
	Status      string
	Enrollments []struct {
		Id      string
		Student string
	}
}

type GetCourseQueryHandler struct {
	CourseRepository     domain.CourseRepository
	EnrollmentRepository domain.EnrollmentRepository
}

// Handle : in a real word scenario, this method would query the database directly to get the course with the enrollments, no need to pass by any repositories
func (h *GetCourseQueryHandler) Handle(query *GetCourseQuery) (*GetCourseQueryResponse, error) {
	course, err := h.CourseRepository.FindById(query.Id)
	if err != nil {
		return nil, err
	}

	enrollments, err := h.EnrollmentRepository.FindByCourseId(query.Id)
	if err != nil {
		return nil, err
	}

	response := &GetCourseQueryResponse{
		Id:      course.State().Id,
		Name:    course.State().Name,
		MaxSize: course.State().MaxSize,
		MinSize: course.State().MinSize,
		Status:  course.State().Status,
	}

	for _, enrollment := range enrollments {
		response.Enrollments = append(response.Enrollments, struct {
			Id      string
			Student string
		}{
			Id:      enrollment.Id,
			Student: enrollment.Student,
		})
	}

	return response, nil
}
