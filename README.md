# SD
# Pré requisitos
##### - GoLang 1.18
# Instalação
##### - git clone git@github.com:gzaninotti/gRPC-PoC-SD.git
#  Funções
##### - Broadcast(): Transmite uma mensagem via servidor.
# Utilização
##### - Primeiro é necessário subir o servidor:
- go run server/main.go
##### Agora, para transmitir uma mensagem:
- go run client/main.go -msg "teste"
- Onde está escrito "teste", pode ser substituido pela mensagem a qual será transmitida.
- Caso nenhuma mensagem for passada como parâmetro, o valor padrão adotado para a mensagem é "Hello world!"

