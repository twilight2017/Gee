package template

import (
	"net/http"
)

type Context struct {
	engine *Engine
}

//支持根据模板名称进行模板渲染
func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text-html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req http.Request) {
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}
