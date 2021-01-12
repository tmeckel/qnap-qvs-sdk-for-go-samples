package clients

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
)

type SenderFactory interface {
	GetSingelton() (*http.Client, error)
	CreateInstance(*cookiejar.Jar) (*http.Client, error)
}

type defaultSenderFactory struct {
	cookies  *cookiejar.Jar
	instance *http.Client
}

var DefaultSenderFactory = defaultSenderFactory{}

func (s *defaultSenderFactory) GetSingelton() (*http.Client, error) {
	// FIXME: make this thread-safe
	err := s.initializeCookies()
	if err != nil {
		return nil, err
	}
	if s.instance == nil {
		s.instance, err = s.CreateInstance(s.cookies)
		if err != nil {
			return nil, err
		}
	}
	return s.instance, nil
}

func (s *defaultSenderFactory) CreateInstance(cookies *cookiejar.Jar) (*http.Client, error) {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// transport.Proxy = func(*http.Request) (*url.URL, error) {
	// 	return url.Parse("http://10.0.0.152:8888")
	// }
	return &http.Client{Jar: cookies, Transport: transport}, nil
}

func (s *defaultSenderFactory) initializeCookies() error {
	if s.cookies == nil {
		_cookies, err := cookiejar.New(nil)
		if err != nil {
			return err
		}
		s.cookies = _cookies
	}
	return nil
}
