//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	AllAuthRecipeUsers = AllAuthRecipeUsers.FromSchema(schema)
	AppIDToUserID = AppIDToUserID.FromSchema(schema)
	Apps = Apps.FromSchema(schema)
	Category = Category.FromSchema(schema)
	DashboardUserSessions = DashboardUserSessions.FromSchema(schema)
	DashboardUsers = DashboardUsers.FromSchema(schema)
	EmailpasswordPswdResetTokens = EmailpasswordPswdResetTokens.FromSchema(schema)
	EmailpasswordUserToTenant = EmailpasswordUserToTenant.FromSchema(schema)
	EmailpasswordUsers = EmailpasswordUsers.FromSchema(schema)
	EmailverificationTokens = EmailverificationTokens.FromSchema(schema)
	EmailverificationVerifiedEmails = EmailverificationVerifiedEmails.FromSchema(schema)
	JwtSigningKeys = JwtSigningKeys.FromSchema(schema)
	KeyValue = KeyValue.FromSchema(schema)
	PasswordlessCodes = PasswordlessCodes.FromSchema(schema)
	PasswordlessDevices = PasswordlessDevices.FromSchema(schema)
	PasswordlessUserToTenant = PasswordlessUserToTenant.FromSchema(schema)
	PasswordlessUsers = PasswordlessUsers.FromSchema(schema)
	RolePermissions = RolePermissions.FromSchema(schema)
	Roles = Roles.FromSchema(schema)
	SessionAccessTokenSigningKeys = SessionAccessTokenSigningKeys.FromSchema(schema)
	SessionInfo = SessionInfo.FromSchema(schema)
	TenantConfigs = TenantConfigs.FromSchema(schema)
	TenantThirdpartyProviderClients = TenantThirdpartyProviderClients.FromSchema(schema)
	TenantThirdpartyProviders = TenantThirdpartyProviders.FromSchema(schema)
	Tenants = Tenants.FromSchema(schema)
	ThirdpartyUserToTenant = ThirdpartyUserToTenant.FromSchema(schema)
	ThirdpartyUsers = ThirdpartyUsers.FromSchema(schema)
	TotpUsedCodes = TotpUsedCodes.FromSchema(schema)
	TotpUserDevices = TotpUserDevices.FromSchema(schema)
	TotpUsers = TotpUsers.FromSchema(schema)
	Transaction = Transaction.FromSchema(schema)
	UserLastActive = UserLastActive.FromSchema(schema)
	UserMetadata = UserMetadata.FromSchema(schema)
	UserRoles = UserRoles.FromSchema(schema)
	UseridMapping = UseridMapping.FromSchema(schema)
}
