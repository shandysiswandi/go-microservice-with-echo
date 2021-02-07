package welcomes

import (
	"go-rest-echo/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type web struct {
	usecase Usecase
}

// NewWeb is
func NewWeb(u Usecase) WelcomeDelivery {
	return &web{u}
}

func (web) Home(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// response
	return c.Success(http.StatusOK, "Welcome to our API", nil)
}

func (w *web) MonitorDatabase(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecases
	mysql := w.usecase.CheckDatabaseMysql()
	postgresql := w.usecase.CheckDatabasePostgresql()
	mongo := w.usecase.CheckDatabaseMongo()

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Databases", map[string]interface{}{
		"mysql":      mysql,
		"postgresql": postgresql,
		"mongo":      mongo,
	})
}

func (w *web) MonitorService(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecases
	sentry := w.usecase.CheckServiceSentry()
	redis := w.usecase.CheckServiceRedis()

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Services", map[string]interface{}{
		"sentry": sentry,
		"redis":  redis,
	})
}