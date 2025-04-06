// Package domain provides the domain model for the application.
package domain

// StandardRoleEnum represents the standard roles in the system
type StandardRoleEnum string

const (
	//SuperAdmin role. Allow us to create a new organization
	SuperAdmin = "18ec638e-8411-486b-b81b-4d4dbfeaf246"
	//CustomerSuccess role. Allow us to login as unicorn
	CustomerSuccess = "57d21b04-ec21-4950-8b3b-28876c252921"
	//OrganizationAdmin role. First user into the organization
	OrganizationAdmin = "b2e1eea3-dd6e-49af-bd16-2bf14bca4492"
	//OrganizationMember role. Depends on the access that he has assigned
	OrganizationMember = "f25f653b-3337-42c7-8497-2bb00b1c6739"
)
