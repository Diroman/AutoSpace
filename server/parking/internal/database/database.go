package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"parking/model"
)

const (
	getUserTemplate = `SELECT (second_name || ' ' || first_name || ' ' || middle_name) as name,
       mobile,
       email,
       password,
       a.address_lat,
       a.address_long,
       (city || ', ' || street || ', '
            || cast(house as VARCHAR) || CASE
                            WHEN litera IS NULL AND campus = 0 AND building = 0 THEN ''
                            WHEN litera IS NOT NULL AND campus = 0 AND building = 0 THEN litera
                            WHEN litera IS NOT NULL AND campus <> 0 AND building = 0 THEN litera || ', ' || cast(campus as VARCHAR)
                            WHEN litera IS NOT NULL AND campus <> 0 AND building <> 0 THEN litera || ', ' || cast(campus as VARCHAR) || ', ' || cast(building as VARCHAR)
                            WHEN litera IS NULL AND campus <> 0 AND building <> 0 THEN ', ' || cast(campus as VARCHAR) || ', ' || cast(building as VARCHAR)
                            WHEN litera IS NULL AND campus = 0 AND building <> 0 THEN ', ' || cast(building as VARCHAR)
                            WHEN litera IS NOT NULL AND campus = 0 AND building <> 0 THEN ', ' || campus
            END)                                    as address
FROM users
         JOIN addresses a on a.id = users.address_id
WHERE users.email = %s;`
)

type Database struct {
	host string
	port int
	Conn *pgx.Conn
}

func NewDatabase(host string, port int, database, user, password string) *Database {
	config := pgx.ConnConfig{
		Host:     host,
		Port:     uint16(port),
		Database: database,
		User:     user,
		Password: password,
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		panic(err)
	}

	return &Database{
		host: host,
		port: port,
		Conn: conn,
	}
}

func (d Database) GetUser(email string) (model.User, error) {
	user := model.User{}

	query := fmt.Sprintf(getUserTemplate, email)
	rows, err := d.Conn.Query(query)
	if err != nil {
		return model.User{}, err
	}

	defer rows.Close()

	has := rows.Next()
	if !has {
		return model.User{}, UserNotFound
	}

	err = rows.Scan(&user.Username, &user.Mobile, &user.Email, &user.Password,
		&user.AddressCoord.Latitude, &user.AddressCoord.Longitude, &user.Address)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
