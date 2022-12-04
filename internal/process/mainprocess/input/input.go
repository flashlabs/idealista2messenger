package input

type Input struct {
	AccessTokenFileLocation, PageAccessTokenFileLocation, GmailUserId, GmailQuery,
	PageId, CredentialsFileLocation string
	Recipients []string
}
