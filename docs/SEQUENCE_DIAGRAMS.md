# 📋 Diagramas de Sequência - Sistema VFinance

## 🏗️ Arquitetura Geral

O sistema VFinance é composto por:
- **API Go** (`/api`) - Backend da aplicação
- **Smart Contract Solidity** (`/solidity`) - Contrato inteligente na blockchain
- **4 Nodes Hyperledger Besu** - Rede blockchain privada

## 🔄 Fluxo Principal: Registro de Contrato

### 1. Inicialização do Sistema

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Contract as Smart Contract
    participant Node1 as Node Besu 1
    participant Node2 as Node Besu 2
    participant Node3 as Node Besu 3
    participant Node4 as Node Besu 4

    Client->>API: POST /api/v1/contracts/register
    API->>API: Validação de dados
    API->>API: Geração de metadataHash
    API->>Contract: Simulação da transação
    Contract->>Node1: CallContract (simulação)
    Node1->>Node2: Propagação da simulação
    Node2->>Node3: Propagação da simulação
    Node3->>Node4: Propagação da simulação
    Node4-->>Node3: Resultado da simulação
    Node3-->>Node2: Resultado da simulação
    Node2-->>Node1: Resultado da simulação
    Node1-->>Contract: Resultado da simulação
    Contract-->>API: tokenId + metadataHash
    API->>Contract: Transação real
    Contract->>Node1: Transação
    Node1->>Node2: Propagação da transação
    Node2->>Node3: Propagação da transação
    Node3->>Node4: Propagação da transação
    Node4->>Node4: Mineração do bloco
    Node4-->>Node3: Bloco minerado
    Node3-->>Node2: Bloco minerado
    Node2-->>Node1: Bloco minerado
    Node1-->>Contract: Receipt da transação
    Contract-->>API: Receipt
    API->>API: Parse dos logs
    API-->>Client: Resposta com tokenId + metadataHash
```

### 2. Detalhamento da Simulação

```mermaid
sequenceDiagram
    participant API as API Go
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu

    API->>Client: RegisterContract()
    Client->>Client: Validação de parâmetros
    Client->>Client: Criação do ABI
    Client->>Contract: Simulação via CallContract
    Contract->>Node: eth_call
    Node->>Node: Execução local
    Node-->>Contract: Resultado da simulação
    Contract-->>Client: tokenId + metadataHash
    Client->>Client: Parse dos valores de retorno
    Client-->>API: tokenId + metadataHash
```

### 3. Detalhamento da Transação Real

```mermaid
sequenceDiagram
    participant API as API Go
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu
    participant Network as Rede Besu

    API->>Client: Execução da transação
    Client->>Contract: RegisterContract(auth, params...)
    Contract->>Node: eth_sendTransaction
    Node->>Network: Propagação da transação
    Network->>Network: Consenso entre nós
    Network->>Node: Bloco minerado
    Node-->>Contract: Receipt da transação
    Contract-->>Client: Receipt
    Client->>Client: Parse dos logs
    Client->>Client: Extração de valores
    Client-->>API: Resultado final
```

## 🔍 Consultas e Leitura de Dados

### 4. Consulta de Contrato por Hash

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Service as Contract Service
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu

    Client->>API: GET /api/v1/contracts/hash/{hash}
    API->>Service: GetContractByHash()
    Service->>Client: GetContractByHash()
    Client->>Contract: GetContractByHash()
    Contract->>Node: eth_call
    Node->>Node: Execução local
    Node-->>Contract: Dados do contrato
    Contract-->>Client: ContractRecord + VehicleCore
    Client-->>Service: Dados do contrato
    Service-->>API: Dados formatados
    API-->>Client: Resposta JSON
```

### 5. Verificação de Existência

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Service as Contract Service
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu

    Client->>API: GET /api/v1/contracts/exists/{hash}
    API->>Service: DoesHashExist()
    Service->>Client: DoesHashExist()
    Client->>Contract: DoesHashExist()
    Contract->>Node: eth_call
    Node->>Node: Execução local
    Node-->>Contract: true/false
    Contract-->>Client: Resultado booleano
    Client-->>Service: Resultado
    Service-->>API: Resultado
    API-->>Client: {"exists": true/false}
```

## 🔧 Operações de Atualização

### 6. Atualização de Status

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Service as Contract Service
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu
    participant Network as Rede Besu

    Client->>API: PUT /api/v1/contracts/{id}/status
    API->>Service: UpdateStatus()
    Service->>Client: UpdateStatus()
    Client->>Contract: UpdateStatus(auth, tokenId, active)
    Contract->>Node: eth_sendTransaction
    Node->>Network: Propagação da transação
    Network->>Network: Consenso entre nós
    Network->>Node: Bloco minerado
    Node-->>Contract: Receipt da transação
    Contract-->>Client: Receipt
    Client-->>Service: Confirmação
    Service-->>API: Sucesso
    API-->>Client: {"success": true}
```

