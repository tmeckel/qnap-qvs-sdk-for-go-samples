package clients

import (
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/Azure/go-autorest/autorest"
)

func setHeader(r *http.Request, key, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)
}

func withCSRFTokenHeader(cookies *cookiejar.Jar) autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, err := p.Prepare(r)
			if err == nil && cookies != nil {
				for _, cookie := range cookies.Cookies(r.URL) {
					if strings.EqualFold("csrftoken", cookie.Name) {
						setHeader(r, http.CanonicalHeaderKey("X-CSRFToken"), cookie.Value)
						break
					}
				}
			}
			return r, err
		})
	}
}

func requestDecorator(baseURI string) autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		decorators := []autorest.PrepareDecorator{
			autorest.WithHeader("referer", baseURI+"/"),
			withCSRFTokenHeader(DefaultSenderFactory.cookies),
		}
		p = autorest.DecoratePreparer(p, decorators...)
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			return p.Prepare(r)
		})
	}
}

func configureClient(client *autorest.Client, baseURI string) error {
	var err error

	client.Sender, err = DefaultSenderFactory.GetSingelton()
	if err != nil {
		return err
	}
	client.RequestInspector = requestDecorator(baseURI)
	client.SendDecorators = []autorest.SendDecorator{
		autorest.DoErrorIfStatusCode(http.StatusInternalServerError),
		autorest.DoCloseIfError(),
	}

	return nil
}
