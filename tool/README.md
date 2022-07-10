## Go语言自动注释工具
mkdir -p $GOPATH/src/golang.org/x  //路径下创建此文件
cd $GOPATH/src/golang.org/x      //切换到此目录
git clone https://github.com/golang/tools.git  //通过git安装 tools
git clone https://github.com/golang/lint.git   //安装 lint 
go get golang.org/x/lint/golint       //然后执行