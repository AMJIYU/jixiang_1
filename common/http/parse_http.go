package http

import (
	"strings"

	"github.com/projectdiscovery/urlutil"
)

//移除http协议头
func TrimProtocol(targetURL string, addDefaultPort bool) string {
	URL := strings.TrimSpace(targetURL)
	if strings.HasPrefix(strings.ToLower(URL), "http://") || strings.HasPrefix(strings.ToLower(URL), "https://") {
		if addDefaultPort {
			URL = AddURLDefaultPort(URL)
			URL = URL[strings.Index(URL, "//")+2:]
		}
	}

	return URL
}
// AddURLDefaultPort 给URI增加默认端口 (80/443)
// eg:
// http://foo.com -> http://foo.com:80
// https://foo.com -> https://foo.com:443
func AddURLDefaultPort(rawURL string) string {
	u, err := urlutil.Parse(rawURL)
	if err != nil {
		return rawURL
	}
	return u.String()
}

// RemoveURLDefaultPort 从URI中移除默认端口 (80/443)
// eg:
// http://foo.com:80 -> http://foo.com
// https://foo.com:443 -> https://foo.com
func RemoveURLDefaultPort(rawURL string) string {
	u, err := urlutil.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	if u.Scheme == urlutil.HTTP && u.Port == "80" || u.Scheme == urlutil.HTTPS && u.Port == "443" {
		u.Port = ""
	}
	return u.String()
}
