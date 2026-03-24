package models

// Non-standard object to make compare easier
type Version struct {
	Major  int
	Middle int
	Minor  int

	AsString string // full version as string
	AsInt    int    // full version as int
}
