package repository

import (
	"context"
	"time"
)

type ActivityRepo interface {
	GetAllActivities(
		ctx context.Context, req GetAllActivitiesRequest) (
		res GetAllActivitiesResponse, httpcode int, err error,
	)
	CreateActivity(
		ctx context.Context, req CreateActivityRequest) (
		res CreateActivityResponse, httpcode int, err error,
	)
	GetOneActivityByID(
		ctx context.Context, req GetOneActivityByIDRequest) (
		res GetOneActivityByIDResponse, httpcode int, err error,
	)
	UpdateOneActivityByID(
		ctx context.Context, req UpdateOneActivityByIDRequest) (
		res UpdateOneActivityByIDResponse, httpcode int, err error,
	)
	DeleteOneActivityByID(
		ctx context.Context, req DeleteOneActivityByIDRequest) (
		res DeleteOneActivityByIDResponse, httpcode int, err error,
	)
	GetLatesActivityID(
		ctx context.Context, req GetLatesActivityIDRequest) (
		res GetLatesActivityIDResponse, httpcode int, err error,
	)
}

type GetAllActivitiesRequest struct{}

type GetAllActivitiesData struct {
	ID        int64
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
type GetAllActivitiesResponse []GetAllActivitiesData

type CreateActivityRequest struct {
	Title string
	Email string
}

type CreateActivityResponse struct {
	ID        int64
	Title     string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetOneActivityByIDRequest struct {
	ID int64
}
type GetOneActivityByIDResponse GetAllActivitiesData

type UpdateOneActivityByIDRequest struct {
	ID    int64
	Title string
}

type UpdateOneActivityByIDResponse GetAllActivitiesData

type DeleteOneActivityByIDRequest struct {
	ID int64
}
type DeleteOneActivityByIDResponse DeleteOneActivityByIDRequest

type GetLatesActivityIDRequest struct{}
type GetLatesActivityIDResponse struct{ ID int64 }
