package database

import (
	"github.com/joshblack/howto-sql-golang/models"
	"net/url"
)

type Store interface {
	GetTeams() ([]*models.Team, error)
	GetTeam(id int) (*models.Team, error)
	CreateTeam(*models.Team) (*models.Team, error)
	UpdateTeam(id int, team *models.Team) (*models.Team, error)
	DeleteTeam(id int) error
}

func NewConnString(user, password, addr, database string) string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   addr,
		Path:   database,
	}

	v := url.Values{}
	v.Set("sslmode", "disable")
	v.Set("connect_timeout", "10")
	u.RawQuery = v.Encode()

	return u.String()
}
