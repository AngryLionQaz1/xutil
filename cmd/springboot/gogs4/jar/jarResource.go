package jar

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"xutil/cmd/springboot/gogs4/bean"
)

const (
	PATH      = "path"
	PORT      = "port"
	JAR       = "jar"
	ARGUMENTS = "arguments"
)

type JarResource struct {
}

func (jr JarResource) Routes() chi.Router {

	router := chi.NewRouter()
	router.Route("/{path}/{port}/{jar}/{arguments}", func(r chi.Router) {
		r.Use(jarCtx)
		r.Post("/", jr.Update)
	})

	return router
}

func jarCtx(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := chi.URLParam(r, PATH)
		port := chi.URLParam(r, PORT)
		jar := chi.URLParam(r, JAR)
		arguments := chi.URLParam(r, ARGUMENTS)
		ctx := context.WithValue(r.Context(), "java", bean.NewJava(path, port, jar, arguments))
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (jr JarResource) Update(w http.ResponseWriter, r *http.Request) {
	project := parameters(r)
	project.Update()
}

//获取参数信息
func parameters(r *http.Request) *bean.Project {
	ctx := r.Context()
	java := (ctx.Value("java")).(bean.Java)
	data, _ := ioutil.ReadAll(r.Body)
	result := make(map[string]interface{})
	jsoniter.Unmarshal(data, &result)
	secret := (result["secret"]).(string)
	fmt.Println(secret)
	repository := (result["repository"]).(map[string]interface{})
	name := (repository["name"]).(string)
	git := (repository["ssh_url"]).(string)
	actuator := "http://127.0.0.1:" + java.Port + "/actuator/shutdown"
	return bean.NewProject(java.Path, name, java.Jar, actuator, git, java.Arguments)
}
