package transform

import (
	"github.com/trevatk/layer-2a/internal/domain/user"
	pb "github.com/trevatk/layer-2a/proto/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewUser transform grpc new user into internal user
func NewUser(in *pb.NewUser) *user.User {
	return &user.User{
		Name:        in.Name,
		Email:       in.Email,
		Phone:       in.Phone,
		CountryCode: in.CountryCode,
	}
}

// User transform internal user into grpc user
func User(in *user.User) *pb.User {
	return &pb.User{
		Id:          in.ID,
		Name:        in.Name,
		Email:       in.Email,
		CountryCode: in.CountryCode,
		Phone:       in.Phone,
		Timestamp:   timestamppb.Now(),
	}
}

// InUser transform grpc user into internal user
func InUser(in *pb.User) *user.User {
	return &user.User{
		ID:          in.Id,
		Name:        in.Name,
		Email:       in.Email,
		Phone:       in.Phone,
		CountryCode: in.CountryCode,
	}
}

// Users transform internal users to grpc users
func Users(in []*user.User) *pb.Users {

	users := &pb.Users{}

	for _, u := range in {

		user := &pb.User{
			Id:          u.ID,
			Name:        u.Name,
			Email:       u.Email,
			CountryCode: u.CountryCode,
			Phone:       u.Phone,
			Timestamp:   timestamppb.Now(),
		}

		users.User = append(users.User, user)
	}

	return users
}
