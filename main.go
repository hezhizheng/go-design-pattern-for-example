package main

import (
	"go-design-pattern-for-example/factory"
	"go-design-pattern-for-example/service/github"
	"go-design-pattern-for-example/strategy"
	"log"
)

func main() {
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

}
