package main

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/Sirupsen/logrus"
	restful "github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)


func report_version(r *restful.Request, w *restful.Response) { io.WriteString(w, "itrain V. 0.0.0") }


type resultType struct {
	Value string
	Date  string
}

func report_inverter(r *restful.Request, w *restful.Response) {
	sd := r.QueryParameter("sd") // startdate, "day/month/year"
	res := resultType{"x", sd}
	w.WriteAsJson(res)
}

func main() {
	ws := new(restful.WebService)
	ws.Path("/itrain").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/version").
		To(report_version).
		Doc("get itrain version number").
		Operation("version"))

	ws.Route(ws.GET("/report_inverter").
		To(report_inverter).
		Doc("get inverter report data").
		Operation("report_inverter").
		Param(ws.QueryParameter("sd", "startdate")))

	restful.Add(ws)

	config := swagger.Config{
		WebServices:    restful.DefaultContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: "/",
		ApiPath:        "/apidocs.json",
		// specifiy where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "./swaggerui"}
	swagger.RegisterSwaggerService(config, restful.DefaultContainer)

	var (
		hostname string = "127.0.0.1"
		port     int    = 8080
	)
	// Listen on hostname:port
	fmt.Printf("Listening on %s:%d...\n", hostname, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", hostname, port), nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}

}
