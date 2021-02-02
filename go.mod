module github.com/2336260845/hippo_recall

go 1.15

require (
	github.com/2336260845/hippo_search v0.0.0-20210202073928-1e422d77f435
	github.com/apache/thrift/lib/go/thrift v0.0.0-20210120171102-e27e82c46ba4
	github.com/elastic/go-elasticsearch v0.0.0
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
)

replace github.com/apache/thrift/lib/go/thrift => ./go_thrift //gomod总是自动拉取最新代码，暂时未找到更好的办法，先替换为本地文件;本地版本为v0.13.0
