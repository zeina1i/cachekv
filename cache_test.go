package cachekv

import "testing"

func Test_HappyPath(t *testing.T) {
	c := NewCache()
	c.Set("foo", []byte("bar"))
	foo, ok := c.Get("foo")
	if ok != true {
		t.Error("Expected ok to be true, got false")
	}
	if string(foo) != "bar" {
		t.Errorf("Expected %s, got %s", "bar", foo)
	}
	c.Del("foo")
	_, ok = c.Get("foo")
	if ok != false {
		t.Error("Expected ok to be false, got true")
	}
}
