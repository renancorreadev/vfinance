#!/bin/bash

# Script de instalação do Hyperledger Besu v25.7.0
# Autor: Script automatizado
# Data: $(date)

set -e  # Para o script se houver erro

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para log colorido
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Variáveis
BESU_VERSION="25.7.0"
BESU_URL="https://github.com/hyperledger/besu/releases/download/${BESU_VERSION}/besu-${BESU_VERSION}.tar.gz"
INSTALL_DIR="/opt/besu"
BIN_DIR="/usr/local/bin"
BESU_USER="besu"

log_info "Iniciando instalação do Hyperledger Besu v${BESU_VERSION}"

# Verificar se está rodando como root
if [[ $EUID -ne 0 ]]; then
   log_error "Este script deve ser executado como root (use sudo)"
   exit 1
fi

# Verificar se Java está instalado
log_info "Verificando instalação do Java..."
if ! command -v java &> /dev/null; then
    log_warning "Java não encontrado. Instalando OpenJDK 17..."

    # Detectar distribuição
    if command -v apt-get &> /dev/null; then
        apt-get update
        apt-get install -y openjdk-17-jdk
    elif command -v yum &> /dev/null; then
        yum install -y java-17-openjdk-devel
    elif command -v dnf &> /dev/null; then
        dnf install -y java-17-openjdk-devel
    else
        log_error "Gerenciador de pacotes não suportado. Instale Java 17+ manualmente."
        exit 1
    fi
else
    JAVA_VERSION=$(java -version 2>&1 | awk -F '"' '/version/ {print $2}' | cut -d'.' -f1)
    if [[ $JAVA_VERSION -lt 17 ]]; then
        log_error "Java 17+ é necessário. Versão atual: $JAVA_VERSION"
        exit 1
    fi
    log_success "Java encontrado: $(java -version 2>&1 | head -n 1)"
fi

# Criar usuário para o Besu (se não existir)
if ! id "$BESU_USER" &>/dev/null; then
    log_info "Criando usuário $BESU_USER..."
    useradd -r -s /bin/false -d $INSTALL_DIR $BESU_USER
else
    log_info "Usuário $BESU_USER já existe"
fi

# Criar diretório de instalação
log_info "Criando diretório de instalação: $INSTALL_DIR"
mkdir -p $INSTALL_DIR
cd /tmp

# Download do Besu
log_info "Fazendo download do Besu v${BESU_VERSION}..."
if ! wget -O besu-${BESU_VERSION}.tar.gz "$BESU_URL"; then
    log_error "Falha no download do Besu"
    exit 1
fi

log_success "Download concluído"

# Verificar integridade do arquivo (opcional)
log_info "Verificando arquivo baixado..."
if [[ ! -f "besu-${BESU_VERSION}.tar.gz" ]]; then
    log_error "Arquivo não encontrado após download"
    exit 1
fi

# Extrair arquivo
log_info "Extraindo Besu..."
tar -xzf besu-${BESU_VERSION}.tar.gz

# Mover para diretório de instalação
log_info "Movendo arquivos para $INSTALL_DIR..."
rm -rf $INSTALL_DIR/*
mv besu-${BESU_VERSION}/* $INSTALL_DIR/

# Configurar permissões
log_info "Configurando permissões..."
chown -R $BESU_USER:$BESU_USER $INSTALL_DIR
chmod +x $INSTALL_DIR/bin/besu

# Criar link simbólico
log_info "Criando link simbólico em $BIN_DIR..."
ln -sf $INSTALL_DIR/bin/besu $BIN_DIR/besu

# Criar diretórios de dados e logs
log_info "Criando diretórios de dados e logs..."
mkdir -p /var/lib/besu
mkdir -p /var/log/besu
chown $BESU_USER:$BESU_USER /var/lib/besu
chown $BESU_USER:$BESU_USER /var/log/besu

# Criar arquivo de configuração básico
log_info "Criando arquivo de configuração básico..."
cat > /etc/besu/config.toml << EOF
# Configuração básica do Besu
data-path="/var/lib/besu"
logging="INFO"

# Configurações de rede
network="mainnet"
sync-mode="FAST"

# Configurações RPC (desabilitadas por padrão por segurança)
# rpc-http-enabled=true
# rpc-http-host="127.0.0.1"
# rpc-http-port=8545

# Configurações WebSocket (desabilitadas por padrão)
# rpc-ws-enabled=true
# rpc-ws-host="127.0.0.1"
# rpc-ws-port=8546
EOF

mkdir -p /etc/besu
chown $BESU_USER:$BESU_USER /etc/besu/config.toml

# Criar serviço systemd
log_info "Criando serviço systemd..."
cat > /etc/systemd/system/besu.service << EOF
[Unit]
Description=Hyperledger Besu Ethereum Client
After=network.target
Wants=network.target

[Service]
Type=simple
User=$BESU_USER
Group=$BESU_USER
ExecStart=$INSTALL_DIR/bin/besu --config-file=/etc/besu/config.toml
Restart=on-failure
RestartSec=5
StandardOutput=journal
StandardError=journal
SyslogIdentifier=besu
KillMode=mixed
TimeoutStopSec=60

# Limites de recursos
LimitNOFILE=65536
LimitNPROC=65536

[Install]
WantedBy=multi-user.target
EOF

# Recarregar systemd
systemctl daemon-reload

# Limpeza
log_info "Limpando arquivos temporários..."
rm -f /tmp/besu-${BESU_VERSION}.tar.gz
rm -rf /tmp/besu-${BESU_VERSION}

# Verificar instalação
log_info "Verificando instalação..."
if $INSTALL_DIR/bin/besu --version &>/dev/null; then
    INSTALLED_VERSION=$($INSTALL_DIR/bin/besu --version | head -n 1)
    log_success "Besu instalado com sucesso: $INSTALLED_VERSION"
else
    log_error "Falha na verificação da instalação"
    exit 1
fi

log_success "========================================"
log_success "Instalação do Besu concluída!"
log_success "========================================"
echo ""
log_info "Próximos passos:"
echo "1. Editar configuração: sudo nano /etc/besu/config.toml"
echo "2. Iniciar serviço: sudo systemctl start besu"
echo "3. Habilitar auto-start: sudo systemctl enable besu"
echo "4. Verificar status: sudo systemctl status besu"
echo "5. Ver logs: sudo journalctl -f -u besu"
echo ""
log_info "Comandos úteis:"
echo "- Versão: besu --version"
echo "- Ajuda: besu --help"
echo "- Configuração: cat /etc/besu/config.toml"
echo ""
log_warning "IMPORTANTE: Revise as configurações de segurança antes de usar em produção!"
