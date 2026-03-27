package models

type Offlineunit struct {
	Id               *uint32        `json:"id" validate:"gt=0" cmp:"skip"`
	Name             *string        `json:"name" validate:"required"`
	Rasystem         *uint32        `json:"rasystem" validate:"required"`
	Address          *uint32        `json:"address" validate:"required"`
	Type             *string        `json:"type" validate:"required"`
	Pos              *string        `json:"pos" validate:"required"`
	Batterytime      *string        `json:"batterytime"`
	Batterystatus    *uint32        `json:"batterystatus"`
	Unitsynctime     *string        `json:"unitsynctime"`
	Unitsyncstatus   *uint32        `json:"unitsyncstatus"`
	Unitedittime     *string        `json:"unitedittime"`
	Lasteventtime    *string        `json:"lasteventtime" cmp:"skip"`
	Rastamp          *Rastamp       `json:"rastamp" validate:"required" cmp:"skip"`
	Accessversion    *uint32        `json:"accessversion" validate:"required" cmp:"skip"`
	Versionstarttime *string        `json:"versionstarttime" cmp:"skip"`
	Readeraccess     []Readeraccess `json:"readeraccess"`
}

type OfflineunitList struct {
	Offlineunits []Offlineunit `json:"offlineunits"`
	Count        *int          `json:"count,omitempty"` // only included with ?count='true'
}
