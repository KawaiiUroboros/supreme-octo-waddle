package clients

import (
	"emergence/commandsAndQueries/models"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
	"github.com/opsgenie/opsgenie-go-sdk-v2/user"
	"strings"
)

//IOpsGenieClient interface for OpsGenieClient
type IOpsGenieClient interface {
	CreateTeam(emails []string) (teamId string, err error)
	CreateIncident(incident *models.Incident, teamId string) error
}

//OpsGenieClient is a client for sending alerts to OpsGenie
type OpsGenieClient struct {
	teamClient  *team.Client
	userClient  *user.Client
	alertClient *alert.Client
}

//NewOpsGenieClient creates a new OpsGenieClient with token
func NewOpsGenieClient() *OpsGenieClient {
	apiKey := "631a9139-205b-40ef-856b-e3651c915884"
	teamClient, err := team.NewClient(&client.Config{
		ApiKey: apiKey,
	})
	userClient, err := user.NewClient(&client.Config{
		ApiKey: apiKey,
	})
	alertClient, err := alert.NewClient(&client.Config{
		ApiKey: "0d38a43b-b78a-4c4d-96c8-34910f61f4b4",
	})
	if err != nil {
		panic(err)
	}
	return &OpsGenieClient{
		alertClient: alertClient,
		teamClient:  teamClient,
		userClient:  userClient,
	}
}

//CreateTeam creates a team with given emails
func (client *OpsGenieClient) CreateTeam(emails []string) (teamId string, err error) {

	//get users for each email
	users := make([]user.CreateRequest, len(emails))
	for i, email := range emails {
		users[i] = *getUser(email)
	}

	//for each email, create a userRequest
	for _, userRequest := range users {
		_, err = client.userClient.Create(nil, &userRequest)
		//if error occurs, return the error
		if err != nil {
			//if user already exists, ignore the error
			if strings.Contains(err.Error(), "already exists") {
				err = nil
				continue
			}
			return "", err
		}
	}

	//for each user, create a team.Member
	members := make([]team.Member, len(users))
	for i, userRequest := range users {
		members[i] = *getMember(userRequest)
	}

	teamResult, err := client.teamClient.Create(nil, &team.CreateTeamRequest{
		Name:        emails[0][:strings.Index(emails[0], "@")],
		Description: "тима с" + emails[0],
		Members:     members,
	})

	return teamResult.Id, err
}

//CreateIncident creates an incident
func (client *OpsGenieClient) CreateIncident(incident *models.Incident, teamId string) error {
	_, err := client.alertClient.Create(nil, &alert.CreateAlertRequest{
		Message:  incident.Description,
		Source:   "source2",
		Priority: alert.P1,

		Responders: []alert.Responder{
			{Type: alert.TeamResponder, Id: teamId},
		},
	})
	return err
}

func getMember(user user.CreateRequest) *team.Member {
	return &team.Member{
		User: team.User{
			Username: user.Username,
		},
		Role: user.Role.RoleName,
	}
}

func getUser(email string) *user.CreateRequest {
	userRes := &user.CreateRequest{
		//take string before @
		Username: email,
		FullName: email[:strings.Index(email, "@")],
		Role: &user.UserRoleRequest{
			RoleName: "user",
		},
	}
	return userRes
}
