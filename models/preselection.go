package models

type Preselection struct {
	Id                *uint32            `json:"id" validate:"gt=0" cmp:"skip"`
	Rastamp           *string            `json:"rastamp" validate:"required" cmp:"skip"`
	Name              *string            `json:"name" validate:"required"`
	Type              *uint32            `json:"type" validate:"required"`
	Fkconnr           *uint32            `json:"fkconnr" validate:"required"`
	Owner             *uint32            `json:"owner" validate:"required"`
	Rasystem          *uint32            `json:"rasystem" validate:"required"`
	Machinegrouptypes []Machinegrouptype `json:"machinegrouptypes"`
}

type PreselectionList struct {
	Preselections []Preselection `json:"preselections"`
	Count         *int           `json:"count,omitempty"` // only included with ?count='true'
}
