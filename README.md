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

2021.9.9


不要用goland来编译，只用来编辑，直接用命令行来编译


项目框架就按照goland来处理

在项目目录下，
使用go mod init 来初始化，创建go.mod文件

然后再用go mod  tidy来拉取需要模块

go build main.go
./main

即可

Go mod生成的最主要的文件 go.mod
每个包后面都跟了一个版本。如果想切换分支的话，后面的版本可以任意切换到需要的分支上，比如


也可以使用本地代码替换远程代码分支。就可以使用下面的
/data/www/go/src/go.uuid 代替远程分支 github.com/satori/go.uuid。
在go.mod最后一行加上下面的代码

replace github.com/satori/go.uuid => /data/www/go/src/go.uuid 
Go mod的使用是不是特别简单。

就是这么简单！



