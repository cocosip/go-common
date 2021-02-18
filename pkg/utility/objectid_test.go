package utility

import "testing"

func TestNewObjectId(t *testing.T) {
	s1:=NewObjectId().String()
	s2:=NewObjectId().String()
	if s1 == s2{
		t.Fatalf("NewObjectId failed,s1:%s,s2:%s\n", s1,s2)
	}
	t.Logf("s1:%s,s2:%s\n", s1, s2)
}