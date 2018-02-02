package main

import (
	"github.com/fabioberger/chrome"
	. "github.com/siongui/godom"
	"net/url"
	"strings"
)

func GetDomain(sUrl string) (domain string, err error) {
	u, err := url.Parse(sUrl)
	if err != nil {
		return
	}
	parts := strings.Split(u.Hostname(), ".")
	domain = parts[len(parts)-2] + "." + parts[len(parts)-1]
	return
}

func exportCookies(cookies []chrome.Cookie, domain string) {
	for _, cookie := range cookies {
		if strings.Contains(cookie.Domain, domain) {
			Document.Write(cookie.Domain)
			Document.Write("\t")

			if cookie.HostOnly {
				Document.Write("FALSE")
			} else {
				Document.Write("TRUE")
			}
			Document.Write("\t")

			Document.Write(cookie.Path)
			Document.Write("\t")

			if cookie.Secure {
				Document.Write("TRUE")
			} else {
				Document.Write("FALSE")
			}
			Document.Write("\t")

			// In the code of JS version:
			//
			//   document.write(cookie.expirationDate ? cookie.expirationDate : "0");
			//
			// but cookie.ExpirationDate is of type int64
			// so we use GopherJS Get directly
			if cookie.Get("expirationDate") == nil {
				Document.Write("0")
			} else {
				Document.Call("write", cookie.ExpirationDate)
			}
			Document.Write("\t")

			Document.Write(cookie.Name)
			Document.Write("\t")

			Document.Write(cookie.Value)
			Document.Write("\n")
		}
	}
}

func main() {
	c := chrome.NewChrome()
	queryInfo := chrome.Object{
		"active":        true,
		"currentWindow": true,
	}
	c.Tabs.Query(queryInfo, func(tabs []chrome.Tab) {
		tab := tabs[0]
		domain, _ := GetDomain(tab.Url)

		c.Cookies.GetAll(chrome.Object{}, func(cookies []chrome.Cookie) {
			Document.Write("<pre>\n")
			Document.Write("# Cookies for domains related to <b>" + domain + "</b>.\n")
			Document.Write("# This content may be pasted into a cookies.txt file and used by wget\n")
			Document.Write("# Example:  wget -x <b>--load-cookies cookies.txt</b> " + tab.Url + "\n")
			Document.Write("#\n")
			exportCookies(cookies, domain)
			Document.Write("</pre>\n")
		})

	})
}
