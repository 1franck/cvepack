package config

import "path/filepath"

var Default = Config{
	Name:                     "CVEPack",
	Version:                  "0.4.0",
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

func FromDefault(config Config) Config {
	if config.Name == "" {
		config.Name = Default.Name
	}
	if config.Version == "" {
		config.Version = Default.Version
	}
	if config.DatabaseRootDir == "" {
		config.DatabaseRootDir = Default.DatabaseRootDir
	}
	if config.DatabaseFileName == "" {
		config.DatabaseFileName = Default.DatabaseFileName
	}
	if config.DatabaseChecksumFileName == "" {
		config.DatabaseChecksumFileName = Default.DatabaseChecksumFileName
	}
	if config.DatabaseUrl == "" {
		config.DatabaseUrl = Default.DatabaseUrl
	}
	if config.DatabaseChecksumUrl == "" {
		config.DatabaseChecksumUrl = Default.DatabaseChecksumUrl
	}
	return config
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
