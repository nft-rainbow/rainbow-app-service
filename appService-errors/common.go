package appService_errors

import "net/http"

func init() {
	// VALIDATION ERRORS
	RainbowAppServiceErrorInfos[ERR_INVALID_REQUEST_COMMON] = RainbowAppServiceErrorInfo{"Invalid request", http.StatusBadRequest}
	RainbowAppServiceErrorInfos[ERR_INVALID_PAGINATION] = RainbowAppServiceErrorInfo{"Invalid page or limit", http.StatusBadRequest}

	// TIME RESTRAIN ERRORS
	RainbowAppServiceErrorInfos[ERR_TIME_EXPIRED_COMMON] = RainbowAppServiceErrorInfo{"This activity has been expired", http.StatusBadRequest}
	RainbowAppServiceErrorInfos[ERR_TIME_NOT_ARRIVED_COMMON] = RainbowAppServiceErrorInfo{"This activity has not been opened", http.StatusBadRequest}

	// INTERNAL SERVER ERRORS
	RainbowAppServiceErrorInfos[ERR_INTERNAL_SERVER_COMMON] = RainbowAppServiceErrorInfo{"Internal Server error", http.StatusInternalServerError}
}

// VALIDATION ERRORS
const (
	ERR_INVALID_REQUEST_COMMON RainbowAppServiceError = http.StatusBadRequest*100 + iota //40000
	ERR_INVALID_PAGINATION
)

// INTERNAL SERVER ERRORS
const (
	ERR_INTERNAL_SERVER_COMMON RainbowAppServiceError = http.StatusInternalServerError*100 + iota //50000
)

// TIME RESTRAIN ERRORS
const (
	ERR_TIME_EXPIRED_COMMON RainbowAppServiceError = 600*100 + iota //60000
	ERR_TIME_NOT_ARRIVED_COMMON
)

func GetAppServiceOthersErrCode(httpStatusCode int) int {
	return httpStatusCode * 100
}
