#!/bin/bash

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Função para mostrar uso
show_usage() {
    echo -e "${BLUE}Uso: $0 [comando]${NC}"
    echo ""
    echo "Comandos disponíveis:"
    echo "  start     - Iniciar todos os nós"
    echo "  stop      - Parar todos os nós"
    echo "  restart   - Reiniciar todos os nós"
    echo "  status    - Mostrar status de todos os nós"
    echo "  logs      - Mostrar logs de todos os nós"
    echo "  logs1     - Mostrar logs do Node-1"
    echo "  logs2     - Mostrar logs do Node-2"
    echo "  logs3     - Mostrar logs do Node-3"
    echo "  logs4     - Mostrar logs do Node-4"
    echo "  enable    - Habilitar todos os serviços"
    echo "  disable   - Desabilitar todos os serviços"
    echo ""
    echo "Exemplos:"
    echo "  $0 start"
    echo "  $0 status"
    echo "  $0 logs1"
}

# Função para verificar se está rodando como root
check_root() {
    if [ "$EUID" -ne 0 ]; then
        echo -e "${RED}❌ Este comando precisa ser executado como root (sudo)${NC}"
        echo "Execute: sudo $0 $1"
        exit 1
    fi
}

# Função para iniciar todos os serviços
start_all() {
    check_root "start"
    echo -e "${GREEN}🚀 Iniciando todos os nós Besu...${NC}"
    systemctl start besu-node-1.service
    systemctl start besu-node-2.service
    systemctl start besu-node-3.service
    systemctl start besu-node-4.service
    echo -e "${GREEN}✅ Todos os nós foram iniciados!${NC}"
}

# Função para parar todos os serviços
stop_all() {
    check_root "stop"
    echo -e "${YELLOW}🛑 Parando todos os nós Besu...${NC}"
    systemctl stop besu-node-4.service
    systemctl stop besu-node-3.service
    systemctl stop besu-node-2.service
    systemctl stop besu-node-1.service
    echo -e "${GREEN}✅ Todos os nós foram parados!${NC}"
}

# Função para reiniciar todos os serviços
restart_all() {
    check_root "restart"
    echo -e "${BLUE}🔄 Reiniciando todos os nós Besu...${NC}"
    systemctl restart besu-node-1.service
    systemctl restart besu-node-2.service
    systemctl restart besu-node-3.service
    systemctl restart besu-node-4.service
    echo -e "${GREEN}✅ Todos os nós foram reiniciados!${NC}"
}

# Função para mostrar status
show_status() {
    echo -e "${BLUE}📊 Status dos serviços Besu:${NC}"
    echo ""
    echo -e "${YELLOW}Node-1:${NC}"
    systemctl status besu-node-1.service --no-pager -l
    echo ""
    echo -e "${YELLOW}Node-2:${NC}"
    systemctl status besu-node-2.service --no-pager -l
    echo ""
    echo -e "${YELLOW}Node-3:${NC}"
    systemctl status besu-node-3.service --no-pager -l
    echo ""
    echo -e "${YELLOW}Node-4:${NC}"
    systemctl status besu-node-4.service --no-pager -l
}

# Função para mostrar logs de todos os nós
show_logs() {
    echo -e "${BLUE}📝 Logs dos últimos 50 eventos de cada nó:${NC}"
    echo ""
    echo -e "${YELLOW}Node-1 logs:${NC}"
    journalctl -u besu-node-1.service --no-pager -n 50
    echo ""
    echo -e "${YELLOW}Node-2 logs:${NC}"
    journalctl -u besu-node-2.service --no-pager -n 50
    echo ""
    echo -e "${YELLOW}Node-3 logs:${NC}"
    journalctl -u besu-node-3.service --no-pager -n 50
    echo ""
    echo -e "${YELLOW}Node-4 logs:${NC}"
    journalctl -u besu-node-4.service --no-pager -n 50
}

# Função para mostrar logs de um nó específico
show_node_logs() {
    local node_num=$1
    check_root "logs$node_num"
    echo -e "${BLUE}📝 Logs do Node-$node_num (pressione Ctrl+C para sair):${NC}"
    journalctl -u besu-node-$node_num.service -f
}

# Função para habilitar todos os serviços
enable_all() {
    check_root "enable"
    echo -e "${GREEN}✅ Habilitando todos os serviços...${NC}"
    systemctl enable besu-node-1.service
    systemctl enable besu-node-2.service
    systemctl enable besu-node-3.service
    systemctl enable besu-node-4.service
    echo -e "${GREEN}✅ Todos os serviços foram habilitados!${NC}"
}

# Função para desabilitar todos os serviços
disable_all() {
    check_root "disable"
    echo -e "${YELLOW}❌ Desabilitando todos os serviços...${NC}"
    systemctl disable besu-node-1.service
    systemctl disable besu-node-2.service
    systemctl disable besu-node-3.service
    systemctl disable besu-node-4.service
    echo -e "${GREEN}✅ Todos os serviços foram desabilitados!${NC}"
}

# Verificar argumentos
if [ $# -eq 0 ]; then
    show_usage
    exit 1
fi

# Processar comandos
case "$1" in
    start)
        start_all
        ;;
    stop)
        stop_all
        ;;
    restart)
        restart_all
        ;;
    status)
        show_status
        ;;
    logs)
        show_logs
        ;;
    logs1)
        show_node_logs 1
        ;;
    logs2)
        show_node_logs 2
        ;;
    logs3)
        show_node_logs 3
        ;;
    logs4)
        show_node_logs 4
        ;;
    enable)
        enable_all
        ;;
    disable)
        disable_all
        ;;
    *)
        echo -e "${RED}❌ Comando inválido: $1${NC}"
        echo ""
        show_usage
        exit 1
        ;;
esac
