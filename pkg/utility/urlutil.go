package utility

import (
	"net/url"
	"strings"
)

//Whether the url is https or not
func IsHttps(rawurl string) (bool,error) {
	uri,err:= url.Parse(rawurl)
	if err!=nil{
		return false, err
	}
	return  strings.ToLower(uri.Scheme) == "https", err
}

//Get url authority
func GetAuthority(rawurl string) string {
	uri, err := url.Parse(rawurl)
	if err != nil {
		return rawurl
	}

	var builder strings.Builder
	builder.WriteString(uri.Scheme)
	builder.WriteString("://")
	builder.WriteString(uri.Host)
	builder.WriteString(uri.Path)
	return builder.String()
}