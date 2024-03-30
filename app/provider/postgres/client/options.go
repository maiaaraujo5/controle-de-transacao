package client

import (
	"log/slog"

	"github.com/maiaaraujo5/controle-de-transacao/app/config"
)

type Options struct {
	URL      string
	Username string
	Password string
	Database string
}

func NewOptions() (*Options, error) {
	options := new(Options)

	err := config.ReadConfigPath("app.provider.postgres", options)
	if err != nil {
		slog.Error("error to read configs from postgres", "err:", err)
		return nil, err
	}

	return options, nil
}
