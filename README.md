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

### Scanner
#### scan Path(s)

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

#### Update CVE database

```bash
cvepack update
```