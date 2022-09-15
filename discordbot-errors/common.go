package discordbot_errors

import "net/http"

func init() {
	// AUTH ERRORS
	discordBotErrorInfos[ERR_AUTHORIZATION_COMMON] = DiscordBotErrorInfo{"Unauthorized", http.StatusUnauthorized}
	discordBotErrorInfos[ERR_AUTHORIZATION_JWT] = DiscordBotErrorInfo{"Unauthorized, invalid JWT token", http.StatusUnauthorized}
	discordBotErrorInfos[ERR_AUTHORIZATION_TOKEN_MISSING] = DiscordBotErrorInfo{"Authorization header is empty", http.StatusUnauthorized}
	discordBotErrorInfos[ERR_AUTHORIZATION_TOKEN_INVALID] = DiscordBotErrorInfo{"Authorization token is invalid", http.StatusUnauthorized}
	discordBotErrorInfos[ERR_AUTHORIZATION_TOKEN_EXPIRED] = DiscordBotErrorInfo{"Token is expired", http.StatusUnauthorized}
	discordBotErrorInfos[ERR_AUTHORIZATION_NOT_KYC] = DiscordBotErrorInfo{"KYC required", http.StatusUnauthorized}

	// VALIDATION ERRORS
	discordBotErrorInfos[ERR_INVALID_REQUEST_COMMON] = DiscordBotErrorInfo{"Invalid request", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_APP_ID] = DiscordBotErrorInfo{"Invalid app id", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_ADDRESS] = DiscordBotErrorInfo{"Invalid address", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_CHAIN] = DiscordBotErrorInfo{"Chain is not supported", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_CONTRACT_TYPE] = DiscordBotErrorInfo{"Contract type is not supported", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_URL] = DiscordBotErrorInfo{"Invalid url", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_METADATA_ID] = DiscordBotErrorInfo{"Invalid metadataId", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_MINT_AMOUNT] = DiscordBotErrorInfo{"Invalid mint amount, mint amount could not be 0", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_MINT_AMOUNT_721] = DiscordBotErrorInfo{"Invalid mint amount, mint amount could not more than 1 for erc 721 contract", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_TOKEN_ID] = DiscordBotErrorInfo{"Invalid token ID", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_CONTRACT_TYPE_UNMATCH] = DiscordBotErrorInfo{"Contract type and contract address not match", http.StatusBadRequest}
	discordBotErrorInfos[ERR_INVALID_PAGINATION] = DiscordBotErrorInfo{"Invalid page or limit", http.StatusBadRequest}

	// CONFLICT ERRORS
	discordBotErrorInfos[ERR_CONFLICT_COMMON] = DiscordBotErrorInfo{"Conflict", http.StatusConflict}
	discordBotErrorInfos[ERR_CONFLICT_COMPANY_EXISTS] = DiscordBotErrorInfo{"Company already exists", http.StatusConflict}

	// RATELIMIT ERRORS
	discordBotErrorInfos[ERR_TOO_MANY_REQUEST_COMMON] = DiscordBotErrorInfo{"Too many requests", http.StatusTooManyRequests}

	// INTERNAL SERVER ERRORS
	discordBotErrorInfos[ERR_INTERNAL_SERVER_COMMON] = DiscordBotErrorInfo{"Internal Server error", http.StatusInternalServerError}
	discordBotErrorInfos[ERR_INTERNAL_SERVER_DB] = DiscordBotErrorInfo{"Database operation error", http.StatusInternalServerError}

	// BUSINESS ERRORS
	discordBotErrorInfos[ERR_BUSINESS_COMMON] = DiscordBotErrorInfo{"Business error", HTTP_STATUS_BUSINESS_ERROR}
	discordBotErrorInfos[ERR_MINT_LIMIT_EXCEEDED] = DiscordBotErrorInfo{"Mint limit exceeded", HTTP_STATUS_BUSINESS_ERROR}
	discordBotErrorInfos[ERR_DEPLOY_LIMIT_EXCEEDED] = DiscordBotErrorInfo{"Deploy limit exceeded", HTTP_STATUS_BUSINESS_ERROR}
	discordBotErrorInfos[ERR_UPLOADE_FILE_LIMIT_EXCEEDED] = DiscordBotErrorInfo{"Uploade file limit exceeded", HTTP_STATUS_BUSINESS_ERROR}

	discordBotErrorInfos[ERR_NO_SPONSOR] = DiscordBotErrorInfo{"Contract has no sponsor", HTTP_STATUS_BUSINESS_ERROR}
	discordBotErrorInfos[ERR_NO_SPONSOR_BALANCE] = DiscordBotErrorInfo{"Contract sponsor balance not enough", HTTP_STATUS_BUSINESS_ERROR}

}

const (
	HTTP_STATUS_BUSINESS_ERROR = 599
)

// AUTH ERRORS
const (
	ERR_AUTHORIZATION_COMMON DiscordBotError = http.StatusUnauthorized*100 + iota //40100
	ERR_AUTHORIZATION_JWT
	ERR_AUTHORIZATION_TOKEN_MISSING
	ERR_AUTHORIZATION_TOKEN_INVALID
	ERR_AUTHORIZATION_TOKEN_EXPIRED
	ERR_AUTHORIZATION_NOT_KYC
)

// VALIDATION ERRORS
const (
	ERR_INVALID_REQUEST_COMMON DiscordBotError = http.StatusBadRequest*100 + iota //40000
	ERR_INVALID_APP_ID
	ERR_INVALID_ADDRESS
	ERR_INVALID_CHAIN
	ERR_INVALID_CONTRACT_TYPE
	ERR_INVALID_URL
	ERR_INVALID_METADATA_ID
	ERR_INVALID_MINT_AMOUNT
	ERR_INVALID_MINT_AMOUNT_721
	ERR_INVALID_TOKEN_ID
	ERR_INVALID_CONTRACT_TYPE_UNMATCH
	ERR_INVALID_PAGINATION
)

// RESOURCE CONFLICT ERRORS
const (
	ERR_CONFLICT_COMMON DiscordBotError = http.StatusConflict*100 + iota //40900
	ERR_CONFLICT_COMPANY_EXISTS
)

// RATELIMIT ERRORS
const (
	ERR_TOO_MANY_REQUEST_COMMON DiscordBotError = http.StatusTooManyRequests*100 + iota //42900
)

// INTERNAL SERVER ERRORS
const (
	ERR_INTERNAL_SERVER_COMMON DiscordBotError = http.StatusInternalServerError*100 + iota //50000
	ERR_INTERNAL_SERVER_DB
	ERR_INTERNAL_SERVER_DB_NOT_FOUND
)

// BUSINESS ERRORS
const (
	ERR_BUSINESS_COMMON DiscordBotError = HTTP_STATUS_BUSINESS_ERROR*100 + iota //60000
	ERR_NO_SPONSOR
	ERR_NO_SPONSOR_BALANCE
	ERR_MINT_LIMIT_EXCEEDED
	ERR_DEPLOY_LIMIT_EXCEEDED
	ERR_UPLOADE_FILE_LIMIT_EXCEEDED
)

func GetDiscordBotOthersErrCode(httpStatusCode int) int {
	return httpStatusCode * 100
}
