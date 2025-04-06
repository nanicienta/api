package errors

// ErrorCode contains the error codes for the application (Different services
type ErrorCode string

const (
	//ErrorUserNotAcceptTermsAndConditions is returned when the user has not accepted the terms and conditions
	ErrorUserNotAcceptTermsAndConditions ErrorCode = "0000001"
	//ErrorInvalidEmail is returned when the email is invalid
	ErrorInvalidEmail ErrorCode = "0000002"
	//ErrorPasswordDoesntHaveTheRequestedFormat is returned when the password does not have the requested format
	ErrorPasswordDoesntHaveTheRequestedFormat ErrorCode = "0000003"
	//ErrorPasswordDoesntMatch is returned when the password does not match
	ErrorPasswordDoesntMatch ErrorCode = "0000004"
	//ErrorCreatingUser is returned when there is an error creating the user (Unexpected error)
	ErrorCreatingUser ErrorCode = "0000005"
	//ErrorAuthenticatingUser is returned when there is an error authenticating the user.
	ErrorAuthenticatingUser ErrorCode = "0000006"
	//ErrorUserAuthenticatedNotFound is returned when the user authenticated is not found
	ErrorUserAuthenticatedNotFound ErrorCode = "0000007"
	//ErrorUserEmailAlreadyExists is returned when the user email already exists
	ErrorUserEmailAlreadyExists ErrorCode = "0000008"
	//ErrorUserCantAddUsersToOrganization is returned when the user cannot add users to the organization
	ErrorUserCantAddUsersToOrganization ErrorCode = "0000009"
	//ErrorUserCantRemoveUsersFromOrganization is returned when the user cannot remove users from the organization
	ErrorUserCantRemoveUsersFromOrganization ErrorCode = "00000020"
	//ErrorUserCantCreateTeamIntoOrganization is returned when the user cannot create a team into the organization
	ErrorUserCantCreateTeamIntoOrganization ErrorCode = "00000021"

	// ErrorTryingToGetTeamsByNameAndSlug indicates an error occurred when fetching teams by name and slug.
	ErrorTryingToGetTeamsByNameAndSlug ErrorCode = "00000022"

	// ErrorConflictTeamExistWithSameNameOrSlug indicates a team exists with the same name or slug.
	ErrorConflictTeamExistWithSameNameOrSlug ErrorCode = "00000023"
	// ErrorCreatingTeam indicates an error occurred while creating a team.
	ErrorCreatingTeam ErrorCode = "00000024"
	// ErrorUserDontHavePrivilegesToAddTeamMembersToTeam indicates the user does not have privileges to add team members to the team.
	ErrorUserDontHavePrivilegesToAddTeamMembersToTeam ErrorCode = "00000025"
	// ErrorUserDontHavePrivilegesToReadOrganizationMembers indicates the user does not have privileges to read organization members.
	ErrorUserDontHavePrivilegesToReadOrganizationMembers ErrorCode = "00000026"
	// ErrorGettingOrganizationMembers indicates an error occurred while getting organization members.
	ErrorGettingOrganizationMembers ErrorCode = "00000027"

	// InternalServerErrorGettingAccounts indicates an unexpected error at the moment to get the
	//accounts
	InternalServerErrorGettingAccounts ErrorCode = "00000028"
	// InternalServerErrorCreatingAccount indicates an unexpected error at the moment to create an
	//account
	InternalServerErrorCreatingAccount ErrorCode = "00000029"

	// InternalServerErrorSendingNotification indicates an unexpected error while sending a notification.
	InternalServerErrorSendingNotification ErrorCode = "00000030"
)
