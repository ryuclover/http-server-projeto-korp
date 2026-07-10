# http-server-projeto-korp

Este é um serviço HTTP simples em Golang.

## Etapa 1: Rodar o servidor localmente

Certifique-se de ter o Go instalado em sua máquina.

```bash
go run main.go
```
Ou se preferir com Go Modules:
```bash
go build -o server
./server
```

## Etapa 2: Usar Docker (standalone)

Caso queira buildar e testar somente o container da aplicação na porta 8080:

```bash
docker build -t http-server-projeto-korp .
docker run -d -p 8080:8080 --name korp-server http-server-projeto-korp
```
*Lembre-se de deletar esse container (`docker rm -f korp-server`) antes de ir para a etapa do Docker Compose para evitar conflitos!*

## Etapa 3: Usar Docker Compose com NGINX

Nesta etapa, temos o `docker-compose.yml` que provisiona a aplicação Go em uma rede interna, sem expor sua porta para a máquina host, e o NGINX atuando como proxy reverso na porta 80.

### Como subir o ambiente:

1. Certifique-se de que a porta `80` (e a `8080` de testes antigos) estão livres no seu sistema. Pare os containers avulsos com `docker rm -f korp-server`.
2. Na raiz do projeto, execute o Docker Compose:

```bash
docker compose up -d
```
*O `--build` é opcional, mas se alterar o código e quiser recompilar, rode `docker compose up -d --build`.*

### Como testar o proxy reverso:

Agora você acessará o NGINX na porta padrão HTTP (80):

```bash
curl http://localhost/projeto-korp
# O ":80" é opcional, mas pode testar assim também:
# curl http://localhost:80/projeto-korp
```

### Resposta esperada

```json
{
  "nome": "Projeto Korp",
  "horario": "2026-07-10T03:07:35Z"
}
```

### Como derrubar o ambiente:

Para parar e remover os containers, bem como a rede criada:

```bash
docker compose down
```
