package common

import (
	"net/http"
	"strings"
)


//声明一个新的数据类型（函数类型）
type FilterHandle func(rw http.ResponseWriter,req *http.Request)error


//拦截器结构体
type filter struct {
	filterMap map[string]FilterHandle
}

func NewFilter() *filter {
	return &filter{filterMap: make(map[string]FilterHandle)}
}

func (f *filter)RegisterFilterUrl(url string,handler FilterHandle)  {
	f.filterMap[url]=handler
}

func (f *filter) GetFilterHandle(url string) FilterHandle{
	return f.filterMap[url]
}

type WebHandle func(rw http.ResponseWriter,req *http.Request)

func (f *filter) Handle(webhandle WebHandle) func(rw http.ResponseWriter,r *http.Request){
	return func(rw http.ResponseWriter, r *http.Request) {
		for path,handle:=range f.filterMap{
			if strings.Contains(r.RequestURI,path) {
				err:=handle(rw,r)
				if err!=nil {
					rw.Write([]byte(err.Error()))
					return
				}
				// 跳出循环
				break
			}
		}
		webhandle(rw,r)
	}
}


