package usecases

import "context"

type ActivityUseCase interface {
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
}

type GetAllActivitiesRequest struct {
}

type GetAllActivitiesData struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GetAllActivitiesResponse []GetAllActivitiesData

type CreateActivityRequest struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type CreateActivityResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type GetOneActivityByIDRequest struct {
	ID int64 `json:"-"`
}
type GetOneActivityByIDResponse GetAllActivitiesData

type UpdateOneActivityByIDRequest struct {
	ID    int64  `json:"-"`
	Title string `json:"title"`
}
type UpdateOneActivityByIDResponse GetAllActivitiesData

type DeleteOneActivityByIDRequest struct{ ID int64 }
type DeleteOneActivityByIDResponse DeleteOneActivityByIDRequest
