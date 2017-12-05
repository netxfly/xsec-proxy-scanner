package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"

	"proxy_scanner/cmd"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "xsec proxy scanner"
	app.Author = "netxfly"
	app.Email = "x@xsec.io"
	app.Version = "20171205"
	app.Usage = "A SOCK4/SOCK4a/SOCK5/HTTP/HTTPS proxy scanner"
	app.Commands = []cli.Command{cmd.Scan, cmd.Dump}
	app.Flags = append(app.Flags, cmd.Scan.Flags...)
	app.Flags = append(app.Flags, cmd.Dump.Flags...)
	app.Run(os.Args)
}
