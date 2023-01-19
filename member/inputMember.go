package member

type InputMember struct {
	Nama     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type LoginMember struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

type CheckEmailIsAvailable struct {
	Email string `binding:"required,email"`
}
