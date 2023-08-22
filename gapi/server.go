package gapi

import (
	"fmt"

	db "github.com/jonathangloria/simplebank/db/sqlc"
	"github.com/jonathangloria/simplebank/pb"
	"github.com/jonathangloria/simplebank/token"
	"github.com/jonathangloria/simplebank/util"
	"github.com/jonathangloria/simplebank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
