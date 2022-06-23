package urlbuilder

import (
	"fmt"
	"net/url"
)

type UrlBuilder struct {
	url url.URL
}

func (u *UrlBuilder) SetPath(path string) *UrlBuilder {
	u.url.Path = path
	return u
}

func (u *UrlBuilder) SetHost(host string) *UrlBuilder {
	u.url.Host = host
	return u
}

func (u *UrlBuilder) SetScheme(scheme string) *UrlBuilder {
	u.url.Scheme = scheme
	return u
}

func (u *UrlBuilder) SetPort(port int) *UrlBuilder {
	u.url.Host = fmt.Sprintf("%s:%d", u.url.Hostname(), port)
	return u
}

func (u *UrlBuilder) SetHostName(hostName string) *UrlBuilder {
	port := u.url.Port()
	if port != "" {
		u.url.Host = hostName
	} else {
		u.url.Host = fmt.Sprintf("%s:%s", hostName, port)
	}
	return u
}

func (u *UrlBuilder) AddQuery(field string, value string) *UrlBuilder {
	query := u.url.Query()
	query.Add(field, value)
	u.url.RawQuery = query.Encode()
	return u
}

func (u *UrlBuilder) GetUrl() *url.URL {
	return &u.url
}

func (u *UrlBuilder) Build() string {
	return u.url.String()
}
