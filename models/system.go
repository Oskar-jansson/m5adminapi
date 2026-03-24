package models

type System struct {
	Servers []struct {
		Id      *uint32 `json:"id" validate:"gt=0" cmp:"skip"`
		Online  *bool   `json:"online" validate:"required"`
		Name    *string `json:"name" validate:"required"`
		Systems []struct {
			Name *string `json:"name" validate:"required"`
		} `json:"systems" validate:"required"`
	} `json:"servers" validate:"required"`
}
