package e2e_tests

import (
	"context"
	"database/sql"
	"github.com/maiaaraujo5/controle-de-transacao/app/fx/server"
	"github.com/maiaaraujo5/controle-de-transacao/app/provider/postgres/model"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"os"
	"testing"
)

type E2ETestSuite struct {
	suite.Suite
	dbConn   *bun.DB
	posgtres testcontainers.Container
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

func (s *E2ETestSuite) SetupSuite() {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16.2-alpine3.18",
		ExposedPorts: []string{"5437:5432"},
		Env:          map[string]string{"POSTGRES_PASSWORD": "pismo", "POSTGRES_USER": "pismo", "POSTGRES_DB": "test"},
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
		Name:         "postgres-test",
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatal(err)
	}

	sqlDb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithAddr("localhost:"+port.Port()),
		pgdriver.WithUser("pismo"),
		pgdriver.WithPassword("pismo"),
		pgdriver.WithDatabase("test"),
		pgdriver.WithInsecure(true),
	))

	s.dbConn = bun.NewDB(sqlDb, pgdialect.New())
	s.posgtres = container

	err = os.Setenv("CONF", "./configs/test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	go server.Start()
}

func (s *E2ETestSuite) TearDownSuite() {
	err := s.posgtres.Terminate(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (s *E2ETestSuite) SetupTest() {
	_, err := s.dbConn.NewCreateTable().Model((*model.Account)(nil)).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.dbConn.NewCreateTable().Model((*model.Transaction)(nil)).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (s *E2ETestSuite) TearDownTest() {
	_, err := s.dbConn.NewDropTable().Model((*model.Account)(nil)).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.dbConn.NewDropTable().Model((*model.Transaction)(nil)).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
