package utility

import "testing"

func TestGetSha1(t *testing.T) {
	s1:=GetSha1([]byte("123456"))
	if s1 != "7c4a8d09ca3762af61e59520943dc26494f8941b"{
		t.Fatalf("GetSha1 failed,s1:%s\n", s1)
	}
}

func TestGetSha256(t *testing.T) {
	s1:=GetSha256([]byte("123456"))
	if s1 != "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92"{
		t.Fatalf("GetSha256 failed,s1:%s\n", s1)
	}
}

func TestGetBufferSha512(t *testing.T) {
	s1:=GetSha512([]byte("111111"))
	if s1 != "b0412597dcea813655574dc54a5b74967cf85317f0332a2591be7953a016f8de56200eb37d5ba593b1e4aa27cea5ca27100f94dccd5b04bae5cadd4454dba67d"{
		t.Fatalf("TestGetBufferSha512 failed,s1:%s\n", s1)
	}
}

func TestGetHmacSha1(t *testing.T) {
	s1:=GetHmacSha1([] byte("123456"),[]byte("123456"))
	if s1 != "74b55b6ab2b8e438ac810435e369e3047b3951d0"{
		t.Fatalf("GetHmacSha1 failed,s1:%s\n", s1)
	}
}

func TestGetHmacSha256(t *testing.T) {
	s1:=GetHmacSha256([] byte("1234567890"),[]byte("999999999"))
	if s1 != "9fd05e347df4bd05dcac4f57897c424dbc496812796e7b5c172a21cf3087b6de"{
		t.Fatalf("GetHmacSha256 failed,s1:%s\n", s1)
	}
}
