# Aplicação para Pagamentos
Esta aplicação tem como objetivo automatizar pagamentos de um determinado serviço setando automaticamente o benefício proposto direto no banco de dados.
## Recursos
* ### Configurar de forma fácil seus próprios produtos
É possível configurar de forma fácil os seus próprios produtos que deseja vender na aplicação com apenas uma function no arquivo _main.js_.
* ### Mercado Pago
A aplicação tem suporte 100% com a plataforma do Mercado Pago.
* ### API
O BackEnd do projeto é uma API com apenas 2 endpoints para Login, identificar quem é o beneficiário e a compra.
* ### Robo
O Robo faz a verificação com minutos a definir de todos os pagamentos em abertos presente no banco de dados.
## Tecnologias utilizadas
Foi utilizado neste projeto Golang em todo o BackEnd e as lingagens padrões no FrontEnd, HTML, CSS e JS com Framework Bootstrap 5.

## Configuração (_config.json_)
Existem dois arquivos de configuração no projeto. Um localizado na pasta raíz do diretório "Robo", e outro na pasta raíz do diretório "Backend".
* ### Backend
```json
{
  "Host": "localhost",
  "Port": "3306",
  "DBname": "BANCO_DE_DADOS",
  "User": "root",
  "Pass": "",
  "PortaRodando": "5000",
  "TabelaPersonagens": "TABELA_PERSONAGENS",
  "TokenMercadoPago": "TOKEN_MERCADO_PAGO"
}
```
* ### Robo

```json
{
  "Host": "localhost",
  "Port": "3306",
  "DBname": "BANCO_DE_DADOS",
  "User": "root",
  "Pass": "",
  "TabelaPersonagens": "TABELA_PERSONAGENS",
  "TokenMercadoPago": "TOKEN_MERCADO_PAGO",
  "MinutosRodando": 3, // INTERVALO DE TEMPO QUE O ROBO IRÁ VARRER O BANCO (MINUTOS)
  "Coins": 1 // VALOR DE CADA COIN
}
```
* ### FrontEnd
```js
geraProduto(idProduto, NOME_PRODUTO, PREÇO_PRODUTO, []BENEFÍCIOS_PRODUTO);
```
Por Exemplo para gerar um Vip Bronze seria: 
```js
geraProduto('vipBronze', 'Vip Bronze', '10,00', ['30 dias de vip', 'Tag no Discord de Vip', 'Tag Vip no servidor', '10 coins', '+1 veículo no inventário'])
```
## Endpoints API

* ### {link}/usuario - POST
Este endpoint é um POST onde deverá ser passado no body um json com nome e nome do personagem como string.
EX:
```json
{
  "nome": "Ellivelton_Narciso"
}
```
e o retorno será o ID deste usuário no banco de dados juntamente com seu token valido por 3 horas, uma compra só é feita e validada caso este token existir e for válido.

* ### {link}/s/compraFeita - POST
Este endpoint é um POST onde deverá ser passado um Baerer Token no header com o token recebido ao logar.
No body deverá ser passado um json com o id recebido no login que corresponde ao id do usuário no banco juntamente com o idProduto comprado.
Ex:

```json
{
  "id": 2,
  "token": SEU_TOKEN_RECEBIDO_AQUI
}
```
## Rodando a aplicação
O FrontEnd pode somente ser hospedado em um servidor Apache na versão mais recente.

Já o Backend depois de configurar o arquivo _config.json_ o mesmo pode ser buildado ou roda-lo como arquivo _.go_ como preferir.
Tanto para build como para roda-lo como arquivo .go é necessário a versão mais recente do Golang disponível para download no <a href="https://go.dev/" target="_blank">site</a> do mesmo.

O Robo funciona da mesma forma que o BackEnd, apos feito o arquivo _config.json_ poderá tanto ser buildado como roda-lo como _.go_.
