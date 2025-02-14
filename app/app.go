package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "server",
			Usage:  "Busca o nome dos Servers de endereços na internet",
			Flags:  flags,
			Action: buscarServer,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	//net
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServer(c *cli.Context) {
	host := c.String("host")

	//net
	nss, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, n := range nss {
		fmt.Println(n)
	}
}
