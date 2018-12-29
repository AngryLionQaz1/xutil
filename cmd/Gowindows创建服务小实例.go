package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/kardianos/service"
)

type services struct {
	log service.Logger
	srv *http.Server
	cfg *service.Config
}

func (srv *services) Start(s service.Service) error {
	if srv.log != nil {
		srv.log.Info("Start run http server")
	}

	lis, err := net.Listen("tcp", ":1789")
	if err != nil {
		return err
	}
	go srv.srv.Serve(lis)
	return nil
}

func (srv *services) Stop(s service.Service) error {
	if srv.log != nil {
		srv.log.Info("Start stop http server")
	}
	return srv.srv.Shutdown(context.Background())
}

func main() {
	File, err := os.Create("http-server.log")
	if err != nil {
		File = os.Stdout
	}
	defer File.Close()

	log.SetOutput(File)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Path)
	})

	var s = &services{srv: &http.Server{Handler: http.DefaultServeMux}, cfg: &service.Config{
		Name:        "GoHttpServer",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}}

	sys := service.ChosenSystem()
	srv, err := sys.New(s, s.cfg)
	if err != nil {
		log.Fatalf("Init service error:%s\n", err.Error())
	}

	s.log, err = srv.SystemLogger(nil)
	if err != nil {
		log.Printf("Set logger error:%s\n", err.Error())
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			err := srv.Install()
			if err != nil {
				log.Fatalf("Install service error:%s\n", err.Error())
			}
		case "uninstall":
			err := srv.Uninstall()
			if err != nil {
				log.Fatalf("Uninstall service error:%s\n", err.Error())
			}
		}
		return
	}
	err = srv.Run()
	if err != nil {
		log.Fatalf("Run programe error:%s\n", err.Error())
	}
}
