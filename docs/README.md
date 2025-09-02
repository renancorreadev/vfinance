# 📚 Documentação do Sistema VFinance

## 🎯 Visão Geral

O sistema VFinance é uma plataforma blockchain para registro e gerenciamento de contratos financeiros de veículos, construída com:

- **Backend**: API Go com integração blockchain
- **Smart Contract**: Solidity na rede Hyperledger Besu
- **Rede**: 4 nodes Besu (3 validators + 1 miner)
- **Bindings**: Go bindings gerados automaticamente via `abigen`

## 📁 Estrutura da Documentação

### 🔄 [Diagramas de Sequência](./SEQUENCE_DIAGRAMS.md)
Documentação detalhada de todos os fluxos do sistema:
- Registro de contratos
- Consultas e leituras
- Operações de atualização
- Registro de marcas e modelos
- Consenso e mineração

### 🏗️ [Diagramas de Arquitetura](./ARCHITECTURE_DIAGRAMS.md)
Visão arquitetural do sistema:
- Topologia da rede
- Estrutura de dados
- Fluxos de autenticação
- Pipeline de deploy

## 🚀 Início Rápido

### 1. Pré-requisitos
```bash
# Go 1.21+
go version

# Foundry
forge --version

# Docker
docker --version
```

### 2. Configuração da Rede Besu
```bash
# Node 1 (Validator)
besu --data-path=/data/node1 \
     --genesis-file=/config/genesis.json \
     --rpc-http-port=8545 \
     --miner-enabled=false

# Node 2 (Validator)
besu --data-path=/data/node2 \
     --genesis-file=/config/genesis.json \
     --rpc-http-port=8546 \
     --miner-enabled=false

# Node 3 (Validator)
besu --data-path=/data/node3 \
     --genesis-file=/config/genesis.json \
     --rpc-http-port=8547 \
     --miner-enabled=false

# Node 4 (Miner)
besu --data-path=/data/node4 \
     --genesis-file=/config/genesis.json \
     --rpc-http-port=8548 \
     --miner-enabled=true
```

### 3. Deploy do Smart Contract
```bash
cd solidity
forge build
forge deploy --rpc-url http://localhost:8545
```

### 4. Configuração da API
```bash
cd api
# Configure o endereço do contrato em .env
echo "CONTRACT_ADDRESS=0x..." > .env

# Execute a API
go run cmd/main.go
```

## 🔧 Principais Funcionalidades

### 📝 Registro de Contratos
- Simulação prévia para obter `tokenId` e `metadataHash`
- Transação real na blockchain
- Parse automático dos logs para extrair valores

### 🔍 Consultas
- Busca por hash de metadata
- Verificação de existência
- Consulta por chassis ou registry ID

### 🔄 Atualizações
- Status do contrato
- Hash de metadata
- Configurações do servidor

### 🏷️ Gestão de Marcas e Modelos
- Registro de marcas
- Registro de modelos
- Mapeamento automático de IDs

## 🌐 Endpoints da API

| Método | Endpoint | Descrição |
|--------|----------|------------|
| POST | `/api/v1/contracts/register` | Registra contrato |
| GET | `/api/v1/contracts/hash/{hash}` | Busca por hash |
| GET | `/api/v1/contracts/exists/{hash}` | Verifica existência |
| PUT | `/api/v1/contracts/{id}/status` | Atualiza status |
| PUT | `/api/v1/contracts/{id}/metadata` | Atualiza metadata |
| POST | `/api/v1/brands` | Registra marca |
| POST | `/api/v1/models` | Registra modelo |

## 📊 Estrutura de Dados

### Smart Contract
```solidity
struct ContractRecord {
    bytes32 registryId;
    bytes32 contractNumber;
    uint32 contractDate;
    bytes32 metadataHash;
    uint32 timestamp;
    address registeredBy;
    bool active;
}

struct VehicleCore {
    bytes32 chassis;
    bytes32 licensePlate;
    uint128 totalValue;
    uint64 brandId;
    uint64 modelId;
}
```

### API Go
```go
type BlockchainContractRecord struct {
    RegistryId     string `json:"registryId"`
    ContractNumber string `json:"contractNumber"`
    ContractDate   uint32 `json:"contractDate"`
    MetadataHash   string `json:"metadataHash"`
    Timestamp      uint32 `json:"timestamp"`
    RegisteredBy   string `json:"registeredBy"`
    Active         bool   `json:"active"`
}
```

## 🔐 Autenticação e Segurança

- **JWT Tokens** para autenticação da API
- **Assinatura de transações** via wallet privada
- **Validação de parâmetros** em todas as operações
- **Controle de acesso** baseado em roles

## 📈 Monitoramento e Métricas

### Blockchain
- Altura dos blocos
- Número de peers
- Pool de transações
- Consumo de gas

### API
- Taxa de requisições
- Tempo de resposta
- Taxa de erros
- Conexões ativas

## 🚨 Troubleshooting

### Erro: "Execution reverted"
- Verifique se os parâmetros estão corretos
- Confirme se o contrato está inicializado
- Verifique se há saldo suficiente para gas

### Erro: "Failed to unpack return values"
- Confirme se os bindings estão atualizados
- Verifique se o ABI está correto
- Execute `go generate` para regenerar bindings

### Erro: "Transaction failed"
- Verifique o status da transação
- Confirme se foi minerada
- Verifique os logs para detalhes

## 🔄 Fluxo de Desenvolvimento

1. **Desenvolvimento**: Modifique o código Go ou Solidity
2. **Testes**: Execute os testes unitários
3. **Build**: Compile o código
4. **Deploy**: Faça deploy do contrato (se necessário)
5. **Configuração**: Atualize as configurações da API
6. **Teste**: Teste a funcionalidade

## 📚 Referências

- [Hyperledger Besu](https://besu.hyperledger.org/)
- [Go Ethereum](https://geth.ethereum.org/)
- [Solidity](https://docs.soliditylang.org/)
- [Foundry](https://book.getfoundry.sh/)

## 🤝 Contribuição

1. Fork o repositório
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](../LICENSE) para mais detalhes.

---

## 📞 Suporte

Para dúvidas ou problemas:
- Abra uma issue no GitHub
- Consulte a documentação técnica
- Entre em contato com a equipe de desenvolvimento


