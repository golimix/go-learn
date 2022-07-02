# 项目初始化
```bash
mkdir {leetcode,spec,tutorial}
cd leetcode
go mod init github.com/24king/go-learn/leetcode
go mod init github.com/24king/go-learn/tutorial
cd ../spec
go mod init github.com/24king/go-learn/spec
cd ../tutorial
go mod init github.com/24king/go-learn/tutorial
```
# 工作空间初始化
```
go work init ./tutorial
go work init ./leetcode
```