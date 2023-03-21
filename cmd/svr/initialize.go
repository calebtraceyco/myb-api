package main

import (
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business/internal/facade"
	"github.com/joho/godotenv"
)

const OpenaiApi = "openAi"

func initializeDAO(appConfig *config.Config) (facade.Service, []error) {
	var errs []error
	if err := godotenv.Load(); err != nil {
		errs = append(errs, err)
	}
	//
	//openAiSvc, err := appConfig.Service(OpenaiApi)
	//if err != nil {
	//	errs = append(errs, err)
	//}

	return facade.Service{}, errs
}
