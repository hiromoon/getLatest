package models

// Version is Version struct
type Version struct {
	Major int
	Minor int
	Patch int
	Alpha int
	Beta  int
	RC    int
}

// IsAlpha is ..
func (v *Version) IsAlpha() bool {
	return v.Alpha != -1
}

// IsBeta is ...
func (v *Version) IsBeta() bool {
	return v.Beta != -1
}

// IsRC is ...
func (v *Version) IsRC() bool {
	return v.Beta != -1
}
