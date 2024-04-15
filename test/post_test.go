package test

import (
	"encoding/json"
	"testing"

	"github.com/levigross/grequests"
	"github.com/stretchr/testify/require"
)

func TestPost(t *testing.T) {
	// t.Skip()
	require := require.New(t)
	bodystr := `
	{
		"name": "inventory-connector",  
		"config": {  
		  "connector.class": "io.debezium.connector.mysql.MySqlConnector",
		  "tasks.max": "1",  
		  "database.hostname": "mysql",  
		  "database.port": "3306",
		  "database.user": "debezium",
		  "database.password": "dbz",
		  "database.server.id": "184054",  
		  "topic.prefix": "dbserver1",  
		  "database.include.list": "inventory",  
		  "schema.history.internal.kafka.bootstrap.servers": "kafka:9092",  
		  "schema.history.internal.kafka.topic": "schema-changes.inventory"  
		}
	}
	`

	var body map[string]interface{}
	json.Unmarshal([]byte(bodystr), &body)
	res, err := grequests.Post("http://localhost:8083/connectors/", &grequests.RequestOptions{
		JSON: body,
	})
	require.NoError(err)
	// require.Equal(201, res.StatusCode)
	t.Log(res.StatusCode)

	t.Log(res.String())

}
