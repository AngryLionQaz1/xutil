package jar

import (
	"fmt"
	"github.com/go-chi/chi"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strings"
	"xutil/cmd/springboot/gogs4/bean"
	"xutil/cmd/springboot/gogs4/util"
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
	router.Route("/{path}{port}/{jar}/{arguments}", func(r chi.Router) {
		r.Post("/", jr.Update)
	})

	return router
}

func (jr JarResource) Update(w http.ResponseWriter, r *http.Request) {
	project := parameters(r)
	project.Update()
	w.Write([]byte("success"))
}

//获取参数信息
func parameters(r *http.Request) *bean.Project {
	path := chi.URLParam(r, PATH)
	port := chi.URLParam(r, PORT)
	jar := chi.URLParam(r, JAR)
	arguments := chi.URLParam(r, ARGUMENTS)
	data, _ := ioutil.ReadAll(r.Body)
	result := make(map[string]interface{})
	jsoniter.Unmarshal(data, &result)
	repository := (result["repository"]).(map[string]interface{})
	name := (repository["name"]).(string)
	git := (repository["clone_url"]).(string)
	actuator := "http://127.0.0.1:" + port + "/actuator/shutdown"
	return bean.NewProject(GetPath(path), name, jar, actuator, git, arguments)
}

//获取路径
func GetPath(path string) string {
	split := strings.Split(path, ",")
	os := util.GetOs()
	builder := strings.Builder{}
	if os == 1 {
		for i := 0; i < len(split); i++ {
			fmt.Fprint(&builder, `/`+split[i])
		}
	} else {
		for i := 0; i < len(split); i++ {
			if i == 0 {
				fmt.Fprint(&builder, split[i]+`:\`)
			} else if i == len(split)-1 {
				fmt.Fprint(&builder, split[i])
			} else {
				fmt.Fprint(&builder, split[i]+`\`)
			}

		}
	}
	return builder.String()

}
