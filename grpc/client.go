package grpc

import (
	"github.com/bombergame/common/logs"
	"google.golang.org/grpc"
)

type Client struct {
	Config     Config
	Components Components
	Conn       *grpc.ClientConn
}

func NewClient(cf Config, cp Components) *Client {
	return &Client{
		Config:     cf,
		Components: cp,
	}
}

func (c *Client) Connect() error {
	var err error
	c.Conn, err = grpc.Dial(
		c.Config.Host+":"+c.Config.Port,
		grpc.WithInsecure(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Disconnect() error {
	return c.Conn.Close()
}

func (c *Client) Logger() *logs.Logger {
	return c.Components.Logger
}
