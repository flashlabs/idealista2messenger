package input

type Params struct {
	AccessTokenFileLocation, PageAccessTokenFileLocation, GmailUserId, GmailQuery,
	PageId, CredentialsFileLocation string
	Recipients []string
}
