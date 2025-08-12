# Guia de Instalação Rápida - VFinance Registry

## 🚀 Início Rápido

### Pré-requisitos
```bash
# Verificar versões
docker --version          # >= 24.0
docker-compose --version  # >= 2.0
go version                # >= 1.21
psql --version            # >= 15.0
```

### 1. Clonar e Preparar Ambiente

```bash
# Clonar repositório
git clone <repository-url>
cd mobx-api

# Configurar variáveis de ambiente
cp .env.example .env
# Editar .env com suas configurações
```

### 2. Exemplo de .env

```bash
# API Configuration
API_PORT=3000
JWT_SECRET=sua_chave_secreta_super_forte_aqui_2024

# Database
DATABASE_URL=postgres://vfinance:senha123@localhost:5432/vfinance?sslmode=disable

# Blockchain
ETHEREUM_RPC=http://localhost:8545
CONTRACT_ADDRESS=0x... # Será preenchido após deploy
PRIVATE_KEY=0x627306090abaB3A6e1400e9345bC60c78a8BEf57

# Security
RATE_LIMIT_WINDOW=900000
RATE_LIMIT_MAX=100
```

### 3. Inicializar Blockchain (4 comandos)

```bash
# Navegar para infra
cd infra/

# Dar permissões
chmod +x *.sh Node-*/run.sh

# Iniciar todos os nós
./start-all-nodes.sh

# Verificar se estão rodando
docker ps | grep besu
```

### 4. Deploy Smart Contract

```bash
cd solidity/

# Instalar Foundry (se necessário)
curl -L https://foundry.paradigm.xyz | bash
foundryup

# Compilar
forge build

# Deploy
forge script script/VFinanceRegistry.s.sol \
  --rpc-url http://localhost:8545 \
  --private-key 0x627306090abaB3A6e1400e9345bC60c78a8BEf57 \
  --broadcast

# Copiar endereço do contrato para .env
```

### 5. Configurar Banco de Dados

```bash
# PostgreSQL via Docker
docker run --name vfinance-postgres \
  -e POSTGRES_DB=vfinance \
  -e POSTGRES_USER=vfinance \
  -e POSTGRES_PASSWORD=senha123 \
  -p 5432:5432 \
  -d postgres:15

# Verificar conexão
psql -h localhost -U vfinance -d vfinance -c "SELECT version();"
```

### 6. Iniciar API

```bash
# Voltar para raiz do projeto
cd ..

# Instalar dependências
go mod download

# Executar migrações e iniciar API
go run cmd/main.go
```

### 7. Testar Instalação

```bash
# Health check
curl http://localhost:3000/health

# Gerar token
curl -X POST http://localhost:3000/api/auth/token \
  -H "Content-Type: application/json" \
  -d '{}'

# Verificar blockchain
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```

## 🐳 Instalação com Docker Compose

### Arquivo docker-compose.yml

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15
    container_name: vfinance-postgres
    environment:
      POSTGRES_DB: vfinance
      POSTGRES_USER: vfinance
      POSTGRES_PASSWORD: senha123
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./init.sql:/docker-entrypoint-initdb.d/init.sql
    ports:
      - "5432:5432"
    networks:
      - vfinance-network

  besu-node-1:
    image: hyperledger/besu:latest
    container_name: besu-node-1
    command: |
      --data-path=/var/lib/besu
      --genesis-file=/var/lib/besu/genesis.json
      --node-private-key-file=/var/lib/besu/key
      --rpc-http-enabled
      --rpc-http-host=0.0.0.0
      --rpc-http-port=8545
      --rpc-http-api=ETH,NET,QBFT,ADMIN,DEBUG,WEB3
      --rpc-http-cors-origins="*"
      --host-allowlist="*"
      --min-gas-price=0
      --p2p-host=0.0.0.0
      --p2p-port=30303
      --logging=INFO
    volumes:
      - ./infra/networkFiles/genesis.json:/var/lib/besu/genesis.json
      - ./infra/networkFiles/keys/0x101b5a9b5d13f532a2d62a1339d2425e26fdb68b/key:/var/lib/besu/key
    ports:
      - "8545:8545"
      - "30303:30303"
    networks:
      - vfinance-network

  api:
    build: .
    container_name: vfinance-api
    environment:
      API_PORT: 3000
      DATABASE_URL: postgres://vfinance:senha123@postgres:5432/vfinance?sslmode=disable
      ETHEREUM_RPC: http://besu-node-1:8545
      CONTRACT_ADDRESS: "0x..." # Preencher após deploy
      PRIVATE_KEY: "0x627306090abaB3A6e1400e9345bC60c78a8BEf57"
      JWT_SECRET: "sua_chave_secreta_super_forte_aqui_2024"
    ports:
      - "3000:3000"
    depends_on:
      - postgres
      - besu-node-1
    networks:
      - vfinance-network

volumes:
  postgres_data:

networks:
  vfinance-network:
    driver: bridge
```

### Execução Docker

```bash
# Subir toda a stack
docker-compose up -d

# Logs
docker-compose logs -f

# Parar
docker-compose down
```

## 🔧 Comandos Úteis

### Blockchain
```bash
# Status dos nós
./infra/manage-services.sh status

# Logs do Node-1
docker logs besu-node-1 -f

# Restart completo
./infra/stop-all-nodes.sh && ./infra/start-all-nodes.sh
```

### API
```bash
# Rebuild e restart
go build -o vfinance-api cmd/main.go
./vfinance-api

# Logs com detalhes
go run cmd/main.go --verbose
```

### Database
```bash
# Conectar ao banco
psql postgres://vfinance:senha123@localhost:5432/vfinance

# Reset completo
docker-compose down -v
docker-compose up -d postgres
```

## 📋 Checklist de Verificação

### ✅ Blockchain
- [ ] 4 nós Besu rodando
- [ ] RPC respondendo na porta 8545
- [ ] Blocos sendo minerados
- [ ] Smart contract deployado

### ✅ API
- [ ] Servidor rodando na porta 3000
- [ ] Conexão com banco funcionando
- [ ] Conexão com blockchain ativa
- [ ] Endpoints respondendo

### ✅ Banco de Dados
- [ ] PostgreSQL ativo
- [ ] Tabelas criadas (migrações)
- [ ] Conexão estabelecida

## 🆘 Troubleshooting

### Problema: Nós não conectam
```bash
# Verificar portas
netstat -tulpn | grep -E "(8545|30303)"

# Verificar logs
docker logs besu-node-1
```

### Problema: API não conecta ao banco
```bash
# Testar conexão manual
psql $DATABASE_URL -c "SELECT 1"

# Verificar logs da API
```

### Problema: Smart contract não deploya
```bash
# Verificar RPC
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"net_version","params":[],"id":1}'

# Verificar saldo da conta
forge script script/CheckBalance.s.sol --rpc-url http://localhost:8545
```

## 📞 Suporte

### Logs Importantes
```bash
# API logs
tail -f /var/log/vfinance-api.log

# Blockchain logs
journalctl -u besu-node-1.service -f

# Database logs
docker logs vfinance-postgres
```

### Monitoramento
- **API Health**: http://localhost:3000/health
- **Node Metrics**: http://localhost:9547/metrics
- **Database**: Conexão via psql

---

**⚡ Setup Completo em ~10 minutos**
**🔄 Última atualização**: Janeiro 2024
