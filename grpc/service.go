package grpc

import (
	"net"

	"github.com/bombergame/common/auth"
	"github.com/bombergame/common/logs"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//Service is a wrapper for grpc server
type Service struct {
	Config     ServiceConfig
	Components ServiceComponents
	Server     *grpc.Server
}

//ServiceConfig contains service configuration parameters
type ServiceConfig struct {
	Host string
	Port string
}

//ServiceComponents contains service components
type ServiceComponents struct {
	Logger      *logs.Logger
	AuthManager auth.AuthenticationManager
}

//NewService creates service instance
func NewService(cf ServiceConfig, cp ServiceComponents) *Service {
	logEntry := logrus.NewEntry(cp.Logger.LogrusLogger())
	grpc_logrus.ReplaceGrpcLogger(logEntry)

	srv := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
			grpc_recovery.UnaryServerInterceptor(),
			grpc_logrus.UnaryServerInterceptor(logEntry),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(
				grpc_ctxtags.WithFieldExtractor(
					grpc_ctxtags.CodeGenRequestFieldExtractor,
				),
			),
			grpc_recovery.StreamServerInterceptor(),
			grpc_logrus.StreamServerInterceptor(logEntry),
		),
	)

	return &Service{
		Config:     cf,
		Components: cp,
		Server:     srv,
	}
}

//Run starts the service
func (srv *Service) Run() error {
	addr := srv.Config.Host + ":" + srv.Config.Port

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	srv.Logger().Info("grpc service running on: " + addr)
	return srv.Server.Serve(lis)
}

//Shutdown forces the service to stop
func (srv *Service) Shutdown() error {
	srv.Logger().Info("grpc service shutdown initialized")
	srv.Server.GracefulStop()
	return nil
}

//Logger returns the service logger
func (srv *Service) Logger() *logs.Logger {
	return srv.Components.Logger
}
