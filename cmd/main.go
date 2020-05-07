package main

import (
	"flag"
	"monitorclientinstall/install"
)

// Author: Gerry
// 2020-05-07

func main() {
	var (
		s string
		v string
		r string
	)
	// 传入引入参数
	flag.StringVar(&s, "s", "nil", "Install related client services, support <mysql|redis|Linux>")
	flag.StringVar(&r, "r", "nil", "Unloading service, support <mysql|redis|Linux>")
	flag.StringVar(&v, "v", "nil", "version : 0.1.4")
	flag.Parse()
	switch {
	case s == "linux":
		install.NodeInstall()
	case s == "redis":
		install.RedisInstall()
	case s == "mysql":
		install.MysqlInstall()
	case s == "kafka":
		install.KafkaInstall()
	case r == "linux":
		install.Noderemove()
	case r == "redis":
		install.Redisremove()
	case r == "mysql":
		install.Mysqlremove()
	case r == "kafka":
		install.Kafkaremove()
	default:
		flag.Usage()
	}
}