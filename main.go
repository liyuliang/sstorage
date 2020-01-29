package main

import (
	"github.com/liyuliang/sstorage/system"
	"github.com/liyuliang/utils/format"
	"github.com/liyuliang/sstorage/service"
	"fmt"
	"os"
	"flag"
	"github.com/liyuliang/sstorage/database"
)

func main() {

	//cli 中获取 -a, --auth 提交到队列中心校验身份
	//cli 中获取 -port, 开启 web 界面(获取 top 信息)
	//程序启动
	//上报 ip、top、启动时间
	//请求爬虫任务队列
	//根据任务队列优先级{队列名称:优先级}, 获取任务
	//if empty { next queue }
	//if current_queue_max_failed { next queue }
	//if no_available_queue { hold on }
	data := format.ToMap(map[string]string{
		system.SystemGateway: g,
	})

	system.Init(data)

	database.Init()
	service.Start()
}

var (
	g string
)

func init() {
	required := []string{"g"}

	flag.StringVar(&g, "g", "", "gateway url")

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })

	for _, req := range required {

		if !seen[req] {
			fmt.Fprintf(os.Stderr, "flag -%s is required \n", req)
			os.Exit(2)
		}
	}
}
