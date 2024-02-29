package vo

type UserCred struct {
	Email string `json:"email`
	Pass  string `json:"pass"`
	Id    int
}

type UserForm struct {
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Gender   string `json:"gender"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
}
