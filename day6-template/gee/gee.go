package gee

import "net/http"
import "html/template"

type Engine struct {
	*RouterGroup
	router        *router
	groups        []*RouterGroup     // store all groups
	htmlTemplates *template.Template //for html render,将所有模板加载进内存
	funcMap       template.FuncMap   //for html render，自定义模板渲染函数
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}

func (engine *Engine) LoadHTMLGlob(pattern string){
	engine.htmlTemplates = template.Must(template.New(""), Funcs(engine.funcMap), ParseGlob(pattern))
}

//create static handler
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.predix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		//Check if file exists and/or if we have permission to access it
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

//server static files
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	//Register GET handlers
	group.GET(urlPattern, handler)
}
