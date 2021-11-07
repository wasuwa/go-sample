package middleware

import "time"

func TimeNow() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return time.Now().In(jst)
}
