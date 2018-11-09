package main

import (
	"flag"

	"github.com/magicsong/glog/glog"
)

func main() {
	flag.Parse()
	glog.V(0).Infoln("这是信息")
	glog.Errorln("这是错误用例")
	glog.Warningln("这是警告用例")
	glog.Fatalln("致命错误，准备退出")
}
