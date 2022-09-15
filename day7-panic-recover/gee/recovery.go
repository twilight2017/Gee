package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	//Callers用来返回调用栈的程序计数器
	n := rumtime.Callers(3, pcs[:]) //skip first 3 caller

	var str strings.Builder
	str.WruteSting(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)  //获取调用该函数的文件名和行号
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
