package randtoken

import (
	"errors"

	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/utils"
)

const (
	//DefaultKeyLength is the default random key length
	DefaultKeyLength = 128
)

//TokenManager provides random token handlers
type TokenManager struct {
	randSeqGen *utils.RandomSequenceGenerator
}

//NewTokenManager creates new random token manager
func NewTokenManager() *TokenManager {
	return &TokenManager{
		randSeqGen: utils.NewRandomSequenceGenerator(),
	}
}

//CreateToken returns new random token
func (m *TokenManager) CreateToken(_ auth.TokenInfo) (string, error) {
	return m.randSeqGen.Next(DefaultKeyLength), nil
}

//ParseToken does nothing. Required by interface
func (m *TokenManager) ParseToken(_ string) (*auth.TokenInfo, error) {
	err := errors.New("simple token cannot be parsed")
	return nil, errs.NewInternalServiceError(err)
}
