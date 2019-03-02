package checkers

import (
	"github.com/hiromoon/getLatest/models"
)

// VersionParser is ...
type VersionParser interface {
	Parse(string) *models.Version
}

// SemanticVersionChecker is ...
type SemanticVersionChecker struct {
	Parser VersionParser
}

// Check is ...
func (checker *SemanticVersionChecker) Check(nowVersion, newVersion string) bool {
	nowV := checker.Parser.Parse(nowVersion)
	newV := checker.Parser.Parse(newVersion)
	// Compair major version.
	if nowV.Major < newV.Major {
		return true
	}
	if nowV.Major > newV.Major {
		return false
	}
	// Compair minor version, if equals major version.
	if nowV.Minor < newV.Minor {
		return true
	}
	if nowV.Minor > newV.Minor {
		return false
	}
	// Compair patch version, if equals major and minor version.
	return nowV.Patch < newV.Patch
}
