package server

type createUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
