#!/bin/bash

echo "🔧 Instalando serviços systemd para Besu QBFT..."
echo "================================================"

# Verificar se está rodando como root
if [ "$EUID" -ne 0 ]; then
    echo "❌ Este script precisa ser executado como root (sudo)"
    echo "Execute: sudo ./install-systemd-services.sh"
    exit 1
fi

# Verificar se o Besu está instalado
if ! command -v besu &> /dev/null; then
    echo "❌ Besu não está instalado ou não está no PATH"
    echo "Instale o Besu primeiro: https://besu.hyperledger.org/en/stable/HowTo/GetStarted/Installation/"
    exit 1
fi

echo "✅ Besu encontrado: $(which besu)"

# Caminhos absolutos
BASE_DIR="/home/ubuntu/besu"
JAVA_HOME="/usr/lib/jvm/java-1.21.0-openjdk-amd64"
USER_SVC="ubuntu"

# Gerar arquivos de serviço systemd
for i in 1 2 3 4; do
  cat > $BASE_DIR/besu-node-$i.service <<EOF
[Unit]
Description=Besu Node $i - QBFT Validator
After=network.target
Wants=network.target

[Service]
Type=simple
User=$USER_SVC
Group=$USER_SVC
WorkingDirectory=$BASE_DIR/Node-$i
Environment="JAVA_HOME=$JAVA_HOME"
ExecStart=$BASE_DIR/Node-$i/run.sh
Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal
SyslogIdentifier=besu-node-$i

[Install]
WantedBy=multi-user.target
EOF
  echo "✅ Gerado: $BASE_DIR/besu-node-$i.service"
done

# Copiar arquivos de serviço para /etc/systemd/system/
echo "📁 Copiando arquivos de serviço..."
for i in 1 2 3 4; do
  cp $BASE_DIR/besu-node-$i.service /etc/systemd/system/
done

# Recarregar systemd
echo "🔄 Recarregando systemd..."
systemctl daemon-reload

# Habilitar serviços para iniciar automaticamente
echo "✅ Habilitando serviços..."
for i in 1 2 3 4; do
  systemctl enable besu-node-$i.service
done

echo "================================================"
echo "🎉 Serviços systemd instalados com sucesso!"
echo ""
echo "📋 Comandos úteis:"
echo ""
echo "🟢 Iniciar todos os serviços:"
for i in 1 2 3 4; do
  echo "sudo systemctl start besu-node-$i.service"
done
echo ""
echo "🔴 Parar todos os serviços:"
for i in 1 2 3 4; do
  echo "sudo systemctl stop besu-node-$i.service"
done
echo ""
echo "📊 Verificar status:"
for i in 1 2 3 4; do
  echo "sudo systemctl status besu-node-$i.service"
done
echo ""
echo "📝 Ver logs:"
for i in 1 2 3 4; do
  echo "sudo journalctl -u besu-node-$i.service -f"
done
echo ""
echo "🔄 Reiniciar serviços:"
for i in 1 2 3 4; do
  echo "sudo systemctl restart besu-node-$i.service"
done
echo ""
echo "🚀 Os serviços iniciarão automaticamente na próxima inicialização do sistema!"
