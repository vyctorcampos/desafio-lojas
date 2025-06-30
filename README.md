
# Gerenciador de lojas
Sistema gerenciador de lojas e estabelecimentos.

## ------------------------------TECNOLOGIAS------------------------
* Echo (Vue 3)
* Nuxt 3 (Vue 3) 
* PostgreSQL 15.x.x
* Gorm
* Docker
* Git
* Docker Compose

## -------------------------------URL------------------------
**Frontend:** [http://localhost:3000](http://localhost:3000)  
**Backend:** [http://localhost:8080](http://localhost:8080)  
**Porta do banco de dados:** 5432

## -------------------------------Docker------------------------
Download da ferramenta: https://www.docker.com/products/docker-desktop/

Abra o terminal (CMD, PowerShell ou Docker Prompt) na raiz do projeto `../desafio-lojas`  
Comando:

```bash
docker-compose up --build -d
```

**Obs:**  
Em caso de problema com o backend, execute:

```bash
docker-compose build --no-cache backend
```

Em caso de problema com o frontend, execute:

```bash
docker compose down --volumes --remove-orphans
```

E depois execute novamente o comando para subir o projeto.

## -------------------------------DATABASE------------------------
PostgreSQL: https://www.postgresql.org/download/  
Recomendação de ferramenta: https://dbeaver.io/download/

Credenciais de conexão:

```go
host     := getEnv("DB_HOST", "localhost")
port     := getEnv("DB_PORT", "5432")
user     := getEnv("DB_USER", "postgres")
password := getEnv("DB_PASSWORD", "123")
dbname   := getEnv("DB_NAME", "desafio_lojas")
```

## -------------------------------ENDPOINT-------------------------

### Autenticação
**POST** - `/login` — Simula login

### Estabelecimentos
**POST** - `/estabelecimentos` — Criar um novo estabelecimento  
**GET** - `/estabelecimentos` — Listar todos os estabelecimentos  
**GET** - `/estabelecimentos/:id` — Buscar um estabelecimento por ID  
**PUT** - `/estabelecimentos/:id` — Editar um estabelecimento  
**DELETE** - `/estabelecimentos/:id` — Remover um estabelecimento  

### Lojas
**POST** - `/lojas` — Criar uma nova loja  
**GET** - `/lojas` — Listar todas as lojas  
**GET** - `/lojas/:id` — Buscar uma loja por ID  
**PUT** - `/lojas/:id` — Editar uma loja  
**DELETE** - `/lojas/:id` — Remover uma loja  
**GET** - `/estabelecimentos/:id/lojas` — Listar lojas de um estabelecimento específico  

## -------------------------------Script-------------------------
**Obs:** Apenas a nível de conhecimento, mas o Gorm já cria automaticamente as tabelas e relacionamentos no banco.

### Tabela de estabelecimentos

```sql
CREATE TABLE estabelecimentos (
    id SERIAL PRIMARY KEY,
    numero_estabelecimento VARCHAR(50) NOT NULL,
    nome VARCHAR(100) NOT NULL,
    razao_social VARCHAR(100),
    endereco VARCHAR(200),
    cidade VARCHAR(100),
    estado VARCHAR(2),
    cep VARCHAR(10),
    numero VARCHAR(20)
);
```

### Tabela de lojas

```sql
CREATE TABLE lojas (
    id SERIAL PRIMARY KEY,
    numero_loja VARCHAR(50) NOT NULL,
    nome VARCHAR(100) NOT NULL,
    razao_social VARCHAR(100),
    endereco VARCHAR(200),
    cidade VARCHAR(100),
    estado VARCHAR(2),
    cep VARCHAR(10),
    numero VARCHAR(20),
    estabelecimento_id INTEGER REFERENCES estabelecimentos(id) ON DELETE RESTRICT
);
```
