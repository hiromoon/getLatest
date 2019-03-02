package parsers

import (
	"regexp"
	"strconv"

	"github.com/hiromoon/getLatest/models"
)

// ElixirVersionParser is ...
type ElixirVersionParser struct{}

// Parse is ...
func (p *ElixirVersionParser) Parse(version string) *models.Version {
	re := regexp.MustCompile("v(\\d+)\\.(\\d+)\\.(\\d+)(-(.+)\\.(\\d+))*")
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
	return &models.Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}
