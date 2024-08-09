package usecases

type CancelEnrollmentCommand struct {
	CourseId     string
	EnrollmentId string
}

type CancelEnrollmentCommandResponse struct {
	CourseId     string
	EnrollmentId string
}

type CancelEnrollmentCommandHandler struct{}

func (h *CancelEnrollmentCommandHandler) Handle(command *CancelEnrollmentCommand) (*CancelEnrollmentCommandResponse, error) {
	return nil, nil
}
