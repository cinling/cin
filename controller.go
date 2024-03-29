package cin

import "github.com/cinling/cin/core/models"

// base controller [websocket/http/console]
type BaseController struct {
	models.BaseController
}

// get application instance
func (controller *BaseController) GetApp() *application {
	return controller.GetAppInterface().(*application)
}
