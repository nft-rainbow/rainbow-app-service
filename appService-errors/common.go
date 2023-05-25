package appService_errors

import "net/http"

const (
	StatusBussinessError = 599
)

func init() {
	// VALIDATION ERRORS
	RainbowAppServiceErrorInfos[ERR_INVALID_REQUEST_COMMON] = RainbowAppServiceErrorInfo{Message: "Invalid request", HttpStatusCode: http.StatusBadRequest}
	RainbowAppServiceErrorInfos[ERR_INVALID_PAGINATION] = RainbowAppServiceErrorInfo{Message: "Invalid page or limit", HttpStatusCode: http.StatusBadRequest}

	// INTERNAL SERVER ERRORS
	RainbowAppServiceErrorInfos[ERR_INTERNAL_SERVER_COMMON] = RainbowAppServiceErrorInfo{Message: "Internal Server error", HttpStatusCode: http.StatusInternalServerError}
	RainbowAppServiceErrorInfos[ERR_TOO_MANY_REQUEST_COMMON] = RainbowAppServiceErrorInfo{Message: "Too many request", HttpStatusCode: http.StatusTooManyRequests}

	// BUSINESS ERRORS
	RainbowAppServiceErrorInfos[ERR_BUSINESS_NOT_BIND_WALLET] = RainbowAppServiceErrorInfo{Message: "Not bind wallet address", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_ACTIVITY_NOT_EXIST] = RainbowAppServiceErrorInfo{Message: "The activity is not exist", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_TIME_EARLY] = RainbowAppServiceErrorInfo{Message: "This activity has not been opened", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_TIME_EXPIRED] = RainbowAppServiceErrorInfo{Message: "This activity has been expired", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_ACTIVITY_MAX_AMOUNT_ARRIVED] = RainbowAppServiceErrorInfo{Message: "This activity max mint amount arrived", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_PERSONAL_MAX_AMOUNT_ARRIVED] = RainbowAppServiceErrorInfo{Message: "Your max mint amount arrived", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_MISS_VISPER] = RainbowAppServiceErrorInfo{Message: "This activity need visper password", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_VISPER_WRONG] = RainbowAppServiceErrorInfo{Message: "Wrong visper password", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSINESS_NO_MINT_PERMISSIION] = RainbowAppServiceErrorInfo{Message: "No mint permission", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BUSNISS_ACTIVITY_CONFIG_WRONG] = RainbowAppServiceErrorInfo{Message: "Activity config wrong", HttpStatusCode: StatusBussinessError}

	RainbowAppServiceErrorInfos[ERR_BIND_ADDRESS_WRONG_FORMAT] = RainbowAppServiceErrorInfo{Message: "Address format is not correct", HttpStatusCode: StatusBussinessError}
	RainbowAppServiceErrorInfos[ERR_BIND_ADDRESS_OTHER] = RainbowAppServiceErrorInfo{Message: "Failed to bind address, unkown reason occured", HttpStatusCode: StatusBussinessError}
}

// VALIDATION ERRORS
const (
	ERR_INVALID_REQUEST_COMMON RainbowAppServiceError = http.StatusBadRequest*100 + iota //40000
	ERR_INVALID_PAGINATION
)

// RATELIMIT ERRORS
const (
	ERR_TOO_MANY_REQUEST_COMMON RainbowAppServiceError = http.StatusTooManyRequests*100 + iota //42900
)

// INTERNAL SERVER ERRORS
const (
	ERR_INTERNAL_SERVER_COMMON RainbowAppServiceError = http.StatusInternalServerError*100 + iota //50000
)

// BUSINESS ERRORS
const (

	// MINT
	ERR_BUSINESS_COMMON RainbowAppServiceError = StatusBussinessError*100 + iota //60000
	// ❌未绑定钱包地址！
	ERR_BUSINESS_NOT_BIND_WALLET
	// ❌此活动不支持或不存在。
	ERR_BUSINESS_ACTIVITY_NOT_EXIST
	// ❌活动1234铸造失败。活动未开始。
	ERR_BUSINESS_TIME_EARLY
	// ❌活动1234铸造失败。活动已过期。
	ERR_BUSINESS_TIME_EXPIRED
	// ❌活动1234铸造失败。NFT已领取完。
	ERR_BUSINESS_ACTIVITY_MAX_AMOUNT_ARRIVED
	// ❌活动1234铸造失败。超过活动领取限制。
	ERR_BUSINESS_PERSONAL_MAX_AMOUNT_ARRIVED
	// ❌活动1234铸造失败。需要领取口令。使用帮助指令查看如何获取领取口令。
	ERR_BUSINESS_MISS_VISPER

	ERR_BUSINESS_VISPER_WRONG
	// ❌活动1234铸造失败。无领取资格。
	ERR_BUSINESS_NO_MINT_PERMISSIION

	ERR_BUSNISS_ACTIVITY_CONFIG_WRONG

	// ❌绑定失败，地址格式不正确。
	ERR_BIND_ADDRESS_WRONG_FORMAT
	// ❌绑定失败，发生未知错误，请重新绑定。
	ERR_BIND_ADDRESS_OTHER
)

/*
@触发用户

⭕️准备铸造活动1234NFT，请耐心等待......

在指定频道@触发用户

@触发用户

⭕活动1234铸造成功！请到区块链浏览器或绑定的钱包查看。

@触发用户

❌未绑定钱包地址！

@触发用户

❌活动1234铸造失败。活动已过期。

@触发用户

❌活动1234铸造失败。活动未开始。

@触发用户

❌活动1234铸造失败。NFT已领取完。

@触发用户

❌活动1234铸造失败。超过活动领取限制。

@触发用户

❌活动1234铸造失败。需要领取口令。使用帮助指令查看如何获取领取口令。

@触发用户

❌活动1234铸造失败。无领取资格。
*/

func GetAppServiceOthersErrCode(httpStatusCode int) int {
	return httpStatusCode * 100
}
