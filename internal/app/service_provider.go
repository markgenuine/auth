package app

import (
	"context"
	"log"

	"github.com/markgenuine/auth/internal/api/auth"
	"github.com/markgenuine/auth/internal/client/db"
	"github.com/markgenuine/auth/internal/client/db/pg"
	"github.com/markgenuine/auth/internal/client/db/transaction"
	"github.com/markgenuine/auth/internal/closer"
	"github.com/markgenuine/auth/internal/config"
	"github.com/markgenuine/auth/internal/config/env"
	"github.com/markgenuine/auth/internal/repository"
	userRepo "github.com/markgenuine/auth/internal/repository/auth"
	"github.com/markgenuine/auth/internal/service"
	userService "github.com/markgenuine/auth/internal/service/auth_v1"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	authService    service.AuthService
	userImpl       *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %s", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = userService.NewService(s.UserRepository(ctx), s.TxManager(ctx))
	}

	return s.authService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *auth.Implementation {
	if s.userImpl == nil {
		s.userImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.userImpl
}
