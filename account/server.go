package account

import (
	"context"
	"fmt"
	"net"

	"github.com/vaxxnsh/go-microservices/account/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//go:generate protoc --go_out=pb --go-grpc_out=pb   --go_opt=paths=source_relative   --go-grpc_opt=paths=source_relative   account.proto
type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	service Service
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	pb.RegisterAccountServiceServer(serv, &grpcServer{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		service:                           s,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostAccount(ctx context.Context, r *pb.PostAccountRequest) (*pb.PostAccountResponse, error) {
	account, err := s.service.PostAccount(ctx, r.Name)
	if err != nil {
		return nil, err
	}

	return &pb.PostAccountResponse{
		Account: &pb.Account{
			Id:   account.ID,
			Name: account.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccount(ctx context.Context, r *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := s.service.GetAccount(ctx, r.Id)

	if err != nil {
		return nil, err
	}

	return &pb.GetAccountResponse{
		Account: &pb.Account{
			Id:   account.ID,
			Name: account.Name,
		},
	}, nil
}

func (s *grpcServer) GetAccounts(ctx context.Context, r *pb.GetAccountsRequest) (*pb.GetAccountsResponse, error) {
	res, err := s.service.GetAccounts(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	accounts := []*pb.Account{}
	for _, p := range res {
		accounts = append(
			accounts,
			&pb.Account{
				Id:   p.ID,
				Name: p.Name,
			},
		)
	}
	return &pb.GetAccountsResponse{Accounts: accounts}, nil
}
