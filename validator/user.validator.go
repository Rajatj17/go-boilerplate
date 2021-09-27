package validator

type CreateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Age       int    `json:"age"`
	Password  string `json:"password"`
}
