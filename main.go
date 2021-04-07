package main

import (
	"go-design-pattern-for-example/factory"
	"go-design-pattern-for-example/service/github"
	"go-design-pattern-for-example/strategy"
	"log"
)

func main() {
	// 简单工厂调用
	c := factory.GetFactoryInstance()
	log.Println("单例 | 简单工厂模式调用", c.FactoryCreate("gitee","gitee token").Put(nil))

	// 策略模式调用
	x := strategy.RepoStrategy(&github.GithubConstruct{
		Token: "github token",
	})

	q := *x.Service

	log.Println("策略调用",q.Put(nil))
}
