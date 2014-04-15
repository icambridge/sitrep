package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/icambridge/framework"
	"github.com/icambridge/genkins"
	"github.com/icambridge/gobucket"
	"github.com/robfig/config"
	"sitrep/controller"
	"sitrep/model"
	"sitrep"
)

func main() {
	flag.Parse()

	hookProcessors := &gobucket.HookObserver{}
	hookProcessors.Add(&sitrep.Unapprove{})

	cfg, _ := config.ReadDefault("config.cfg")

	bitbucketUser, _ := cfg.String("bitbucket", "username")
	bitbucketPass, _ := cfg.String("bitbucket", "password")

	jenkinsUser, _ := cfg.String("jenkins", "username")
	jenkinsHost, _ := cfg.String("jenkins", "hostname")
	jenkinsToken, _ := cfg.String("jenkins", "token")

	glog.Info(jenkinsHost)

	mysqlUsername, _ := cfg.String("mysql", "username")
	mysqlPassword, _ := cfg.String("mysql", "password")
	mysqlDatabase, _ := cfg.String("mysql", "database")
	mysqlDsn := fmt.Sprintf("%s:%s@/%s", mysqlUsername, mysqlPassword, mysqlDatabase)

	jenkins := genkins.NewClient(jenkinsHost, jenkinsUser, jenkinsToken)
	bitbucket := gobucket.NewClient(bitbucketUser, bitbucketPass)
	memClient := memcache.New("127.0.0.1:11211")

	db, _ := sql.Open("mysql", mysqlDsn)

	buildModel := model.BuildModel{Db: db}


	base := controller.Base{Memcache: memClient, Jenkins: jenkins, Bitbucket: bitbucket}

	router := framework.NewRouter()
	router.RegisterController(controller.Home{Base: base})
	router.RegisterController(controller.Jenkins{Base: base, Model: buildModel})

	app := framework.NewApp(9090)
	app.RegisterRouter(router)
	glog.Infof("Application started and running on port %d", 9090)
	glog.Flush()
	app.Start()
}
