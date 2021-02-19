package utility

import "testing"

func TestGetMd5(t *testing.T) {
	s1 := GetMd5([]byte("123123123"))
	if s1 != "f5bb0c8de146c67b44babbf4e6584cc0" {
		t.Fatalf("GetBufferMd5 failed,s1:%s\n", s1)
	}
}
