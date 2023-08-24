package packagist

import (
	"encoding/json"
	"github.com/1franck/cvepack/internal/common"
)

func readComposerLockFile(filePath string) (*composerLockFile, error) {
	content, err := common.ReadAllFile(filePath)
	if err != nil {
		return nil, err
	}

	var composerLockFile composerLockFile
	if err := json.Unmarshal(content, &composerLockFile); err != nil {
		return nil, err
	}

	return &composerLockFile, nil
}
