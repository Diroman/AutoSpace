package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"parking/model"
	"strings"
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
		WHERE users.email = '%s';`
	getCameraImage = `SELECT image
						FROM cams
						WHERE id = %v;`
	getAllCameras = `SELECT cams.id, (city || ', ' || street || ', д.'
                || cast(house as VARCHAR) || CASE
                         WHEN litera IS NULL AND campus = 0 THEN ''
                         WHEN litera IS NOT NULL AND campus = 0 THEN litera
                         WHEN litera IS NOT NULL AND campus <> 0 THEN litera || ', к.' || cast(campus as VARCHAR)
                         WHEN litera IS NULL AND campus <> 0 THEN ', ' || cast(campus as VARCHAR)
                END)                                    as address
				FROM cams JOIN addresses a on a.id = cams.address_id;`
	insertFrame = `UPDATE cams SET image = '%s', free_space = %v WHERE id = %v;`
	getFreeSpace = `SELECT coord
					FROM parking_fields
					 JOIN cams c on c.id = parking_fields.cam_no
					 JOIN users on c.address_id = users.address_id
					WHERE users.id = %v;`
	updateFreeSpace = `UPDATE parking_fields
						SET coord = ARRAY [%s]
						WHERE cam_no = (SELECT cams.id
						FROM cams
								 JOIN users u on cams.address_id = u.address_id
						WHERE u.id = %v);`
	//getUserInfo = ``
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

func (d Database) GetAllCameras() (model.Cameras, error) {
	rows, err := d.Conn.Query(getAllCameras)
	if err != nil {
		return model.Cameras{}, err
	}

	defer rows.Close()

	cameras := model.Cameras{}
	for rows.Next() {
		camera := model.Camera{}
		err = rows.Scan(&camera.Id, &camera.Address)
		if err != nil {
			return model.Cameras{}, err
		}

		cameras.Cameras = append(cameras.Cameras, camera)
	}

	return cameras, nil
}

func (d Database) GetFrame(id int) (string, error) {
	query := fmt.Sprintf(getCameraImage, id)
	rows, err := d.Conn.Query(query)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	has := rows.Next()
	if !has {
		return "", NotFound
	}

	var photo string
	err = rows.Scan(&photo)
	if err != nil {
		return "", err
	}

	return photo, nil
}

func (d Database) SaveFrame(id, count int, frame string) error {
	query := fmt.Sprintf(insertFrame, frame, count, id)
	_, err := d.Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (d Database) UpdateFreeSpace(id int, spaces []interface{}) error {
	params := make([]string, 0, len(spaces))
	for i := range spaces {
		params = append(params, fmt.Sprintf("$%v", i+1))
	}

	query := fmt.Sprintf(updateFreeSpace, strings.Join(params, ", "), id)
	_, err := d.Conn.Exec(query, spaces...)
	if err != nil {
		return err
	}

	return nil
}

func (d Database) GetParkingSpace(id int) ([]float64, error) {
	query := fmt.Sprintf(getFreeSpace, id)
	rows, err := d.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	has := rows.Next()
	if !has {
		return nil, NotFound
	}

	var spaces []float64
	err = rows.Scan(&spaces)
	if err != nil {
		return nil, err
	}

	return spaces, nil
}