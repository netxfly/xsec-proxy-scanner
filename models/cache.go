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

package models

import (
	"github.com/patrickmn/go-cache"
	"github.com/urfave/cli"

	"proxy_scanner/util"

	"encoding/gob"
	"fmt"
	"os"
)

func init() {
	gob.Register(ProxyInfo{})
}

func SaveProxies(err error, isProxy bool, proxyInfo ProxyInfo) () {
	if err == nil && isProxy {
		k := fmt.Sprintf("%v://%v:%v", proxyInfo.Protocol, proxyInfo.Addr, proxyInfo.Port)
		//util.Log.Debug(k)
		CACHE_PROXIES.Set(k, true, cache.NoExpiration)
	}
}

func SaveProxiesToFile() (error) {
	return CACHE_PROXIES.SaveFile("xsec_proxies.db")
}

func CacheStatus() (count int, items map[string]cache.Item) {
	count = CACHE_PROXIES.ItemCount()
	items = CACHE_PROXIES.Items()
	return count, items
}

func ProxiesNum() () {
	util.Log.Infof("Total proxies: %v", CACHE_PROXIES.ItemCount())
}

func LoadProxiesFromFile() {
	CACHE_PROXIES.LoadFile("xsec_proxies.db")
	ProxiesNum()
}

func Dump(ctx *cli.Context) (err error) {
	LoadProxiesFromFile()

	if ctx.IsSet("file") {
		DUMP_FILENAME = ctx.String("file")
	}
	err = DumpToFile(DUMP_FILENAME)
	if err != nil {
		util.Log.Fatalf("Dump proxies to file err, Err: %v", err)
	}
	return err
}

func DumpToFile(filename string) (err error) {
	file, err := os.Create(filename)
	if err == nil {
		_, items := CacheStatus()
		for k := range items {
			file.WriteString(fmt.Sprintf("%v\n", k))
		}
	}
	return err
}
