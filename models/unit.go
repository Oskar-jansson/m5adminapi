package models

type Unit struct {
	Id            *uint32        `json:"id" validate:"gt=0" cmp:"skip"`
	Name          *string        `json:"name" validate:"required"`
	Parent        *uint32        `json:"parent" validate:"required"`
	Pos           *string        `json:"pos" validate:"required"`
	Rastamp       *Rastamp       `json:"rastamp" validate:"required" cmp:"skip"`
	Rasystem      *uint32        `json:"rasystem" validate:"required"`
	Connumber     *uint32        `json:"connumber" validate:"required"`
	State         *uint32        `json:"state" validate:"required"`
	Type          *string        `json:"type" validate:"required"`
	Apidooraccess *uint32        `json:"apidooraccess"`
	Changedby     *ChangedBy     `json:"changedby" validate:"required" cmp:"skip"`
	Changeddate   *ChangedDate   `json:"changeddate" validate:"required" cmp:"skip"`
	Createdby     *CreatedBy     `json:"createdby" validate:"required" cmp:"skip"`
	Createddate   *CreatedDate   `json:"createddate" validate:"required" cmp:"skip"`
	Accessgroups  []Accessgroup  `json:"accessgroups"`
	Functions     []Function     `json:"functions"`
	Readeraccess  []Readeraccess `json:"readeraccess"`
}

type UnitList struct {
	Units []Unit `json:"units"`
	Count *int   `json:"count,omitempty"` // only included with ?count='true'
}
