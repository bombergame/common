package jwt

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/consts"
)

func TestTokenManagerUnit(t *testing.T) {
	manager := NewTokenManager(consts.EmptyString)
	if manager == nil {
		t.Error("token manager not created")
	}

	info := auth.TokenInfo{
		ProfileID:  strconv.FormatInt(rand.Int63(), 10),
		UserAgent:  "some-user-agent",
		ExpireTime: time.Now().Format(auth.ExpireTimeFormat),
	}

	token, err := manager.CreateToken(info)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	pInfo, err := manager.ParseToken(token)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if info != *pInfo {
		t.Error("tokens differ")
	}

	_, err = manager.ParseToken("some_invalid_token")
	if err == nil {
		t.Error("invalid token parsed")
	}
}
