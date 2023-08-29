package osv

import "time"

// @see https://ossf.github.io/osv-schema/

type Osv struct {
	SchemaVersion    string           `json:"schema_version"`
	ID               string           `json:"id"`
	Modified         time.Time        `json:"modified"`
	Published        time.Time        `json:"published"`
	Withdrawn        *time.Time       `json:"withdrawn,omitempty"`
	Aliases          []string         `json:"aliases"`
	Related          []string         `json:"related"`
	Summary          string           `json:"summary"`
	Details          string           `json:"details"`
	Severity         []Severity       `json:"severity"`
	Affected         []Affected       `json:"affected"`
	References       []References     `json:"references"`
	Credits          []Credits        `json:"credits,omitempty"`
	DatabaseSpecific DatabaseSpecific `json:"database_specific"`
}

type Severity struct {
	Type  string `json:"type"`
	Score string `json:"score"`
}

type Affected struct {
	Package           AffectedPackage           `json:"package"`
	Severity          []Severity                `json:"severity,omitempty"`
	Versions          []string                  `json:"versions,omitempty"`
	Ranges            []AffectedRanges          `json:"ranges"`
	DatabaseSpecific  AffectedDatabaseSpecific  `json:"database_specific"`
	EcosystemSpecific AffectedEcosystemSpecific `json:"ecosystem_specific"`
}

type AffectedPackage struct {
	Ecosystem string  `json:"ecosystem"`
	Name      string  `json:"name"`
	Purl      *string `json:"purl,omitempty"`
}

type AffectedRanges struct {
	Type             string                         `json:"type"`
	Repo             *string                        `json:"repo,omitempty"`
	Events           []AffectedRangesEvent          `json:"events"`
	DatabaseSpecific AffectedRangesDatabaseSpecific `json:"database_specific"`
}

type AffectedRangesEvent struct {
	Introduced   *string `json:"introduced,omitempty"`
	Fixed        *string `json:"fixed,omitempty"`
	LastAffected *string `json:"last_affected,omitempty"`
	Limit        *string `json:"limit,omitempty"`
}

type AffectedDatabaseSpecific Unknown

type AffectedRangesDatabaseSpecific Unknown

type AffectedEcosystemSpecific Unknown

type References struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type Credits struct {
	Name    string   `json:"name"`
	Contact []string `json:"contact"`
	Type    string   `json:"type"`
}
type DatabaseSpecific struct {
	CweIds           []string  `json:"cwe_ids"`
	Severity         string    `json:"severity"`
	GithubReviewed   bool      `json:"github_reviewed"`
	GithubReviewedAt time.Time `json:"github_reviewed_at"`
	NvdPublishedAt   time.Time `json:"nvd_published_at"`
	Unknown
}

type Unknown struct {
	Extras map[string]interface{} `json:"-"`
}
