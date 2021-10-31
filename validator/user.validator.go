package validator

type CreateUser struct {
	FirstName string `json:"firstName" binding:"required,alpha,min=1"`
	LastName  string `json:"lastName" binding:"required,alpha,min=1"`
	Username  string `json:"username" binding:"required,alphanum,min=4"`
	Age       int    `json:"age" binding:"required,gte=12"`
	Password  string `json:"password" binding:"required,alphanum,min=8"`
}
