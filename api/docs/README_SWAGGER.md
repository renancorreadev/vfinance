# 📚 MobX API - Documentação Swagger

Este documento explica como usar e configurar a documentação Swagger da MobX API.

## 🚀 Acesso à Documentação

### URLs Disponíveis:
- **Swagger UI**: `http://localhost:3000/swagger/index.html`
- **Redirect**: `http://localhost:3000/docs` (redireciona para Swagger UI)
- **OpenAPI JSON**: `http://localhost:3000/swagger/doc.json`
- **OpenAPI YAML**: Disponível em `/docs/swagger.yaml`

## 🔧 Configuração

### 1. Dependências Instaladas:
```go
require (
    github.com/swaggo/gin-swagger v1.6.0
    github.com/swaggo/swag v1.16.2
)
```

### 2. Arquivos Criados:
- `docs/swagger.yaml` - Especificação OpenAPI 3.0.3 completa
- `docs/docs.go` - Configuração Go para integração com gin-swagger
- `docs/index.html` - Interface customizada do Swagger UI
- `docs/README_SWAGGER.md` - Este arquivo de documentação

### 3. Integração no Servidor:
O Swagger foi integrado no arquivo `internal/server/server.go`:
```go
import (
    "vfinance-api/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

// Rotas do Swagger
s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
s.router.GET("/docs", func(c *gin.Context) {
    c.Redirect(301, "/swagger/index.html")
})
```

## 📖 Funcionalidades da Documentação

### 🔐 Autenticação
A documentação inclui suporte completo para autenticação JWT:
1. **Obter Token**: Use o endpoint `POST /api/auth/token`
2. **Autorizar**: Clique no botão "Authorize" no Swagger UI
3. **Inserir Token**: Format: `Bearer {seu_token_jwt}`

### 📊 Endpoints Documentados

#### **Autenticação (`/api/auth`)**
- `POST /token` - Gerar token JWT
- `GET /validate` - Validar token JWT

#### **Contratos (`/api/contracts`)**
- `POST /` - Registrar novo contrato (🔒 JWT obrigatório)
- `GET /{regConId}` - Buscar por Registry ID
- `GET /token/{tokenId}` - Buscar por Token ID (🆕 NOVO)
- `GET /chassis/{chassis}` - Buscar por chassi (🆕 NOVO)
- `GET /active` - Listar contratos ativos com paginação
- `GET /hash/{hash}` - Buscar por hash dos metadados
- `GET /stats` - Estatísticas do sistema
- `GET /metadata-url/{hash}` - URL dos metadados por hash (🆕 NOVO)
- `GET /metadata-url/registry/{registryId}` - URL por Registry ID (🆕 NOVO)

#### **Metadados (`/api/metadata`)**
- `POST /{hash}` - Armazenar metadados (🔒 JWT obrigatório)
- `GET /{hash}` - Buscar metadados
- `PUT /{hash}` - Atualizar metadados (🔒 JWT obrigatório)
- `DELETE /{hash}` - Remover metadados (🔒 JWT obrigatório)

#### **Sistema (`/`)**
- `GET /health` - Health check da API

### 🎯 Recursos Avançados

#### **1. Modelos de Dados Completos**
Todos os schemas estão definidos com:
- Tipos de dados corretos
- Descrições detalhadas
- Exemplos práticos
- Campos obrigatórios marcados

#### **2. Exemplos Interativos**
- **Try it out**: Teste direto na interface
- **Curl commands**: Comandos prontos para copiar
- **Request/Response samples**: Exemplos reais

#### **3. Validações**
- Parâmetros obrigatórios
- Formatos de dados (UUID, datas, etc.)
- Limites de paginação

#### **4. Códigos de Resposta**
Todas as respostas possíveis documentadas:
- `200` - Sucesso
- `201` - Criado
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `409` - Conflict
- `500` - Internal Server Error

## 🎨 Interface Customizada

### Características da UI:
- **Design Responsivo**: Funciona em desktop e mobile
- **Tema Personalizado**: Cores azuis corporativas
- **Header Customizado**: Informações da API em destaque
- **Quick Start Guide**: Guia rápido na página inicial
- **Features Cards**: Destaque das principais funcionalidades

### Features em Destaque:
- 🔗 **Blockchain**: Hyperledger Besu com contratos UUPS
- 🎫 **ERC721**: Sistema de registry não-transferível
- 🔐 **Segurança**: JWT e rate limiting
- 📊 **Rastreabilidade**: Múltiplas formas de busca

## 🚦 Como Usar (Passo a Passo)

### 1. **Iniciar a API**
```bash
cd api
go run cmd/main.go
```

### 2. **Acessar Documentação**
Abra o navegador em: `http://localhost:3000/docs`

### 3. **Autenticar**
1. Clique em `POST /api/auth/token`
2. Clique "Try it out"
3. Execute a requisição
4. Copie o token retornado
5. Clique no botão "Authorize" (🔒 no topo)
6. Cole o token no formato: `Bearer {token}`
7. Clique "Authorize"

### 4. **Testar Endpoints**
Agora você pode testar todos os endpoints protegidos!

## 🔄 Atualizações Automáticas

### Regenerar Documentação:
Se você adicionar novos endpoints ou modificar existentes:

```bash
# Instalar swag CLI (se não tiver)
go install github.com/swaggo/swag/cmd/swag@latest

# Regenerar documentação
swag init -g cmd/main.go -o docs/
```

### Arquivos Afetados:
- `docs/docs.go` - Código Go atualizado
- `docs/swagger.json` - Especificação JSON atualizada
- `docs/swagger.yaml` - Especificação YAML (manual)

## 📝 Notas Importantes

### Versionamento:
- **Versão Atual**: 2.0.0
- **OpenAPI Version**: 3.0.3
- **Compatibilidade**: Todos os endpoints da API v2

### Configuração de Produção:
Para produção, altere no `server.go`:
```go
docs.SwaggerInfo.Host = "seu-dominio.com"
docs.SwaggerInfo.Schemes = []string{"https"}
```

### Segurança:
- Em produção, considere desabilitar o Swagger ou restringir acesso
- Use HTTPS sempre que possível
- Tokens JWT têm expiração configurável

## 🐛 Troubleshooting

### Problemas Comuns:

1. **Swagger não carrega**
   - Verifique se as dependências foram instaladas: `go mod tidy`
   - Certifique-se que a porta está correta

2. **Endpoint não aparece**
   - Verifique se está em `docs.go`
   - Regenere com `swag init`

3. **Autorização não funciona**
   - Formato correto: `Bearer {token}`
   - Token deve estar válido (não expirado)

4. **CORS errors**
   - Configure CORS no servidor se acessando de domínio diferente

### Logs Úteis:
O Swagger UI inclui logs no console do navegador para debug.

---

**🎉 Documentação completa e funcional criada com sucesso!**

A MobX API agora possui uma documentação Swagger profissional, interativa e fácil de usar.