### 7. Atualização de Metadata

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Service as Contract Service
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu
    participant Network as Rede Besu

    Client->>API: PUT /api/v1/contracts/{id}/metadata
    API->>Service: UpdateMetadataHash()
    Service->>Client: UpdateMetadataHash()
    Client->>Contract: UpdateMetadataHash(auth, tokenId, newHash)
    Contract->>Node: eth_sendTransaction
    Node->>Network: Propagação da transação
    Network->>Network: Consenso entre nós
    Network->>Node: Bloco minerado
    Node-->>Contract: Receipt da transação
    Contract-->>Client: Receipt
    Client-->>Service: Confirmação
    Service-->>API: Sucesso
    API-->>Client: {"success": true}
```

## 🏷️ Registro de Marcas e Modelos

### 8. Registro de Marca

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Service as Contract Service
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu
    participant Network as Rede Besu

    Client->>API: POST /api/v1/brands
    API->>Service: RegisterBrand()
    Service->>Client: RegisterBrand()
    Client->>Contract: RegisterBrand(auth, brandName)
    Contract->>Node: eth_sendTransaction
    Node->>Network: Propagação da transação
    Network->>Network: Consenso entre nós
    Network->>Node: Bloco minerado
    Node-->>Contract: Receipt da transação
    Contract-->>Client: Receipt
    Client->>Client: Parse dos logs
    Client-->>Service: brandId
    Service-->>API: brandId
    API-->>Client: {"brandId": 123}
```

### 9. Registro de Modelo

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Service as Contract Service
    participant Client as Blockchain Client
    participant Contract as Smart Contract
    participant Node as Node Besu
    participant Network as Rede Besu

    Client->>API: POST /api/v1/models
    API->>Service: RegisterModel()
    Service->>Client: RegisterModel()
    Client->>Contract: RegisterModel(auth, modelName)
    Contract->>Node: eth_sendTransaction
    Node->>Network: Propagação da transação
    Network->>Network: Consenso entre nós
    Network->>Node: Bloco minerado
    Node-->>Contract: Receipt da transação
    Contract-->>Client: Receipt
    Client->>Client: Parse dos logs
    Client-->>Service: modelId
    Service-->>API: modelId
    API-->>Client: {"modelId": 456}
```

## 🌐 Rede Hyperledger Besu

### 10. Topologia da Rede

```mermaid
graph TB
    subgraph "Rede VFinance"
        subgraph "Node 1 (Validator)"
            N1[Besu Node 1<br/>Porta 8545<br/>Validator]
        end

        subgraph "Node 2 (Validator)"
            N2[Besu Node 2<br/>Porta 8546<br/>Validator]
        end

        subgraph "Node 3 (Validator)"
            N3[Besu Node 3<br/>Porta 8547<br/>Validator]
        end

        subgraph "Node 4 (Miner)"
            N4[Besu Node 4<br/>Porta 8548<br/>Miner]
        end

        subgraph "API Go"
            API[API Server<br/>Porta 8080]
        end
    end

    API --> N1
    API --> N2
    API --> N3
    API --> N4

    N1 <--> N2
    N1 <--> N3
    N1 <--> N4
    N2 <--> N3
    N2 <--> N4
    N3 <--> N4
```

### 11. Consenso e Mineração

```mermaid
sequenceDiagram
    participant N1 as Node 1
    participant N2 as Node 2
    participant N3 as Node 3
    participant N4 as Node 4 (Miner)

    Note over N1,N4: Transação recebida
    N1->>N2: Propagação da transação
    N1->>N3: Propagação da transação
    N1->>N4: Propagação da transação

    N2->>N3: Propagação da transação
    N2->>N4: Propagação da transação
    N3->>N4: Propagação da transação

    Note over N4: Mineração do bloco
    N4->>N1: Bloco minerado
    N4->>N2: Bloco minerado
    N4->>N3: Bloco minerado

    Note over N1,N4: Validação do bloco
    N1->>N1: Validação local
    N2->>N2: Validação local
    N3->>N3: Validação local
```

## 📊 Estrutura de Dados

### 12. Modelos de Dados

```mermaid
classDiagram
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

    class BlockchainContractRecord {
        +RegistryId string
        +ContractNumber string
        +ContractDate uint32
        +MetadataHash string
        +Timestamp uint32
        +RegisteredBy string
        +Active bool
    }

    class BlockchainVehicleCore {
        +Chassis string
        +LicensePlate string
        +TotalValue string
        +BrandId uint64
        +ModelId uint64
    }

    ContractRecord ||--|| VehicleCore
    BlockchainContractRecord ||--|| BlockchainVehicleCore
```

## 🔐 Autenticação e Autorização

### 13. Fluxo de Autenticação

```mermaid
sequenceDiagram
    participant Client as Cliente
    participant API as API Go
    participant Auth as Auth Service
    participant Wallet as Wallet Manager

    Client->>API: Request com credenciais
    API->>Auth: Validação de credenciais
    Auth->>Wallet: Geração de transação
    Wallet->>Wallet: Assinatura da transação
    Wallet-->>Auth: Transação assinada
    Auth-->>API: Token de acesso
    API-->>Client: JWT Token
