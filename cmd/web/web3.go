package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

type Java struct {
	Path string
	Name string
	Jar  string
}

func main() {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	r.Mount("/java", javaResource{}.Routes())

	http.ListenAndServe(":3001", r)

}

type javaResource struct {
}

func (jr javaResource) Routes() chi.Router {

	router := chi.NewRouter()
	router.Route("/{path}/{name}/{jar}", func(r chi.Router) {
		r.Use(JavaCtx)
		r.Post("/", jr.Update)
	})

	return router
}

func JavaCtx(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := chi.URLParam(r, "path")
		name := chi.URLParam(r, "name")
		jar := chi.URLParam(r, "jar")
		java := Java{Path: path, Name: name, Jar: jar}
		ctx := context.WithValue(r.Context(), "java", java)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})

}

func (jr javaResource) Update(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	java := (ctx.Value("java")).(Java)
	fmt.Println(java.Path)
	data, _ := ioutil.ReadAll(r.Body)
	result := make(map[string]interface{})
	jsoniter.Unmarshal(data, &result)
	fmt.Println(result["secret"])
	i := (result["repository"]).(map[string]interface{})
	fmt.Println(i["ssh_url"])
	fmt.Println(i["clone_url"])

}
