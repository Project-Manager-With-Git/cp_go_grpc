package main

import (
	"os"

	sdk "cp_go_grpc/echo_sdk"
	serv "cp_go_grpc/echo_serv"

	log "github.com/Golang-Tools/loggerhelper"
	s "github.com/Golang-Tools/schema-entry-go"
)

func main() {
	root, err := s.New(&s.EntryPointMeta{Name: "cp_go_grpc", Usage: "cp_go_grpc serv|test [options] "})
	if err != nil {
		log.Error("init root node err", log.Dict{"err": err.Error()})
		os.Exit(2)
	}
	serv, err := s.New(&s.EntryPointMeta{Name: "serv", Usage: "cp_go_grpc serv [options]"}, &serv.ServNode)
	if err != nil {
		log.Error("create serv node get error", log.Dict{"err": err.Error()})
		os.Exit(2)
	}
	test, err := s.New(&s.EntryPointMeta{Name: "test", Usage: "cp_go_grpc test [options]"}, &sdk.TestNode)
	if err != nil {
		log.Error("create test node get error", log.Dict{"err": err.Error()})
		os.Exit(2)
	}
	root.RegistSubNode(serv)
	root.RegistSubNode(test)
	root.Parse(os.Args)
}
