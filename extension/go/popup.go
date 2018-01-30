package main

import (
	"github.com/fabioberger/chrome"
)

func main() {
	c := chrome.NewChrome()
	c.BrowserAction.OnClicked(func(tab chrome.Tab) {
		o := chrome.Object{
			"file": "cc.js",
		}
		c.Tabs.ExecuteScript(tab.Id, o, nil)
	})
}
