package config

import "path/filepath"

var Default = Config{
	Name:                     "CVEPack",
	Version:                  "0.0.1",
	DatabaseRootDir:          "./cvepack-database-main",
	DatabaseFileName:         "advisories.db",
	DatabaseChecksumFileName: "db.checksum",
	DatabaseUrl:              "https://github.com/1franck/cvepack-database/archive/refs/heads/main.zip",
	DatabaseChecksumUrl:      "https://raw.githubusercontent.com/1franck/cvepack-database/main/db.checksum",
}

type Config struct {
	Name                     string
	Version                  string
	DatabaseRootDir          string
	DatabaseFileName         string
	DatabaseChecksumFileName string
	DatabaseUrl              string
	DatabaseChecksumUrl      string
}

func (config *Config) DatabaseFilePath() string {
	return filepath.Join(config.DatabaseRootDir, config.DatabaseFileName)
}

func (config *Config) DatabaseChecksumFilePath() string {
	return filepath.Join(config.DatabaseRootDir, config.DatabaseChecksumFileName)
}

func (config *Config) NameAndVersion() string {
	return config.Name + " v" + config.Version
}