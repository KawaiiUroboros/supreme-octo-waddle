package integration

import (
	commandModels "emergence/commandsAndQueries/models"
	"emergence/integration/clients"
	repoModels "emergence/integration/models"
	"github.com/ahmetb/go-linq/v3"
	"log"
)

//IRepository interface for creating Slack channels and storing them in the database
type IRepository interface {
	CreateChannelsByUserIdsAndIntervals(newChatOptions *[]commandModels.NewChatOptions) (userIdToChannelId *map[string]string, err error)
	DeleteChats(externalUserIds *[]string) (deletedChatIds *[]string, err error)
	// NotifyChats sending notifications to chats that have exceeded the interval
	NotifyChats() (err error)
}

// Repository implementation of the IRepository interface for the database and Slack
type Repository struct {
	db clients.IPostgresClient
	sl clients.ISlackClient
}

//NewRepository creates a new Repository
func NewRepository(db clients.IPostgresClient, sl clients.ISlackClient) *Repository {
	return &Repository{db: db, sl: sl}
}

// CreateChannelsByUserIdsAndIntervals creates a channels for the users with the given userIDs and store the channelIds in the database with the given intervals and user ids
func (r *Repository) CreateChannelsByUserIdsAndIntervals(newChatOptions *[]commandModels.NewChatOptions) (userIdToChannelId *map[string]string, err error) {
	var ExternalUserIdChannelNames []string
	linq.From(*newChatOptions).SelectT(func(newChatOption commandModels.NewChatOptions) string {
		return newChatOption.ExternalUserId
	}).Distinct().ToSlice(&ExternalUserIdChannelNames)
	userIdToChannelIdMap, err := r.sl.CreateChannels(&ExternalUserIdChannelNames)
	if err != nil {
		return nil, err
	}

	//store the channelIds, externalUserIds, intervals in the database and return the map
	var activeChannels []repoModels.ActiveChannel
	linq.From(*newChatOptions).SelectT(func(newChatOption commandModels.NewChatOptions) repoModels.ActiveChannel {
		return repoModels.ActiveChannel{
			ChannelId:      (*userIdToChannelIdMap)[newChatOption.ExternalUserId],
			ExternalUserId: newChatOption.ExternalUserId,
			Interval:       newChatOption.Interval,
		}
	}).ToSlice(&activeChannels)
	err = r.db.AddSlackChannelForUserToDb(&activeChannels)
	if err != nil {
		return nil, err
	}
	return userIdToChannelIdMap, nil
}

// DeleteChats deletes the chats for the given externalUserIds
func (r *Repository) DeleteChats(externalUserIds *[]string) (deletedChatIds *[]string, err error) {
	var activeChannels []repoModels.ActiveChannel
	err = r.db.GetActiveChannelsByExternalUserIds(externalUserIds, &activeChannels)
	if err != nil {
		return nil, err
	}
	var channelIds []string
	linq.From(activeChannels).SelectT(func(activeChannel repoModels.ActiveChannel) string {
		//return activeChannel.ChannelId
		return activeChannel.ChannelId
	}).ToSlice(&channelIds)
	err = r.sl.DeleteChannels(&channelIds)
	if err != nil {
		return nil, err
	}
	err = r.db.DeleteActiveChannelsByExternalUserIds(externalUserIds)
	if err != nil {
		return nil, err
	}
	return &channelIds, nil
}

// NotifyChats sends notifications to the chats that have exceeded the interval every minute
func (r *Repository) NotifyChats() (err error) {
	var activeChannels []repoModels.ActiveChannel
	err = r.db.GetActiveChannels(&activeChannels)
	//if activeChannels is empty, log the error
	if len(activeChannels) == 0 {
		log.Println("No active channels found")
	}
	if err != nil {
		return err
	}
	var channelIds []string
	linq.From(activeChannels).SelectT(func(activeChannel repoModels.ActiveChannel) string {
		return activeChannel.ChannelId
	}).ToSlice(&channelIds)
	err = r.sl.NotifyChannels(&channelIds)
	if err != nil {
		log.Println(err)
		return err
	}
	//for active channel, update the last_confirmation field
	err = r.db.UpdateBeginDateForActiveChannels(&activeChannels)
	if err != nil {
		return err
	}
	return nil
}
