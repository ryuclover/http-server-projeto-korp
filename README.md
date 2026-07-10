# HTTP Server - Projeto Korp 🚀

Este projeto é uma demonstração de infraestrutura DevOps moderna, cobrindo desde o desenvolvimento de uma API simples em **Go** até o empacotamento com **Docker/Docker Compose**, observabilidade com **Prometheus/Grafana** e automação de provisionamento com **Ansible**.

---

## 🛠️ Tecnologias e Ferramentas
* **Linguagem:** Go (servidor HTTP nativo e expose de métricas Prometheus)
* **Containers & Redes:** Docker & Docker Compose
* **Proxy Reverso:** NGINX (atuando como gateway na porta 80 e encaminhando o tráfego interno)
* **Observabilidade:** Prometheus (coleta de métricas da API) e Grafana (visualização em dashboard provisionado automaticamente)
* **Automação/IaC:** Ansible (Playbook pronto para instalar dependências, configurar o host e subir a stack)

---

## 📐 Estrutura e Arquitetura

O tráfego é centralizado no NGINX para evitar exposição direta das portas dos serviços da stack ao host.

```text
[Cliente] 
   │
   ▼ (Porta 80)
┌────────────────────────┐
│         NGINX          │
└──────────┬─────────────┘
           │
           ├──────────────► [ /projeto-korp ] ──► App Go (Porta 8080)
           │                                            │
           │                                            ▼ (Métricas /metrics)
           │                                      ┌───────────┐
           │                                      │Prometheus │ (Porta 9090)
           │                                      └─────▲─────┘
           │                                            │ (Scrape)
           └──────────────► [ Grafana Dashboard ] ──────┘ (Porta 3000)
```

---

## 🚀 Como Executar o Projeto

### Etapa 1: Rodar o servidor Go localmente
Se quiser testar a aplicação pura na sua máquina:
```bash
go run main.go
```
Ou compile o binário:
```bash
go build -o server
./server
```
A API ficará disponível em `http://localhost:8080/projeto-korp`.

### Etapa 2: Executar com Docker Compose (NGINX + App + Observabilidade)
Para subir toda a estrutura (Nginx, App, Prometheus e Grafana):
```bash
docker compose up -d --build
```

**Validar se o ambiente subiu com sucesso:**
* **API Principal (via Proxy Nginx na porta 80):**
  ```bash
  curl http://localhost/projeto-korp
  ```
  *Resposta esperada:* `{"nome":"Projeto Korp","horario":"2026-07-10T18:32:14Z"}`
* **Prometheus:** `http://localhost:9090` (Endpoint `/api/v1/targets` deve mostrar o target `app:8080` como **UP**)
* **Grafana:** `http://localhost:3000` (Login: `admin` / Senha: `admin`). O Dashboard `HTTP Server Projeto Korp` já virá pronto e configurado exibindo os gráficos de requisições.

Derrubar o ambiente:
```bash
docker compose down
```

---

## 🤖 Provisionamento Automatizado com Ansible

O playbook configura todo o servidor de destino automaticamente. Ele instala o Docker, configura as regras de grupos, cria a pasta do projeto em `/opt/http-server-projeto-korp`, transfere os arquivos de configuração e inicializa os containers via Docker Compose.

**Como rodar a automação:**
1. Instale a dependência de Ansible necessária:
   ```bash
   ansible-galaxy collection install -r ansible/requirements.yml
   ```
2. Execute o playbook localmente:
   ```bash
   ansible-playbook -i ansible/inventory.ini ansible/playbook.yml
   ```

---

## 🌟 Destaques do Projeto (Diferenciais DevOps)
* **Provisionamento As-Code do Grafana:** O datasource do Prometheus e o Dashboard de monitoramento da aplicação são injetados automaticamente no container através de arquivos de configuração em `grafana/provisioning`, sem necessidade de configuração manual via interface.
* **Resiliência do NGINX:** Configuração utilizando DNS dinâmico e resolver interno do Docker (`127.0.0.11`) para evitar falhas/crash do proxy caso o container da aplicação Go demore alguns segundos extras para iniciar.
* **Segurança de Portas:** Apenas as portas necessárias para acesso externo (`80` para o proxy, `9090` para o Prometheus e `3000` para o Grafana) estão expostas para o host. A aplicação Go fica protegida dentro da rede virtual isolada do Docker.
