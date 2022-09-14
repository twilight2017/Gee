package gee

RouterGroup struct{
	prefix string //前缀
	middlerwares []Handlerfunc //support middleware
	parent *RouterGroup //当前分组的父亲
	engine *Engine //所有分组的engine，负责协调所有资源
}

Engine struct{
	*RouterGroup
	router *router
	groups []*RouterGroup
}

// New is the constructor of gee.Engine
func New() *Engine{
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouteGroup{engine:engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
}

//Group is defined to create a new RouteGroup
//remember all groups share the same Engine instance
//新建一个子组，并加入到组列表中
func (group *RouterGroup) Group(prefix string) *RouterGroup{
	engine := group.engine
	newGroup = &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
}

func(group *RouterGroup) addRoute(method string, comp string, handler Handlerfunc){
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler Handlerfunc){
	group.addRoute("GET", pattern, handler)
}

//POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler Handlerfunc){
	group.addRoute("POST", pattern, handler)
}