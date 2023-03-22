package main

import (
	"github.com/NYTimes/gziphandler"
	config "github.com/calebtracey/config-yaml"
	"github.com/calebtracey/mind-your-business-api/internal/facade"
	"github.com/calebtracey/mind-your-business-api/internal/routes"
	"github.com/calebtracey/mind-your-business-api/internal/routes/endpoints"
	log "github.com/sirupsen/logrus"
)

const configPath = "dev_config.yaml"

type Application struct {
	Initializer InitializerI
	Config      *config.Config
}

func init() {
	log.Infoln("initializing...")
}

//	@title			Mind Your Business API
//	@version		1.0
//	@description	This is a development MYB server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:6080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	defer panicQuit()

	appConfig := config.New(configPath)
	appRouter := &endpoints.Router{Service: new(facade.Service)}

	if err := new(Initializer).Database(appConfig, appRouter.Service.(*facade.Service)); err != nil {
		log.Errorf("failed to initialize database: %s", err)
		panicQuit()
	}

	log.Fatal(listenAndServe(appConfig.Port.Value, gziphandler.GzipHandler(
		routes.Handler{Router: appRouter}.Routes(),
	)),
	)
}

func panicQuit() {
	if r := recover(); r != nil {
		log.Errorf("I panicked and am quitting: %v", r)
		log.Error("I should be alerting someone...")
	}
}
