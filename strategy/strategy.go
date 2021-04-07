package strategy

import "go-design-pattern-for-example/factory"

type RepoStrategyConstruct struct {
	Service *factory.RepoAPI
}

func RepoStrategy(api factory.RepoAPI) *RepoStrategyConstruct {
	return &RepoStrategyConstruct{
		Service: &api,
	}
}