package core

import (
	"insomnia/internal/db"
	"log"
)

type EndpointRequest struct {
	// EndpointId
}

// Updates all the endpoints with last_run and next_run as NULL
func UpdateUnmonitoredEndpoints() {
	// Get all the endpoints with the last_run and next_run as NULL and update them
	query := `
	UPDATE TARGET
	SET
		LAST_RUN = NOW(),
		NEXT_RUN = NOW() + "interval"
	WHERE
		LAST_RUN IS NULL AND
		NEXT_RUN IS NULL
	`
	db := db.DB
	_, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	log.Println("updated unmonitored endpoints ...")
}

func GetMonitoredEndpoints() []db.Target {
	/**
	1) Get all the endpoints from the target table - Now that everything is updated
	*/

	return db.GetTargetDB()
}
