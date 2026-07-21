package response

type UserResponse struct {
	UUID  string
	Name  string
	Email string
	Role  string
}

type RegisterResponse struct {
	User UserResponse
}

type LoginResponse struct {
	User  UserResponse
	Token string
}
