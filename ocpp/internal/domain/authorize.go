package domain

type AuthorizeRequest struct {
	IdTag string `json:"idTag"`
}

type IdTagInfo struct {
	Status      AuthorizationStatus `json:"status"`
	ExpiryDate  string              `json:"expiryDate,omitempty"`
	ParentIdTag string              `json:"parentIdTag,omitempty"`
}

type AuthorizationStatus string

const (
	AuthorizationStatusAccepted     AuthorizationStatus = "Accepted"
	AuthorizationStatusBlocked      AuthorizationStatus = "Blocked"
	AuthorizationStatusExpired      AuthorizationStatus = "Expired"
	AuthorizationStatusInvalid      AuthorizationStatus = "Invalid"
	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

type AuthorizeResponse struct {
	IdTagInfo IdTagInfo `json:"idTagInfo"`
}
