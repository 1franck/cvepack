package core

var insertVulnerabilityQuery = `
    INSERT INTO vulnerabilities 
    (id, modified, published, withdrawn, aliases, related, summary, details, severity, refs, credits, database_specific)
    VALUES (?,?,?,?,?,?,?,?,?,?,?,?)
`

var countVulnerabilityQuery = `
    SELECT COUNT(*) FROM vulnerabilities
`

var insertAffectedQuery = `
    INSERT INTO affected
    (vulnerability_id, package_ecosystem, package_name, package_purl, severity, versions, ecosystem_specific, database_specific)
    VALUES (?,?,?,?,?,?,?,?)
`

var insertAffectedRangesQuery = `
    INSERT INTO affected_ranges
    (vulnerability_id, affected_id, range_type, range_repo, database_specific, e_introduced, e_fixed, e_last_affected, e_limit)
    VALUES (?,?,?,?,?,?,?,?,?)
`

var searchAffectedPackageQuery = `
    SELECT a.*
    FROM affected a
    JOIN affected_ranges ar ON a.id = ar.affected_id
    WHERE a.package_name = ? AND a.package_ecosystem = ?
`
