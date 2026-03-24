package models

type Department struct {
	Id             *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
	Departmentname *string `json:"departmentname" validate:"required"`
	Rastamp        *string `json:"rastamp" validate:"required" cmp:"skip"`
	Departmentdata *string `json:"departmentdata"`
	Users          []User  `json:"users"`
}

type DepartmentInput struct {
	Departmentname *string `json:"departmentname,omitempty"`
	Rastamp        *string `json:"rastamp,omitempty"`
	Departmentdata *string `json:"departmentdata,omitempty"`
}

type DepartmentList struct {
	Departments []Department `json:"departments"`
	Count       *int         `json:"count,omitempty"` // only included with ?count='true'
}
