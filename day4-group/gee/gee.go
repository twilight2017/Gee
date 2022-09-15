package gee

import "net/http"
import "strings"

//Handler defines the request handler used by gee
type HandlerFunc func(*Context)

type router struct {
	roots    map[string]*node       //每种请求方式的trie树根节点
	handlers map[string]HandlerFunc //存储每种请求方式的HandlerFunc
}

type RouterGroup struct {
	prefix      string        //前缀
	middlewares []HandlerFunc //support middleware
	parent      *RouterGroup  //当前分组的父亲
	engine      *Engine       //所有分组的engine，负责协调所有资源
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New is the constructor of gee.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
}

//Group is defined to create a new RouteGroup
//remember all groups share the same Engine instance
//新建一个子组，并加入到组列表中
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

//POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

//Use is defined to add middleware to the group
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := NewContext(w, req)
	c.handlers = middlewares //获得该前缀下的中间件列表
	engine.router.handle(c)
}
