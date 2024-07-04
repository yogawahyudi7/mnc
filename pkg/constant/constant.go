package constant

const (
	TimeFormatYMD    = "2006-01-02"
	TimeFormatYMDHMS = "2006-01-02 15:04:05"

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	InvalidExpiredToken = "invalid or expired refresh token"
	InvalidRequestBody  = "invalid request body"

	UpperCaseSuccess = "SUCCESS"
	UpperCaseDebit   = "DEBIT"
	UpperCaseCredit  = "CREDIT"

	AuthorizationHeader = "Authorization"
	UserContext         = "user"
	InvalidToken        = "invalid token"

	TransferType = "transfer"
	TopUpType    = "topup"
	PaymentType  = "payment"

	TopUpRemarks = "Top Up"

	TransferIdSnakeCase = "transfer_id"
	TopUpIdSnakeCase    = "top_up_id"
	PaymentIdSnakeCase  = "payment_id"
)
