# Exemplos de Uso da API - VFinance Registry

## 📋 Índice
1. [Autenticação](#autenticação)
2. [Registro de Contratos](#registro-de-contratos)
3. [Consultas](#consultas)
4. [Gerenciamento de Metadados](#gerenciamento-de-metadados)
5. [Casos de Uso Completos](#casos-de-uso-completos)
6. [Códigos de Erro](#códigos-de-erro)

---

## 🔐 Autenticação

### Gerar Token JWT

```bash
curl -X POST http://localhost:3000/api/auth/token \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin"
  }'
```

**Resposta:**
```json
{
  "success": true,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_in": 86400
}
```

### Validar Token

```bash
curl -X GET http://localhost:3000/api/auth/validate \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Resposta:**
```json
{
  "success": true,
  "user_id": "admin",
  "role": "admin",
  "exp": 1705420800
}
```

---

## 📝 Registro de Contratos

### Exemplo Completo de Registro

```bash
curl -X POST http://localhost:3000/api/contracts/ \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "regConId": "CONT-2024-001",
    "numeroContrato": "FIN123456789",
    "dataContrato": "2024-01-15",
    "vehicleData": {
      "regConId": "CONT-2024-001",
      "numeroContrato": "FIN123456789",
      "dataContrato": "2024-01-15",
      "cnpjAgenteFinanceiro": "12.345.678/0001-90",
      "nomeAgenteFinanceiro": "Banco Exemplo S.A.",
      "enderecoAgenteFinanceiro": "Av. Paulista, 1000",
      "numeroEnderecoAgenteFinanceiro": "1000",
      "complementoEnderecoAgenteFinanceiro": "Andar 15",
      "bairroEnderecoAgenteFinanceiro": "Bela Vista",
      "nomeMunicipioEnderecoAgenteFinanceiro": "São Paulo",
      "ufEnderecoAgenteFinanceiro": "SP",
      "cepEnderecoAgenteFinanceiro": "01310-100",
      "telefoneAgenteFinanceiro": "(11) 3000-0000",
      "emailAgenteFinanceiro": "contato@bancoexemplo.com.br",
      "cpfCnpjProprietario": "123.456.789-00",
      "nomeProprietario": "João Silva Santos",
      "enderecoProprietario": "Rua das Flores, 123",
      "numeroEnderecoProprietario": "123",
      "bairroEnderecoProprietario": "Jardins",
      "nomeMunicipioProprietario": "São Paulo",
      "ufEnderecoProprietario": "SP",
      "cepEnderecoProprietario": "01401-000",
      "telefoneProprietario": "(11) 99999-9999",
      "emailProprietario": "joao.silva@email.com",
      "veiculoZeroKm": false,
      "chassiVeiculo": "9BWZZZ377VT004251",
      "chassiRemarcadoVeiculo": "",
      "placaVeiculo": "ABC-1234",
      "tipoPlacaVeiculo": "MERCOSUL",
      "ufPlacaVeiculo": "SP",
      "renavamVeiculo": "12345678901",
      "anoFabricacaoVeiculo": "2022",
      "anoModeloVeiculo": "2023",
      "numeroRestricaoVeiculo": "REST001",
      "especieVeiculo": "PASSAGEIRO",
      "marcaVeiculo": "TOYOTA",
      "modeloVeiculo": "COROLLA",
      "tipoRestricacaoContrato": "ALIENACAO_FIDUCIARIA",
      "ufRegistroContrato": "SP",
      "cnpjResponsavelPeloRegistro": "12.345.678/0001-90",
      "valorTotalContrato": "45000.00",
      "valorParcelaContrato": "1250.00",
      "quantidadeParcelasContrato": "36",
      "taxaJurosMesContrato": "1.99",
      "taxaJurosMesAnoContrato": "23.88",
      "possuiJurosMultaContrato": "SIM",
      "taxaJurosMultaContrato": "2.00",
      "possuiJurosMoraDiaContrato": "SIM",
      "taxaJurosMoraDiaContrato": "0.033",
      "valorCustoRegistroContrato": "150.00",
      "valorIofContrato": "450.00",
      "dataVencimentoPrimeiraParcelaContrato": "2024-02-15",
      "dataVencimentoUltimaParcelaContrato": "2027-02-15",
      "dataLiberacaoCreditoContrato": "2024-01-15",
      "cidadeLiberacaoCreditoContrato": "São Paulo",
      "ufLiberacaoCreditoContrato": "SP",
      "indiceCorrecaoContrato": "IPCA",
      "numeroGrupoConsorcioContrato": "",
      "numeroCotaConsorcioContrato": "",
      "indicativoPenalidadeContrato": "SIM",
      "penalidadeContrato": "20.00",
      "indicativoComissaoContrato": "SIM",
      "comissaoContrato": "5.00",
      "categoriaVeiculo": "AUTOMOVEL"
    }
  }'
```

**Resposta de Sucesso:**
```json
{
  "success": true,
  "message": "Contrato registrado com sucesso",
  "regConId": "CONT-2024-001",
  "metadataHash": "0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
  "txHash": "0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890"
}
```

### Exemplo Mínimo (Campos Obrigatórios)

```bash
curl -X POST http://localhost:3000/api/contracts/ \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "regConId": "CONT-2024-002",
    "numeroContrato": "FIN987654321",
    "dataContrato": "2024-01-16",
    "vehicleData": {
      "cnpjAgenteFinanceiro": "98.765.432/0001-10",
      "nomeAgenteFinanceiro": "Banco Minimal LTDA",
      "cpfCnpjProprietario": "987.654.321-00",
      "nomeProprietario": "Maria Oliveira",
      "chassiVeiculo": "9BWZZZ377VT004252",
      "placaVeiculo": "XYZ-5678",
      "valorTotalContrato": "25000.00"
    }
  }'
```

---

## 🔍 Consultas

### Buscar Contrato por ID

```bash
curl -X GET http://localhost:3000/api/contracts/CONT-2024-001
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
      "metadataHash": "0x1234567890abcdef...",
      "timestamp": 1705334400,
      "registeredBy": "0x627306090abaB3A6e1400e9345bC60c78a8BEf57",
      "active": true
    },
    "offChain": {
      "cnpjAgenteFinanceiro": "12.345.678/0001-90",
      "nomeAgenteFinanceiro": "Banco Exemplo S.A.",
      "cpfCnpjProprietario": "123.456.789-00",
      "nomeProprietario": "João Silva Santos",
      "chassiVeiculo": "9BWZZZ377VT004251",
      "placaVeiculo": "ABC-1234",
      "valorTotalContrato": "45000.00"
      // ... todos os outros campos
    }
  }
}
```

### Listar Contratos Ativos

```bash
# Primeira página (10 registros)
curl -X GET "http://localhost:3000/api/contracts/active?offset=0&limit=10"

# Segunda página
curl -X GET "http://localhost:3000/api/contracts/active?offset=10&limit=10"

# Limite máximo (100 registros)
curl -X GET "http://localhost:3000/api/contracts/active?offset=0&limit=100"
```

**Resposta:**
```json
{
  "success": true,
  "data": [
    "CONT-2024-001",
    "CONT-2024-002",
    "CONT-2024-003",
    "CONT-2024-004",
    "CONT-2024-005"
  ]
}
```

### Buscar por Hash de Metadados

```bash
curl -X GET http://localhost:3000/api/contracts/hash/0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
```

### Estatísticas Gerais

```bash
curl -X GET http://localhost:3000/api/contracts/stats
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

---

## 📊 Gerenciamento de Metadados

### Buscar Metadados por Hash

```bash
curl -X GET http://localhost:3000/api/metadata/0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
```

**Resposta:**
```json
{
  "success": true,
  "data": {
    "cnpjAgenteFinanceiro": "12.345.678/0001-90",
    "nomeAgenteFinanceiro": "Banco Exemplo S.A.",
    "enderecoAgenteFinanceiro": "Av. Paulista, 1000",
    "cpfCnpjProprietario": "123.456.789-00",
    "nomeProprietario": "João Silva Santos",
    "chassiVeiculo": "9BWZZZ377VT004251",
    "placaVeiculo": "ABC-1234",
    "valorTotalContrato": "45000.00"
    // ... todos os campos detalhados
  }
}
```

### Atualizar Metadados (Requer Autenticação)

```bash
curl -X PUT http://localhost:3000/api/metadata/0x1234567890abcdef... \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -H "Content-Type: application/json" \
  -d '{
    "cnpjAgenteFinanceiro": "12.345.678/0001-90",
    "nomeAgenteFinanceiro": "Banco Exemplo S.A. - Filial Centro",
    "telefoneAgenteFinanceiro": "(11) 3000-0001",
    "cpfCnpjProprietario": "123.456.789-00",
    "nomeProprietario": "João Silva Santos",
    "telefoneProprietario": "(11) 88888-8888",
    "chassiVeiculo": "9BWZZZ377VT004251",
    "placaVeiculo": "ABC-1234",
    "valorTotalContrato": "44500.00"
  }'
```

---

## 💼 Casos de Uso Completos

### Caso 1: Registro e Consulta Completa

```bash
# 1. Obter token
TOKEN=$(curl -s -X POST http://localhost:3000/api/auth/token \
  -H "Content-Type: application/json" \
  -d '{}' | jq -r '.token')

# 2. Registrar contrato
RESPONSE=$(curl -s -X POST http://localhost:3000/api/contracts/ \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "regConId": "CONT-2024-EXEMPLO",
    "numeroContrato": "FIN999888777",
    "dataContrato": "2024-01-20",
    "vehicleData": {
      "cnpjAgenteFinanceiro": "11.222.333/0001-44",
      "nomeAgenteFinanceiro": "Financeira Exemplo",
      "cpfCnpjProprietario": "111.222.333-44",
      "nomeProprietario": "Pedro Exemplo",
      "chassiVeiculo": "9BWZZZ377VT999888",
      "placaVeiculo": "EXE-1234",
      "valorTotalContrato": "35000.00"
    }
  }')

echo "Registro: $RESPONSE"

# 3. Extrair dados da resposta
REG_CON_ID=$(echo $RESPONSE | jq -r '.regConId')
METADATA_HASH=$(echo $RESPONSE | jq -r '.metadataHash')
TX_HASH=$(echo $RESPONSE | jq -r '.txHash')

# 4. Aguardar confirmação (opcional)
sleep 5

# 5. Consultar contrato registrado
curl -s -X GET "http://localhost:3000/api/contracts/$REG_CON_ID" | jq '.'

# 6. Consultar por hash
curl -s -X GET "http://localhost:3000/api/contracts/hash/$METADATA_HASH" | jq '.'

# 7. Consultar metadados diretamente
curl -s -X GET "http://localhost:3000/api/metadata/$METADATA_HASH" | jq '.'
```

### Caso 2: Validação de Integridade

```bash
# 1. Buscar contrato
CONTRACT_DATA=$(curl -s -X GET http://localhost:3000/api/contracts/CONT-2024-001)

# 2. Extrair hash
HASH=$(echo $CONTRACT_DATA | jq -r '.data.onChain.metadataHash')

# 3. Buscar metadados
METADATA=$(curl -s -X GET "http://localhost:3000/api/metadata/$HASH")

# 4. Verificar consistência
echo "Hash do contrato: $HASH"
echo "Metadados encontrados: $(echo $METADATA | jq '.success')"

# 5. Validar campos específicos
CHASSI_BLOCKCHAIN=$(echo $CONTRACT_DATA | jq -r '.data.offChain.chassiVeiculo')
CHASSI_METADATA=$(echo $METADATA | jq -r '.data.chassiVeiculo')

if [ "$CHASSI_BLOCKCHAIN" = "$CHASSI_METADATA" ]; then
  echo "✅ Integridade verificada - Chassi: $CHASSI_BLOCKCHAIN"
else
  echo "❌ Inconsistência detectada!"
fi
```

### Caso 3: Busca e Paginação

```bash
# 1. Obter estatísticas
STATS=$(curl -s -X GET http://localhost:3000/api/contracts/stats)
TOTAL=$(echo $STATS | jq '.data.totalContracts')

echo "Total de contratos: $TOTAL"

# 2. Paginar através de todos os contratos ativos
OFFSET=0
LIMIT=10

while true; do
  CONTRACTS=$(curl -s -X GET "http://localhost:3000/api/contracts/active?offset=$OFFSET&limit=$LIMIT")

  # Verificar se retornou dados
  COUNT=$(echo $CONTRACTS | jq '.data | length')

  if [ $COUNT -eq 0 ]; then
    echo "Fim da paginação"
    break
  fi

  echo "Página $(($OFFSET / $LIMIT + 1)): $COUNT contratos"
  echo $CONTRACTS | jq '.data[]'

  OFFSET=$((OFFSET + LIMIT))
done
```

---

## ❌ Códigos de Erro

### Erros de Autenticação

```json
// 401 - Token ausente
{
  "error": "Token de autorização necessário"
}

// 401 - Token inválido
{
  "error": "Token inválido"
}
```

### Erros de Validação

```json
// 400 - Dados obrigatórios ausentes
{
  "error": "regConId, numeroContrato e dataContrato são obrigatórios"
}

// 400 - Parâmetros inválidos
{
  "error": "Hash é obrigatório"
}

// 400 - Formato JSON inválido
{
  "error": "invalid character '}' looking for beginning of object key string"
}
```

### Erros de Negócio

```json
// 409 - Contrato já existe
{
  "error": "Contrato já registrado na blockchain"
}

// 404 - Contrato não encontrado
{
  "error": "Contrato não encontrado"
}

// 404 - Metadados não encontrados
{
  "error": "Metadados não encontrados"
}
```

### Erros de Sistema

```json
// 500 - Erro de blockchain
{
  "error": "Erro ao conectar com blockchain: connection refused"
}

// 500 - Erro de banco de dados
{
  "error": "Erro ao salvar no banco de dados"
}

// 503 - Serviço indisponível
{
  "error": "Blockchain temporariamente indisponível"
}
```

---

## 🧪 Scripts de Teste

### Teste de Carga Básico

```bash
#!/bin/bash
# test_load.sh

TOKEN=$(curl -s -X POST http://localhost:3000/api/auth/token -d '{}' | jq -r '.token')

for i in {1..100}; do
  curl -s -X POST http://localhost:3000/api/contracts/ \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d "{
      \"regConId\": \"LOAD-TEST-$i\",
      \"numeroContrato\": \"LOAD$i\",
      \"dataContrato\": \"2024-01-$(printf %02d $((i % 28 + 1)))\",
      \"vehicleData\": {
        \"cnpjAgenteFinanceiro\": \"00.000.000/0001-$i\",
        \"nomeAgenteFinanceiro\": \"Banco Teste $i\",
        \"cpfCnpjProprietario\": \"000.000.00$i-00\",
        \"nomeProprietario\": \"Teste $i\",
        \"chassiVeiculo\": \"TEST$i\",
        \"placaVeiculo\": \"TST-$i\",
        \"valorTotalContrato\": \"$((i * 1000)).00\"
      }
    }" &

  # Limitar concorrência
  if [ $((i % 10)) -eq 0 ]; then
    wait
  fi
done

wait
echo "Teste de carga concluído!"
```

### Validação de Integridade em Lote

```bash
#!/bin/bash
# validate_integrity.sh

# Buscar todos os contratos ativos
CONTRACTS=$(curl -s "http://localhost:3000/api/contracts/active?offset=0&limit=100" | jq -r '.data[]')

for CONTRACT_ID in $CONTRACTS; do
  echo "Validando: $CONTRACT_ID"

  # Buscar dados completos
  FULL_DATA=$(curl -s "http://localhost:3000/api/contracts/$CONTRACT_ID")

  # Verificar se dados on-chain e off-chain estão consistentes
  ON_CHAIN_HASH=$(echo $FULL_DATA | jq -r '.data.onChain.metadataHash')

  # Buscar metadados pelo hash
  METADATA=$(curl -s "http://localhost:3000/api/metadata/$ON_CHAIN_HASH")

  if [ "$(echo $METADATA | jq -r '.success')" = "true" ]; then
    echo "✅ $CONTRACT_ID: Integridade OK"
  else
    echo "❌ $CONTRACT_ID: Falha na integridade!"
  fi
done
```

---

**📚 Documentação Completa**: Ver `ARQUITETURA_COMPLETA.md`
**🚀 Instalação Rápida**: Ver `GUIA_INSTALACAO_RAPIDA.md`
**🔧 Atualizado**: Janeiro 2024
