package routes

import (
	"go-app/controllers"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func LoadRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
	http.Handle("/metrics", promhttp.Handler())
}
