package service

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
func (s *UserService) GetUsers() string {
	return "List users"
}
