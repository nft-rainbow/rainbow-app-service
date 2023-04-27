package utils

import (
	"context"
	"net/http"
	"time"

	openapiclient "github.com/nft-rainbow/rainbow-sdk-go"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SendBatchBurnNFTRequest(token string, dto openapiclient.ServicesBurnBatchDto) ([]openapiclient.ModelsBurnTask, error) {
	logrus.Info("Start to Batch burn")
	resp, _, err := newClient().BurnsApi.BurnBatch(context.Background()).Authorization(token).BurnBatchDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SendCustomMintRequest(token string, dto openapiclient.ServicesCustomMintDto) (*openapiclient.ModelsMintTask, error) {
	logrus.Info("Start to mint")
	resp, _, err := newClient().MintsApi.CustomMint(context.Background()).Authorization(token).CustomMintDto(dto).Execute()
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SendCreateMetadataRequest(token string, dto openapiclient.ServicesMetadataDto) (*openapiclient.ModelsExposedMetadata, error) {
	resp, _, err := newClient().MetadataApi.CreateMetadata(context.Background()).Authorization(token).MetadataInfo(dto).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetTokenInfo(id int32, token string) (string, string, int32, error) {
	resp, _, err := newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return "", "", 0, err
	}

	for *resp.Status == 0 {
		resp, _, err = newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
		if err != nil {
			return "", "", 0, err
		}
		time.Sleep(3 * time.Second)
	}
	return *resp.TokenId, *resp.Hash, resp.GetStatus(), nil
}

func GetMintDetail(id int32, token string) (string, string, int32, error) {
	resp, _, err := newClient().MintsApi.GetMintDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return "", "", 0, err
	}

	return *resp.TokenId, *resp.Hash, resp.GetStatus(), nil
}

func GetBurnInfo(id int32, token string) (int32, string, error) {
	resp, _, err := newClient().BurnsApi.GetBurnDetail(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return 0, "", err
	}
	for *resp.Status == 0 && *resp.Hash == "" {
		resp, _, err = newClient().BurnsApi.GetBurnDetail(context.Background(), id).Authorization(token).Execute()
		if err != nil {
			return 0, "", err
		}
		time.Sleep(3 * time.Second)
	}
	return *resp.Status, *resp.Hash, nil
}

func GetContractInfo(id int32, token string) (*openapiclient.ModelsContract, error) {
	logrus.Info("Start to get contract information")
	resp, _, err := newClient().ContractApi.GetContractInfo(context.Background(), id).Authorization(token).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func GetContractProfile(address string, token string) (*openapiclient.ModelsContractRuntimeProfile, error) {
	logrus.Info("Start to get contract profile")
	resp, _, err := newClient().ContractApi.GetContractProfile(context.Background(), address).Authorization(token).Execute()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func newClient() *openapiclient.APIClient {
	configuration := openapiclient.NewConfiguration()
	configuration.HTTPClient = http.DefaultClient
	configuration.Servers = openapiclient.ServerConfigurations{
		{
			URL: viper.GetString("rainbowOpenApi") + "/v1",
		},
	}
	apiClient := openapiclient.NewAPIClient(configuration)
	return apiClient
}
