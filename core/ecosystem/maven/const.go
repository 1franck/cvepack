package maven

const (
	EcosystemLanguage = "Java"
	EcosystemName     = "Maven"
	PomXml            = "pom.xml"
)

func EcosystemTitle() string {
	return EcosystemLanguage + " (" + EcosystemName + ")"
}
