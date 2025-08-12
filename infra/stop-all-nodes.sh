#!/bin/bash

echo "🛑 Parando todos os nós da rede Besu QBFT..."
echo "================================================"

# Parar todos os processos Besu
echo "⏹️  Parando processos..."
pkill -f "besu.*Node-1" 2>/dev/null || true
pkill -f "besu.*Node-2" 2>/dev/null || true
pkill -f "besu.*Node-3" 2>/dev/null || true
pkill -f "besu.*Node-4" 2>/dev/null || true

echo "✅ Todos os nós foram parados!"
echo ""
echo "📊 Processos Besu ativos:"
ps aux | grep besu | grep -v grep
