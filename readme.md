# 一个例子对比实现PHP与Golang的三种设计模式：单例、简单工厂、策略

## 案例 | 需求 是什么？

有段时间一直在弄自用的图床工具，主要希望满足这几个要求：

1、免费

2、国内访问快

3、上传简单

找了一圈，发现使用代码托管平台 github|gitee 还是比较靠谱的，接着就有了以下两个图床工具

Go：[repo-image-hosting](https://github.com/hezhizheng/repo-image-hosting)

PHP：[repo-storage](https://github.com/hezhizheng/repo-storage)

## 如何实现(应用设计模式)？
大部分情况下我们在针对只有一种服务的情况下直接定义好类跟方法接着实现它就完事了：

```php
// 1、定义服务类
class Gitee
{
    private $token;

    public function __construct(string $token)
    {
        $this->token = $token;
    }

    public function put(array $putData)
    {
        // 处理入参
        // 请求第三方接口上传文件
        // 返回结果集
    }
}

// 2、 controller 层调用
$serve = new Gitee("token"); 
$serve->put([...]);
```

此时如果新增多一种服务我们会去添加一个新的类去实现对应的功能

```php
// 1、定义服务类
class Github
{
    private $token;

    public function __construct(string $token)
    {
        $this->token = $token;
    }

    public function put(array $putData)
    {
        // 处理入参
        // 请求第三方接口上传文件
        // 返回结果集
    }
}
```

那么如果我们要替换服务的话，就需要在controller层中使用的类给替换掉，我们使用简单工厂模式，改变这种现状。

```php
// 定义工厂类
class Facotry
{
    public function create($type,$token){
        switch ($type)
        {
        case "github":
        return new Github($token);
        case "gitee":
        return new Gitee($token);
        default :
        throw new Exception("not support type");
        }
    }
}
// 2、 controller 层调用, 这样在替换服务的时候我们只需要修改入参就好了
$serve = (new Facotry)->create("github","token");
$serve->put([...]);
```
如果新增其他服务的话，直接在工厂类中实现对应的类型就好了。
一直在switch里面case，不够优雅？再改写一下....

```php
// 定义工厂类
class Facotry
{
    private $map = [
    "github" => 'Github',
    "gitee" => 'Gitee'
];
    public function create($type,$token){
        return new $this->map[$type]($token);
    }
}
```
至此，一个简单的工厂就实现了，那么单例在哪？我们在工厂类里面稍作修改
```php
// 定义工厂类
class Facotry
{
    private static $singleton = null;
    private  $map = [
    "github" => 'Github',
    "gitee" => 'Gitee'
];

     public static function singleton()
    {
        if ( self::$singleton === null )
        {
            self::$singleton = new self();
        }
        return self::$singleton;
    }
    
    public function create($type,$token){
        return new $this->map[$type]($token);
    }
}
// 单例调用
Facotry::singleton()->create('github','token');
```
这样就实现了单例，避免我们在重复调用静态方法造成对资源的消耗

最后来到策略模式，这里需要定义interface，新增策略类Strategy
```php
// 定义接口

interface StorehouseInterface
{
    public function put(array $putData);
}

//  实际的服务类 github、gitee 分别 使用 implements 关键字 去实现接口定义的方法

class Github implements StorehouseInterface
{
    private $token;

    public function __construct(string $token)
    {
        $this->token = $token;
    }

    public function put(array $putData)
    {
        // 处理入参
        // 请求第三方接口上传文件
        // 返回结果集
    }
}

class Gitee implements StorehouseInterface
{
    private $token;

    public function __construct(string $token)
    {
        $this->token = $token;
    }

    public function put(array $putData)
    {
        // 处理入参
        // 请求第三方接口上传文件
        // 返回结果集
    }
}

// 添加策略类

class StoreStrategy
{
    public $serve;
    
    // 在构造函数中通过入参的方式注入刚刚定义的接口
    public function __construct(StorehouseInterface $storehouse)
    {
        $this->serve = $storehouse;
    }
}

// 最后在controller 中使用策略模式调用，需要替换其他服务只需要注入已经实现interface的对应服务类即可
$strategy = new StoreStrategy(new Github('token'));
$strategy->serve->put([...]);
```

## Golang 对应实现
### 为方便理解，定义了相关目录结构，具体代码可参考 [github](https://github.com/hezhizheng/go-design-pattern-for-example)
```go
.
|-- factory
|   `-- factory.go                     # 定义简单工厂类(结构体)
|-- go.mod
|-- main.go                            # 用于调用定义方法的主程序
|-- readme.md
|-- service                            # 业务实现类 
|   |-- gitee
|   |   `-- gitee.go
|   `-- github
|       `-- github.go
`-- strategy
    `-- strategy.go                    # 定义策略类(结构体)
```


### 简单工厂+单例

```go

//Singleton 是单例模式类
type SingletonFactory struct{}

// 定义接口
type RepoAPI interface {
	Put(data map[string]interface{}) map[string]interface{}
}

var singleton *SingletonFactory
var once sync.Once

// GetInstance 用于获取单例模式对象，多协程的场景下不是线程安全的，使用 sync.Once 来实现
func GetFactoryInstance() *SingletonFactory {
	once.Do(func() {
		singleton = &SingletonFactory{}
	})

	return singleton
}

// 简单工厂 创建对应的服务类
func (t *SingletonFactory) FactoryCreate(cate, token string) RepoAPI {
	if cate == "github" {
		return &github.GithubConstruct{
			Token: token,
		}
	} else if cate == "gitee" {
		return &gitee.GiteeConstruct{
			Token: token,
		}
	}
	return nil
}

```

### 策略模式

```go
// 策略模式类，定义 service 属性 实际为具体的业务类
type RepoStrategyConstruct struct {
    Service *factory.RepoAPI
}

// 定义策略方法 入参为指定的接口(具体的业务类)
func RepoStrategy(api factory.RepoAPI) *RepoStrategyConstruct {
        return &RepoStrategyConstruct{
        Service: &api,
    }
}
```

### 相关业务/服务类
```go
// 都要实现 接口 RepoAPI 定义的方法

type GiteeConstruct struct {
	Token string
}

func (t *GiteeConstruct) Put(data map[string]interface{})  map[string]interface{} {
	log.Println("gitee",data,t.Token)
	return nil
}

// ===========================================================

type GithubConstruct struct {
	Token string
}

func (t *GithubConstruct) Put(data map[string]interface{})  map[string]interface{} {
	log.Println("github",data,t.Token)
	return nil
}


// =======================控制层调用=======================

// 简单工厂调用
log.Println("单例 | 简单工厂模式调用")
instance := factory.GetFactoryInstance()
create := instance.FactoryCreate("gitee","gitee-token")
create.Put(nil)

// 策略模式调用
log.Println("策略模式调用")
repoStrategy := strategy.RepoStrategy(&github.GithubConstruct{
	Token: "github-token",
})
service := *repoStrategy.Service
service.Put(nil)

```