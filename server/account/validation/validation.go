package validation

type RegistValid struct {
	Phone    string `validate:"max=11,min=11,required"`
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type LoginValid struct {
	PhoneEmail string `validate:"required"`
	Password   string `valid:"required"`
}

type DeleteByIdValid struct {
	Id int `validate:"required"`
}

type AddUserValid struct {
	Phone       string `validate:"required,max=11,min=11"`
	Username    string `validate:"required"`
	Password    string `validate:"required"`
	UserType    string
	Declaration string
	Integral	int
	Level		string
	Email       string
	Avatar      string
}

type UpdateUserValid struct {
	Id          int    `validate:"required"`
	Phone       string `validate:"max=11,min=11"`
	Username    string
	Password    string
	UserType    string
	Match       string
	Declaration string
	Email       string
	Integral	int
	Level		string
	Avatar      string
}

type DynamicValid struct {
	Content string `validate:"max=500"`
	ImgPath string
}

type CommentsValid struct {
	Id         int
	UserId     int `validate:"required"`
	DynamicId  int `validate:"required"`
	CommentsId int
	Content    string `validate:"max=500"`
	ImgPath    string `validate:"max=200"`
}
