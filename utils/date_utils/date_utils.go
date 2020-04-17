package date_utils

import "time"

const(
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDbLayout   = "2006-01-02 15:04:05"
)

//GetNow a..
func GetNow() time.Time{
	return time.Now().UTC()
}

// GetNowString a..
func GetNowString() string{
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}