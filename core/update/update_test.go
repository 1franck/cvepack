package update

import (
	"cvepack/core"
	"cvepack/core/config"
	"github.com/h2non/gock"
	"net/http"
	"strings"
	"testing"
)

type IsNeededTestCase struct {
	UpdateAvailable         bool
	Reason                  core.ErrorMsg
	ExpectedUpdateAvailable bool
	ExpectedReason          core.ErrorMsg
	ExpectedReasonContains  string
}

func assertUpdateIsAvailable(t *testing.T, testCase IsNeededTestCase) {
	if testCase.ExpectedUpdateAvailable {
		if !testCase.UpdateAvailable {
			t.Errorf("Should need update: got '%s'", testCase.Reason)
		}
	} else {
		if testCase.UpdateAvailable {
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

func Test_Update_IsAvailable_HappyPath(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checksum").
		Reply(200).
		BodyString("6b466ca3ac976d32e380e072f461f7ac38ca528a788dbd37587965e93aa08e4d")

	// create http server with one route
	http.ListenAndServe("127.0.0.1", nil)

	conf := config.FromDefault(config.Config{
		DatabaseRootDir: "./testdata",
	})

	var updateAvailable, reason = IsAvailable(conf)
	assertUpdateIsAvailable(t, IsNeededTestCase{
		UpdateAvailable:         updateAvailable,
		Reason:                  reason,
		ExpectedUpdateAvailable: false,
		ExpectedReason:          core.EmptyError,
	})
}

func Test_Update_IsAvailable_DbFolderNotFound(t *testing.T) {
	conf := config.Config{
		DatabaseRootDir: "./unknown_folder",
	}
	var updateAvailable, reason = IsAvailable(conf)

	assertUpdateIsAvailable(t, IsNeededTestCase{
		UpdateAvailable:         updateAvailable,
		Reason:                  reason,
		ExpectedUpdateAvailable: true,
		ExpectedReason:          ErrorDatabaseFolderNotFound,
	})
}

func Test_Update_IsAvailable_DbFileNotFound(t *testing.T) {
	conf := config.FromDefault(config.Config{
		DatabaseRootDir:  "./testdata",
		DatabaseFileName: "./unknown.db",
	})
	var updateAvailable, reason = IsAvailable(conf)

	assertUpdateIsAvailable(t, IsNeededTestCase{
		UpdateAvailable:         updateAvailable,
		Reason:                  reason,
		ExpectedUpdateAvailable: true,
		ExpectedReason:          ErrorDatabaseFileNotFound,
	})
}

func Test_Update_IsAvailable_DbChecksumFileNotFound(t *testing.T) {
	conf := config.FromDefault(config.Config{
		DatabaseRootDir:          "./testdata",
		DatabaseChecksumFileName: "./unknown.checksum",
	})
	var updateAvailable, reason = IsAvailable(conf)

	assertUpdateIsAvailable(t, IsNeededTestCase{
		UpdateAvailable:         updateAvailable,
		Reason:                  reason,
		ExpectedUpdateAvailable: true,
		ExpectedReason:          ErrorDatabaseChecksumFileNotFound,
	})
}

func Test_Update_IsAvailable_DatabaseServerChecksumFileInvalid(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.DisableNetworking()
	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checkssum").
		Reply(500).
		BodyString("")

	// create http server with one route
	http.ListenAndServe("127.0.0.1", nil)

	conf := config.FromDefault(config.Config{
		DatabaseRootDir: "./testdata",
	})
	var updateAvailable, reason = IsAvailable(conf)

	assertUpdateIsAvailable(t, IsNeededTestCase{
		UpdateAvailable:         updateAvailable,
		Reason:                  reason,
		ExpectedUpdateAvailable: false,
		ExpectedReasonContains:  "error checking server database checksum:",
	})
}

func Test_Update_IsAvailable_DatabaseServerChecksumMismatch(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	gock.New("https://raw.githubusercontent.com").
		Get("/1franck/cvepack-database/main/db.checksum").
		Reply(200).
		BodyString("6b466ca3ac976d32e380e072f461f7ac38ca528a788dbd37587965e93aa08e4d")

	// create http server with one route
	http.ListenAndServe("127.0.0.1", nil)

	conf := config.FromDefault(config.Config{
		DatabaseRootDir:          "./testdata",
		DatabaseChecksumFileName: "db.checksum.wrong",
	})
	var updateAvailable, reason = IsAvailable(conf)

	assertUpdateIsAvailable(t, IsNeededTestCase{
		UpdateAvailable:         updateAvailable,
		Reason:                  reason,
		ExpectedUpdateAvailable: true,
		ExpectedReason:          ErrorDatabaseChecksumMismatch,
	})
}
