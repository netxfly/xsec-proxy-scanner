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
	"github.com/urfave/cli"
	"github.com/sirupsen/logrus"

	"proxy_scanner/util"
	"proxy_scanner/models"

	"sync"
	"time"
)

var (
	DEBUG_MODE = false
	SCAN_NUM   = 1000
	PROXY_FILE = "iplist.txt"
	TIMEOUT    = 5
)

type HttpProxyFunc func(ip string, port int, protocol string) (err error, isProxy bool, proxyInfo models.ProxyInfo)

type SockProxyFunc func(ip string, port int, protocol int) (err error, isProxy bool, proxyInfo models.ProxyInfo)

var (
	httpProxyFunc HttpProxyFunc = CheckHttpProxy
	sockProxyFunc SockProxyFunc = CheckSockProxy
)

func Scan(ctx *cli.Context) (err error) {
	if ctx.IsSet("debug") {
		DEBUG_MODE = ctx.Bool("debug")
	}

	if DEBUG_MODE {
		util.Log.Logger.Level = logrus.DebugLevel
	}

	if ctx.IsSet("timeout") {
		TIMEOUT = ctx.Int("timeout")
	}

	if ctx.IsSet("scan_num") {
		SCAN_NUM = ctx.Int("scan_num")
	}

	if ctx.IsSet("filename") {
		PROXY_FILE = ctx.String("filename")
	}

	startTime := time.Now()

	proxyAddrList := util.ReadProxyAddr(PROXY_FILE)
	proxyNum := len(proxyAddrList)
	util.Log.Infof("%v proxies will be check", proxyNum)

	scanBatch := proxyNum / SCAN_NUM
	for i := 0; i < scanBatch; i++ {
		util.Log.Debugf("Scanning %v batches", i+1)
		proxies := proxyAddrList[i*SCAN_NUM:(i+1)*SCAN_NUM]
		CheckProxy(proxies)
	}

	util.Log.Debugf("Scanning The last batches(%v)", scanBatch+1)
	if proxyNum%SCAN_NUM > 0 {
		proxies := proxyAddrList[SCAN_NUM*scanBatch:proxyNum]
		CheckProxy(proxies)
	}

	count, _ := models.CacheStatus()
	models.SaveProxiesToFile()
	models.DumpToFile(models.DUMP_FILENAME)
	util.Log.Infof("Scan proxies Done, Found %v proxies, used time: %v", count, time.Since(startTime))

	return err
}

func CheckProxy(proxyAddr []util.ProxyAddr) {
	var wg sync.WaitGroup
	wg.Add(len(proxyAddr) * (len(HttpProxyProtocol) + len(SockProxyProtocol)))

	for _, addr := range proxyAddr {
		for _, proto := range HttpProxyProtocol {
			go func(ip string, port int, protocol string) {
				defer wg.Done()
				models.SaveProxies(httpProxyFunc(ip, port, protocol))
			}(addr.IP, addr.Port, proto)
		}

		for proto := range SockProxyProtocol {
			go func(ip string, port int, protocol int) {
				defer wg.Done()
				models.SaveProxies(sockProxyFunc(ip, port, protocol))
			}(addr.IP, addr.Port, proto)
		}
	}
	wg.Wait()
}
