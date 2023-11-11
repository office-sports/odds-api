package data

import (
	"github.com/office-sports/ttapp-api/data/queries"
	"github.com/office-sports/ttapp-api/models"
)

// GetUserByLogin returns user model for requested login
func GetUserByLogin(login string) (*models.User, error) {
	u := new(models.User)
	err := models.DB.QueryRow(queries.GetUserDataQuery(), login).
		Scan(&u.ID, &u.Login, &u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}
