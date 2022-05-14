package clients

import (
	"database/sql"
	"emergence/commandsAndQueries/models"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

//IPostgresClient interface
type IPostgresClient interface {
	AddTeam(id string, city string) error
	GetTeamByCity(city string) (team *models.Team, err error)
	GetCityAndAddressByExternalUserId(externalUserId string) (address string, city string, err error)
	SetTeamIsBlockedFalse(teamId string) error
}

//PostgresClient is a client for Postgres implementation of IPostgresClient
type PostgresClient struct {
	db *sql.DB
}

// NewPostgresClient returns a new PostgresClient with the given connection string
func NewPostgresClient() *PostgresClient {
	connectionString := "postgres://ttuzpbnnzexpnh:3aae25a4574b6a3a47781a0afe4c306a0461d9e7010cad46d69a9ba3933ffea6@ec2-52-18-116-67.eu-west-1.compute.amazonaws.com:5432/df82hje0i8jjqa"
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

// AddTeam adds a team to the database and city to the database
func (p *PostgresClient) AddTeam(id string, city string) error {
	_, err := p.db.Exec("INSERT INTO teams (id, city, false) VALUES ($1, $2)", id, city)
	if err != nil {
		return err
	}
	return nil
}

// GetTeamByCity returns a team from the database by city
func (p *PostgresClient) GetTeamByCity(city string) (team *models.Team, err error) {
	var id string
	var teamCity string
	err = p.db.QueryRow("SELECT id, city, is_blocked FROM teams WHERE city = $1 AND is_blocked = false", city).Scan(&id, &teamCity)
	//if id is empty, then take the first team from the database
	if id == "" {
		err = p.db.QueryRow("SELECT id, city, is_blocked FROM teams WHERE city = $1", city).Scan(&id, &teamCity)
	}
	// set selected team is_blocked true
	_, err = p.db.Exec("UPDATE teams SET is_blocked = true WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &models.Team{Id: id, City: teamCity}, nil
}

// GetCityAndAddressByExternalUserId returns the city and address of the user by external user id
func (p *PostgresClient) GetCityAndAddressByExternalUserId(externalUserId string) (address string, city string, err error) {
	err = p.db.QueryRow("SELECT address, city FROM users WHERE external_user_id = $1", externalUserId).Scan(&address, &city)
	if err != nil {
		return "", "", err
	}
	return address, city, nil
}

// SetTeamIsBlockedFalse sets the team is_blocked to false
func (p *PostgresClient) SetTeamIsBlockedFalse(teamId string) error {
	_, err := p.db.Exec("UPDATE teams SET is_blocked = false WHERE id = $1", teamId)
	if err != nil {
		return err
	}
	return nil
}
