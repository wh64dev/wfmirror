package auth

type Account struct {
	Username string
	Password string
}

type AccountData struct {
	Id       string
	Username string
}

func (acc *Account) GetId() string {
	return ""
}
