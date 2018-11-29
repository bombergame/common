package registry

import (
	"github.com/bombergame/common/logs"
	consul "github.com/hashicorp/consul/api"
	"log"
	"time"
)

type Client struct {
	Config      ClientConfig
	Components  ClientComponents
	ConsulAgent *consul.Agent
}

type ClientConfig struct {
	ServiceName     string
	ServiceTTL      time.Duration
	RegistryAddress string
}

type ClientComponents struct {
	Logger *logs.Logger
}

func NewClient(cf ClientConfig, cp ClientComponents) *Client {
	instance := &Client{
		Config:     cf,
		Components: cp,
	}

	conf := consul.DefaultConfig()
	conf.Address = cf.RegistryAddress

	c, err := consul.NewClient(conf)
	if err != nil {
		panic(err)
	}
	instance.ConsulAgent = c.Agent()

	serviceDef := &consul.AgentServiceRegistration{
		Name: cf.ServiceName,
		Check: &consul.AgentServiceCheck{
			TTL: cf.ServiceTTL.String(),
		},
	}

	if err := instance.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		panic(err)
	}

	return instance
}

type StatusCheck func() (bool, error)

func (instance *Client) UpdateTTL(check StatusCheck) {
	ticker := time.NewTicker(instance.Config.ServiceTTL / 2)
	for range ticker.C {
		instance.update(check)
	}
}

func (instance *Client) update(check StatusCheck) {
	ok, err := check()
	if !ok {
		if agentErr := instance.ConsulAgent.FailTTL(
			"service:"+instance.Config.ServiceName, err.Error(),
		); agentErr != nil {
			log.Print(agentErr)
		}
	} else {
		if agentErr := instance.ConsulAgent.PassTTL(
			"service:"+instance.Config.ServiceName, err.Error(),
		); agentErr != nil {
			log.Print(agentErr)
		}
	}
}
