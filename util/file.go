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

package util

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type ProxyAddr struct {
	IP   string
	Port int
}

func ReadProxyAddr(fileName string) (sliceProxyAddr []ProxyAddr) {
	proxyFile, err := os.Open(fileName)
	if err != nil {
		Log.Fatalf("Open proxy file err, %v", err)
	}

	defer proxyFile.Close()

	scanner := bufio.NewScanner(proxyFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ipPort := scanner.Text()
		t := strings.Split(ipPort, ":")
		ip := t[0]
		port, err := strconv.Atoi(t[1])
		if err == nil {
			proxyAddr := ProxyAddr{IP: ip, Port: port}
			sliceProxyAddr = append(sliceProxyAddr, proxyAddr)
		}
	}

	return sliceProxyAddr
}
