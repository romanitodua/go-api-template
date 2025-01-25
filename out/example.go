package out

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Example struct {
}

func (a Example) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Example) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}
