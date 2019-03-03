package parsers

import (
	"regexp"
	"strconv"

	"github.com/hiromoon/getLatest/models"
)

// ErlangVersionParser is ...
type ErlangVersionParser struct{}

// Parse is ...
func (p *ErlangVersionParser) Parse(version string) *models.Version {
	re := regexp.MustCompile("OTP-(\\d+)\\.(\\d+)\\.(\\d+)(-(.*)(\\d+))*")
	v := re.FindAllStringSubmatch(version, -1)
	if v == nil {
		return &models.Version{
			Major: -1,
			Minor: -1,
			Patch: -1,
		}
	}

	major, err := strconv.Atoi(v[0][1])
	if err != nil {
		panic(err)
	}
	minor, err := strconv.Atoi(v[0][2])
	if err != nil {
		panic(err)
	}
	patch, err := strconv.Atoi(v[0][3])
	if err != nil {
		panic(err)
	}
	var alpha int
	var beta int
	var rc int
	switch v[0][5] {
	case "alpha":
		a, err := strconv.Atoi(v[0][6])
		if err != nil {
			panic(err)
		}
		alpha = a
		beta = -1
		rc = -1
	case "beta":
		b, err := strconv.Atoi(v[0][6])
		if err != nil {
			panic(err)
		}
		alpha = -1
		beta = b
		rc = -1
	case "rc":
		r, err := strconv.Atoi(v[0][6])
		if err != nil {
			panic(err)
		}
		alpha = -1
		beta = -1
		rc = r
	default:
		alpha = -1
		beta = -1
		rc = -1
	}
	return &models.Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Alpha: alpha,
		Beta:  beta,
		RC:    rc,
	}
}
