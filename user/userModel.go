package user

type user struct {
	Username   string        `binding:"required" json:"username, string"`
	FirstName  string        `binding:"required" json:"firstName, string"`
	Lastname   string        `binding:"required" json:"lastName, string"`
	Email      string        `binding:"required" json:"email, string"`
	Attributes userAttribute `binding:"required" json:"attributes"`
}

type userAttribute struct {
	DepartmentNm string `json:"departmentNm, string"`
	Position     string `json:"position, string"`
	PhoneNumber  string `json:"phoneNumber, string"`
}

func InitUserInfo() *user {
	user := user{}
	return &user
}
