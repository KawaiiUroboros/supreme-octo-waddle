package server

import (
	"context"
	api_pb "emergence/api"
	createChatsMediator "emergence/commandsAndQueries/CreateChat"
	deleteChatsMediator "emergence/commandsAndQueries/DeleteChat"
	"emergence/integration"
	"github.com/baranius/godiator"
	"github.com/izumin5210/grapi/pkg/grapiserver"
)

// NotyfierServiceServer is a composite interface of api_pb.NotyfierServiceServer and grapiserver.Server.
type NotyfierServiceServer interface {
	api_pb.NotyfierServiceServer
	grapiserver.Server
}

// NewNotyfierServiceServer creates a new NotyfierServiceServer instance.
func NewNotyfierServiceServer(repository integration.IRepository) NotyfierServiceServer {
	return &notyfierServiceServerImpl{
		mediator: godiator.GetInstance(),
		repo:     repository,
	}
}

type notyfierServiceServerImpl struct {
	mediator godiator.IGodiator
	repo     integration.IRepository
}

func (s *notyfierServiceServerImpl) GetNotification(ctx context.Context, request *api_pb.GetNotificationRequest) (*api_pb.GetNotificationResponse, error) {
	//TODO implement me
	return &api_pb.GetNotificationResponse{Answer: "Ok"}, nil
}

func (s *notyfierServiceServerImpl) DeleteChats(ctx context.Context, request *api_pb.DeleteChatsRequest) (*api_pb.DeleteChatsResponse, error) {
	result, err := s.mediator.Send(&deleteChatsMediator.DeleteChatsCommand{
		ExternalUserIds: &request.UserIds,
	})
	if err != nil {
		return nil, err
	}
	response := result.(*deleteChatsMediator.DeleteChatsResult)
	return &api_pb.DeleteChatsResponse{
		DeletedChats: *response.DeletedChatIds,
	}, nil
}

func (s *notyfierServiceServerImpl) CreateChats(ctx context.Context, req *api_pb.CreateChatsRequest) (*api_pb.CreateChatsResponse, error) {

	//turn int 32 in int for req.ExternalUserIdToInterval and map[string]int32 to map[string]int
	var externalUserIdToInterval = make(map[string]int)
	for k, v := range req.ExternalUserIdToInterval {
		externalUserIdToInterval[k] = int(v)
	}

	request := &createChatsMediator.CreateChatCommand{ExternalUserIdToInterval: &externalUserIdToInterval}

	result, err := s.mediator.Send(request)

	if err != nil {
		return nil, err
	}

	response := result.(*createChatsMediator.CreateChatResult)

	return &api_pb.CreateChatsResponse{ExternalUserIdToChatUrl: *response.ExternalUserIdToChatUrl}, nil
}
