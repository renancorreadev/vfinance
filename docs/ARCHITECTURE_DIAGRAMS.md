# 🏗️ Diagramas de Arquitetura - Sistema VFinance

## 🌐 Arquitetura da Rede Blockchain

```mermaid
graph TB
    subgraph "Cliente"
        C[Cliente Web/Mobile]
    end

    subgraph "API Layer"
        API[API Go Server<br/>Porta 8080]
        AUTH[Auth Service]
        SERVICE[Contract Service]
    end

    subgraph "Blockchain Layer"
        CLIENT[Blockchain Client]
        BINDINGS[Go Bindings]
        CONTRACT[Smart Contract<br/>VFinanceRegistry]
    end

    subgraph "Hyperledger Besu Network"
        N1[Node 1<br/>Validator<br/>Porta 8545]
        N2[Node 2<br/>Validator<br/>Porta 8546]
        N3[Node 3<br/>Validator<br/>Porta 8547]
        N4[Node 4<br/>Miner<br/>Porta 8548]
    end

    C --> API
    API --> AUTH
    API --> SERVICE
    SERVICE --> CLIENT
    CLIENT --> BINDINGS
    BINDINGS --> CONTRACT
    CONTRACT --> N1
    CONTRACT --> N2
    CONTRACT --> N3
    CONTRACT --> N4

    N1 <--> N2
    N1 <--> N3
    N1 <--> N4
    N2 <--> N3
    N2 <--> N4
    N3 <--> N4
```

## 🔄 Fluxo de Registro de Contrato

```mermaid
sequenceDiagram
    participant C as Cliente
    participant API as API Go
    participant S as Service
    participant BC as Blockchain Client
    participant SC as Smart Contract
    participant N as Node Besu

    C->>API: POST /contracts/register
    API->>S: RegisterContract()
    S->>BC: RegisterContract()

    Note over BC: Simulação
    BC->>SC: CallContract (simulação)
    SC->>N: eth_call
    N-->>SC: Resultado
    SC-->>BC: tokenId + metadataHash

    Note over BC: Transação Real
    BC->>SC: RegisterContract (transação)
    SC->>N: eth_sendTransaction
    N-->>SC: Receipt
    SC-->>BC: Receipt

    BC->>BC: Parse logs
    BC-->>S: Resultado final
    S-->>API: Sucesso
    API-->>C: {"tokenId": 123, "metadataHash": "0x..."}
```

## 📊 Estrutura de Dados

```mermaid
classDiagram
    class APIRequest {
        +RegistryId string
        +ContractNumber string
        +ContractDate uint32
        +Chassis string
        +LicensePlate string
        +TotalValue string
        +BrandName string
        +ModelName string
    }

    class ContractRecord {
        +RegistryId [32]byte
        +ContractNumber [32]byte
        +ContractDate uint32
        +MetadataHash [32]byte
        +Timestamp uint32
        +RegisteredBy address
        +Active bool
    }

    class VehicleCore {
        +Chassis [32]byte
        +LicensePlate [32]byte
        +TotalValue *big.Int
        +BrandId uint64
        +ModelId uint64
    }

    class APIResponse {
        +TokenId *big.Int
        +MetadataHash string
        +Success bool
        +Message string
    }

    APIRequest --> ContractRecord
    ContractRecord --> VehicleCore
    ContractRecord --> APIResponse
```

## 🔐 Fluxo de Autenticação

```mermaid
sequenceDiagram
    participant C as Cliente
    participant API as API Go
    participant AUTH as Auth Service
    participant W as Wallet Manager

    C->>API: Login Request
    API->>AUTH: Validate Credentials
    AUTH->>W: Generate Transaction
    W->>W: Sign Transaction
    W-->>AUTH: Signed Transaction
    AUTH-->>API: JWT Token
    API-->>C: Access Token

    Note over C,API: Subsequent Requests
    C->>API: API Request + Token
    API->>AUTH: Validate Token
    AUTH-->>API: Valid/Invalid
    API-->>C: Response
```

## 📝 Eventos da Blockchain

```mermaid
sequenceDiagram
    participant SC as Smart Contract
    participant N as Node Besu
    participant API as API Go
    participant C as Cliente

    Note over SC: Event Emitted
    SC->>N: ContractRegistered Event
    N->>API: Event Log
    API->>API: Parse Event
    API->>C: WebSocket/SSE Update
    C->>C: UI Update
```

## 🚀 Deploy Pipeline

```mermaid
graph LR
    subgraph "Development"
        CODE[Source Code]
        TEST[Tests]
    end

    subgraph "Build"
        BUILD[Go Build]
        COMPILE[Solidity Compile]
        BIND[Generate Bindings]
    end

    subgraph "Deploy"
        DEPLOY[Deploy Contract]
        CONFIG[Configure API]
        START[Start Services]
    end

    CODE --> TEST
    TEST --> BUILD
    BUILD --> COMPILE
    COMPILE --> BIND
    BIND --> DEPLOY
    DEPLOY --> CONFIG
    CONFIG --> START
```

## 📋 Configuração dos Nodes

### Node 1 (Validator)
```bash
besu --data-path=/data/node1 \
     --genesis-file=/config/genesis.json \
     --p2p-port=30303 \
     --rpc-http-port=8545 \
     --miner-enabled=false
```

### Node 2 (Validator)
```bash
besu --data-path=/data/node2 \
     --genesis-file=/config/genesis.json \
     --p2p-port=30304 \
     --rpc-http-port=8546 \
     --miner-enabled=false
```

### Node 3 (Validator)
```bash
besu --data-path=/data/node3 \
     --genesis-file=/config/genesis.json \
     --p2p-port=30305 \
     --rpc-http-port=8547 \
     --miner-enabled=false
```

### Node 4 (Miner)
```bash
besu --data-path=/data/node4 \
     --genesis-file=/config/genesis.json \
     --p2p-port=30306 \
     --rpc-http-port=8548 \
     --miner-enabled=true
```

## 🔧 Endpoints da API

| Método | Endpoint | Descrição |
|--------|----------|------------|
| POST | `/api/v1/contracts/register` | Registra contrato |
| GET | `/api/v1/contracts/hash/{hash}` | Busca por hash |
| GET | `/api/v1/contracts/exists/{hash}` | Verifica existência |
| PUT | `/api/v1/contracts/{id}/status` | Atualiza status |
| PUT | `/api/v1/contracts/{id}/metadata` | Atualiza metadata |
| POST | `/api/v1/brands` | Registra marca |
| POST | `/api/v1/models` | Registra modelo |

## 📊 Métricas de Monitoramento

- **Block Height**: Altura atual da blockchain
- **Peer Count**: Nós conectados
- **Transaction Pool**: Transações pendentes
- **Gas Used**: Gás consumido por bloco
- **API Response Time**: Tempo de resposta da API
- **Error Rate**: Taxa de erros
- **Active Connections**: Conexões ativas


