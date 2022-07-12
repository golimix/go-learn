// 日志管理框架测试
//
package logrus_test

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestXx(t *testing.T) {
	log.WithFields(log.Fields{"animal": "walrus"}).Info("Info信息*****")
}
