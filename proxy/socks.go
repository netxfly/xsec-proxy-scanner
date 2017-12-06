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
	"h12.me/socks"

	"proxy_scanner/models"
	"proxy_scanner/util"

	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"strings"
)

var (
	//SockProxyProtocol = []int{socks.SOCKS4, socks.SOCKS4A, socks.SOCKS5}
	SockProxyProtocol = map[int]string{socks.SOCKS4: "SOCKS4", socks.SOCKS4A: "SOCKS4A", socks.SOCKS5: "SOCKS5"}
)

func CheckSockProxy(ip string, port, protocol int) (err error, isProxy bool, proxyInfo models.ProxyInfo) {
	proxyInfo.Addr = ip
	proxyInfo.Port = port
	proxyInfo.Protocol = SockProxyProtocol[protocol]

	proxy := fmt.Sprintf("%v:%v", ip, port)
	dialSocksProxy := socks.DialSocksProxy(protocol, proxy)
	tr := &http.Transport{Dial: dialSocksProxy}
	httpClient := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(TIMEOUT) * time.Second}
	util.Log.Debugf("Checking proxy: %v", fmt.Sprintf("%v://%v:%v", SockProxyProtocol[protocol], ip, port))
	resp, err := httpClient.Get(WebUrl)
	if err == nil {
		if resp.StatusCode == http.StatusOK {
			body, err := ioutil.ReadAll(resp.Body)
			if err == nil && strings.Contains(string(body), "网易免费邮箱") {
				isProxy = true
			}
		}
	}
	return err, isProxy, proxyInfo
}
