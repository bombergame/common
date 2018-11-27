package auth

type UserInfo struct {
	ProfileID int64  `mapstructure:"profile_id"`
	UserAgent string `mapstructure:"user_agent"`
}

type TokenManager interface {
	CreateToken(info UserInfo) (string, error)
	ParseToken(token string) (*UserInfo, error)
}
