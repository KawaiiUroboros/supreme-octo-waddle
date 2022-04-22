package server

import (
	"context"
	createIncedentCommand "emergence/commandsAndQueries/CreateIncident"
	commandsAndQueries "emergence/commandsAndQueries/CreateTeam"
	getCityAndAddressByExternalUserIdCommand "emergence/commandsAndQueries/GetCityAndAddressByExternalUserId"
	"emergence/commandsAndQueries/models"
	"emergence/integration"
	"github.com/baranius/godiator"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api_pb "emergence/api"
)

// HospitalServiceServer is a composite interface of api_pb.HospitalServiceServer and grapiserver.Server.
type HospitalServiceServer interface {
	api_pb.HospitalServiceServer
	grapiserver.Server
}

// NewHospitalServiceServer creates a new HospitalServiceServer instance.
func NewHospitalServiceServer(repository integration.IRepository) HospitalServiceServer {
	return &hospitalServiceServerImpl{
		mediator: godiator.GetInstance(),
		repo:     repository,
	}
}

type hospitalServiceServerImpl struct {
	mediator godiator.IGodiator
	repo     integration.IRepository
}

func (s *hospitalServiceServerImpl) GetCityAndAddressByExternalUserId(ctx context.Context, request *api_pb.GetCityAndAddressByExternalUserIdRequest) (*api_pb.GetCityAndAddressByExternalUserIdResponse, error) {
	//create command and send it to mediator
	command := getCityAndAddressByExternalUserIdCommand.GetCityAndAddressByExternalUserIdCommand{
		ExternalUserId: request.ExternalUserId,
	}
	result, err := s.mediator.Send(&command)
	if err != nil {
		return &api_pb.GetCityAndAddressByExternalUserIdResponse{},
			status.Error(codes.Internal, err.Error())
	}
	//get result from mediator
	cityAndAddress := result.(*getCityAndAddressByExternalUserIdCommand.GetCityAndAddressByExternalUserIdCommandResult)
	//create response
	response := &api_pb.GetCityAndAddressByExternalUserIdResponse{
		City:    cityAndAddress.City,
		Address: cityAndAddress.Address,
	}
	return response, nil
}

func (s *hospitalServiceServerImpl) CreateIncident(ctx context.Context, request *api_pb.CreateIncidentRequest) (*api_pb.CreateIncidentResponse, error) {
	command := &createIncedentCommand.CreateIncidentCommand{
		Incident: models.Incident{
			City:        request.City,
			Description: request.Description,
		},
	}
	_, err := s.mediator.Send(command)
	if err != nil {
		return &api_pb.CreateIncidentResponse{Success: false}, status.Error(codes.Internal, err.Error())
	}
	return &api_pb.CreateIncidentResponse{
		Success: true,
	}, nil
}

func (s *hospitalServiceServerImpl) CreateTeam(ctx context.Context, req *api_pb.CreateTeamRequest) (*api_pb.CreateTeamResponse, error) {
	createTeamCommand := &commandsAndQueries.CreateTeamCommand{
		City:   req.City,
		Emails: req.Emails,
	}
	if createTeamCommand == nil {
		return nil, status.Errorf(codes.InvalidArgument, "createTeamCommand must be set")
	}

	result, err := s.mediator.Send(createTeamCommand)

	if err != nil {
		return &api_pb.CreateTeamResponse{
			Success: false,
		}, status.Errorf(codes.Internal, "failed to create team: %v", err)
	}

	response := result.(*commandsAndQueries.CreateTeamResult)

	return &api_pb.CreateTeamResponse{
		Success: response.Success,
	}, nil
}
