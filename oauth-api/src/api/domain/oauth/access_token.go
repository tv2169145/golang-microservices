package oauth

import "time"

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	Expires     int64 `json:"expires"`	// 2020-01-18T23:00:00Z
}

func (at *AccessToken) IsExpired() bool {
	zone, _ := time.LoadLocation("Asia/Taipei")
	return time.Unix(at.Expires, 0).In(zone).Before(time.Now().In(zone))
}