```

## 📝 Eventos da Blockchain

### 14. Eventos do Contrato

```mermaid
sequenceDiagram
    participant Contract as Smart Contract
    participant Node as Node Besu
    participant API as API Go
    participant Client as Cliente

    Note over Contract: Evento emitido
    Contract->>Node: Event ContractRegistered
    Node->>API: Log do evento
    API->>API: Parse do evento
    API->>Client: WebSocket/SSE
    Client->>Client: Atualização da UI
```

## 🚀 Deploy e Configuração

### 15. Processo de Deploy

```mermaid
sequenceDiagram
    participant Dev as Desenvolvedor
    participant Foundry as Foundry
    participant Contract as Smart Contract
    participant Network as Rede Besu
    participant API as API Go

    Dev->>Foundry: forge build
    Foundry->>Contract: Compilação
    Dev->>Foundry: forge deploy
    Foundry->>Network: Deploy do contrato
    Network-->>Foundry: Endereço do contrato
    Foundry-->>Dev: Endereço
    Dev->>API: Configuração do endereço
    API->>API: Inicialização do cliente
```

## 📋 Resumo dos Endpoints

### API Endpoints

| Método | Endpoint | Descrição |
|--------|----------|------------|
| POST | `/api/v1/contracts/register` | Registra novo contrato |
| GET | `/api/v1/contracts/hash/{hash}` | Busca contrato por hash |
| GET | `/api/v1/contracts/exists/{hash}` | Verifica existência |
| PUT | `/api/v1/contracts/{id}/status` | Atualiza status |
| PUT | `/api/v1/contracts/{id}/metadata` | Atualiza metadata |
| POST | `/api/v1/brands` | Registra marca |
| POST | `/api/v1/models` | Registra modelo |

### Smart Contract Functions

| Função | Tipo | Descrição |
|--------|------|-----------|
| `registerContract` | Write | Registra contrato |
| `getContractByHash` | View | Busca contrato |
| `doesHashExist` | View | Verifica existência |
| `updateStatus` | Write | Atualiza status |
| `updateMetadataHash` | Write | Atualiza metadata |
| `registerBrand` | Write | Registra marca |
| `registerModel` | Write | Registra modelo |

## 🔧 Configuração dos Nodes

### Node 1 (Validator)
```bash
besu --data-path=/data/node1 \
     --genesis-file=/config/genesis.json \
     --bootnodes=enode://... \
     --p2p-port=30303 \
     --rpc-http-port=8545 \
     --rpc-http-cors-origins="*" \
     --rpc-http-apis=ETH,NET,WEB3,DEBUG,ADMIN \
     --host-allowlist="*" \
     --miner-enabled=false
```

### Node 2 (Validator)
```bash
besu --data-path=/data/node2 \
     --genesis-file=/config/genesis.json \
     --bootnodes=enode://... \
     --p2p-port=30304 \
     --rpc-http-port=8546 \
     --rpc-http-cors-origins="*" \
     --rpc-http-apis=ETH,NET,WEB3,DEBUG,ADMIN \
     --host-allowlist="*" \
     --miner-enabled=false
```

### Node 3 (Validator)
```bash
besu --data-path=/data/node3 \
     --genesis-file=/config/genesis.json \
     --bootnodes=enode://... \
     --p2p-port=30305 \
     --rpc-http-port=8547 \
     --rpc-http-cors-origins="*" \
     --rpc-http-apis=ETH,NET,WEB3,DEBUG,ADMIN \
     --host-allowlist="*" \
     --miner-enabled=false
```

### Node 4 (Miner)
```bash
besu --data-path=/data/node4 \
     --genesis-file=/config/genesis.json \
     --bootnodes=enode://... \
     --p2p-port=30306 \
     --rpc-http-port=8548 \
     --rpc-http-cors-origins="*" \
     --rpc-http-apis=ETH,NET,WEB3,DEBUG,ADMIN \
     --host-allowlist="*" \
     --miner-enabled=true \
     --min-gas-price=0
```

## 📊 Métricas e Monitoramento

### Métricas dos Nodes
- **Block Height**: Altura atual da blockchain
- **Peer Count**: Número de nós conectados
- **Transaction Pool**: Transações pendentes
- **Gas Used**: Gás consumido por bloco
- **Block Time**: Tempo médio entre blocos

### Métricas da API
- **Request Rate**: Taxa de requisições por segundo
- **Response Time**: Tempo médio de resposta
- **Error Rate**: Taxa de erros
- **Active Connections**: Conexões ativas
- **Blockchain Calls**: Chamadas para a blockchain

---

## 📚 Referências

- [Hyperledger Besu Documentation](https://besu.hyperledger.org/)
- [Go Ethereum Documentation](https://geth.ethereum.org/docs/)
- [Solidity Documentation](https://docs.soliditylang.org/)
- [Foundry Book](https://book.getfoundry.sh/)


