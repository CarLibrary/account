package serializer

import (
	"github.com/CarLibrary/account/model"
	pb "github.com/CarLibrary/proto/account"
)

func BuildUser(user *model.User) *pb.User {
	return &pb.User{
		Id:       int32(user.ID),
		Username: user.Username,
	}
}

func BuildUserInfo(user *model.User) *pb.UserInfoResponse {
	return &pb.UserInfoResponse{
		Id:       int32(user.ID),
		Username: user.Username,
		HeadUrl:  user.Head,
		Sign:     user.Sign,
	}
}