#!/bin/bash

echo "Iniciando Node-1 do Besu QBFT..."

# Configurações específicas do Node-1
NODE_ID=1
RPC_PORT=8545
P2P_PORT=30303
KEY_FILE="data/key"
WS_PORT=6174
METRICS_PORT=9547

# Comando Besu com configurações QBFT e gas zero
/usr/local/bin/besu \
  --data-path=data \
  --genesis-file=data/genesis.json \
  --node-private-key-file=data/key \
  --rpc-http-enabled \
  --rpc-http-port=$RPC_PORT \
  --rpc-http-host=0.0.0.0 \
  --rpc-http-api=ETH,NET,QBFT,ADMIN,DEBUG,WEB3 \
  --rpc-http-cors-origins="*" \
  --host-allowlist="*" \
  --rpc-ws-enabled \
  --rpc-ws-port=$WS_PORT \
  --rpc-ws-host=0.0.0.0 \
  --rpc-ws-api=ETH,NET,QBFT,ADMIN,DEBUG,WEB3 \
  --metrics-enabled \
  --metrics-host=0.0.0.0 \
  --metrics-port=$METRICS_PORT \
  --min-gas-price=0 \
  --p2p-host=0.0.0.0 \
  --p2p-port=$P2P_PORT \
  --profile=ENTERPRISE \
  --logging=INFO

echo "Node-1 iniciado na porta RPC: $RPC_PORT, P2P: $P2P_PORT"
echo "Para ver logs: sudo journalctl -u besu-node-1.service -f"
echo "Para parar: sudo systemctl stop besu-node-1.service"
