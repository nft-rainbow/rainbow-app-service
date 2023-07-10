package enums

type BatchMintStatus int

const (
	BATCH_MINT_STATUS_INIT BatchMintStatus = iota
	BATCH_MINT_STATUS_CREATING_WALLET
	BATCH_MINT_STATUS_CREATE_WALLET_DONE
	BATCH_MINT_STATUS_MINT
	BATCH_MINT_STATUS_FAIL
)
