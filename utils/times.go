package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type HTime struct {
	time.Time
}

var (
	formatTime = "2006-01-02 15:04:05"
)

func (t HTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

// UnmarshalJSON 实现 JSON 反序列化
func (t *HTime) UnmarshalJSON(data []byte) error {
	// 移除引号
	str := string(data)
	if len(str) < 2 || str[0] != '"' || str[len(str)-1] != '"' {
		return fmt.Errorf("invalid time format: %s", str)
	}
	str = str[1 : len(str)-1]
	
	// 解析时间
	parsed, err := time.Parse(formatTime, str)
	if err != nil {
		return err
	}
	
	t.Time = parsed
	return nil
}

// Value 实现 driver.Valuer 接口，用于数据库写入
func (t HTime) Value() (driver.Value, error) {
	return t.Time, nil
}

// Scan 实现 sql.Scanner 接口，用于数据库读取
func (t *HTime) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}
	
	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case string:
		parsed, err := time.Parse(formatTime, v)
		if err != nil {
			return err
		}
		t.Time = parsed
		return nil
	case []byte:
		parsed, err := time.Parse(formatTime, string(v))
		if err != nil {
			return err
		}
		t.Time = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into HTime", value)
	}
}