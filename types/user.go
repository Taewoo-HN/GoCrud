package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"users"`
}

type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age"binding:"required"`
}

type CreateUserResponse struct {
	*ApiResponse
}

func (c *CreateUserRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`

	Age int64 `json:"updatedAge"binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserRequest struct {
	Name string `json:"name" binding:"required"`
}

func (request DeleteUserRequest) ToUser() *User {
	return &User{
		Name: request.Name,
	}
}

type DeleteUserResponse struct {
	*ApiResponse
}
