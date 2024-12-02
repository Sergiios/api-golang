# **API Golang - Gerenciamento de Centrais**

Este é um projeto de API desenvolvido em Go utilizando as seguintes tecnologias:

- **Fiber**: Um framework web rápido e flexível.
- **GORM**: ORM para interação com o banco de dados.
- **SQLite**: Banco de dados utilizado para persistência.
- **Swagger**: Para documentação interativa da API.
- **Testify**: Framework para testes unitários.

---

## **Sumário**
- [Funcionalidades](#funcionalidades)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Instalação](#instalação)
- [Execução do Projeto](#execução-do-projeto)
- [Documentação com Swagger](#documentação-com-swagger)
- [Testes Unitários](#testes-unitários)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Contribuindo](#contribuindo)
- [Melhorias Futuras](#melhorias-futuras)

---

## **Funcionalidades**

A API oferece um CRUD para gerenciamento de "Centrais", com as seguintes operações:

- **Criar Central**: Adiciona uma nova central no sistema.
- **Listar Centrais**: Retorna todas as centrais cadastradas.
- **Buscar Central por ID**: Retorna uma central específica pelo ID.
- **Atualizar Central**: Atualiza os dados de uma central existente.
- **Deletar Central**: Remove uma central do sistema.

---

## **Tecnologias Utilizadas**

- [Golang](https://golang.org/)
- [Fiber](https://gofiber.io/)
- [GORM](https://gorm.io/)
- [SQLite](https://sqlite.org/)
- [Swagger](https://swagger.io/)
- [Testify](https://github.com/stretchr/testify)

---

## **Instalação**

1. Clone o repositório:
   ```bash
   git clone https://github.com/Sergiios/api-golang.git
   cd api-golang
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure o banco de dados:
   - O projeto utiliza o SQLite por padrão (`database.db`).

---

## **Execução do Projeto**

Para rodar o projeto localmente, utilize o comando:
```bash
go run main.go
```

O servidor estará disponível em:
```
http://localhost:3000
```

---

## **Documentação com Swagger**

A API utiliza o Swagger para fornecer uma documentação interativa. Acesse a documentação em:
```
http://localhost:3000/swagger/index.html
```

### **Gerando a Documentação Swagger**

Certifique-se de que o Swagger está configurado corretamente no seu projeto. Para gerar ou atualizar a documentação, execute:
```bash
swag init
```

Caso ainda não tenha o Swagger instalado:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

---

## **Testes Unitários**

O projeto possui testes unitários cobrindo os seguintes componentes:

1. **Handler**:
   - Testa as rotas e validações da API.
2. **UseCase**:
   - Valida a lógica de negócio e interações com o repositório.
3. **Repositório**:
   - Simula operações de banco de dados usando SQLite em memória.
4. **Utils**:
   - Testa funções auxiliares, como formatação de erros de validação.
5. **Config**:
   - Testa a inicialização do banco de dados.

### **Rodando os Testes**

Para executar todos os testes, utilize o comando:
```bash
go test ./... -v
```

### **Cobertura de Testes**

Para visualizar a cobertura de testes:
```bash
go test ./... -cover
```

Ou gere um relatório em HTML:
```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## **Estrutura do Projeto**

```
.
├── cmd/
│   ├── main.go          # Arquivo principal
├── internal/
│   ├── config/          # Configuração do banco de dados
│   ├── domain/          # Definição das entidades
│   ├── handler/         # Rotas e controladores
│   ├── repository/      # Interação com o banco de dados
│   ├── usecase/         # Regras de negócio
│   ├── utils/           # Funções auxiliares         
├── go.mod               # Dependências do projeto
└── swagger/             # Arquivos de documentação gerados pelo Swagger
```

---

## **Contribuindo**

Contribuições são bem-vindas! Siga os passos abaixo para colaborar:

1. Faça um fork do repositório.
2. Crie uma branch com sua feature:
   ```bash
   git checkout -b minha-feature
   ```
3. Commit suas alterações:
   ```bash
   git commit -m "feat: descrição da feature"
   ```
4. Envie sua branch:
   ```bash
   git push origin minha-feature
   ```
5. Abra um Pull Request.

---

---

## **Melhorias Futuras**

Aqui estão algumas ideias e melhorias que podem ser implementadas no futuro para aprimorar o projeto:

### **Funcionalidades**
- **Autenticação e Autorização**:
  - Implementar autenticação de usuários com tokens JWT.
  - Adicionar diferentes níveis de autorização para proteger endpoints (ex.: admin, user).
  
- **Busca Avançada**:
  - Implementar filtros, ordenação e paginação para o endpoint de listagem de centrais.
  - Adicionar suporte a buscas por nome, MAC ou IP.

- **Notificações**:
  - Integrar um sistema de notificações para alertar eventos importantes, como erros ou atualizações.

- **Logs e Monitoramento**:
  - Adicionar logs estruturados para depuração e análise.
  - Integrar ferramentas de monitoramento, como Prometheus e Grafana.

### **Melhorias Técnicas**
- **Banco de Dados**:
  - Migrar para um banco de dados relacional mais robusto em produção, como PostgreSQL.
  - Configurar migrações de banco de dados usando uma biblioteca como [golang-migrate](https://github.com/golang-migrate/migrate).

- **Testes**:
  - Expandir os testes unitários para cobrir cenários adicionais.
  - Adicionar testes de integração para validar fluxos completos entre diferentes camadas.
  - Configurar testes automatizados em um pipeline de CI/CD.

- **Documentação**:
  - Melhorar a documentação com exemplos detalhados de uso da API.
  
- **Escalabilidade**:
  - Adicionar suporte a caching (ex.: Redis) para melhorar o desempenho em endpoints de leitura.
  - Suporte para executar em ambientes com múltiplas réplicas (horizontal scaling).

### **DevOps**
- **Contêinerização**:
  - Adicionar um `Dockerfile` para facilitar a execução e deploy do projeto.
  - Criar um `docker-compose` para gerenciar o banco de dados e a aplicação juntos.

- **CI/CD**:
  - Configurar um pipeline automatizado para testes, build e deploy usando GitHub Actions ou GitLab CI/CD.

### **Interface do Usuário**
- **Frontend**:
  - Desenvolver uma interface gráfica para consumir a API.
  - Criar dashboards para visualizar dados das centrais em tempo real.

---
