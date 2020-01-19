package oauth

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessToken_IsExpired(t *testing.T) {
	zone, _ := time.LoadLocation("Asia/Taipei")
	h, _ := time.ParseDuration("-1h")
	oldToken := &AccessToken{
		AccessToken: "USR_123",
		UserId: 123,
		Expires: time.Now().In(zone).Add(1 * h).Unix(),
	}
	bool := oldToken.IsExpired()
	assert.EqualValues(t, true, bool)
}
