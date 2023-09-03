# GOSM - Go Site Monitor

O GOSM é projeto em desenvolvimento de um monitorador de sites feito usando GO (também conhecida como Golang).

Seu  objetivo é fornecer ao usuário um sistema de monitoramento de sites customizável.

## 🔧 Instalação

Caso deseje baixar o código e rodá-lo locamente para testes, é necessário instalar o Go. Durante o desenvolvimento, foi usada a versão mais recente na época (1.21.0).

Tendo o código fonte localmente, entre no diretório do aplicativo pelo terminal e digite:
```
go run main.go
```
Por padrão, o executável é compilado para o sistema operacional que está sendo usado, mas é possível compilar versões para outros sistemas operacionais como MacOS, Linux, entre outros. Para tirar dúvidas sobre o processo, [confira o tutorial da Digital Ocean](https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures-pt).

## 🔨 Funcionalidades do projeto
 - `Monitoramento`: Realizar o monitoramento
 - `Configurações do monitoramento`: Permite ersonalizar a repetição e o tempo entre cada verificação
 - `Sites`: Permite cadastrar/remover sites para o processo de monitoramento
 - `Logs`: Visualizar/limpar logs

## ✔️ Técnicas e tecnologias utilizadas

As técnicas e tecnologias utilizadas pra isso são:

- `GO 1.21.0`
- `Visual Studio Code`
- `Escrita/leitura/remoção de arquivo de texto`
- `Solicitar e ler response HTTP`
- `Organização em módulos`

## 📋 Implementações futuras

- Salvar/remover/ler os sites de um arquivo de texto local
- Transformar em uma web application, usando HTML, talvez Boostrap.

## 📁 Acesso ao projeto

Você pode [acessar o código fonte do projeto inicial](https://github.com/Kallrish/monitoramento) ou [baixá-lo](https://github.com/Kallrish/monitoramento/archive/refs/heads/main.zip).

# Autores

| [<img loading="lazy" src="https://avatars.githubusercontent.com/u/87496595?s=400&u=16a3d4a78219ed8520bda2a300dfccccfa02853e&v=4" width=115><br><sub>Jonatas Ribeiro</sub>](https://github.com/Kallrish) |
| :---: |