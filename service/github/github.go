package github

import "log"

type GithubConstruct struct {
	Token string
}


func (t *GithubConstruct) Put(data map[string]interface{})  map[string]interface{} {
	log.Println("github",data,t.Token)
	return nil
}