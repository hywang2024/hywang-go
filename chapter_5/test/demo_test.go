package test

import "testing"

//单元测试  // go test -v
func TestGetArea(t *testing.T) {
	area := GetArea(40, 50)
	if area != 2000 {
		t.Error("测试失败")
	}
}

//性能测试  // go test -bench="."
func BenchmarkGetArea(t *testing.B) {
	for i := 0; i < t.N; i++ {
		GetArea(40, 50)
	}
}
