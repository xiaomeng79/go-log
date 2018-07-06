# go-log
#### go封装的日志 logrus zap,并且增加了分布式日志追踪，日志格式化成json,日志大小切割

#### 获取
    `go get -u github.com/xiaomeng79/go-log`
   
    
#### 快速开始，初始化zap,记录日志(推荐)


    `
    //全局初始化一次即可
    var Log = zaplog.New()
    //使用context包来分布式跟踪日志
    Log.Info(context.Background(),"test")
    `
#### 快速开始，初始化logrous,记录日志(备选)


    `
    //全局初始化一次即可
    var Log = zaplog.New()
    //使用context包来分布式跟踪日志
    Log.Info(context.Background(),"test")
    `
  
> 注意： 每个日志组件下的option文件都有配置项,每个日志都有默认配置

#### 配置日志(参考example下的文件)

    `
    var Log = zaplog.New(
    	zaplog.WithLogPath("tmp/log/"),
    	zaplog.WithLogName("test"),
    	zaplog.WithMaxAge(7),
    	zaplog.WithMaxSize(100),
    	zaplog.WithIsStdOut(true),
    )
    `
#### 日志使用(见example下示例)

    `
    //普通info日志
	ZapLog.Info(context.Background(),"test")
	//错误日志，打印错误栈信息
	ZapLog.Error(context.Background(),"inside error")
	//模拟新建一个错误日志类型,打印错误的详细信息
	err := errors.New("this is a test error")
	ZapLog.WarnO(context.Background(),err,"inside error")
	//模拟一个http请求,打印请求的一些信息,包括请求头，请求体，延迟，请求类型，方法，响应等
	cb := &curl.CurlBuilder{}
	c :=cb.SetMethod("GET").SetUrl("https://www.baidu.com/").SetHeader("Content-Type","application/json").Build()
	//执行请求
	c.Do()
	ZapLog.InfoO(context.Background(),c,"请求百度")
`



    
