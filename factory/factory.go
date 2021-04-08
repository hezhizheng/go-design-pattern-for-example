package factory

import (
	"go-design-pattern-for-example/service/gitee"
	"go-design-pattern-for-example/service/github"
	"sync"
)

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