package filter_builder

import (
	"fmt"
	"time"
	"web_server/http_context"
)

// 为了实现对于责任链的处理
type FilterBuilder func(f Filter) Filter

// 业务处理函数
type Filter func(c http_context.HttpContext)

var _ FilterBuilder = TimeFilterBuilder

// test 创建一个用于记录运行时间的FilterBuilder
func TimeFilterBuilder(f Filter) Filter {
	return func(c http_context.HttpContext) {
		start := time.Now().UnixNano()
		f(c)
		end := time.Now().UnixNano()
		fmt.Printf("run time is:%d\n", end-start)
	}
}
