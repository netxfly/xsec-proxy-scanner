/*

Copyright (c) 2017 xsec.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THEq
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

*/

package proxy

import (
	"proxy_scanner/models"
	"proxy_scanner/util"

	"fmt"
	"time"

	"net/http"
	"net/url"
	"io/ioutil"
	"strings"
)

var (
	HttpProxyProtocol = []string{"http", "https"}
	WebUrl            = "http://email.163.com/"
)

func CheckHttpProxy(ip string, port int, protocol string) (err error, isProxy bool, proxyInfo models.ProxyInfo) {
	proxyInfo.Addr = ip
	proxyInfo.Port = port
	proxyInfo.Protocol = protocol

	rawProxyUrl := fmt.Sprintf("%v://%v:%v", protocol, ip, port)
	proxyUrl, err := url.Parse(rawProxyUrl)
	if err == nil {
		Transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		client := &http.Client{Transport: Transport, Timeout: time.Duration(TIMEOUT) * time.Second}
		util.Log.Debugf("Checking proxy: %v", rawProxyUrl)
		resp, err := client.Get(WebUrl)
		if err == nil {
			if resp.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				if err == nil && strings.Contains(string(body), "网易免费邮箱") {
					isProxy = true
				}

			}
		}
	}
	return err, isProxy, proxyInfo
}
