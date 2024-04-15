package test

import (
	"testing"

	"github.com/levigross/grequests"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	require := require.New(t)
	url := "http://localhost:8083/connectors/"
	resp, err := grequests.Get(url, nil)
	require.NoError(err)
	defer resp.Close()
	require.Equal(200, resp.StatusCode)
	t.Log(resp.String())
}

func TestGetTasks(t *testing.T) {
	require := require.New(t)
	url := "http://localhost:8083/connectors/inventory-connector"
	resp, err := grequests.Get(url, nil)
	require.NoError(err)
	defer resp.Close()
	require.Equal(200, resp.StatusCode)
	t.Log(resp.String())
	/*
	{
		"name": "inventory-connector",
		"config": {
			"connector.class": "io.debezium.connector.mysql.MySqlConnector",
			"database.user": "debezium",
			"topic.prefix": "dbserver1",
			"schema.history.internal.kafka.topic": "schema-changes.inventory",
			"database.server.id": "184054",
			"database.hostname": "mysql",
			"tasks.max": "1",
			"database.password": "dbz",
			"name": "inventory-connector",
			"schema.history.internal.kafka.bootstrap.servers": "kafka:9092",
			"database.include.list": "inventory",
			"database.port": "3306"
		},
		"tasks": [
			{
				"connector": "inventory-connector",
				"task": 0
			}
		],
		"type": "source"
	}
	*/
}
