# Aplicação de Linha de Comando para Busca de IPs e Nomes de Servidores

Esta é uma aplicação de linha de comando (CLI) escrita em Go que permite buscar IPs e nomes de servidores de um determinado host na internet.

### Pré-requisitos

- Go instalado (versão 1.16 ou superior recomendada).

### Instalação

1. Clone este repositório: go get github.com/urfave/cli
2. Navegue até o diretório do projeto;
3. Compile o projeto;

## Como usar
A aplicação possui dois comandos principais: ip e server.

### Buscar IPs
Para buscar os IPs associados a um host, use o comando ip:
      ./busca-ips-servers ip --host google.com.br
      
Isso retornará uma lista de IPs associados ao host google.com.br.

### Buscar Nomes de Servidores
Para buscar os nomes dos servidores associados a um host, use o comando server:
      ./busca-ips-servers server --host google.com.br
      
Isso retornará uma lista de nomes de servidores associados ao host google.com.br.

### Flags
--host: Especifica o host para o qual deseja buscar informações. O valor padrão é google.com.br.
