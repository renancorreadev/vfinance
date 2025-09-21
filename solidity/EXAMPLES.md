# 🛠️ Exemplos Práticos - VFinance Registry

## 📋 Índice

1. [Setup Inicial](#setup-inicial)
2. [Registro de Contratos](#registro-de-contratos)
3. [Consultas de Dados](#consultas-de-dados)
4. [Sistema de Metadados](#sistema-de-metadados)
5. [Gestão de Marcas/Modelos](#gestão-de-marcasmodelos)
6. [Administração](#administração)
7. [Integração Web3](#integração-web3)
8. [Scripts Forge](#scripts-forge)

---

## 🚀 Setup Inicial

### **1. Configurar Environment**

```bash
# .env
PRIVATE_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
RPC_URL=http://localhost:8545
CONTRACT_ADDRESS=0x5FbDB2315678afecb367f032d93F642f64180aa3
METADATA_BASE_URL=https://api.vfinance.com.br
```

### **2. Verificar Contrato**

```bash
# Testar conexão
forge script script/RegisterContract.s.sol --rpc-url $RPC_URL -s "listExistingContracts()"

# Verificar configuração
cast call $CONTRACT_ADDRESS "name()" --rpc-url $RPC_URL
cast call $CONTRACT_ADDRESS "symbol()" --rpc-url $RPC_URL
cast call $CONTRACT_ADDRESS "owner()" --rpc-url $RPC_URL
```

---

## 📝 Registro de Contratos

### **Exemplo 1: Registro Básico**

```javascript
// Via Web3.js
const Web3 = require('web3');
const web3 = new Web3('http://localhost:8545');

const contract = new web3.eth.Contract(ABI, CONTRACT_ADDRESS);

async function registerBasicContract() {
    const result = await contract.methods.registerContract(
        "VFIN-2024-001",                    // registryId único
        "FIN123456789",                     // número do contrato
        Math.floor(Date.now() / 1000),      // data atual
        "9BWZZZ377VT004251",                // chassi
        "BRA2E19",                          // placa
        web3.utils.toWei("50000", "ether"), // R$ 50.000
        "VOLKSWAGEN",                       // marca
        "GOLF TSI 1.4"                     // modelo
    ).send({
        from: API_SERVER_ADDRESS,
        gas: 500000
    });

    console.log('Token ID:', result.events.ContractRegistered.returnValues.tokenId);
    console.log('Metadata Hash:', result.events.ContractRegistered.returnValues.metadataHash);
}
```

### **Exemplo 2: Registro via Cast**

```bash
# Usando Foundry Cast
cast send $CONTRACT_ADDRESS \
    "registerContract(string,string,uint32,string,string,uint128,string,string)" \
    "VFIN-2024-002" \
    "FIN987654321" \
    1703980800 \
    "9AHZZZ377VT004252" \
    "BRA3F28" \
    "45000000000000000000000" \
    "HONDA" \
    "CIVIC LX 1.6" \
    --private-key $PRIVATE_KEY \
    --rpc-url $RPC_URL
```

### **Exemplo 3: Registro com Validação**

```solidity
// Script Solidity completo
contract RegisterWithValidation is Script {
    function run() external {
        vm.startBroadcast();

        VFinanceRegistry registry = VFinanceRegistry(payable(CONTRACT_ADDRESS));

        // Verificar se já existe
        require(!registry.doesContractExist("VFIN-2024-003"), "Contrato já existe");

        // Registrar
        (uint256 tokenId, bytes32 metadataHash) = registry.registerContract(
            "VFIN-2024-003",
            "FIN456789123",
            uint32(block.timestamp),
            "9CGZZZ377VT004253",
            "BRA4G37",
            120000 ether,
            "BMW",
            "320I SPORT 2.0"
        );

        // Verificar resultado
        require(tokenId > 0, "Token ID inválido");
        require(metadataHash != bytes32(0), "Hash de metadata inválido");

        console.log("Contrato registrado com sucesso!");
        console.log("Token ID:", tokenId);

        vm.stopBroadcast();
    }
}
```

### **Exemplo 4: Batch Registration**

```javascript
// Registrar múltiplos contratos
const contracts = [
    {
        registryId: "VFIN-2024-004",
        contractNumber: "FIN111222333",
        chassis: "9DJZZZ377VT004254",
        licensePlate: "BRA5H46",
        totalValue: "35000",
        brand: "TOYOTA",
        model: "COROLLA GLI 1.8"
    },
    {
        registryId: "VFIN-2024-005",
        contractNumber: "FIN444555666",
        chassis: "9EKZZZ377VT004255",
        licensePlate: "BRA6I55",
        totalValue: "65000",
        brand: "AUDI",
        model: "A3 SEDAN 1.4"
    }
];

async function batchRegister() {
    for (const contractData of contracts) {
        try {
            const result = await contract.methods.registerContract(
                contractData.registryId,
                contractData.contractNumber,
                Math.floor(Date.now() / 1000),
                contractData.chassis,
                contractData.licensePlate,
                web3.utils.toWei(contractData.totalValue, "ether"),
                contractData.brand,
                contractData.model
            ).send({
                from: API_SERVER_ADDRESS,
                gas: 500000
            });

            console.log(`✅ ${contractData.registryId} registrado: Token ${result.events.ContractRegistered.returnValues.tokenId}`);
        } catch (error) {
            console.log(`❌ Erro ao registrar ${contractData.registryId}:`, error.message);
        }
    }
}
```

---

## 🔍 Consultas de Dados

### **Exemplo 1: Buscar por Token ID**

```javascript
async function getContractByTokenId(tokenId) {
    const result = await contract.methods.getContract(tokenId).call();

    const contractData = result[0]; // ContractRecord
    const vehicleData = result[1];  // VehicleCore

    console.log('=== Dados do Contrato ===');
    console.log('Registry ID Hash:', contractData.registryId);
    console.log('Contract Number:', web3.utils.hexToUtf8(contractData.contractNumber));
    console.log('Active:', contractData.active);
    console.log('Registered By:', contractData.registeredBy);

    console.log('=== Dados do Veículo ===');
    console.log('Chassis:', web3.utils.hexToUtf8(vehicleData.chassis));
    console.log('License Plate:', web3.utils.hexToUtf8(vehicleData.licensePlate));
    console.log('Total Value:', web3.utils.fromWei(vehicleData.totalValue, 'ether'), 'ETH');
    console.log('Brand ID:', vehicleData.brandId);
    console.log('Model ID:', vehicleData.modelId);
}
```

### **Exemplo 2: Buscar por Registry ID**

```bash
# Via Cast
cast call $CONTRACT_ADDRESS \
    "getContractByRegistryId(string)" \
    "VFIN-2024-001" \
    --rpc-url $RPC_URL
```

```javascript
// Via Web3
async function getByRegistryId(registryId) {
    try {
        const result = await contract.methods.getContractByRegistryId(registryId).call();
        console.log('Contrato encontrado:', result);
        return result;
    } catch (error) {
        console.log('Contrato não encontrado:', error.message);
        return null;
    }
}
```

### **Exemplo 3: Buscar por Chassi**

```javascript
async function getByChassisWithDetails(chassis) {
    const result = await contract.methods.getContractByChassis(chassis).call();

    if (result) {
        const contractData = result[0];
        const vehicleData = result[1];

        // Buscar nomes das marcas e modelos
        const brandName = await contract.methods.getBrandName(vehicleData.brandId).call();
        const modelName = await contract.methods.getModelName(vehicleData.modelId).call();

        return {
            registryId: web3.utils.hexToUtf8(contractData.registryId),
            chassis: web3.utils.hexToUtf8(vehicleData.chassis),
            licensePlate: web3.utils.hexToUtf8(vehicleData.licensePlate),
            brand: brandName,
            model: modelName,
            totalValue: web3.utils.fromWei(vehicleData.totalValue, 'ether'),
            active: contractData.active
        };
    }
    return null;
}
```

### **Exemplo 4: Listar Contratos Ativos**

```javascript
async function listActiveContracts(page = 0, pageSize = 10) {
    const offset = page * pageSize;
    const activeTokenIds = await contract.methods.getActiveContracts(offset, pageSize).call();

    const contracts = [];
    for (const tokenId of activeTokenIds) {
        const contractData = await getContractByTokenId(tokenId);
        contracts.push({
            tokenId,
            ...contractData
        });
    }

    return contracts;
}

// Uso
const firstPage = await listActiveContracts(0, 5);
console.log('Primeiros 5 contratos ativos:', firstPage);
```

### **Exemplo 5: Verificar Existência**

```bash
# Verificar se contrato existe
cast call $CONTRACT_ADDRESS \
    "doesContractExist(string)" \
    "VFIN-2024-001" \
    --rpc-url $RPC_URL

# Verificar total de contratos
cast call $CONTRACT_ADDRESS \
    "totalSupply()" \
    --rpc-url $RPC_URL
```

---

## 🌐 Sistema de Metadados

### **Exemplo 1: Acessar Token URI**

```javascript
async function getTokenMetadataUri(tokenId) {
    const tokenUri = await contract.methods.tokenURI(tokenId).call();
    console.log('Token URI:', tokenUri);

    // Buscar metadados do servidor
    const response = await fetch(tokenUri);
    const metadata = await response.json();

    console.log('Metadados completos:', metadata);
    return metadata;
}
```

### **Exemplo 2: URL Direta por Hash**

```javascript
async function getMetadataByHash(metadataHash) {
    const metadataUrl = await contract.methods.getMetadataUrl(metadataHash).call();

    console.log('URL dos metadados:', metadataUrl);

    // Buscar dados
    try {
        const response = await fetch(metadataUrl);
        if (response.ok) {
            const metadata = await response.json();
            return metadata;
        }
    } catch (error) {
        console.log('Erro ao buscar metadados:', error.message);
    }
    return null;
}
```

### **Exemplo 3: Metadados por Registry ID**

```bash
# Via Cast
cast call $CONTRACT_ADDRESS \
    "getMetadataUrlByRegistryId(string)" \
    "VFIN-2024-001" \
    --rpc-url $RPC_URL
```

### **Exemplo 4: Atualizar Hash de Metadados**

```javascript
async function updateMetadataHash(tokenId, newHash) {
    // Apenas API server pode fazer isso
    const result = await contract.methods.updateMetadataHash(tokenId, newHash).send({
        from: API_SERVER_ADDRESS,
        gas: 100000
    });

    console.log('Hash atualizado:', result.events.MetadataHashUpdated.returnValues);
}
```

---

## 🏷️ Gestão de Marcas/Modelos

### **Exemplo 1: Registrar Nova Marca**

```javascript
async function registerNewBrand(brandName) {
    // Apenas owner pode fazer isso
    const result = await contract.methods.registerBrand(brandName).send({
        from: OWNER_ADDRESS,
        gas: 100000
    });

    const brandId = result.events.BrandRegistered.returnValues.brandId;
    console.log(`Marca '${brandName}' registrada com ID: ${brandId}`);
    return brandId;
}
```

### **Exemplo 2: Registrar Novo Modelo**

```bash
# Via Cast (como owner)
cast send $CONTRACT_ADDRESS \
    "registerModel(string)" \
    "NOVO MODELO 2024" \
    --private-key $OWNER_PRIVATE_KEY \
    --rpc-url $RPC_URL
```

### **Exemplo 3: Consultar Marcas Existentes**

```javascript
async function listBrands() {
    // Marcas pré-registradas são IDs 1-10
    const brands = {};

    for (let i = 1; i <= 15; i++) { // Incluir algumas extras
        try {
            const brandName = await contract.methods.getBrandName(i).call();
            if (brandName && brandName !== '') {
                brands[i] = brandName;
            }
        } catch (error) {
            // Brand ID não existe
            break;
        }
    }

    console.log('Marcas registradas:', brands);
    return brands;
}
```

### **Exemplo 4: Consultar Modelos por Marca**

```javascript
async function getModelsByBrand(brandId) {
    // Esta é uma funcionalidade que pode ser implementada
    // Atualmente precisamos iterar pelos contratos

    const activeContracts = await contract.methods.getActiveContracts(0, 100).call();
    const models = new Set();

    for (const tokenId of activeContracts) {
        const result = await contract.methods.getContract(tokenId).call();
        const vehicleData = result[1];

        if (vehicleData.brandId == brandId) {
            const modelName = await contract.methods.getModelName(vehicleData.modelId).call();
            models.add(modelName);
        }
    }

    return Array.from(models);
}
```

---

## ⚙️ Administração

### **Exemplo 1: Atualizar Configurações**

```javascript
async function updateServerConfig(newMetadataUrl, newApiServer) {
    // Apenas owner
    const result = await contract.methods.updateServerConfig(
        newMetadataUrl,
        newApiServer
    ).send({
        from: OWNER_ADDRESS,
        gas: 150000
    });

    console.log('Configurações atualizadas:', result.events.ServerConfigUpdated.returnValues);
}
```

### **Exemplo 2: Atualizar Status de Contrato**

```bash
# Desativar contrato
cast send $CONTRACT_ADDRESS \
    "updateStatus(uint256,bool)" \
    1 \
    false \
    --private-key $API_SERVER_PRIVATE_KEY \
    --rpc-url $RPC_URL
```

### **Exemplo 3: Upgrade do Contrato**

```solidity
// Script de upgrade
contract UpgradeContract is Script {
    function run() external {
        vm.startBroadcast();

        // Deploy nova implementação
        VFinanceRegistryV3 newImplementation = new VFinanceRegistryV3();

        // Fazer upgrade
        VFinanceRegistry registry = VFinanceRegistry(payable(CONTRACT_ADDRESS));
        registry.upgradeToAndCall(address(newImplementation), "");

        console.log("Upgrade realizado para:", address(newImplementation));

        vm.stopBroadcast();
    }
}
```

### **Exemplo 4: Monitoramento de Eventos**

```javascript
// Monitorar eventos em tempo real
function monitorContractEvents() {
    // Evento de registro
    contract.events.ContractRegistered({
        fromBlock: 'latest'
    }, (error, event) => {
        if (error) {
            console.log('Erro:', error);
            return;
        }

        console.log('Novo contrato registrado:');
        console.log('Token ID:', event.returnValues.tokenId);
        console.log('Registry ID Hash:', event.returnValues.registryIdHash);
        console.log('Timestamp:', event.returnValues.timestamp);
    });

    // Evento de atualização
    contract.events.MetadataHashUpdated({
        fromBlock: 'latest'
    }, (error, event) => {
        if (error) return;

        console.log('Metadata atualizada:');
        console.log('Token ID:', event.returnValues.tokenId);
        console.log('Novo Hash:', event.returnValues.newHash);
    });
}
```

---

## 🌐 Integração Web3

### **Exemplo 1: Frontend React**

```jsx
import Web3 from 'web3';
import { useState, useEffect } from 'react';

function VFinanceContract() {
    const [web3, setWeb3] = useState(null);
    const [contract, setContract] = useState(null);
    const [account, setAccount] = useState('');

    useEffect(() => {
        initWeb3();
    }, []);

    async function initWeb3() {
        if (window.ethereum) {
            const web3Instance = new Web3(window.ethereum);
            await window.ethereum.request({ method: 'eth_requestAccounts' });

            const accounts = await web3Instance.eth.getAccounts();
            const contractInstance = new web3Instance.eth.Contract(ABI, CONTRACT_ADDRESS);

            setWeb3(web3Instance);
            setContract(contractInstance);
            setAccount(accounts[0]);
        }
    }

    async function registerContract(contractData) {
        if (!contract) return;

        try {
            const result = await contract.methods.registerContract(
                contractData.registryId,
                contractData.contractNumber,
                Math.floor(Date.now() / 1000),
                contractData.chassis,
                contractData.licensePlate,
                web3.utils.toWei(contractData.totalValue, 'ether'),
                contractData.brand,
                contractData.model
            ).send({ from: account });

            console.log('Contrato registrado:', result);
            return result;
        } catch (error) {
            console.error('Erro ao registrar:', error);
            throw error;
        }
    }

    return (
        <div>
            <h2>VFinance Registry</h2>
            <p>Conta conectada: {account}</p>
            {/* Formulário de registro aqui */}
        </div>
    );
}
```

### **Exemplo 2: Backend Node.js**

```javascript
const Web3 = require('web3');
const express = require('express');

class VFinanceService {
    constructor() {
        this.web3 = new Web3(process.env.RPC_URL);
        this.contract = new this.web3.eth.Contract(ABI, process.env.CONTRACT_ADDRESS);
        this.account = this.web3.eth.accounts.privateKeyToAccount(process.env.PRIVATE_KEY);
        this.web3.eth.accounts.wallet.add(this.account);
    }

    async registerContract(contractData) {
        const result = await this.contract.methods.registerContract(
            contractData.registryId,
            contractData.contractNumber,
            contractData.contractDate,
            contractData.chassis,
            contractData.licensePlate,
            this.web3.utils.toWei(contractData.totalValue.toString(), 'ether'),
            contractData.brand,
            contractData.model
        ).send({
            from: this.account.address,
            gas: 500000
        });

        return {
            tokenId: result.events.ContractRegistered.returnValues.tokenId,
            metadataHash: result.events.ContractRegistered.returnValues.metadataHash,
            transactionHash: result.transactionHash
        };
    }

    async getContract(tokenId) {
        const result = await this.contract.methods.getContract(tokenId).call();

        return {
            contractData: result[0],
            vehicleData: result[1]
        };
    }
}

// API Express
const app = express();
const vfinanceService = new VFinanceService();

app.post('/api/contracts/register', async (req, res) => {
    try {
        const result = await vfinanceService.registerContract(req.body);
        res.json({ success: true, data: result });
    } catch (error) {
        res.status(500).json({ success: false, error: error.message });
    }
});

app.get('/api/contracts/:tokenId', async (req, res) => {
    try {
        const result = await vfinanceService.getContract(req.params.tokenId);
        res.json({ success: true, data: result });
    } catch (error) {
        res.status(500).json({ success: false, error: error.message });
    }
});
```

---

## 🔨 Scripts Forge

### **Exemplo 1: Script de Deploy Completo**

```bash
# Deploy + Configuração + Registro de teste
forge script script/FullSetup.s.sol --broadcast --rpc-url $RPC_URL --verify
```

### **Exemplo 2: Script de Verificação**

```bash
# Verificar integridade dos dados
forge script script/VerifyData.s.sol --rpc-url $RPC_URL
```

### **Exemplo 3: Script de Migração**

```bash
# Migrar dados de contrato antigo para novo
forge script script/MigrateData.s.sol --broadcast --rpc-url $RPC_URL
```

### **Exemplo 4: Script de Backup**

```bash
# Exportar todos os dados para arquivo JSON
forge script script/ExportData.s.sol --rpc-url $RPC_URL > backup.json
```

---

## 📊 Exemplos de Análise

### **Exemplo 1: Relatório de Contratos**

```javascript
async function generateReport() {
    const totalSupply = await contract.methods.totalSupply().call();
    const activeContracts = await contract.methods.getActiveContracts(0, totalSupply).call();

    let totalValue = 0;
    const brandStats = {};

    for (const tokenId of activeContracts) {
        const result = await contract.methods.getContract(tokenId).call();
        const vehicleData = result[1];

        totalValue += parseInt(vehicleData.totalValue);

        const brandName = await contract.methods.getBrandName(vehicleData.brandId).call();
        brandStats[brandName] = (brandStats[brandName] || 0) + 1;
    }

    console.log('=== RELATÓRIO VFINANCE ===');
    console.log('Total de contratos:', totalSupply);
    console.log('Contratos ativos:', activeContracts.length);
    console.log('Valor total:', web3.utils.fromWei(totalValue.toString(), 'ether'), 'ETH');
    console.log('Distribuição por marca:', brandStats);
}
```

### **Exemplo 2: Auditoria de Eventos**

```javascript
async function auditEvents(fromBlock = 0) {
    const events = await contract.getPastEvents('allEvents', {
        fromBlock,
        toBlock: 'latest'
    });

    console.log('=== AUDITORIA DE EVENTOS ===');
    console.log('Total de eventos:', events.length);

    const eventTypes = {};
    events.forEach(event => {
        eventTypes[event.event] = (eventTypes[event.event] || 0) + 1;
    });

    console.log('Por tipo:', eventTypes);

    return events;
}
```

---

## 🚨 Tratamento de Erros

### **Exemplo 1: Erros Comuns**

```javascript
async function safeRegisterContract(contractData) {
    try {
        const result = await contract.methods.registerContract(...contractData).send({
            from: account,
            gas: 500000
        });
        return { success: true, data: result };
    } catch (error) {
        if (error.message.includes('Contract already exists')) {
            return { success: false, error: 'DUPLICATE_REGISTRY_ID' };
        } else if (error.message.includes('Chassis already registered')) {
            return { success: false, error: 'DUPLICATE_CHASSIS' };
        } else if (error.message.includes('Only API server')) {
            return { success: false, error: 'UNAUTHORIZED' };
        } else if (error.message.includes('gas')) {
            return { success: false, error: 'INSUFFICIENT_GAS' };
        } else {
            return { success: false, error: 'UNKNOWN_ERROR', details: error.message };
        }
    }
}
```

### **Exemplo 2: Retry Logic**

```javascript
async function registerWithRetry(contractData, maxRetries = 3) {
    for (let attempt = 1; attempt <= maxRetries; attempt++) {
        try {
            const result = await safeRegisterContract(contractData);
            if (result.success) {
                return result;
            } else if (result.error === 'DUPLICATE_REGISTRY_ID' || result.error === 'DUPLICATE_CHASSIS') {
                // Não tentar novamente para erros de duplicação
                return result;
            }
        } catch (error) {
            if (attempt === maxRetries) {
                throw error;
            }
            console.log(`Tentativa ${attempt} falhou, tentando novamente...`);
            await new Promise(resolve => setTimeout(resolve, 1000 * attempt));
        }
    }
}
```

---

## 📝 Próximos Passos

1. **Integração com API Go**: Conectar estes exemplos com o backend
2. **Frontend Dashboard**: Criar interface para gestão
3. **Monitoramento**: Implementar alerts e métricas
4. **Backup/Restore**: Scripts para backup dos dados
5. **Analytics**: Dashboards de análise financeira

---

Para mais exemplos, consulte:
- [README.md](README.md) - Documentação completa
- [Testes](test/VFinanceRegistryTest.t.sol) - Casos de uso nos testes
- [Scripts](script/) - Scripts de exemplo prontos


