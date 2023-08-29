package database

var DbSchema = `
CREATE TABLE IF NOT EXISTS vulnerabilities (
    id TEXT PRIMARY KEY,
    modified TEXT,
    published TEXT,
    withdrawn TEXT,
    aliases TEXT,
    related TEXT,
    summary TEXT,
    details TEXT,
    severity TEXT,
    refs TEXT,
    credits TEXT,
    database_specific TEXT
);

CREATE TABLE IF NOT EXISTS affected (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    vulnerability_id TEXT,
    package_ecosystem TEXT,
    package_name TEXT,
    package_purl TEXT,
    severity TEXT,
    versions TEXT,
    ecosystem_specific TEXT,
    database_specific TEXT,
    FOREIGN KEY (vulnerability_id) REFERENCES vulnerabilities(id)
);


CREATE TABLE IF NOT EXISTS affected_ranges (
    vulnerability_id TEXT,
    affected_id INTEGER,
    range_type TEXT,
    range_repo TEXT,
    database_specific TEXT,
    e_introduced TEXT,
    e_fixed TEXT,
    e_last_affected TEXT,
    e_limit TEXT,
    FOREIGN KEY (vulnerability_id) REFERENCES vulnerabilities(id),
    FOREIGN KEY (affected_id) REFERENCES affected(id)
);
`
