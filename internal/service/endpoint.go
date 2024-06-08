package service

type createUserRequest struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type createUserResponse struct {
	Message string `json:"message"`
}
