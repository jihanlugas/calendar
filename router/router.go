package router

import (
	"encoding/json"
	"fmt"

	"github.com/jihanlugas/calendar/app/auth"
	"github.com/jihanlugas/calendar/app/company"
	"github.com/jihanlugas/calendar/app/event"
	"github.com/jihanlugas/calendar/app/listener"
	"github.com/jihanlugas/calendar/app/photo"
	"github.com/jihanlugas/calendar/app/product"
	"github.com/jihanlugas/calendar/app/property"
	"github.com/jihanlugas/calendar/app/propertyprice"
	"github.com/jihanlugas/calendar/app/propertytimeline"
	"github.com/jihanlugas/calendar/app/unit"
	"github.com/jihanlugas/calendar/app/user"
	"github.com/jihanlugas/calendar/app/usercompany"
	"github.com/jihanlugas/calendar/app/websocket"
	"github.com/jihanlugas/calendar/config"
	"github.com/jihanlugas/calendar/constant"
	"github.com/jihanlugas/calendar/db"
	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/model"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/ws"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"net/http"

	_ "github.com/jihanlugas/calendar/docs"
)

func Init() *echo.Echo {
	router := websiteRouter()

	hubManager := ws.NewHubManager(router.Validator)

	listener.StartEventListener(hubManager)

	// repositories
	photoRepository := photo.NewRepository()
	userRepository := user.NewRepository()
	companyRepository := company.NewRepository()
	usercompanyRepository := usercompany.NewRepository()
	productRepository := product.NewRepository()
	propertyRepository := property.NewRepository()
	propertytimelineRepository := propertytimeline.NewRepository()
	unitRepository := unit.NewRepository()
	propertypriceRepository := propertyprice.NewRepository()
	eventRepository := event.NewRepository()

	// usecases
	authUsecase := auth.NewUsecase(userRepository, companyRepository, usercompanyRepository)
	photoUsecase := photo.NewUsecase(photoRepository)
	userUsecase := user.NewUsecase(userRepository, usercompanyRepository)
	companyUsecase := company.NewUsecase(companyRepository, usercompanyRepository)
	productUsecase := product.NewUsecase(productRepository)
	propertyUsecase := property.NewUsecase(propertyRepository, propertytimelineRepository, unitRepository, propertypriceRepository)
	unitUsecase := unit.NewUsecase(unitRepository)
	eventUsecase := event.NewUsecase(eventRepository)

	// handlers
	authHandler := auth.NewHandler(authUsecase)
	photoHandler := photo.NewHandler(photoUsecase)
	companyHandler := company.NewHandler(companyUsecase)
	userHandler := user.NewHandler(userUsecase)
	propertyHandler := property.NewHandler(propertyUsecase)
	productHandler := product.NewHandler(productUsecase)
	unitHandler := unit.NewHandler(unitUsecase)
	eventHandler := event.NewHandler(eventUsecase)
	websocketHandler := websocket.NewHandler(hubManager)

	if config.Debug {
		router.GET("/", func(c echo.Context) error {
			return response.Success(http.StatusOK, "Welcome", nil).SendJSON(c)
		})
		router.GET("/swg/*", echoSwagger.WrapHandler)
	}

	routerAuth := router.Group("/auth")
	routerAuth.POST("/sign-in", authHandler.SignIn)
	routerAuth.POST("/sign-out", authHandler.SignOut)
	routerAuth.GET("/init", authHandler.Init, checkTokenMiddleware)
	routerAuth.GET("/refresh-token", authHandler.RefreshToken, checkTokenMiddleware)

	routerPhoto := router.Group("/photo")
	routerPhoto.GET("/:id", photoHandler.GetById)

	routerUser := router.Group("/user", checkTokenMiddleware)
	routerUser.GET("", userHandler.Page)
	routerUser.POST("", userHandler.Create)
	routerUser.POST("/change-password", userHandler.ChangePassword)
	routerUser.PUT("/:id", userHandler.Update)
	routerUser.GET("/:id", userHandler.GetById)
	routerUser.DELETE("/:id", userHandler.Delete)

	routerCompany := router.Group("/company", checkTokenMiddleware)
	routerCompany.PUT("/:id", companyHandler.Update)

	routerProperty := router.Group("/property", checkTokenMiddleware)
	routerProperty.GET("", propertyHandler.Page)
	routerProperty.POST("", propertyHandler.Create)
	routerProperty.PUT("/:id", propertyHandler.Update)
	routerProperty.GET("/:id", propertyHandler.GetById)
	routerProperty.DELETE("/:id", propertyHandler.Delete)
	routerProperty.POST("/get-price", propertyHandler.GetPrice)

	routerProduct := router.Group("/product", checkTokenMiddleware)
	routerProduct.GET("", productHandler.Page)
	routerProduct.POST("", productHandler.Create)
	routerProduct.PUT("/:id", productHandler.Update)
	routerProduct.GET("/:id", productHandler.GetById)
	routerProduct.DELETE("/:id", productHandler.Delete)

	routerUnit := router.Group("/unit", checkTokenMiddleware)
	routerUnit.GET("", unitHandler.Page)
	routerUnit.POST("", unitHandler.Create)
	routerUnit.PUT("/:id", unitHandler.Update)
	routerUnit.GET("/:id", unitHandler.GetById)
	routerUnit.DELETE("/:id", unitHandler.Delete)

	routerEvent := router.Group("/event", checkTokenMiddleware)
	routerEvent.GET("", eventHandler.Page)
	routerEvent.GET("/timeline", eventHandler.Timeline)
	routerEvent.POST("", eventHandler.Create)
	routerEvent.PUT("/:id", eventHandler.Update)
	routerEvent.GET("/:id", eventHandler.GetById)
	routerEvent.DELETE("/:id", eventHandler.Delete)

	routerWebsocket := router.Group("/ws")
	routerWebsocket.GET("", websocketHandler.Serve)

	return router

}

