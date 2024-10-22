package do

type OrganizationUnitsEx struct {
	OrganizationUnits
	Children []*OrganizationUnitsEx
}
