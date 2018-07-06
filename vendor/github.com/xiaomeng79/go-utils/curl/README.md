#### 实现一个CURL

##### 使用方法1：
    `	//以下设置不不要的可以不设置
     	curl := New() //创建对象
     	curl.SetMethod("GET") //设置请求方法
     	curl.SetUrl("https://www.baidu.com/") //设置Url
     	curl.SetHeader("Content-Type","application/json") //设置请求类型
     	curl.SetBody("")//设置请求体，
     	curl.AddHeader("test","test01")//增加请求头test的值
     	curl.AddHeader("test","test02")//增加请求头
     	err := curl.Do()`


###### curl请求的详细信息

        
    `{"method":"GET","url":"https://xxx.test.com/","status":200,"delay":321,"request":{"Header":{"Content-Type":["application/json"],"Test":["test01","test02"]},"Body":"","ContentType":"application/json","ContentLength":0},"response":{"Header":{"Connection":["keep-alive"],"Content-Length":["8"],"Content-Type":["text/html; charset=utf-8"],"Date":["Thu, 28 Jun 2018 04:21:12 GMT"],"Server":["nginx"]},"Body":"shop api","ContentType":"text/html; charset=utf-8","ContentLength":8}}
`

##### 使用方法2(推荐)：

    `	//curl也可以通过链式构建,
     	cb := &CurlBuilder{}
     	curl :=cb.SetMethod("GET").SetUrl("https://www.baidu.com/").SetHeader("Content-Type","application/json").Build()
     	//执行请求
     	err := curl.Do()`