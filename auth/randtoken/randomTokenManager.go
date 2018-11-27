package randtoken

import (
	"errors"
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/utils"
)

const (
	DefaultKeyLength = 128
)

type TokenManager struct {
	randSeqGen *utils.RandomSequenceGenerator
}

func NewTokenManager() *TokenManager {
	return &TokenManager{
		randSeqGen: utils.NewRandomSequenceGenerator(),
	}
}

func (m *TokenManager) CreateToken(info auth.UserInfo) (string, error) {
	return m.randSeqGen.Next(DefaultKeyLength), nil
}

func (m *TokenManager) ParseToken(token string) (*auth.UserInfo, error) {
	err := errors.New("simple token cannot be parsed")
	return nil, errs.NewServiceError(err)
}