func httpErrorHandler(err error, c echo.Context) {
	var errorResponse *response.Response
	code := http.StatusInternalServerError
	switch e := err.(type) {
	case *echo.HTTPError:
		// Handle pada saat URL yang di request tidak ada. atau ada kesalahan server.
		code = e.Code
		errorResponse = &response.Response{
			Status:  false,
			Message: fmt.Sprintf("%v", e.Message),
			Code:    code,
		}
	case *response.Response:
		errorResponse = e
	default:
		// Handle error dari panic
		code = http.StatusInternalServerError
		if config.Debug {
			errorResponse = &response.Response{
				Status:  false,
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
		} else {
			errorResponse = &response.Response{
				Status:  false,
				Message: response.ErrorInternalServer,
				Code:    http.StatusInternalServerError,
			}
		}
	}

	js, err := json.Marshal(errorResponse)
	if err == nil {
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, js)
	} else {
		b := []byte("{status: false, code: 500, message: \"unresolved error\"}")
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, b)
	}
}

func checkTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var err error

		userLogin, err := jwt.ExtractClaims(c.Request().Header.Get(constant.AuthHeaderKey))
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, err.Error()).SendJSON(c)
		}

		conn, closeConn := db.GetConnection()
		defer closeConn()

		var user model.User
		err = conn.Where("id = ? ", userLogin.UserID).First(&user).Error
		if err != nil {
			return response.ErrorForce(http.StatusUnauthorized, response.ErrorMiddlewareUserNotFound).SendJSON(c)
		}

		if user.PassVersion != userLogin.PassVersion {
			return response.ErrorForce(http.StatusUnauthorized, response.ErrorMiddlewarePassVersion).SendJSON(c)
		}

		c.Set(constant.TokenUserContext, userLogin)
		return next(c)
	}
}
