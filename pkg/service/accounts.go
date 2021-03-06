package service

import (
	"encoding/json"
	"time"

	"github.com/Optum/dce-cli/client/operations"
	"github.com/Optum/dce-cli/configs"
	observ "github.com/Optum/dce-cli/internal/observation"
	utl "github.com/Optum/dce-cli/internal/util"
)

type AccountsService struct {
	Config      *configs.Root
	Observation *observ.ObservationContainer
	Util        *utl.UtilContainer
}

func (s *AccountsService) AddAccount(accountID, adminRoleARN string) {
	params := &operations.PostAccountsParams{
		Account: operations.PostAccountsBody{
			ID:           &accountID,
			AdminRoleArn: &adminRoleARN,
		},
	}
	params.SetTimeout(5 * time.Second)
	_, err := ApiClient.PostAccounts(params, nil)
	if err != nil {
		log.Fatalln("err: ", err)
	} else {
		log.Infoln("Account added to DCE accounts pool")
	}
}

func (s *AccountsService) RemoveAccount(accountID string) {
	params := &operations.DeleteAccountsIDParams{
		ID: accountID,
	}
	params.SetTimeout(5 * time.Second)
	_, err := ApiClient.DeleteAccountsID(params, nil)
	if err != nil {
		log.Fatalln("err: ", err)
	} else {
		log.Infoln("Account removed from DCE accounts pool")
	}
}

func (s *AccountsService) GetAccount(accountID string) {
	params := &operations.GetAccountsIDParams{
		ID: accountID,
	}
	params.SetTimeout(5 * time.Second)
	res, err := ApiClient.GetAccountsID(params, nil)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	jsonPayload, err := json.MarshalIndent(res.GetPayload(), "", "\t")
	if err != nil {
		log.Fatalln("err: ", err)
	}
	if _, err := Out.Write(jsonPayload); err != nil {
		log.Fatalln("err: ", err)
	}
}

// ListAccounts lists the accounts
func (s *AccountsService) ListAccounts() {
	params := &operations.GetAccountsParams{}
	params.SetTimeout(5 * time.Second)
	res, err := ApiClient.GetAccounts(params, nil)
	if err != nil {
		log.Fatalln("err: ", err)
	}
	jsonPayload, err := json.MarshalIndent(res.GetPayload(), "", "\t")
	if err != nil {
		log.Fatalln("err: ", err)
	}
	if _, err := Out.Write(jsonPayload); err != nil {
		log.Fatalln("err: ", err)
	}
}
