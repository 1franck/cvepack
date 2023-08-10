package update

import (
	"github.com/1franck/cvepack/internal"
	"github.com/1franck/cvepack/internal/config"
	"github.com/1franck/cvepack/internal/update"
	"github.com/h2non/gock"
	"net/http"
	"strings"
	"testing"
)

type IsNeededTestCase struct {
	UpdateNeeded           bool
	Reason                 internal.ErrorMsg
	ExpectedUpdateNeeded   bool
	ExpectedReason         internal.ErrorMsg
	ExpectedReasonContains string
}

func assertUpdateNeeded(t *testing.T, testCase IsNeededTestCase) {
	if testCase.ExpectedUpdateNeeded {
		if !testCase.UpdateNeeded {
			t.Errorf("Should need update: got '%s'", testCase.Reason)
		}
	} else {
		if testCase.UpdateNeeded {
			t.Errorf("Should not need update: got '%s'", testCase.Reason)
		}
	}

	if testCase.ExpectedReasonContains == "" {
		if testCase.Reason != testCase.ExpectedReason {
			t.Errorf("Reason should be '%s', got: %s", testCase.ExpectedReason, testCase.Reason)
		}
	} else {
		if !strings.Contains(testCase.Reason.ToString(), testCase.ExpectedReasonContains) {
			t.Errorf("Reason should contain '%s', got: %s", testCase.ExpectedReasonContains, testCase.Reason)
		}
	}
}

func Test_Update_IsNeeded_HappyPath(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checksum").
		Reply(200).
		BodyString("6b466ca3ac976d32e380e072f461f7ac38ca528a788dbd37587965e93aa08e4d")

	// create http server with one route
	http.ListenAndServe("", nil)

	conf := config.FromDefault(config.Config{
		DatabaseRootDir: "./_fixtures",
	})

	var updateNeeded, reason = update.IsNeeded(conf)
	assertUpdateNeeded(t, IsNeededTestCase{
		UpdateNeeded:         updateNeeded,
		Reason:               reason,
		ExpectedUpdateNeeded: false,
		ExpectedReason:       internal.EmptyError,
	})
}

func Test_Update_IsNeeded_DbFolderNotFound(t *testing.T) {
	conf := config.Config{
		DatabaseRootDir: "./unknown_folder",
	}
	var updateNeeded, reason = update.IsNeeded(conf)

	assertUpdateNeeded(t, IsNeededTestCase{
		UpdateNeeded:         updateNeeded,
		Reason:               reason,
		ExpectedUpdateNeeded: true,
		ExpectedReason:       update.ErrorDatabaseFolderNotFound,
	})
}

func Test_Update_IsNeeded_DbFileNotFound(t *testing.T) {
	conf := config.FromDefault(config.Config{
		DatabaseRootDir:  "./_fixtures",
		DatabaseFileName: "./unknown.db",
	})
	var updateNeeded, reason = update.IsNeeded(conf)

	assertUpdateNeeded(t, IsNeededTestCase{
		UpdateNeeded:         updateNeeded,
		Reason:               reason,
		ExpectedUpdateNeeded: true,
		ExpectedReason:       update.ErrorDatabaseFileNotFound,
	})
}

func Test_Update_IsNeeded_DbChecksumFileNotFound(t *testing.T) {
	conf := config.FromDefault(config.Config{
		DatabaseRootDir:          "./_fixtures",
		DatabaseChecksumFileName: "./unknown.checksum",
	})
	var updateNeeded, reason = update.IsNeeded(conf)

	assertUpdateNeeded(t, IsNeededTestCase{
		UpdateNeeded:         updateNeeded,
		Reason:               reason,
		ExpectedUpdateNeeded: true,
		ExpectedReason:       update.ErrorDatabaseChecksumFileNotFound,
	})
}

func Test_Update_IsNeeded_DatabaseServerChecksumFileInvalid(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.DisableNetworking()
	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checkssum").
		Reply(500).
		BodyString("")

	// create http server with one route
	http.ListenAndServe("", nil)

	conf := config.FromDefault(config.Config{
		DatabaseRootDir: "./_fixtures",
	})
	var updateNeeded, reason = update.IsNeeded(conf)

	assertUpdateNeeded(t, IsNeededTestCase{
		UpdateNeeded:           updateNeeded,
		Reason:                 reason,
		ExpectedUpdateNeeded:   false,
		ExpectedReasonContains: "error checking server database checksum:",
	})
}

func Test_Update_IsNeeded_DatabaseServerChecksumMismatch(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checksum").
		Reply(200).
		BodyString("6b466ca3ac976d32e380e072f461f7ac38ca528a788dbd37587965e93aa08e4d")

	// create http server with one route
	http.ListenAndServe("", nil)

	conf := config.FromDefault(config.Config{
		DatabaseRootDir:          "./_fixtures",
		DatabaseChecksumFileName: "db.checksum.wrong",
	})
	var updateNeeded, reason = update.IsNeeded(conf)

	assertUpdateNeeded(t, IsNeededTestCase{
		UpdateNeeded:         updateNeeded,
		Reason:               reason,
		ExpectedUpdateNeeded: true,
		ExpectedReason:       update.ErrorDatabaseChecksumMismatch,
	})
}
