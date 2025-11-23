package core

import (
	"insomnia/internal"
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
	DB := internal.DB
	_, err := DB.Query(query)
	if err != nil {
		panic(err)
	}

	log.Println("updated unmonitored endpoints ...")
}

func GetMonitoredEndpoints() []EndpointRequest {
	/**
	1) Get all the endpoints from the target table
	*/

	return []EndpointRequest{}
}
