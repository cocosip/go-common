package utility

import "testing"

func TestIsHttps(t *testing.T) {
	url1 := "http://127.0.0.100/file"
	url2 := "https://baidu.com"

	b1, err := IsHttps(url1)
	if err != nil {
		t.Errorf("url: %s is wrong url,err:%s.\n", url1, err)
	}

	if b1 {
		t.Fatalf("IsHttps failed,url %s \n", url1)
	}

	b2, err := IsHttps(url2)
	if err != nil {
		t.Errorf("url: %s is wrong url,err:%s.\n", url2, err)
	}

	if !b2 {
		t.Fatalf("IsHttps failed,url %s \n", url2)
	}
}

func TestGetAuthority(t *testing.T) {
	url1 := "http://www.kingsoft.com/bucket1/studyuid/seriesuid/instance/100.dcm?id=300&name=4#50"
	authority := GetAuthority(url1)
	if authority != "http://www.kingsoft.com/bucket1/studyuid/seriesuid/instance/100.dcm" {
		t.Fatalf("GetAuthority failed,rawurl %s,authority:%s \n", url1, authority)
	}
}
