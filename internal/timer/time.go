package timer

import "time"

func GetNowTime() time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai") // 考量時區
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d) // 解析出持續時間
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil // 目前時間 + 持續時間
}
