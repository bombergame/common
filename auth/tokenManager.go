package auth

import (
	"time"
)

const (
	ExpireTimeFormat = time.RFC3339
)

type UserInfo struct {
	ProfileID  int64  `mapstructure:"profile_id"`
	UserAgent  string `mapstructure:"user_agent"`
	ExpireTime string `mapstructure:"expire_time"`
}

type TokenManager interface {
	CreateToken(info UserInfo) (string, error)
	ParseToken(token string) (*UserInfo, error)
}
