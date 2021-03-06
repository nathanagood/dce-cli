package service

import (
	"github.com/Optum/dce-cli/configs"
	observ "github.com/Optum/dce-cli/internal/observation"
	utl "github.com/Optum/dce-cli/internal/util"
)

// ServiceContainer is a service that injects its config and util into other services
type ServiceContainer struct {
	Config      *configs.Root
	Observation *observ.ObservationContainer
	Util        *utl.UtilContainer
	Deployer
	Accounter
	Leaser
	Initer
	Authenticater
	Usager
}

var log observ.Logger
var ApiClient utl.APIer
var Out observ.OutputWriter

// New returns a new ServiceContainer given config
func New(config *configs.Root, observation *observ.ObservationContainer, util *utl.UtilContainer) *ServiceContainer {

	log = observation.Logger
	ApiClient = util.APIer
	Out = observation.OutputWriter

	serviceContainer := ServiceContainer{
		Config:        config,
		Observation:   observation,
		Util:          util,
		Deployer:      &DeployService{Config: config, Util: util},
		Accounter:     &AccountsService{Config: config, Util: util},
		Leaser:        &LeasesService{Config: config, Util: util},
		Initer:        &InitService{Config: config, Util: util},
		Authenticater: &AuthService{Config: config, Util: util},
		Usager:        &UsageService{Config: config, Util: util},
	}

	return &serviceContainer
}

type DeployOverrides struct {
	AWSRegion                         string
	GlobalTags                        []string
	Namespace                         string
	BudgetNotificationFromEmail       string
	BudgetNotificationBCCEmails       []string
	BudgetNotificationTemplateHTML    string
	BudgetNotificationTemplateText    string
	BudgetNotificationTemplateSubject string
	DCEVersion                        string
	// Location of the DCE terraform module
	DCEModulePath string
}
type Deployer interface {
	Deploy(input *DeployConfig) error
	PostDeploy(input *DeployConfig) error
}

type Usager interface {
	GetUsage(startDate, endDate float64)
}

type Accounter interface {
	AddAccount(accountID, adminRoleARN string)
	RemoveAccount(accountID string)
	GetAccount(accountID string)
	ListAccounts()
}

type LeaseLoginOptions struct {
	CliProfile  string
	OpenBrowser bool
	PrintCreds  bool
}

type Leaser interface {
	CreateLease(principalID string, budgetAmount float64, budgetCurrency string, email []string, expiresOn string)
	EndLease(accountID, principalID string)
	LoginByID(leaseID string, opts *LeaseLoginOptions)
	Login(opts *LeaseLoginOptions)
	ListLeases(acctID, principalID, nextAcctID, nextPrincipalID, leaseStatus string, pagLimit int64)
	GetLease(leaseID string)
}

type Initer interface {
	InitializeDCE()
}

type Authenticater interface {
	Authenticate() error
}
