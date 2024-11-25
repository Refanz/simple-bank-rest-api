package constant

const (
	PaymentApi  = "/api/payments"
	LoginApi    = "/api/auth/login"
	LogoutApi   = "/api/auth/logout"
	CustomerApi = "/api/customers"

	ServerRunningMessage = "server is running on port 8080.."

	GetAllCustomersSuccessMessage    = "successfully get all customers"
	TransferSuccessMessage           = "customer Transfer to merchant"
	TransferSuccessWithAmountMessage = "customer successfully transfers to merchant in the amount of"
	LoginSuccessMessage              = "login successfully"
	LogoutSuccessMessage             = "logout successfully"

	AuthorizationRequiredMessage = "authorization header is required"
	CredentialsErrorMessage      = "credentials is not valid"

	CookieName         = "user_session"
	CookieErrorMessage = "missing session information"

	UserNotFoundErrorMessage = "login error, user is not found"
)
