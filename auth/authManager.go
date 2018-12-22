package auth

const (
	//NotAuthorizedProfileID is the id of a non-authorized user
	NotAuthorizedProfileID = -1
)

//AuthenticationManager provides user authentication method
type AuthenticationManager interface {
	GetProfileInfo(authToken string, userAgent string) (*ProfileInfo, error)
}
