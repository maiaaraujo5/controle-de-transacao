package e2e_tests

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/model"
	"io"
	"net/http"
	"strings"
)

func (s *E2ETestSuite) TestEndToEnd_Recover_Account_Successfully() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/v1/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"account_id":1,"document_number":"123"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Try_Recover_Account_that_does_not_exists_in_database() {

	req, err := http.NewRequest(echo.GET, "http://localhost:8080/v1/accounts/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusNotFound, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"code":404,"description":"no account was found"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}
