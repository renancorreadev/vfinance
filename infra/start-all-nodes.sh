#!/bin/bash

echo "🚀 Iniciando rede Besu QBFT com 4 nós..."
echo "================================================"

# Função para verificar se o processo está rodando
check_process() {
    local node_name=$1
    if pgrep -f "besu.*$node_name" > /dev/null; then
        return 0
    else
        return 1
    fi
}

# Função para aguardar o nó estar pronto
wait_for_node() {
    local node_name=$1
    local rpc_port=$2
    echo "⏳ Aguardando $node_name estar pronto..."
    
    for i in {1..30}; do
        if curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"net_listening","params":[],"id":1}' \
            http://localhost:$rpc_port > /dev/null 2>&1; then
            echo "✅ $node_name está pronto!"
            return 0
        fi
        sleep 2
    done
    echo "❌ Timeout aguardando $node_name"
    return 1
}

# Parar processos existentes se houver
echo "🧹 Parando processos existentes..."
pkill -f "besu.*Node-1" 2>/dev/null || true
pkill -f "besu.*Node-2" 2>/dev/null || true
pkill -f "besu.*Node-3" 2>/dev/null || true
pkill -f "besu.*Node-4" 2>/dev/null || true

# Iniciar Node-1
echo "🔧 Iniciando Node-1..."
cd Node-1
./run.sh
cd ..

# Aguardar Node-1 estar pronto
wait_for_node "Node-1" 8545

# Iniciar Node-2
echo "🔧 Iniciando Node-2..."
cd Node-2
./run.sh
cd ..

# Aguardar Node-2 estar pronto
wait_for_node "Node-2" 8546

# Iniciar Node-3
echo "🔧 Iniciando Node-3..."
cd Node-3
./run.sh
cd ..

# Aguardar Node-3 estar pronto
wait_for_node "Node-3" 8547

# Iniciar Node-4
echo "🔧 Iniciando Node-4..."
cd Node-4
./run.sh
cd ..

# Aguardar Node-4 estar pronto
wait_for_node "Node-4" 8548

echo "================================================"
echo "🎉 Rede Besu QBFT iniciada com sucesso!"
echo ""
echo "📊 Status dos nós:"
echo "Node-1: http://localhost:8545"
echo "Node-2: http://localhost:8546"
echo "Node-3: http://localhost:8547"
echo "Node-4: http://localhost:8548"
echo ""
echo "🔍 Para ver logs de um nó específico:"
echo "tail -f Node-1/besu-node-1.log"
echo "tail -f Node-2/besu-node-2.log"
echo "tail -f Node-3/besu-node-3.log"
echo "tail -f Node-4/besu-node-4.log"
echo ""
echo "🛑 Para parar todos os nós:"
echo "./stop-all-nodes.sh"
echo ""
echo "📝 Para testar a conexão:"
echo "curl -X POST -H \"Content-Type: application/json\" --data '{\"jsonrpc\":\"2.0\",\"method\":\"net_listening\",\"params\":[],\"id\":1}' http://localhost:8545"
