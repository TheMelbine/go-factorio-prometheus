package factorio

import (
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/gorcon/rcon"
)

type RCONClient struct {
	conn   *rcon.Conn
	logger *log.Logger
}

func NewRCONClient(address, password string) (*RCONClient, error) {
	logger := log.WithPrefix("rcon")
	logger.Debug("Creating new RCON client", "address", address, "password", password)

	conn, err := rcon.Dial(address, password)
	if err != nil {
		return nil, err
	}

	return &RCONClient{conn, logger}, nil
}

func (c *RCONClient) Send(cmd string) (string, error) {
	logger := c.logger.With("reqid", uuid.NewString())
	logger.Debug("Sending command", "command", cmd)

	resp, err := c.conn.Execute(cmd)
	if err != nil {
		logger.Error("Failed to execute command", "error", err)
		return "", err
	}
	logger.Debug("Received response", "response", resp)

	return resp, nil
}

func (c *RCONClient) Close() error {
	return c.conn.Close()
}
