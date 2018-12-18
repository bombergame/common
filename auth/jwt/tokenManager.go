package jwt

import (
	"time"

	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/consts"
	"github.com/bombergame/common/errs"
	"github.com/bombergame/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

const (
	//DefaultKeyLength is the default length of encryption key
	DefaultKeyLength = 64
)

//TokenManager provides JWT token handlers
type TokenManager struct {
	key []byte
}

//NewTokenManager creates new JWT token manager
func NewTokenManager(key string) *TokenManager {
	randSeqGen := utils.NewRandomSequenceGenerator()

	if key == consts.EmptyString {
		key = randSeqGen.Next(DefaultKeyLength)
	}

	return &TokenManager{
		key: []byte(key),
	}
}

//CreateToken creates new JWT token
func (m *TokenManager) CreateToken(info auth.TokenInfo) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"profile_id":  info.ProfileID,
		"user_agent":  info.UserAgent,
		"expire_time": time.Now().Format(auth.ExpireTimeFormat),
	})
	return t.SignedString(m.key)
}

//ParseToken returns authorized user info by token
func (m *TokenManager) ParseToken(token string) (*auth.TokenInfo, error) {
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

	info := &auth.TokenInfo{}

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result: info, TagName: "mapstructure",
	})
	if err != nil {
		return nil, errs.NewInternalServiceError(err)
	}

	if err := decoder.Decode(claims); err != nil {
		return nil, invFmtErr
	}

	return info, nil
}
