package success

import (
	"fmt"
	"runtime"

	"github.com/carlescere/scheduler"
)

func Main() {
	// 5秒に1回 success!! と出力させる
	scheduler.Every(5).Seconds().Run(printSuccess)
	runtime.Goexit()
}

func printSuccess() {
	fmt.Println("success!!")
}
