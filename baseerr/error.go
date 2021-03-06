package baseerr

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	_ error = new(AppError)
)

type AppError struct {
	StatusCode int                    `json:"-"`
	Err        error                  `json:"error"`
	Data       map[string]interface{} `json:"data"`
}

func New(statusCode int, err error) AppError {
	return AppError{
		StatusCode: statusCode,
		Err:        err,
		Data:       make(map[string]interface{}),
	}
}

func (a AppError) Error() string {
	return fmt.Sprintf("AppError: %s", a.Err)
}

func (a AppError) Unwrap() error {
	return a.Err
}

func (a AppError) WithData(key string, value interface{}) AppError {
	a.Data[key] = value
	return a
}

func (a AppError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Err  string                 `json:"error"`
		Data map[string]interface{} `json:"data"`
	}{
		Err:  a.Err.Error(),
		Data: a.Data,
	})
}

func (a *AppError) UnmarshalJSON(data []byte) error {
	var tempS = struct {
		Err  string                 `json:"error"`
		Data map[string]interface{} `json:"data"`
	}{}

	if err := json.Unmarshal(data, &tempS); err != nil {
		return err
	}

	a.Data = tempS.Data
	a.Err = errors.New(tempS.Err)
	return nil
}
