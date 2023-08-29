package semver

import "github.com/axllent/semver"

func IsVersionAffectedByFixedVersion(ver string, introducedVer string, fixedVer string) bool {
	if semver.Compare(ver, introducedVer) >= 0 &&
		semver.Compare(ver, fixedVer) < 0 {
		return true
	}
	return false
}

func IsVersionInRange(ver string, minVer string, maxVer string) bool {
	if semver.Compare(ver, minVer) >= 0 &&
		semver.Compare(ver, maxVer) <= 0 {
		return true
	}
	return false
}
