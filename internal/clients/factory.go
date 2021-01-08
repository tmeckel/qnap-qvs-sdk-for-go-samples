package clients

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
)

var cookies *cookiejar.Jar

func SenderFactory(renengotiation tls.RenegotiationSupport) (*http.Client, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// FIXME: make this thread-safe
	if nil == cookies {
		_cookies, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		cookies = _cookies
	}

	// transport.Proxy = func(*http.Request) (*url.URL, error) {
	// 	return url.Parse("http://10.0.0.152:8888")
	// }

	return &http.Client{Jar: cookies, Transport: transport}, nil
}
