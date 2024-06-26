package dto

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
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
