// 日志管理框架测试
//
package logrus_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestXx(t *testing.T) {
	logrus.WithFields(logrus.Fields{"animal": "walrus"}).Info("Info信息*****")
}
