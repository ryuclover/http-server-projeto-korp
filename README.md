# Projeto Korp - Servidor Go com Docker, Prometheus, Grafana e Ansible

Fala galera! Subi aqui o projeto do desafio da Korp. Tá bem completinho: tem a aplicação Go, Nginx fazendo proxy reverso na porta 80, Prometheus/Grafana pra monitorar tudo e um playbook do Ansible pra subir o ambiente automático.

## Como rodar o projeto local (sem docker nem nada)

Se você tiver o Go instalado e quiser testar rapidão na máquina:

```bash
go run main.go
```

Aí a API vai escutar na porta `8080` (ex: `http://localhost:8080/projeto-korp`).
Dá pra gerar o binário também se quiser:
```bash
go build -o server
./server
```

---

## Como rodar com Docker Compose

Essa parte é legal porque a aplicação Go fica escondida (só exposta internamente na rede do docker) e o Nginx cuida de receber os acessos na porta `80` e mandar pra ela.

Pra subir os containers (App Go, Nginx, Prometheus e Grafana):

```bash
docker compose up -d --build
```

**Para testar se deu bom:**
```bash
# O endpoint da aplicação via Nginx
curl http://localhost/projeto-korp

# Prometheus (porta 9090)
curl http://localhost:9090/-/ready

# Grafana (porta 3000)
curl http://localhost:3000/api/health
```

E para limpar tudo:
```bash
docker compose down
```

---

## Como rodar a automação com Ansible

Fiz um playbook pra automatizar o provisionamento. Ele instala o docker, as dependências do python, joga as configs pro diretório `/opt/http-server-projeto-korp` e sobe o compose sozinho.

**Pré-requisitos:**
Ter o Ansible instalado (se estiver no Ubuntu/WSL, dá pra rodar `pipx install ansible` e instalar a collection do docker com `ansible-galaxy collection install -r ansible/requirements.yml`).

**Rodando o playbook:**
```bash
# Rodar na própria máquina (localhost)
ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
```

Se precisar rodar como sudo e pedir senha, adiciona a flag `-K` no final.

Qualquer dúvida manda um alô!
