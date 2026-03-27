package models

type Function struct {
	Id            *uint32      `json:"id" validate:"gt=0" cmp:"skip"`
	Fkaccessgroup *uint32      `json:"fkaccessgroup"`
	Fkunit        *uint32      `json:"fkunit"`
	Times         *string      `json:"times" validate:"required"`
	Timeperiods   *string      `json:"timeperiods"`
	Activedays    *string      `json:"activedays"`
	Type          *uint32      `json:"type" validate:"required"`
	Comment       *string      `json:"comment"`
	Rastamp       *Rastamp     `json:"rastamp" validate:"required" cmp:"skip"`
	Unit          *Unit        `json:"unit"`
	Accessgroup   *Accessgroup `json:"Accessgroup"`
}

type FunctionList struct {
	Functions []Function `json:"functions"`
	Count     *int       `json:"count,omitempty"` // only included with ?count='true'
}

type FunctionInput struct {
	Fkaccessgroup *uint32      `json:"fkaccessgroup,omitempty"`
	Fkunit        *uint32      `json:"fkunit,omitempty"`
	Times         *string      `json:"times,omitempty"`
	Timeperiods   *string      `json:"timeperiods,omitempty"`
	Activedays    *string      `json:"activedays,omitempty"`
	Type          *uint32      `json:"type,omitempty"`
	Comment       *string      `json:"comment,omitempty"`
	Rastamp       *Rastamp     `json:"rastamp,omitempty"`
	Unit          *Unit        `json:"unit,omitempty"`
	Accessgroup   *Accessgroup `json:"Accessgroup,omitempty"`
}
