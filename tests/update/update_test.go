package update

import (
	"github.com/1franck/cvepack/internal/config"
	"github.com/1franck/cvepack/internal/update"
	"github.com/h2non/gock"
	"net/http"
	"testing"
)

var testConfig = config.Config{
	DatabaseRootDir: "./_fixtures",
}

func Test_Update_IsNeeded_HappyPath(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checksum").
		Reply(200).
		BodyString("6b466ca3ac976d32e380e072f461f7ac38ca528a788dbd37587965e93aa08e4d")

	// create http server with one route
	http.ListenAndServe("", nil)

	conf := config.FromDefault(testConfig)
	var updateNeeded, reason = update.IsNeeded(conf)
	var expectedReason = "Database is up to date"

	if updateNeeded {
		t.Errorf("Should not need update: got '%s'", reason)
	}

	if reason != expectedReason {
		t.Errorf("Reason should be '%s', got: %s", expectedReason, reason)
	}
}
