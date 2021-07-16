package account

import (

	"context"
	"github.com/CarLibrary/account/cache"
	"github.com/CarLibrary/account/model"
	"github.com/CarLibrary/account/serializer"
	"github.com/CarLibrary/account/util"
	pb "github.com/CarLibrary/proto/account"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServiceServer struct {
	pb.UnsafeAccountServiceServer

}

//注册
func (s *AccountServiceServer) Signup(ctx context.Context, re *pb.SignupRequset) (*pb.User, error) {
	var u = model.User{
		Username: re.GetUsername(),
		Password: re.GetPassword(),
		Head:     re.GetHeadUrl(),
		Sign:     re.GetSgin(),
	}
	res,err:=u.Signup()
	if err != nil {
		return &pb.User{},status.Error(codes.Aborted,err.Error())
	}
	return serializer.BuildUser(res), status.Error(codes.OK,"ok")

}

//登录
func (s *AccountServiceServer)Login(ctx context.Context, re *pb.LoginRequest) (*pb.TokenResponse, error){
	var u = model.User{
		Username: re.GetUsername(),
		Password: re.GetPassword(),
	}
	res,err:=u.Login()

	//生成token，并存在redis里面
	tokenstr,err:=util.CreateToken(res.ID,res.Username)
	if err != nil {
		return &pb.TokenResponse{Token: ""}, status.Error(codes.Aborted,err.Error())
	}
	err=cache.SetToken(tokenstr,int32(res.ID))
	if err!=nil {
		return &pb.TokenResponse{Token: ""}, status.Error(codes.Aborted,err.Error())
	}

	//ok
	return &pb.TokenResponse{Token: tokenstr}, status.Error(codes.OK,"ok")

}

//查看个人信息
func (s *AccountServiceServer)GetUserInfo(ctx context.Context, re *pb.InfoRequest) (*pb.UserInfoResponse, error){


	user,err:=model.GetUserInfo(re.GetId())
	if err != nil {
		return &pb.UserInfoResponse{}, status.Error(codes.Aborted,err.Error())
	}
	return serializer.BuildUserInfo(user),status.Error(codes.OK,"ok")

}

//token校验
func (s *AccountServiceServer)CheckToken(ctx context.Context, re *pb.TokenRequest) (*pb.User, error){

	uid,err:=cache.GetToken(re.GetToken())
	if err != nil {
		return &pb.User{}, status.Error(codes.Aborted,err.Error())
	}
	
	user,err:=model.GetUserInfo(uid)
	if err != nil {
		return &pb.User{}, status.Error(codes.Aborted,err.Error())
	}
	return serializer.BuildUser(user), nil
	

}