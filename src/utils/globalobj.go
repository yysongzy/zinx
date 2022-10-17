package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"zinx/src/ziface"
)

type GlobalObj struct {
	TCPServer      ziface.IServer
	Name           string
	Version        string
	IP             string
	Port           uint32
	MaxConn        uint32
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:           "Eric Server",
		Version:        "v0.1",
		IP:             "127.0.0.1",
		Port:           7777,
		MaxConn:        3,
		MaxPackageSize: 4096,
	}

	//reload()
}

func reload() {
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "/demo/server/conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
