package nuget

const (
	EcosystemLanguage = ".Net"
	EcosystemName     = "NuGet"
)

func EcosystemTitle() string {
	return EcosystemLanguage + " (" + EcosystemName + ")"
}
