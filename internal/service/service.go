package service

type Authentication interface {
	SignIn()
}

type Profile interface {
	GetDetail()
}
