package models

// ActiveChannel struct for active channel consist of channel id user id and interval to send messages
type ActiveChannel struct {
	ChannelId      string
	ExternalUserId string
	Interval       int
}
