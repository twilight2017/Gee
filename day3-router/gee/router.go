package gee

type router struct {
	roots    map[string]*node       //每种请求方式的trie树根节点
	handlers map[string]HandlerFunc //存储每种请求方式的HandlerFunc
}

//roots key eg, roots['GET'] roots['POST']
//handlers key eg, handlers['GET-/p/:lang/doc'], handlers['POST-/p/book']

//router construct function
func NewRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//Only one * is allowed,将字符串解析成字符列表
func parsePattern(pattern string) []string {
	vs := string.Split(pattern, '/')

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
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	//从第0层开始搜索parse后的parts序列
	//返回search成功后的完整路劲
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
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
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
