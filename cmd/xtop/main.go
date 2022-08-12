package main

import (
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"

	//"log"
	//"os"
	//"runtime"

	//"net/http"
	//_ "net/http/pprof"

	"github.com/xhsky/xtop/global"
	"github.com/xhsky/xtop/internal/terminal"
	"github.com/xhsky/xtop/internal/web"
	_ "github.com/xhsky/xtop/pkg/logger"
)

var (
	//t uint // 间隔时间, 单位秒
	d bool // 后台web模式
	l bool // 后台web模式
	//s bool // 终端显示
)

func init() {
	flag.BoolVar(&d, "d", false, "deamon mode")
	flag.BoolVar(&l, "l", false, "list debug")
	//flag.BoolVar(&s, "s", true, "terminal mode")
	flag.UintVar(&global.T, "t", 1, "interval seconds")
}

func main() {
	flag.Parse()

	if l == true {
		runtime.GOMAXPROCS(1)
		runtime.SetMutexProfileFraction(1)
		runtime.SetBlockProfileRate(1)

		go func() {
			// 启动一个 http server, 注意 pprof 相关的 handler 已经自动注册过了
			if err := http.ListenAndServe(":6060", nil); err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}()
	}

	if d == true {
		web.Show()
	} else {
		terminal.Show()
	}
}
