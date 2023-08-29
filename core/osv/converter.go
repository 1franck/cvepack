package osv

import (
	"encoding/json"
	"fmt"
)

func encode(field any) string {
	content, err := json.Marshal(field)
	if err != nil {
		return fmt.Sprintf("<error:%s>", err)
	}
	return string(content)
}

func (o *Osv) AliasesJson() string {
	return encode(o.Aliases)
}

func (o *Osv) RelatedJson() string {
	if o.Related == nil {
		return encode([]string{})
	}
	return encode(o.Related)
}

func (o *Osv) SeverityJson() string {
	return encode(o.Severity)
}

func (o *Osv) ReferencesJson() string {
	return encode(o.References)
}

func (o *Osv) CreditsJson() string {
	if o.Credits == nil {
		return encode([]string{})
	}
	return encode(o.Credits)
}

func (o *Osv) DatabaseSpecificJson() string {
	return encode(o.DatabaseSpecific)
}

func (a *Affected) SeverityJson() string {
	if a.Severity == nil {
		return encode([]string{})
	}
	return encode(a.Severity)
}

func (a *Affected) VersionsJson() string {
	if a.Versions == nil {
		return encode([]string{})
	}
	return encode(a.Versions)
}

func (a *Affected) EcosystemSpecificJson() string {
	return encode(a.EcosystemSpecific)
}

func (a *Affected) DatabaseSpecificJson() string {
	return encode(a.DatabaseSpecific)
}

func (a *AffectedRanges) DatabaseSpecificJson() string {
	return encode(a.DatabaseSpecific)
}
