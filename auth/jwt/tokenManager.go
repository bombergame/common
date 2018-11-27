package jwt

import (
	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

const (
	DefaultKeyLength  = 64
	DefaultSaltLength = 32
)

type TokenManager struct {
	key        []byte
	randSeqGen *utils.RandomSequenceGenerator
}

func NewTokenManager(key string) *TokenManager {
	randSeqGen := utils.NewRandomSequenceGenerator()

	if key == consts.EmptyString {
		key = randSeqGen.Next(DefaultKeyLength)
	}

	return &TokenManager{
		key:        []byte(key),
		randSeqGen: randSeqGen,
	}
}

func (m *TokenManager) CreateToken(info auth.UserInfo) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"profile_id": info.ProfileID,
		"user_agent": info.UserAgent,
		"rand_salt":  m.randSeqGen.Next(DefaultSaltLength),
	})
	return t.SignedString(m.key)
}

func (m *TokenManager) ParseToken(token string) (*auth.UserInfo, error) {
	invFmtErr := errs.NewInvalidFormatError("wrong token")

	t, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invFmtErr
		}
		return m.key, nil
	})

	if err != nil || !t.Valid {
		return nil, invFmtErr
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return nil, invFmtErr
	}

	info := &auth.UserInfo{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: info, TagName: "mapstructure",
	})
	if err != nil {
		return nil, errs.NewServiceError(err)
	}

	if err := decoder.Decode(claims); err != nil {
		return nil, invFmtErr
	}

	return info, nil
}