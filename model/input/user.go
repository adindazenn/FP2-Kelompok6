package input

type UserRegisterInput struct {
	Username string `json:"username" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Password string `json:"password" valid:"required,length(6|20)"`
	Age      int    `json:"age" valid:"required"`
}

type UserLoginInput struct {
	Email    string `json:"email" valid:"required"`
	Password string `json:"password" valid:"required"`
}

type UserUpdateInput struct {
	Email    string `json:"email" valid:"email"`
	Username string `json:"username"`
	Password string `json:"password" valid:"length(6|20)"`
}

type UserUpdateID struct {
	ID int `uri:"id" binding:"required"`
}

type UserDeleteID struct {
	ID int `uri:"id" binding:"required"`
}
