package services

type UserService interface {
}

type ServiceConnector struct {
}

func NewUserService() UserService {
	return &ServiceConnector{}
}
