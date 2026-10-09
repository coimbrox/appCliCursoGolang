package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Retornar a app de linha de comando pronta pra ser executada

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App Cli"
	app.Usage = "buscar ip e nome de servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "coimbrox.github.io",
		}}

	app.Commands = []cli.Command{

		{
			Name:   "ip",
			Usage:  "Buscar IP de endereços",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Buscar Servidores de la servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidores := range servidores {
		fmt.Println(servidores.Host)
	}
}
