package grpc

import (
	"github.com/bombergame/common/logs"
	"google.golang.org/grpc"
)

//Client is a wrapper for grpc client
type Client struct {
	Config     ClientConfig
	Components ClientComponents
	Conn       *grpc.ClientConn
}

//ClientConfig contains the client configuration parameters
type ClientConfig struct {
	ServiceHost string
	ServicePort string
}

//ClientComponents contains the client components
type ClientComponents struct {
	Logger *logs.Logger
}

//NewClient creates client instance
func NewClient(cf ClientConfig, cp ClientComponents) *Client {
	return &Client{
		Config:     cf,
		Components: cp,
	}
}

//Connect establishes connection to the corresponding grpc service
func (c *Client) Connect() error {
	var err error
	c.Conn, err = grpc.Dial(
		c.Config.ServiceHost+":"+c.Config.ServicePort,
		grpc.WithInsecure(),
	)
	return err
}

//Disconnect closes the connection to the service
func (c *Client) Disconnect() error {
	return c.Conn.Close()
}

//Logger returns the client logger
func (c *Client) Logger() *logs.Logger {
	return c.Components.Logger
}
