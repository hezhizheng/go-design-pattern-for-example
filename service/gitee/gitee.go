package gitee

import "log"

type GiteeConstruct struct {
	Token string
}


func (t *GiteeConstruct) Put(data map[string]interface{})  map[string]interface{} {
	log.Println("gitee",data,t.Token)
	return nil
}