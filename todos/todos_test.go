package todos

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Setup
	// ...

	// Run tests
	exitCode := m.Run()

	// Teardown
	// ...

	// Exit
	os.Exit(exitCode)
}
