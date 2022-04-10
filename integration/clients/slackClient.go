package clients

import (
	"github.com/slack-go/slack"
	"log"
	"strings"
)

// ISlackClient interface for slack client with slack.Client
type ISlackClient interface {
	// CreateChannels Create public channels using userIds as names for chats
	CreateChannels(channelNames *[]string) (userIdToChannelId *map[string]string, err error)
	//DeleteChannels Delete public channels using userIds as names for chats
	DeleteChannels(channelIds *[]string) error
	// NotifyChannels Send message to channel
	NotifyChannels(channelIds *[]string) error
}

// NewSlackClient creates a new slack client
func NewSlackClient() *SlackClient {
	return &SlackClient{
		client: slack.New(
			"xoxb-3098103420019-3278674908690-Jm5l2HujsRUvoN1qpiA9lIDN",
			slack.OptionDebug(true)),
	}
}

// SlackClient implementation of ISlackClient interface
type SlackClient struct {
	client *slack.Client
}

//CreateChannels Create public channels by iterating over given channel names, where every userId will become a conversation name, and return map of userId to channelId
func (s *SlackClient) CreateChannels(channelNames *[]string) (userIdToChannelId *map[string]string, err error) {
	userIdToChannelId = &map[string]string{}

	for _, userIdAsChannelName := range *channelNames {
		channel, err := s.client.CreateConversation(strings.ToLower(userIdAsChannelName), false)
		if err != nil {
			if err.Error() == "name_taken" {
				// channel already exists, get channelId
				log.Printf("Channel %s already exists", userIdAsChannelName)
				userChannels, _, _ := s.client.GetConversationsForUser(&slack.GetConversationsForUserParameters{
					UserID: userIdAsChannelName,
				})
				// get first channelId from channels if userChannels is not null or empty where channel is not archived and is public and channel name is userIdAsChannelName
				for _, userChannel := range userChannels {
					if !userChannel.IsArchived && userChannel.IsChannel && userChannel.Name == strings.ToLower(userIdAsChannelName) {
						(*userIdToChannelId)[userIdAsChannelName] = userChannel.ID
						break
					}
				}
				continue
			}
			return nil, err
		}
		//add user to a channel
		_, err = s.client.InviteUsersToConversation(channel.ID, userIdAsChannelName)
		if err != nil {
			return nil, err
		}
		(*userIdToChannelId)[userIdAsChannelName] = channel.ID
	}
	return userIdToChannelId, nil
}

//DeleteChannels Delete public channels by iterating over given channel ids, where every userId will become a conversation name, and return map of userId to channelId
func (s *SlackClient) DeleteChannels(channelIds *[]string) error {
	for _, channelId := range *channelIds {
		_, err := s.client.RenameConversation(channelId, strings.ToLower(channelId+"_deleted"))
		if err != nil {
			return err
		}
	}
	return nil
}

// NotifyChannels Send message to channel
func (s *SlackClient) NotifyChannels(channelIds *[]string) error {
	for _, channelId := range *channelIds {
		// send a message with an interactive button
		_, _, err := s.client.PostMessage(channelId, slack.MsgOptionAttachments(
			slack.Attachment{
				Title: "Ежедневное уведомление",
				Fields: []slack.AttachmentField{
					slack.AttachmentField{
						Title: "спасибо что живой",
						Value: "правда спасибо",
						Short: false,
					},
				},
				Actions: []slack.AttachmentAction{
					slack.AttachmentAction{
						Name: "action",
						Text: "Пожалуйста!",
						Type: "button",
						// url with userId as query parameter
						URL: "localhost:3000/GetNotification/" + channelId,
					},
				},
			},
		))

		if err != nil {
			return err
		}
	}
	return nil
}
