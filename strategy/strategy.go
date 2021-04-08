package strategy

import "go-design-pattern-for-example/factory"

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