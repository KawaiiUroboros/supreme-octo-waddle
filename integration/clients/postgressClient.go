package clients

import (
	"database/sql"
	"emergence/integration/models"
	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/pkg/errors"
	"log"
	"time"
)

//IPostgresClient interface
type IPostgresClient interface {
	AddSlackChannelForUserToDb(activeChannels *[]models.ActiveChannel) error
	CloseConnection() error
	// GetActiveChannelsByExternalUserIds GetAllActiveChannels() ([]models.ActiveChannel, error)
	GetActiveChannelsByExternalUserIds(externalUserIds *[]string, i *[]models.ActiveChannel) error
	//DeleteActiveChannelsByExternalUserIds DeleteActiveChannelByExternalUserIds(externalUserIds *[]string) error
	DeleteActiveChannelsByExternalUserIds(externalUserIds *[]string) error
	// GetActiveChannels GetAllActiveChannels() ([]models.ActiveChannel, error)
	GetActiveChannels(activeChannels *[]models.ActiveChannel) error
	// UpdateBeginDateForActiveChannels for active channel, update the last_confirmation field
	UpdateBeginDateForActiveChannels(channels *[]models.ActiveChannel) error
}

//PostgresClient is a client for Postgres implementation of IPostgresClient
type PostgresClient struct {
	db *sql.DB
}

// AddSlackChannelForUserToDb adds a Slack channel for a user implementation of IPostgresClient
func (p *PostgresClient) AddSlackChannelForUserToDb(activeChannels *[]models.ActiveChannel) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Error while starting transaction")
		return err
	}
	//roolback if error
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	for _, channel := range *activeChannels {
		//variable that contains the timestamp of current time now
		now := time.Now()
		//if channel id is empty then return error channel id cannot be empty
		if channel.ChannelId == "" {
			return errors.New("channel id cannot be empty")
		}

		//add channel to db with external_user_id, interval date_begin timestamp with timezone as date now if external_user_id is not in db
		_, err = tx.Exec("INSERT INTO users (external_user_id, begin_date, channel_id, notification_interval, last_confirmation, is_deleted) VALUES($1,$2,$3,$4,$5,$6) ON CONFLICT ON CONSTRAINT users_channel_id_key DO NOTHING",
			channel.ExternalUserId,
			now,
			channel.ChannelId,
			channel.Interval,
			now,
			false)

		if err != nil {
			log.Println("Error while inserting channel")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error while committing transaction")
		return err
	}

	return nil
}

// NewPostgresClient returns a new PostgresClient with the given connection string
func NewPostgresClient() *PostgresClient {
	connectionString := "postgres://uhjxadtqnmvedz:827b87d9c520ea5e6c07d04522a1efa923be6740fe2332174bc7511532e52e0a@ec2-63-35-156-160.eu-west-1.compute.amazonaws.com:5432/d5qq14sdoip2f"
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresClient{db: db}
}

// CloseConnection closes the database connection
func (p *PostgresClient) CloseConnection() error {
	err := p.db.Close()
	if err != nil {
		return err
	}
	return nil
}

// GetActiveChannelsByExternalUserIds returns all active channels for a user
func (p *PostgresClient) GetActiveChannelsByExternalUserIds(externalUserIds *[]string, activeChannels *[]models.ActiveChannel) error {
	for _, externalUserId := range *externalUserIds {
		rows, err := p.db.Query("SELECT channel_id, notification_interval FROM users WHERE external_user_id=$1", externalUserId)
		if err != nil {
			log.Println("Error while querying for active channels")
			return err
		}
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {
				log.Println("Error while closing rows")
			}
		}(rows)

		for rows.Next() {
			var channelId string
			var interval int
			err = rows.Scan(&channelId, &interval)
			if err != nil {
				log.Println("Error while scanning for active channels")
				return err
			}
			*activeChannels = append(*activeChannels, models.ActiveChannel{ChannelId: channelId, Interval: interval})
		}
	}
	return nil
}

// DeleteActiveChannelsByExternalUserIds deletes all active channels for a user where flag is_deleted is true
func (p *PostgresClient) DeleteActiveChannelsByExternalUserIds(externalUserIds *[]string) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Error while starting transaction")
		return err
	}
	//rollback if error
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()
	for _, externalUserId := range *externalUserIds {
		_, err = tx.Exec("DELETE FROM users WHERE external_user_id=$1", externalUserId)
		if err != nil {
			log.Println("Error while deleting active channels")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error while committing transaction")
		return err
	}

	return nil
}

// GetActiveChannels channels where the notification interval In minutes exceeds begin_date timestamp and the channel is not deleted
func (p *PostgresClient) GetActiveChannels(activeChannels *[]models.ActiveChannel) error {
	//cast ten minutes to interval in postgres

	rows, err := p.db.Query("SELECT channel_id, notification_interval, begin_date FROM users WHERE is_deleted=false AND make_interval(0,0,0,0,0,notification_interval) < (now() - begin_date)")
	if err != nil {
		log.Println("Error while querying for active channels")
		return err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error while closing rows")
		}
	}(rows)

	for rows.Next() {
		var channelId string
		var interval int
		var lastConfirmation time.Time
		err = rows.Scan(&channelId, &interval, &lastConfirmation)
		if err != nil {
			log.Println("Error while scanning for active channels")
			return err
		}
		*activeChannels = append(
			*activeChannels,
			models.ActiveChannel{
				ChannelId: channelId,
				Interval:  interval})
	}
	return nil
}

//UpdateBeginDateForActiveChannels updates the begin_date timestamp for all active channels
func (p *PostgresClient) UpdateBeginDateForActiveChannels(activeChannels *[]models.ActiveChannel) error {
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("Error while starting transaction")
		return err
	}
	//rollback if error
	defer func() {
		if r := recover(); r != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()
	for _, activeChannel := range *activeChannels {
		_, err = tx.Exec("UPDATE users SET begin_date=now() WHERE channel_id=$1", activeChannel.ChannelId)
		if err != nil {
			log.Println("Error while updating last confirmation for active channels")
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error while committing transaction")
		return err
	}

	return nil
}
