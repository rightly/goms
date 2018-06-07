package manager

import (
	"github.com/labstack/echo"
	"net/http"
)

type api struct {
	name        string
	get		string
	handlerFunc echo.HandlerFunc
}

func newApi() *api {

	return &api{
		name:"",
	}
}

// setDefaultHeader 는 모든 response 의 default 로 추가 / 제거 되어야하는 헤더를 설정한다.
func setDefaultHeader(c *echo.Context) {
	ctx := *c
	header := ctx.Response().Header()

	header.Set("Server", "goms-manager/0.1")
	header.Del("Vary")
	header.Del("X-Request-ID")

}

////
// goms handlerFunc List
////

// metricHandler 는 metric 관련 handler를 등록한다.
func metricHandler(e *echo.Echo)  {
	apiName := "/metrics"

	e.PUT(apiName, createMetric)
	e.GET(apiName, findMetric)
	e.HEAD(apiName, metric)
}

////
// echo handlerFunc List
////

// index 는 goms manager admin 페이지를 리턴한다.
func index(c echo.Context) error {
	return c.File("index.html")
}

// Metrics 수집 관련 Handler
// GET, PUT, HEAD

// createMetric 은 각 서버로부터 metric을 PUT 한다.
func createMetric(c echo.Context) error {

	reqId := c.Response().Header().Get("X-Request-Id")
	remote := c.Request().Header.Get("X-Server")
	message := remote + " metrics created success"
	res := setResponse(reqId, message)

	setDefaultHeader(&c)
	c.Response().Header().Set("Cache-Control", "no-cache, no-store")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	return c.JSON(http.StatusCreated, res)
}

// findMetric 은 수집한 metric 을 리턴한다.
func findMetric(c echo.Context) error {
	setDefaultHeader(&c)

	return c.String(http.StatusOK, c.Request().Header.Get("X-Server"))
}

// metric 은 HEAD / OPTIONS 등의 요청에 응답한다.
func metric(c echo.Context) error {
	setDefaultHeader(&c)

	return c.String(http.StatusOK, "")
}

// errorHandler 는 goms manager 의 http error 를 핸들링한다
func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	reqId := c.Response().Header().Get("X-Request-Id")
	remote := c.Request().Header.Get("X-Server")
	res := setResponse(reqId, "")

	setDefaultHeader(&c)
	switch code {
	case 400:
		res.Message = remote + "request is bad request"
	case 404:
		res.Message = "please check api resource name"
	default:
		res.Message = "manager is not available"
	}
	c.JSON(code, res)
}