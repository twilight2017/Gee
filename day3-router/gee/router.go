package gee

import "net/http"
import "strings"

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	// rquest info
	Path   string
	Method string
	Params map[string]string

	//response info
	StatusCode int
}
type HandlerFunc func(*Context)
type router struct {
	roots    map[string]*node       //每种请求方式的trie树根节点
	handlers map[string]HandlerFunc //存储每种请求方式的HandlerFunc
}

//roots key eg, roots['GET'] roots['POST']
//handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

//router construct function
//该构造函数在这里只做内存分配
func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//Only one * is allowed,将字符串解析成字符列表
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break //符合精确匹配原则,直接break是一旦符合精确匹配原则，不需要再解析后面的部分
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern) //将路径字符串解析成字符串序列

	key := method + "-" + pattern
	//去roots中进行匹配
	_, ok := r.roots[method]
	if !ok {
		//没有该键值对时新建这一条路由映射规则
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0) //在trie树中从第0层插入这条匹配规则，这里也体现了parts只是pattern的拆分列表
	r.handlers[key] = handler                 //在handler映射中按key添加这条映射
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	//1.将path拆成parts
	/*
	  2.在内存中申请一个map来存放参数
	  3.去roots中拿到该方法的trie树
	*/
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	//从第0层开始搜索parse后的parts序列
	//返回search成功后的完整路劲
	n := root.search(searchParts, 0)
	//递归查找节点，全部查到时返回的最后一个节点不为空

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index] //将:lang存储为python这样的映射过程
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params //将解析出来的路由参数赋值给c.Params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c) //按参数调用执行器
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
