package internal

type UserInfo struct {
	UserName string                 `json:"username" gorm:"column:username;primaryKey"`
	Password string                 `json:"password" gorm:"column:password"`
	AuthKey  string                 `json:"authkey" gorm:"column:authkey"`
	Mata     map[string]interface{} `json:"meta" gorm:"column:meta;type:json"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type AuthCode struct {
	AuthKey string `json:"authkey" gorm:"column:authkey"`
}

func (AuthCode) TableName() string {
	return "auth_code"
}
