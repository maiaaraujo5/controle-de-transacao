package e2e_tests

import (
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/model"
	response2 "github.com/maiaaraujo5/controle-de-transacao/app/server/rest/response"
	"io"
	"net/http"
	"strings"
	"time"
)

func (s *E2ETestSuite) TestEndToEnd_Create_Transaction_With_Cash_Purchase() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	requestStr := `{"account_id":1,"operation_type_id":1,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	res := new(response2.Transaction)
	err = json.Unmarshal(byteBody, res)
	s.NoError(err)

	res.EventDate = time.Date(2024, 04, 1, 1, 1, 1, 1, time.UTC)

	body, err := json.Marshal(res)
	s.NoError(err)

	s.Equal(`{"id":1,"account_id":1,"operation_type_id":1,"amount":-126.85,"event_date":"2024-04-01T01:01:01.000000001Z"}`, strings.Trim(string(body), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Create_Transaction_With_Installment_Purchase() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	requestStr := `{"account_id":1,"operation_type_id":2,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	res := new(response2.Transaction)
	err = json.Unmarshal(byteBody, res)
	s.NoError(err)

	res.EventDate = time.Date(2024, 04, 1, 1, 1, 1, 1, time.UTC)

	body, err := json.Marshal(res)
	s.NoError(err)

	s.Equal(`{"id":1,"account_id":1,"operation_type_id":2,"amount":-126.85,"event_date":"2024-04-01T01:01:01.000000001Z"}`, strings.Trim(string(body), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Create_Transaction_With_Withdraw() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	requestStr := `{"account_id":1,"operation_type_id":3,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	res := new(response2.Transaction)
	err = json.Unmarshal(byteBody, res)
	s.NoError(err)

	res.EventDate = time.Date(2024, 04, 1, 1, 1, 1, 1, time.UTC)

	body, err := json.Marshal(res)
	s.NoError(err)

	s.Equal(`{"id":1,"account_id":1,"operation_type_id":3,"amount":-126.85,"event_date":"2024-04-01T01:01:01.000000001Z"}`, strings.Trim(string(body), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Create_Transaction_With_Payment() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	requestStr := `{"account_id":1,"operation_type_id":4,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	res := new(response2.Transaction)
	err = json.Unmarshal(byteBody, res)
	s.NoError(err)

	res.EventDate = time.Date(2024, 04, 1, 1, 1, 1, 1, time.UTC)

	body, err := json.Marshal(res)
	s.NoError(err)

	s.Equal(`{"id":1,"account_id":1,"operation_type_id":4,"amount":126.85,"event_date":"2024-04-01T01:01:01.000000001Z"}`, strings.Trim(string(body), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}

func (s *E2ETestSuite) TestEndToEnd_Try_To_Create_Transaction_Without_One_Invalid_Account() {
	requestStr := `{"account_id":1,"operation_type_id":4,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/transactions", strings.NewReader(requestStr))
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

func (s *E2ETestSuite) TestEndToEnd_Try_To_Create_Transaction_With_Invalid_Operation_Type_Id() {
	ctx := context.Background()

	_, err := s.dbConn.NewInsert().Model(&model.Account{DocumentNumber: "123"}).Exec(ctx)
	s.NoError(err)

	requestStr := `{"account_id":1,"operation_type_id":100,"amount":126.85}`

	req, err := http.NewRequest(echo.POST, "http://localhost:8080/v1/transactions", strings.NewReader(requestStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.NoError(err)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusPreconditionFailed, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)

	s.Equal(`{"code":412,"description":"operation is not valid"}`, strings.Trim(string(byteBody), "\n"))
	err = response.Body.Close()
	s.NoError(err)
}
