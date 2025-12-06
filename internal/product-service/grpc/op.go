package grpc

import (
	"context"
	"database/sql"
	"errors"
	"github.com/autumnterror/breezynotes/pkg/log"
	"github.com/autumnterror/volha-backend/internal/product-service/psql"
	productsRPC "github.com/autumnterror/volha-backend/pkg/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type txStarter interface {
	psql.SqlRepo
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

type ServerAPI struct {
	productsRPC.UnimplementedProductsServer
	db txStarter
}

func Register(
	server *grpc.Server,
	db txStarter,
) {
	productsRPC.RegisterProductsServer(server, &ServerAPI{
		db: db,
	})
}

func (s *ServerAPI) repo() psql.Repo {
	return psql.Driver{Driver: s.db}
}

func (s *ServerAPI) run(
	ctx context.Context,
	op string,
	action func(repo psql.Repo) (any, error),
) (any, error) {
	return handleCRUDResponse(ctx, op, func() (any, error) {
		return action(s.repo())
	})
}

func (s *ServerAPI) runTx(
	ctx context.Context,
	op string,
	action func(repo psql.Repo) (any, error),
) (any, error) {
	return handleCRUDResponse(ctx, op, func() (any, error) {
		return s.withTx(ctx, op, action)
	})
}

func (s *ServerAPI) withTx(
	ctx context.Context,
	op string,
	action func(repo psql.Repo) (any, error),
) (any, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	repo := psql.Driver{Driver: tx}

	res, err := action(repo)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, errors.Join(err, rollbackErr)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return res, nil
}

func handleCRUDResponse(ctx context.Context, op string, action func() (any, error)) (any, error) {
	type resV struct {
		res any
		err error
	}
	res := make(chan resV, 1)
	go func() {
		r, err := action()
		res <- resV{r, err}
	}()
	select {
	case <-ctx.Done():
		log.Error(op, "Context dead", ctx.Err())
		return nil, status.Error(codes.DeadlineExceeded, "Context dead")
	case r := <-res:
		if r.err != nil {
			log.Error(op, "", r.err)
			switch {
			case errors.Is(r.err, psql.ErrUnknownType):
				return nil, status.Error(codes.Unimplemented, r.err.Error())
			case errors.Is(r.err, psql.ErrInvalidType):
				return nil, status.Error(codes.InvalidArgument, r.err.Error())
			case errors.Is(r.err, psql.ErrNotFound):
				return nil, status.Error(codes.NotFound, r.err.Error())
			case errors.Is(r.err, psql.ErrAlreadyExists):
				return nil, status.Error(codes.AlreadyExists, r.err.Error())
			case errors.Is(r.err, psql.ErrForeignKey):
				return nil, status.Error(codes.FailedPrecondition, r.err.Error())
			default:
				return nil, status.Error(codes.Internal, "check logs")
			}
		}
		log.Green(op)
		if r.res != nil {
			return r.res, nil
		}
		return nil, nil
	}
}
