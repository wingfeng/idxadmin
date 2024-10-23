package entity

type OrganizationUnitsEx struct {
	//	Parent *OrganizationUnitsEx
	OrganizationUnits
	Children []*OrganizationUnitsEx `json:"children"`
}
