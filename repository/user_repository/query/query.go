package query

import _ "embed"

var (
	//go:embed create_schema.sql
	CreateSchema string
	//go:embed create_table.sql
	CreateTable string

	//go:embed activity/get_all_activity.sql
	GetAllActivities string
	//go:embed activity/create_activity.sql
	CreateActivity string
	//go:embed activity/get_one_by_id.sql
	GetOneActivityByID string
	//go:embed activity/update_one_by_id.sql
	UpdateOneActivityByID string
	//go:embed activity/delete_one_by_id.sql
	DeleteOneActivityByID string
)
