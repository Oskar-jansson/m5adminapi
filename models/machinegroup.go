package models

type Machinegroup struct {
	Id                 *uint32           `json:"id" validate:"gt=0" cmp:"skip"`
	Rastamp            *string           `json:"rastamp" validate:"required" cmp:"skip"`
	Name               *string           `json:"name" validate:"required"`
	Fkmachinegrouptype *uint32           `json:"fkmachinegrouptype" validate:"required"`
	Unit               *uint32           `json:"unit" validate:"required"`
	Fkconnr            *uint32           `json:"fkconnr" validate:"required"`
	Rasystem           *uint32           `json:"rasystem" validate:"required"`
	Fkdebpricelist     *uint32           `json:"fkdebpricelist" validate:"required"`
	Machinegrouptype   *Machinegrouptype `json:"machinegrouptype"`
	Accessgroups       []Accessgroup     `json:"accessgroups"`
}

type MachinegroupList struct {
	Machinegroups []Machinegroup `json:"machinegroups"`
	Count         *int           `json:"count,omitempty"` // only included with ?count='true'
}
