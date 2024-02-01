# :space_invader: CVEPack

CVEPack is a tool to detect vulnerabilities in packages from various ecosystems. 

It uses a [compiled](https://github.com/1franck/cvepack-database) version of [GitHub Advisory Database](https://github.com/github/advisory-database) as source for detecting CVEs.

#### Ecosystems detected with their package managers and lock files:

- [x] NPM (Node.js)
  - package-lock.json
  - yarn.lock
  - pnpm-lock.yaml
  - /node_modules
- [x] Go
  - go.sum    
- [x] Packagist (PHP)
  - composer.lock 
- [x] Crates.io (Rust)
  - Cargo.lock
- [x] RubyGems (Ruby)
  - Gemfile.lock
- [x] PyPI (Python)
  - poetry.lock
  - pdm.lock
- [x] NuGet (.Net)
  - .sln
  - .csproj
- [x] Maven (Java)
  - pom.xml

### Scanner
#### scan path(s)

```bash
cvepack scan <path1> [<path2> ...]
```

![scan_cmd.png](./screenshots/scan_cmd.png)

#### scan GitHub url(s) with -u/--url

```bash
cvepack scan -u <url1> [<url2> ...]

ex: $ cvepack scan -u github.com/1franck/cvepack
```

#### scan commands flags
| Flag | Description                |
| --- |----------------------------|
| -d, --details | Show CVE details           |
| -u, --url | Scan GitHub repository url |
| -s, --silent | Silent mode                |
| -o, --output | Result output file         |

### Search a package

```bash
cvepack search <package name>
```

![search_cmd.png](./screenshots/search_cmd.png)

#### Update CVE [database](https://github.com/1franck/cvepack-database)

```bash
cvepack update
```

### Build from source

```bash
make
```