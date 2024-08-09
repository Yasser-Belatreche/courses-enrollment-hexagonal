package usecases

type EnrollCommand struct {
	Id      string
	Student string
}

type EnrollCommandResponse struct {
	CourseId     string
	EnrollmentId string
}

type EnrollCommandHandler struct{}

func (h *EnrollCommandHandler) Handle(command *EnrollCommand) (*EnrollCommandResponse, error) {
	return nil, nil
}
