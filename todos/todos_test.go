package todos

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

const MONGO_INITDB_ROOT_USERNAME = "root"
const MONGO_INITDB_ROOT_PASSWORD = "password"

func TestMain(m *testing.M) {
	// Setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=" + MONGO_INITDB_ROOT_USERNAME,
		"MONGO_INITDB_ROOT_PASSWORD=" + MONGO_INITDB_ROOT_PASSWORD,
	}
	resource, err := pool.Run("mongo", "5.0", environmentVariables)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	err = pool.Retry(func() error {
		var err error
		db, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://%s:%s@localhost:%s", MONGO_INITDB_ROOT_USERNAME, MONGO_INITDB_ROOT_PASSWORD, resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		return db.Ping(context.TODO(), nil)
	})
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// seed data

	// Run tests
	exitCode := m.Run()

	// Teardown
	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// Exit
	os.Exit(exitCode)
}
