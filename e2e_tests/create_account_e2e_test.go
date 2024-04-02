package e2e_tests

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/model"
	"io"
	"net/http"
	"strings"
)

func (s *E2ETestSuite) Test_EndToEnd_Create_Account() {
	requestStr := `{"document_number": "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"account_id":1,"document_number":"123"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) Test_EndToEnd_Try_To_Create_Account_with_a_document_number_already_exists() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	requestStr := `{"document_number": "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusConflict, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"code":409,"description":"an account with this document number already exists"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Try_To_Create_Account_Without_Document_Number() {
	requestStr := `{}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusUnprocessableEntity, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"code":422,"description":"The server understands the content type of the request entity but was unable to process the contained instructions.","validationError":[{"path":"CreateAccount.DocumentNumber","field":"DocumentNumber","value":"","message":"{DocumentNumber} is a required field with type string"}]}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Try_To_Create_Account_With_An_Malformed_Json_Body() {
	requestStr := `{"document_number:" "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusBadRequest, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"code":400,"description":"Syntax error: offset=21, error=invalid character '\"' after object key"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Try_To_Create_Account_With_An_Incorrect_Content_Type() {
	requestStr := `{"document_number": "123"}`
	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/accounts", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusUnsupportedMediaType, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.Equal(`{"code":415,"description":"Unsupported Media Type"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}
