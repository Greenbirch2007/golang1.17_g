# golang1.17_g

2021.9.2 最新下载golang第三方库的方法


文档上面的命令都过时了！

用示例代码，在golang里面跑一下，比如安装echo 

go get github.com/labstack/echo

实际上去执行了 

go get -t github.com/labstack/echo命令

最终是下载到了gopath的src里面也就是项目的目录里面
而不是goroot，不是go的sdk的src里面

这次搞清楚！先下载诸多模块！玩起来

测试出的有效方法，实践出真知！

爱情也是一样吧，试错成本也要明确
