package activity_repo

import (
	"coding-test-be/repository"
	"coding-test-be/repository/activity_repository/query"
	"context"
	"log"
	"net/http"
	"time"
)

func (x *PostgreSQLConn) GetAllActivities(
	ctx context.Context, req repository.GetAllActivitiesRequest) (
	res repository.GetAllActivitiesResponse, httpcode int, err error,
) {
	query := query.GetAllActivities
	args := List{}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.GetAllActivitiesData{}
		err := rows.Scan(&data.ID, &data.Title, &data.Email, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		res = append(res, data)
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) CreateActivity(
	ctx context.Context, req repository.CreateActivityRequest) (
	res repository.CreateActivityResponse, httpcode int, err error,
) {
	var id int64
	createdTime := time.Now()

	query := query.CreateActivity
	args := List{
		req.Title,
		req.Email,
		createdTime,
		createdTime,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
		}
	}

	res = repository.CreateActivityResponse{
		ID:        id,
		Title:     req.Title,
		Email:     req.Email,
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) GetOneActivityByID(
	ctx context.Context, req repository.GetOneActivityByIDRequest) (
	res repository.GetOneActivityByIDResponse, httpcode int, err error,
) {
	query := query.GetOneActivityByID
	args := List{
		req.ID,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		data := repository.GetOneActivityByIDResponse{}
		err := rows.Scan(&data.ID, &data.Title, &data.Email, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			log.Println(err)
		}
		res = data
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) UpdateOneActivityByID(
	ctx context.Context, req repository.UpdateOneActivityByIDRequest) (
	res repository.UpdateOneActivityByIDResponse, httpcode int, err error,
) {
	var (
		email       string
		createdTime time.Time
	)

	updatedTime := time.Now()

	query := query.UpdateOneActivityByID
	args := List{
		req.ID,
		req.Title,
		updatedTime,
	}

	rows, err := x.tc.Query(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&email, &createdTime)
		if err != nil {
			log.Println(err)
		}
	}

	res = repository.UpdateOneActivityByIDResponse{
		ID:        req.ID,
		Title:     req.Title,
		Email:     email,
		CreatedAt: createdTime,
		UpdatedAt: updatedTime,
	}

	return res, httpcode, err
}

func (x *PostgreSQLConn) DeleteOneActivityByID(
	ctx context.Context, req repository.DeleteOneActivityByIDRequest) (
	res repository.DeleteOneActivityByIDResponse, httpcode int, err error,
) {
	query := query.DeleteOneActivityByID
	args := List{
		req.ID,
	}

	_, err = x.tc.Exec(query, args...)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	res = repository.DeleteOneActivityByIDResponse{ID: req.ID}

	return res, httpcode, err
}
