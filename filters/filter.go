package filters

import (
	"bo-web/boContext"
	"fmt"
	"time"
)

type FilterBuilder func(next Filter) Filter

type Filter func(c *boContext.Context)

// 检查后面是否实现前面
var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	//我希望在你的下一个实现里面主动调用filter
	return func(c *boContext.Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Println("take time", end-start)
	}
}
