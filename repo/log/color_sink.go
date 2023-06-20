package log

import (
	"github.com/gookit/color"
	"go.uber.org/zap"
	"net/url"
)

type ColorSink struct{}

// Sync 实现Sync方法以实现Sink接口。因为只是控制台打印，不涉及缓存同步问题，所以直接return
func (c ColorSink) Sync() error { return nil }

// Close 实现Close方法以实现Sink接口。因为只是控制台打印，不涉及关闭对象问题，所以直接return
func (c ColorSink) Close() error { return nil }

// Write 实现Write方法以实现Sink接口。使用带颜色的控制台输出信息
func (c ColorSink) Write(p []byte) (n int, err error) {
	color.Red.Printf("%s", string(p))
	return len(p), nil
}

// 定义工厂函数
func colorSink(url *url.URL) (sink zap.Sink, err error) {
	// 工厂函数中，定义了必须接收一个 *url.URL 参数，但是我们需求比较简单，暂时用不到，可以直接忽略
	//实例化一个ColorSink对象，该对象实现了Sink接口
	c := ColorSink{}
	return c, nil
}
