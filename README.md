# :space_invader: CVEPack

CVEPack is a tool to detect automatically vulnerabilities in packages from various ecosystems. 

It uses a [compiled](https://github.com/1franck/cvepack-database) version of [GitHub Advisory Database](https://github.com/github/advisory-database) as source for detecting CVEs.

#### Ecosystems supported

- [x] NPM (Node.js)
- [x] Go
- [x] Packagist (PHP)
- [x] Crates.io (Rust)
- [x] RubyGems (Ruby)
- [x] PyPI (Python)
- [x] NuGet (.Net)
- [x] Maven (Java)

## Usage

#### Scan folder(s)

```bash
cvepack scan [-d|--details] <folder path> [<folder path> ...]
```

![scan_cmd.png](./screenshots/scan_cmd.png)

#### Search a package

```bash
cvepack search <package name>
```

![search_cmd.png](./screenshots/search_cmd.png)

#### Update CVE database

```bash
cvepack update
```