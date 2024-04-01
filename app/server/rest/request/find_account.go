package request

import "github.com/labstack/echo/v4"

type FindAccount struct {
	AccountID int64 `param:"accountId"`
}

func NewFindAccount(c echo.Context) (*FindAccount, error) {
	f := new(FindAccount)
	err := c.Bind(f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
