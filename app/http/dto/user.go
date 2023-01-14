package dto

type GetUserDto struct {
	Id int
}

type AddUserDto struct {
	Username string
	Password string
	Email    string
}
