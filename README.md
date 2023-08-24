# CVEPACK

CVEPack is a tool to detect vulnerabilities in packages. 

It uses [GitHub Advisory Database](https://github.com/github/advisory-database) as source for detecting CVEs.

Be aware, this is a proof of concept and a work in progress.

#### Ecosystems supported

- [x] NPM
- [x] Go
- [x] Packagist
- [ ] NuGet

## Usage

#### Scan a folder

```bash
$ cvepack scan <folder path>
```

#### Update CVE database

```bash
$ cvepack update
```