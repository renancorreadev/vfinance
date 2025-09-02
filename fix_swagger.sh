#!/bin/bash

echo "🔧 Corrigindo todos os schemas do Swagger..."

# Corrigir schemas complexos para tipos simples que funcionam
echo "📝 Corrigindo schemas de resposta..."

# Blockchain Handler
sed -i 's/object{success=boolean,data=object{[^}]*}}/map[string]interface{}/g' internal/handlers/blockchain_handler.go
sed -i 's/object{error=string}/map[string]interface{}/g' internal/handlers/blockchain_handler.go
sed -i 's/array{object{[^}]*}}/[]map[string]interface{}/g' internal/handlers/blockchain_handler.go
sed -i 's/pagination=object{[^}]*}/pagination=map[string]interface{}/g' internal/handlers/blockchain_handler.go
sed -i 's/vehicleData=object{[^}]*}/vehicleData=map[string]interface{}/g' internal/handlers/blockchain_handler.go

# Contract Handler
sed -i 's/object{success=boolean,data=object{[^}]*}}/map[string]interface{}/g' internal/handlers/contract_handler.go
sed -i 's/object{error=string}/map[string]interface{}/g' internal/handlers/contract_handler.go
sed -i 's/array{object{[^}]*}}/[]map[string]interface{}/g' internal/handlers/contract_handler.go
sed -i 's/pagination=object{[^}]*}/pagination=map[string]interface{}/g' internal/handlers/contract_handler.go
sed -i 's/vehicleData=object{[^}]*}/vehicleData=map[string]interface{}/g' internal/handlers/contract_handler.go

# Metadata Handler
sed -i 's/object{success=boolean,data=object{[^}]*}}/map[string]interface{}/g' internal/handlers/metadata_handler.go
sed -i 's/object{error=string}/map[string]interface{}/g' internal/handlers/metadata_handler.go
sed -i 's/vehicleData=object{[^}]*}/vehicleData=map[string]interface{}/g' internal/handlers/metadata_handler.go

# Auth Handler
sed -i 's/object{token=string,success=boolean,message=string}/map[string]interface{}/g' internal/handlers/auth_handler.go
sed -i 's/object{valid=boolean,success=boolean,message=string,user_id=string,role=string,exp=integer}/map[string]interface{}/g' internal/handlers/auth_handler.go
sed -i 's/object{valid=boolean,success=boolean,message=string,error=string}/map[string]interface{}/g' internal/handlers/auth_handler.go

# Corrigir body parameters
echo "📝 Corrigindo body parameters..."
sed -i 's/body map\[string\]interface{} true "Dados do contrato"/body object true "Dados do contrato"/g' internal/handlers/contract_handler.go
sed -i 's/body map\[string\]interface{} true "Dados do veículo"/body object true "Dados do veículo"/g' internal/handlers/metadata_handler.go

echo "✅ Todos os schemas foram corrigidos!"
echo "🚀 Agora você pode regenerar o Swagger com: swag init -g cmd/main.go"
