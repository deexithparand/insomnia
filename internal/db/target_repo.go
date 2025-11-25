package db

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TargetRepo struct {
	ID       uuid.UUID     `json:"id" db:"id"`
	GroupID  uuid.UUID     `json:"groupid" db:"groupid"`
	Endpoint string        `json:"endpoint" db:"endpoint"`
	Interval time.Duration `json:"interval" db:"interval"`
	LastRun  *time.Time    `json:"last_run" db:"last_run"`
	NextRun  *time.Time    `json:"next_run" db:"next_run"`
}

func GetTargetDB() []TargetRepo {

	var (
		targetRepoList    []TargetRepo
		intervalInSeconds float64
	)

	targetTableSelectQuery := `
	SELECT
		id,
		groupid,
		endpoint,
		EXTRACT(EPOCH FROM interval) AS interval_seconds,
		last_run,
		next_run
	FROM target;
	`

	rows, err := DB.Query(targetTableSelectQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var target TargetRepo
		fmt.Println(target.ID)
		err := rows.Scan(&target.ID, &target.GroupID, &target.Endpoint, &intervalInSeconds, &target.LastRun, &target.NextRun)
		if err != nil {
			panic(err)
		}

		target.Interval = time.Duration(intervalInSeconds) * time.Second

		targetRepoList = append(targetRepoList, target)

		fmt.Println(target.ID, target.GroupID, target.Endpoint, target.Interval, target.LastRun, target.NextRun)
	}

	fmt.Println("Queried data successfully from the target group")

	// return targetRepoList
	return []TargetRepo{}
}
