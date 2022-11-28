# PoC gRPC SD-PUCSP
# Pré requisitos
##### - GoLang 1.18+ : 
https://go.dev/dl//
# Instalação
##### - Baixar o projeto do Git:
```git clone git@github.com:gzaninotti/gRPC-PoC-SD.git```
#  Funções
##### - Broadcast(msg): Transmite uma mensagem (msg) via servidor.
# Utilização
##### - Primeiro é necessário subir o servidor:
- ```go run server/main.go```
##### Agora, para transmitir uma mensagem:
- ```go run client/main.go -msg "teste"```
- Onde está escrito "teste", pode ser substituido pela mensagem a qual será transmitida.
- Caso nenhuma mensagem for passada como parâmetro, o valor padrão adotado para a mensagem é "Hello world!"