package cratesio

const (
	EcosystemLanguage = "Rust"
	EcosystemName     = "crates.io"
	CargoFile         = "Cargo.toml"
	CargoLockFile     = "Cargo.lock"
)

func EcosystemTitle() string {
	return EcosystemLanguage + " (" + EcosystemName + ")"
}
