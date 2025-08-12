# Sistema de Registro de Financiamentos Automotivos em Blockchain

## Índice
1. [Visão Geral](#visão-geral)
2. [Arquitetura da Solução](#arquitetura-da-solução)
3. [Contrato Inteligente VFinanceRegistry](#contrato-inteligente-vfinanceregistry)
4. [API Backend](#api-backend)
5. [Infraestrutura Blockchain](#infraestrutura-blockchain)
6. [Fluxos de Funcionamento](#fluxos-de-funcionamento)
7. [Guia de Implantação](#guia-de-implantação)
8. [Referência da API](#referência-da-api)
9. [Segurança e Considerações](#segurança-e-considerações)

---

## Visão Geral

### Objetivo
Sistema distribuído para registro imutável de contratos de financiamento automotivos utilizando tecnologia blockchain, garantindo transparência, auditabilidade e integridade dos dados.

### Principais Características
- **Dados On-Chain**: Informações essenciais do contrato registradas diretamente na blockchain
- **Dados Off-Chain**: Metadados detalhados armazenados em banco PostgreSQL
- **Hash de Integridade**: Sistema de hash único para vincular dados on-chain e off-chain
- **API RESTful**: Interface completa para interação com o sistema
- **Rede Privada**: Infraestrutura Hyperledger Besu com 4 nós em consenso QBFT

### Tecnologias Utilizadas
- **Blockchain**: Hyperledger Besu v23.10+
- **Consenso**: QBFT (Byzantine Fault Tolerant)
- **Smart Contract**: Solidity v0.8.19
- **Backend**: Go v1.21+ com Gin Framework
- **Database**: PostgreSQL 15+
- **Autenticação**: JWT
- **Deploy**: Docker & Docker Compose

---

## Arquitetura da Solução

### Diagrama de Arquitetura
```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend/     │    │   API Backend   │    │   Blockchain    │
│   Cliente       │────│   (Go/Gin)      │────│   Hyperledger   │
│                 │    │                 │    │   Besu QBFT     │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │                       │
                                │                       │
                       ┌─────────────────┐     ┌─────────────────┐
                       │   PostgreSQL    │     │ Smart Contract  │
                       │   (Metadados)   │     │ VFinanceRegistry│
                       └─────────────────┘     └─────────────────┘
```

### Componentes Principais

#### 1. Smart Contract VFinanceRegistry
- **Localização**: `solidity/src/VFinanceRegistry.sol`
- **Função**: Registro imutável de dados primários do contrato
- **Dados Armazenados**: regConId, numeroContrato, dataContrato, metadataHash, timestamp

#### 2. API Backend (Go)
- **Localização**: `internal/`
- **Responsabilidades**:
  - Interface entre frontend e blockchain
  - Gerenciamento de metadados off-chain
  - Autenticação e autorização
  - Validação de dados

#### 3. Infraestrutura Blockchain
- **Localização**: `infra/`
- **Configuração**: 4 nós Hyperledger Besu
- **Consenso**: QBFT com 2 segundos por bloco
- **Network ID**: 1337

---

## Contrato Inteligente VFinanceRegistry

### Estrutura Principal

#### ContractRecord
```solidity
struct ContractRecord {
    string regConId;           // ID único do contrato
    string numeroContrato;     // Número do contrato
    string dataContrato;       // Data do contrato
    bytes32 metadataHash;      // Hash para buscar metadados
    uint256 timestamp;         // Timestamp do registro
    address registeredBy;      // Endereço que registrou
    bool active;              // Status ativo/inativo
}
```

### Funcionalidades Principais

#### 1. Registro de Contratos
```solidity
function registerContract(
    string calldata regConId,
    string calldata numeroContrato,
    string calldata dataContrato
) external onlyOwnerOrApiServer returns (bytes32 metadataHash)
```

**Características:**
- Gera hash único automaticamente
- Valida duplicatas
- Emite evento `ContractRegistered`
- Retorna hash para armazenamento de metadados

#### 2. Consulta de Contratos
```solidity
function getContract(string calldata regConId)
    external view returns (ContractRecord memory)
```

#### 3. Busca por Hash
```solidity
function getContractByHash(bytes32 metadataHash)
    external view returns (ContractRecord memory)
```

#### 4. URL dos Metadados
```solidity
function tokenURI(bytes32 metadataHash)
    public view returns (string memory)
```

### Eventos Principais
- `ContractRegistered`: Emitido no registro de novo contrato
- `MetadataHashUpdated`: Emitido na atualização de hash
- `ContractStatusUpdated`: Emitido na mudança de status

### Controles de Acesso
- **Owner**: Controle total do contrato
- **API Server**: Autorizado para registros e atualizações
- **Público**: Acesso de leitura

---

## API Backend

### Estrutura de Diretórios
```
internal/
├── blockchain/     # Cliente para interação com blockchain
├── config/         # Configurações da aplicação
├── database/       # Conexão e migrações do banco
├── handlers/       # Controladores HTTP
├── middleware/     # Middlewares (auth, rate limit)
├── models/         # Modelos de dados
├── server/         # Configuração do servidor
└── services/       # Lógica de negócio
```

### Componentes Principais

#### 1. Blockchain Client (`blockchain/client.go`)
Responsável pela comunicação com a blockchain:
- Conexão com nó Besu via RPC
- Interação com smart contract
- Assinatura de transações
- Decodificação de eventos

**Métodos Principais:**
- `RegisterContract()`: Registra contrato na blockchain
- `GetContract()`: Busca dados on-chain
- `GetActiveContracts()`: Lista contratos ativos

#### 2. Services Layer

##### Contract Service (`services/contract_service.go`)
- Orquestra operações entre blockchain e banco de dados
- Combina dados on-chain e off-chain
- Gerencia ciclo de vida dos contratos

##### Metadata Service (`services/metadata_service.go`)
- Armazena e recupera metadados detalhados
- Gera e valida hashes de integridade
- Gerencia dados do veículo e financiamento

#### 3. Models (`models/models.go`)

##### VehicleData
Estrutura completa com mais de 60 campos incluindo:
- Dados do agente financeiro
- Informações do proprietário
- Detalhes do veículo (chassi, placa, modelo)
- Termos do contrato (valores, parcelas, taxas)

##### ContractRecord
Mapeamento dos dados on-chain do smart contract.

### Endpoints da API

#### Autenticação
- `POST /api/auth/token` - Gerar token JWT
- `GET /api/auth/validate` - Validar token

#### Contratos
- `POST /api/contracts/` - Registrar novo contrato
- `GET /api/contracts/:regConId` - Buscar contrato por ID
- `GET /api/contracts/active` - Listar contratos ativos
- `GET /api/contracts/hash/:hash` - Buscar por hash
- `GET /api/contracts/stats` - Estatísticas gerais

#### Metadados
- `POST /api/metadata/:hash` - Armazenar metadados
- `GET /api/metadata/:hash` - Recuperar metadados
- `PUT /api/metadata/:hash` - Atualizar metadados
- `DELETE /api/metadata/:hash` - Remover metadados

---

## Infraestrutura Blockchain

### Configuração da Rede

#### Genesis Block (`networkFiles/genesis.json`)
```json
{
  "config": {
    "chainId": 1337,
    "berlinBlock": 0,
    "qbft": {
      "blockperiodseconds": 2,
      "epochlength": 30000,
      "requesttimeoutseconds": 4
    }
  }
}
```

**Características:**
- **Chain ID**: 1337 (rede privada)
- **Consenso QBFT**: Tolerante a falhas bizantinas
- **Tempo de bloco**: 2 segundos
- **Época**: 30.000 blocos
- **Timeout**: 4 segundos

### Topologia de Nós

#### Node-1 (Bootstrap/RPC Principal)
- **RPC Port**: 8545
- **P2P Port**: 30303
- **WebSocket**: 6174
- **Métricas**: 9547
- **Função**: Nó principal para conexões RPC

#### Node-2, Node-3, Node-4 (Validadores)
- **Portas**: Incrementais a partir do Node-1
- **Função**: Nós validadores para consenso QBFT
- **Bootnode**: Conectam ao Node-1

### Scripts de Controle

#### Inicialização Individual
```bash
# Node-1
./infra/Node-1/run.sh

# Node-2
./infra/Node-2/run.sh

# Node-3
./infra/Node-3/run.sh

# Node-4
./infra/Node-4/run.sh
```

#### Controle Centralizado
```bash
# Iniciar todos os nós
./infra/start-all-nodes.sh

# Parar todos os nós
./infra/stop-all-nodes.sh
```

---

## Fluxos de Funcionamento

### 1. Registro de Contrato

```mermaid
sequenceDiagram
    participant C as Cliente
    participant A as API
    participant B as Blockchain
    participant D as Database

    C->>A: POST /api/contracts/
    A->>A: Validar dados
    A->>B: registerContract()
    B->>B: Gerar hash único
    B->>B: Armazenar on-chain
    B-->>A: txHash + metadataHash
    A->>D: Armazenar metadados
    A->>D: Criar registro local
    A-->>C: Resposta de sucesso
```

### 2. Consulta de Contrato

```mermaid
sequenceDiagram
    participant C as Cliente
    participant A as API
    participant B as Blockchain
    participant D as Database

    C->>A: GET /api/contracts/:id
    A->>B: getContract()
    B-->>A: Dados on-chain
    A->>D: Buscar metadados
    D-->>A: Dados off-chain
    A->>A: Combinar dados
    A-->>C: Contrato completo
```

### 3. Validação de Integridade

O sistema garante integridade através de:
1. **Hash único** gerado pelo smart contract
2. **Vinculação** entre dados on-chain e off-chain
3. **Validação** automática de consistência
4. **Imutabilidade** dos registros blockchain

---

## Guia de Implantação

### Pré-requisitos
- Docker 24.0+
- Docker Compose 2.0+
- Go 1.21+
- PostgreSQL 15+
- Besu 23.10+

### 1. Configuração do Ambiente

#### Variáveis de Ambiente (.env)
```bash
# API Configuration
API_PORT=3000
JWT_SECRET=your_super_secret_key

# Database
DATABASE_URL=postgres://user:password@localhost:5432/vfinance?sslmode=disable

# Blockchain
ETHEREUM_RPC=http://localhost:8545
CONTRACT_ADDRESS=0x... # Endereço do contrato deployado
PRIVATE_KEY=0x... # Chave privada para transações

# Security
RATE_LIMIT_WINDOW=900000
RATE_LIMIT_MAX=100
```

### 2. Inicialização da Blockchain

#### Preparar Infraestrutura
```bash
cd infra/
chmod +x *.sh
chmod +x Node-*/run.sh
```

#### Iniciar Nós
```bash
# Iniciar todos os nós
./start-all-nodes.sh

# Verificar status
docker ps | grep besu
```

#### Verificar Conectividade
```bash
# Testar RPC
curl -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```

### 3. Deploy do Smart Contract

#### Compilação (Foundry)
```bash
cd solidity/
forge build
```

#### Deploy
```bash
forge script script/VFinanceRegistry.s.sol \
  --rpc-url http://localhost:8545 \
  --private-key $PRIVATE_KEY \
  --broadcast
```

### 4. Configuração do Backend

#### Instalação de Dependências
```bash
go mod download
```

#### Migrações do Banco
```bash
# O sistema executa migrações automaticamente no startup
go run cmd/main.go
```

#### Inicialização
```bash
# Desenvolvimento
go run cmd/main.go

# Produção
go build -o vfinance-api cmd/main.go
./vfinance-api
```

### 5. Docker Compose (Produção)

```yaml
version: '3.8'
services:
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: vfinance
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  api:
    build: .
    ports:
      - "3000:3000"
    environment:
      DATABASE_URL: postgres://user:password@postgres:5432/vfinance?sslmode=disable
      ETHEREUM_RPC: http://besu-node-1:8545
    depends_on:
      - postgres
      - besu-node-1

volumes:
  postgres_data:
```

---

## Referência da API

### Autenticação

#### Gerar Token
```http
POST /api/auth/token
Content-Type: application/json

{
  "username": "admin",
  "password": "password"
}
```

**Resposta:**
```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIs...",
  "expires_in": 86400
}
```

### Contratos

#### Registrar Contrato
```http
POST /api/contracts/
Authorization: Bearer {token}
Content-Type: application/json

{
  "regConId": "CONT-2024-001",
  "numeroContrato": "FIN123456789",
  "dataContrato": "2024-01-15",
  "vehicleData": {
    "cnpjAgenteFinanceiro": "12.345.678/0001-90",
    "nomeAgenteFinanceiro": "Banco Exemplo S.A.",
    "cpfCnpjProprietario": "123.456.789-00",
    "nomeProprietario": "João Silva",
    "chassiVeiculo": "9BWZZZ377VT004251",
    "placaVeiculo": "ABC-1234",
    "valorTotalContrato": "45000.00",
    // ... outros campos
  }
}
```

**Resposta:**
```json
{
  "success": true,
  "message": "Contrato registrado com sucesso",
  "regConId": "CONT-2024-001",
  "metadataHash": "0x...",
  "txHash": "0x..."
}
```

#### Buscar Contrato
```http
GET /api/contracts/CONT-2024-001
```

**Resposta:**
```json
{
  "success": true,
  "data": {
    "onChain": {
      "regConId": "CONT-2024-001",
      "numeroContrato": "FIN123456789",
      "dataContrato": "2024-01-15",
      "metadataHash": "0x...",
      "timestamp": 1705334400,
      "registeredBy": "0x...",
      "active": true
    },
    "offChain": {
      "cnpjAgenteFinanceiro": "12.345.678/0001-90",
      // ... dados completos do veículo
    }
  }
}
```

#### Listar Contratos Ativos
```http
GET /api/contracts/active?offset=0&limit=10
```

#### Estatísticas
```http
GET /api/contracts/stats
```

**Resposta:**
```json
{
  "success": true,
  "data": {
    "totalContracts": 1250,
    "activeContracts": 1180
  }
}
```

### Metadados

#### Buscar Metadados
```http
GET /api/metadata/0x1234567890abcdef...
```

**Resposta:**
```json
{
  "success": true,
  "data": {
    "cnpjAgenteFinanceiro": "12.345.678/0001-90",
    "nomeAgenteFinanceiro": "Banco Exemplo S.A.",
    // ... todos os campos de VehicleData
  }
}
```

---

## Segurança e Considerações

### Segurança da Blockchain

#### Consenso QBFT
- **Tolerância a Falhas**: Suporta até (n-1)/3 nós maliciosos
- **Finalidade**: Blocos são finais imediatamente
- **Performance**: 2 segundos por bloco

#### Controles de Acesso
- **Owner**: Controle administrativo total
- **API Server**: Permissões para registros
- **Modifiers**: `onlyOwner`, `onlyOwnerOrApiServer`

### Segurança da API

#### Autenticação JWT
- Tokens com expiração de 24 horas
- Chave secreta configurável via ambiente
- Validação em endpoints protegidos

#### Rate Limiting
- Limite configurável por janela de tempo
- Proteção contra ataques DDoS
- Headers informativos de limite

#### Validação de Dados
- Validação rigorosa de entrada
- Sanitização de parâmetros
- Tratamento de erros padronizado

### Considerações de Performance

#### Blockchain
- **TPS**: ~30 transações por segundo
- **Latência**: 2-4 segundos para confirmação
- **Storage**: Dados mínimos on-chain

#### API
- Pool de conexões com banco otimizado
- Cache de metadados frequentes
- Paginação em consultas grandes

#### Banco de Dados
- Índices otimizados em campos de busca
- JSONB para dados estruturados
- Particionamento por data (recomendado)

### Monitoramento

#### Métricas Blockchain
- Altura dos blocos
- Latência de consenso
- Status dos nós validadores

#### Métricas API
- Tempo de resposta
- Taxa de erro
- Throughput de requests

#### Alertas
- Nós offline
- Falhas de consenso
- Erros de aplicação
- Uso de recursos

---

## Roadmap e Melhorias

### Curto Prazo
- [ ] Implementar cache Redis
- [ ] Adicionar métricas Prometheus
- [ ] Criar dashboard Grafana
- [ ] Implementar backup automatizado

### Médio Prazo
- [ ] API Gateway com Kong
- [ ] Implementar IPFS para documentos
- [ ] Sistema de notificações
- [ ] Auditoria avançada

### Longo Prazo
- [ ] Integração com outras blockchains
- [ ] Sistema de tokens para stake
- [ ] Governança descentralizada
- [ ] Compliance LGPD/GDPR

---

## Conclusão

O Sistema de Registro de Financiamentos Automotivos em Blockchain oferece uma solução robusta e escalável para registro imutável de contratos, combinando as vantagens da tecnologia blockchain com a flexibilidade de uma arquitetura híbrida on-chain/off-chain.

A implementação garante transparência, auditabilidade e integridade dos dados, atendendo aos requisitos regulatórios do setor financeiro automotivo brasileiro.

---

**Versão**: 1.0
**Data**: Janeiro 2024
**Autores**: Equipe de Desenvolvimento VFinance
**Revisão**: Em desenvolvimento
