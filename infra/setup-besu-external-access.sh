#!/bin/bash

echo "🔧 Configurando acesso externo para Besu (HTTP + WebSocket)..."

# Verificar se o nginx está instalado
if ! command -v nginx &> /dev/null; then
    echo "❌ Nginx não está instalado. Instalando..."
    sudo apt update && sudo apt install -y nginx
fi

# Criar configuração do nginx
echo "📝 Criando configuração do nginx..."
cat > besu-nginx.conf << 'EOF'
server {
    listen 80;
    server_name 144.22.179.183;

    # Configuração para RPC HTTP do Besu
    location / {
        proxy_pass http://127.0.0.1:8545;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Configurações específicas para JSON-RPC
        proxy_http_version 1.1;
        proxy_set_header Connection "";
        proxy_buffering off;
        proxy_read_timeout 300s;
        proxy_connect_timeout 75s;

        # Headers para CORS
        add_header Access-Control-Allow-Origin *;
        add_header Access-Control-Allow-Methods "GET, POST, OPTIONS";
        add_header Access-Control-Allow-Headers "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization";

        # Tratamento de OPTIONS para CORS
        if ($request_method = 'OPTIONS') {
            add_header Access-Control-Allow-Origin *;
            add_header Access-Control-Allow-Methods "GET, POST, OPTIONS";
            add_header Access-Control-Allow-Headers "DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization";
            add_header Access-Control-Max-Age 1728000;
            add_header Content-Type 'text/plain; charset=utf-8';
            add_header Content-Length 0;
            return 204;
        }
    }

    # Configuração para WebSocket
    location /ws {
        proxy_pass http://127.0.0.1:6174;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # Timeouts para WebSocket
        proxy_read_timeout 3600s;
        proxy_connect_timeout 75s;
        proxy_send_timeout 3600s;
    }
}
EOF

# Instalar configuração
echo "📦 Instalando configuração do nginx..."
sudo cp besu-nginx.conf /etc/nginx/sites-available/besu
sudo ln -sf /etc/nginx/sites-available/besu /etc/nginx/sites-enabled/

# Testar configuração
echo "🧪 Testando configuração do nginx..."
sudo nginx -t

if [ $? -eq 0 ]; then
    echo "✅ Configuração OK! Recarregando nginx..."
    sudo systemctl reload nginx

    # Verificar se o Node-1 está rodando
    if sudo systemctl is-active --quiet besu-node-1.service; then
        echo "✅ Node-1 está rodando"
    else
        echo "⚠️  Node-1 não está rodando. Iniciando..."
        sudo systemctl start besu-node-1.service
        sleep 10
    fi

    # Verificar portas
    echo "🔍 Verificando portas..."
    if netstat -tlnp | grep -q ":8545"; then
        echo "✅ Porta 8545 (HTTP) está escutando"
    else
        echo "❌ Porta 8545 não está escutando"
    fi

    if netstat -tlnp | grep -q ":6174"; then
        echo "✅ Porta 6174 (WebSocket) está escutando"
    else
        echo "❌ Porta 6174 não está escutando"
    fi

    echo ""
    echo "🎉 Configuração completa!"
    echo ""
    echo "📡 Endpoints disponíveis:"
    echo "   HTTP RPC: http://144.22.179.183"
    echo "   WebSocket: ws://144.22.179.183/ws"
    echo ""
    echo "🧪 Teste HTTP:"
    echo "   curl -X POST -H \"Content-Type: application/json\" --data '{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":1}' http://144.22.179.183"
    echo ""
    echo "🧪 Teste WebSocket:"
    echo "   node test-websocket.js"

else
    echo "❌ Erro na configuração do nginx!"
    exit 1
fi
