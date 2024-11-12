package dto

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

// 1.在使用 gorm 框架时，用到 MarshalJSON / Scan / Value
// 2.在使用 Gin 框架时，使用 ShouldBindJSON 绑定参数会调用 UnmarshalJSON，context.JSON 返回 json 时会调用 MarshalJSON

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		return nil
	}
	parseInLocation, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	*t = LocalTime(parseInLocation)
	return err
}

func (t LocalTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return "0000-00-00 00:00:00", nil
	}
	return time.Time(t), nil
}

func (t LocalTime) ToTime() time.Time {
	return time.Time(t)
}

func (t LocalTime) ToString() string {
	if t.ToTime().IsZero() {
		return ""
	}
	return time.Time(t).Format(time.DateTime)
}
