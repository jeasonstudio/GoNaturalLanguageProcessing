# GoNaturalLanguageProcessing
Natural language processing by Golang

## 使用

### 安装官方 Go 语言支持库并启动服务
```Golang
go get github.com/davemeehan/Neo4j-GO
```

```bash
$ neo4j console
$ chrome http://localhost:7474/browser/
```

### 函数说明

```Golang
// 通用 neo4j 配置
type Neo4j struct {
	Method     string // which http method
	StatusCode int    // last http status code received
	URL        string
	Username   string
	Password   string
}
```

```Golang
// 错误
type Error struct {
	List map[int]error
	Code int
}
```

```Golang
// 当来自 neo4j 的数据被存储时使用
type NeoTemplate struct {
	ID                  uint64
	Relationships       string
	RelationshipsOut    string
	RelationshipsIn     string
	RelationshipsAll    string
	RelationshipsCreate string
	Data                map[string]interface{}
	Traverse            string
	Property            string
	Properties          string
	Self                string
	Extensions          map[string]interface{}
	Start               string        // relationships & traverse // returns both obj & string
	End                 string        // relationships & traverse // returns both obj & string
	Type                string        // relationships & traverse
	Indexed             string        // index related
	Length              string        // traverse framework
	Nodes               []interface{} // traverse framework
	TRelationships      []interface{} // traverse framework
}
```

```Golang
// 链接 url 类似： http://127.0.0.1:7474/db/data
// @parameter user passwd 用户名 密码
// @returns (*Neo4j, error)
func NewNeo4j(u string, user string, passwd string) (*Neo4j, error)
```

```Golang
// 返回属性值的字符串和作为错误引发的任何错误
func (this *Neo4j) GetProperty(id uint64, name string) (string, error)
```

```Golang
// CreateNode 由 map[string]string 生成 neo4j 节点，或返回错误
func (this *Neo4j) CreateNode(data map[string]string) (tmp *NeoTemplate, err error) {
```

```Golang
// GetNode 由 id 返回节点 NeoTemplate 或错误
func (this *Neo4j) GetNode(id uint64) (tmp *NeoTemplate, err error)
```

```Golang
// DelNode 根据节点 id 删除节点
func (this *Neo4j) DelNode(id uint64) error
```

```Golang
// DelProperty 删除 id 对应节点，或可扩展为删除 id 对应节点对应属性
func (this *Neo4j) DelProperty(id uint64, s string) error
```

```Golang
// CreateProperty 根据节点 id 添加属性kv，replace 应该为 FALSE ，除非你想删除其他所有属性
func (this *Neo4j) CreateProperty(id uint64, data map[string]string, replace bool) error
```

```Golang
// GetProperties 通过 id 获取节点类型和错误
func (this *Neo4j) GetProperties(id uint64) (tmp *NeoTemplate, err error)
```

```Golang
// GetProperty 根据节点 id 和属性名获取属性值和错误
func (this *Neo4j) GetProperty(id uint64, name string) (string, error)
```

```Golang
// SetProperty 根据节点 id 修改属性内容，除非你想删除其他属性 ，否则 replace 应该为 false
func (this *Neo4j) SetProperty(id uint64, data map[string]string, replace bool) error
```

```Golang
// CreateRelationship 创建两节点关系 src -> dst 返回任何错误 
// @parameter: data 关系 kv map[string]string
// @parameter: rType 关系名 string
func (this *Neo4j) CreateRelationship(src uint64, dst uint64, data map[string]string, rType string) error
```

```Golang
// CreateIdx
// @parameter: idxType relationship or node
func (this *Neo4j) CreateIdx(id uint64, key string, value string, cat string, idxType string) error
```

```Golang
// Traverse 类似关系型数据库中的建表，通过一种算法，使所有节点分别位于各自图下
// @parameter: {
// 		id: 节点
//		returnType: 	
//		order:
// 		uniqueness: 
//		relationships:
// 		depth: 节点深度
//		prune: 修剪
//		filter: 过滤
//}
// 返回节点类型数组
func (this *Neo4j) Traverse(id uint64, returnType string, order string, uniqueness string, relationships map[string]string, depth int, prune map[string]string, filter map[string]string) (map[int]*NeoTemplate, error)
```

```Golang
// TraversePath
func (this *Neo4j) TraversePath(src uint64, dst uint64, relationships map[string]string, depth uint, algo string, paths bool) (map[int]*NeoTemplate, error)
```