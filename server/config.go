package main

import (
	"os"
	"time"
	"log"
	"gopkg.in/yaml.v2"
)

var config Config
type Faissconfig struct {
	Description string
	Metric string
	Nprobe int
	Dimension int
	Syncinterval time.Duration
	Collections []string
}

type Dbconfig struct {
	Capacity uint64
}

type Replicaonfig struct {
	Id int
	Listen string
	Primary string
}

type Config struct {
	Process struct {
		Logfile string
		Pidfile string
		Daemon bool
	}
	Http struct {
		MaxConnections int
		Port int
		HttpTimeout int
	}
	Db struct {
		Dbpath string
		Faiss Faissconfig
		Metadb Dbconfig
		Datadb Dbconfig
		Iddb Dbconfig
		Oplogdb Dbconfig
	}
	Oplog struct {
		Term int
	}
	Feature struct {
		Listen string
	}
	Replica Replicaonfig
}

func loadConfig() {
	configFile := "config.yml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	data, err := ReadFile(configFile)
	if err != nil {
    log.Fatalf("loadConfig() err1 %v", err)
	}
	config = Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
    log.Fatalf("loadConfig() err %v", err)
	}
}
