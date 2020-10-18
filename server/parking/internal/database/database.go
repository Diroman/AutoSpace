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
	insertFrame  = `UPDATE cams SET image = '%s', free_space = %v WHERE id = %v;`
	getFreeSpace = `SELECT coord
					FROM parking_fields
					 JOIN cams c on c.id = parking_fields.cam_no
					 JOIN users on c.address_id = users.address_id
					WHERE users.id = %v;`
	updateFreeSpace = `UPDATE parking_fields
						SET coord = '{%s}'
						WHERE cam_no IN (SELECT cams.id
						FROM cams
								 JOIN users u on cams.address_id = u.address_id
						WHERE u.id = %v);`
	getCameraInfo = `SELECT hight, gps_coord_x, gps_coord_y, horizon_angle
						FROM cams
						WHERE id = %v;`
	getCamerasSpace = `SELECT max(c.gps_coord_x) as lat,
       max(c.gps_coord_y) as long,
       max(c.free_space)  as free,
       CASE
           WHEN max(pf.coord) IS NOT NULL THEN array_length(max(pf.coord), 1)
           ELSE 0
           END            as total
		FROM cams c
         LEFT JOIN users on c.address_id = users.address_id
         LEFT JOIN parking_fields pf on c.id = pf.cam_no
		WHERE users.id = %v
		GROUP BY c.id;`
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
	for _, val := range spaces {
		params = append(params, fmt.Sprintf("%v", val))
	}

	query := fmt.Sprintf(updateFreeSpace, strings.Join(params, ", "), id)
	_, err := d.Conn.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (d Database) GetCameraInfo(id int) (model.CameraInfo, error) {
	query := fmt.Sprintf(getCameraInfo, id)
	rows, err := d.Conn.Query(query)
	if err != nil {
		return model.CameraInfo{}, err
	}

	defer rows.Close()

	has := rows.Next()
	if !has {
		return model.CameraInfo{}, NotFound
	}

	var camera model.CameraInfo
	err = rows.Scan(&camera.Height, &camera.Latitude, &camera.Longitude, &camera.HorizonAngle)
	if err != nil {
		return model.CameraInfo{}, err
	}

	return camera, nil
}

//func (d Database) GetParkingSpace(id int) ([]float64, error) {
//	query := fmt.Sprintf(getFreeSpace, id)
//	rows, err := d.Conn.Query(query)
//	if err != nil {
//		return nil, err
//	}
//
//	defer rows.Close()
//
//	has := rows.Next()
//	if !has {
//		return nil, NotFound
//	}
//
//	var spaces []float64
//	err = rows.Scan(&spaces)
//	if err != nil {
//		return nil, err
//	}
//
//	return spaces, nil
//}

func (d Database) GetParkingSpace(id int) ([]model.Space, error) {
	query := fmt.Sprintf(getCamerasSpace, id)
	rows, err := d.Conn.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var spaces []model.Space
	for rows.Next() {
		space := model.Space{}
		err = rows.Scan(&space.Lat, &space.Long, &space.Free, &space.Total)
		if err != nil {
			return nil, err
		}

		spaces = append(spaces, space)
	}

	return spaces, nil
}
