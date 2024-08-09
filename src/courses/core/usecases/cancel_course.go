package usecases

type CancelCourseCommand struct {
	Id string
}

type CancelCourseCommandResponse struct {
	Id string
}

type CancelCourseCommandHandler struct{}

func (h *CancelCourseCommandHandler) Handle(command *CancelCourseCommand) (*CancelCourseCommandResponse, error) {
	return nil, nil
}
