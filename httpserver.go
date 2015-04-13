package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strings"
)

const (
	// Route vars are gorilla/mux paths variables.
	FileRouteVar     = "file"
	TemplateRouteVar = "template"
	TypeRouteVar     = "type"

	WebAssetsDir = "assets"
	TemplatesDir = "templates"
)

func StartHTTPServer(port int, staticsPath string) chan error {
	r := mux.NewRouter()

	// This is the asset sub-router. It routes the "/assets" path prefix.
	// Assets are found in sub-directories under /assets (i.e. css, js...)
	assetsRouter := r.PathPrefix("/assets").Methods("GET").Subrouter()
	assetsRouter.Handle("/{"+TypeRouteVar+"}/{"+FileRouteVar+"}", http.StripPrefix("/assets/", http.FileServer(http.Dir(strings.Join([]string{staticsPath, WebAssetsDir}, "/")))))

	thandler := NewTemplateHandler(staticsPath)
	r.Handle("/", thandler)
	r.Handle("/{"+TemplateRouteVar+"}", thandler)

	http.Handle("/", r)

	done := make(chan error)
	go withPanicLogging(func() {
		Log.WithField("port", port).Info("HTTP listen and serve.")
		done <- http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	})
	return done
}

type TemplateArgs struct {
}

type TemplateHandler struct {
	baseTemplate *template.Template
}

func newTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"localHostname": GetLocalHostname,
	}
}

func NewTemplateHandler(staticsPath string) *TemplateHandler {
	funcs := newTemplateFuncMap()
	return &TemplateHandler{baseTemplate: template.Must(template.New("base").Funcs(funcs).ParseGlob(strings.Join([]string{staticsPath, TemplatesDir, "*.html"}, "/")))}
}

func (handler *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	template := vars[TemplateRouteVar]

	if template == "" {
		template = "index.html"
	}

	args := &TemplateArgs{}

	if err := handler.baseTemplate.ExecuteTemplate(w, template, args); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
