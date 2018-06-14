package trace

import (
	"bytes"
	"testing"
)

//red-green testing
//先写一个单元测试，此时编译出错。然后再根据出错内容去写(尽可能少的) product code，让单元测试通过
//这种 TDD(test driven development) 方法可以保证所写的测试代码是有意义的
//尽可能少的代码是为了保证测试代码的 coverage
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")
		if buf.String() != "Hello trace package.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	tracer := Off()
	if tracer == nil {
		t.Error("Return from New should not be nil")
	}
	tracer.Trace("something")
}
