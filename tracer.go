//跟踪记录信息，并输出至指定的 io.Writer

package trace

//pkg 的包名最好与文件夹名称相同。在引用时是输入文件夹名来引用，如果包名不相同容易造成混淆

import (
	"fmt"
	"io"
)

// Tracer 表示对象可以在代码中追踪事件
type Tracer interface {
	//...interface{} 意味着可以接受任意个数任意类型的参数
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// New 新建一个 tracer ，可以向 io.Writer 写入信息
// 尽量使用 Go 内建的接口类型，可以使代码可复用性高，更优雅
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

// Off 新建一个无操作的 tracer
func Off() Tracer {
	return &nilTracer{}
}
