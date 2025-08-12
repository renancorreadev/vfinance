#!/bin/bash

echo "Iniciando Node-4 do Besu QBFT..."

# Configurações específicas do Node-4
NODE_ID=4
RPC_PORT=8548
P2P_PORT=30306
KEY_FILE="data/key"
WS_PORT=6177
METRICS_PORT=9550

# Comando Besu com configurações QBFT e gas zero
besu \
  --data-path=data \
  --genesis-file=data/genesis.json \
  --node-private-key-file=data/key \
  --bootnodes=enode://9dd363a8a495616bd4cc6de1229815b69f34f2eb029a6f157e1c8056f23ad08b3fe8de1c778264fe89be798cd7b5d29666458819a4f7f111a838ed8e300debe5@127.0.0.1:30303 \
  --rpc-http-enabled \
  --rpc-http-port=$RPC_PORT \
  --rpc-http-api=ETH,NET,QBFT,ADMIN,DEBUG \
  --rpc-http-cors-origins="*" \
  --rpc-ws-enabled \
  --rpc-ws-port=$WS_PORT \
  --metrics-enabled \
  --metrics-host=0.0.0.0 \
  --metrics-port=$METRICS_PORT \
  --rpc-ws-host=0.0.0.0 \
  --min-gas-price=0 \
  --p2p-host=0.0.0.0 \
  --p2p-port=$P2P_PORT \
  --profile=ENTERPRISE \
  --logging=INFO \

echo "Node-4 iniciado na porta RPC: $RPC_PORT, P2P: $P2P_PORT"
echo "Para ver logs: tail -f besu-node-4.log"
echo "Para parar: pkill -f 'besu.*Node-4'"
