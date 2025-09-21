// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import "../src/VFinanceRegistry.sol";

/**
 * @title InitializeProxyScript
 * @dev Script para executar initialize do smart contract proxy já deployado
 * @notice Use este script quando o proxy já foi deployado mas não foi inicializado
 */
contract InitializeProxyScript is Script {
    // Configurações para inicialização
    string constant NAME = "VFinance Registry";
    string constant SYMBOL = "VREG";
    string constant METADATA_BASE_URL = "http://144.22.179.183:3000";

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        address payable proxyAddress = payable(0xF4074b3685f8F40f2DCA83742DAB19912A0eb2C3);

        console.log("=== INITIALIZE PROXY SCRIPT ===");
        console.log("Deployer account:", deployer);
        console.log("Deployer balance:", deployer.balance);
        console.log("Proxy address:", proxyAddress);

        vm.startBroadcast(deployerPrivateKey);

        // Criar instância do contrato no endereço do proxy
        VFinanceRegistry registry = VFinanceRegistry(proxyAddress);

        // Verificar se já foi inicializado
        try registry.getVersion() returns (string memory version) {
            console.log("Contract already initialized with version:", version);

            // Verificar se está realmente inicializado (owner não é zero)
            address currentOwner = registry.owner();
            console.log("Current owner:", currentOwner);

            if (currentOwner == address(0)) {
                console.log("Contract partially initialized - owner is zero address");
                console.log("Attempting to reinitialize...");

                // Tentar reinitializar (pode falhar se já foi inicializado)
                try registry.initialize(
                    NAME,
                    SYMBOL,
                    METADATA_BASE_URL,
                    deployer
                ) {
                    console.log("Reinitialization successful!");
                } catch {
                    console.log("Reinitialization failed - contract may be locked");
                    console.log("Trying to update server config instead...");

                    // Tentar atualizar configuração do servidor
                    try registry.updateServerConfig(METADATA_BASE_URL, deployer) {
                        console.log("Server config updated successfully!");
                    } catch {
                        console.log("Failed to update server config");
                    }
                }
            } else {
                console.log("Contract fully initialized");
                console.log("API Server:", registry.apiServerAddress());
                console.log("Metadata Base URL:", registry.metadataBaseUrl());

                // Tentar obter total supply com tratamento de erro
                try registry.totalSupply() returns (uint256 supply) {
                    console.log("Total Supply:", supply);
                } catch {
                    console.log("Total Supply: Error reading (contract may be partially initialized)");
                }
            }
        } catch {
            console.log("Contract not initialized yet, proceeding with initialization...");

            // Executar initialize
            registry.initialize(
                NAME,
                SYMBOL,
                METADATA_BASE_URL,
                deployer // API server address
            );

            console.log("Contract initialized successfully!");
        }

        vm.stopBroadcast();

        // Verificar estado final
        console.log("\n=== FINAL STATE ===");
        console.log("Contract name:", registry.name());
        console.log("Contract symbol:", registry.symbol());
        console.log("Contract version:", registry.getVersion());
        console.log("Owner:", registry.owner());
        console.log("API Server:", registry.apiServerAddress());
        console.log("Metadata Base URL:", registry.metadataBaseUrl());
        console.log("Total Supply:", registry.totalSupply());
        console.log("Max Supply:", registry.MAX_SUPPLY());
    }
}

/**
 * @title ReinitializeProxyScript
 * @dev Script para reconfigurar um contrato já inicializado
 * @notice Use este script para atualizar configurações do contrato
 */
contract ReinitializeProxyScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);
        address payable proxyAddress = payable(vm.envAddress("CONTRACT_ADDRESS"));

        console.log("=== REINITIALIZE PROXY SCRIPT ===");
        console.log("Deployer account:", deployer);
        console.log("Proxy address:", proxyAddress);

        vm.startBroadcast(deployerPrivateKey);

        VFinanceRegistry registry = VFinanceRegistry(proxyAddress);

        // Verificar se é o owner
        require(registry.owner() == deployer, "Only owner can reinitialize");

        // Atualizar configurações do servidor
        string memory newMetadataBaseUrl = "http://144.22.179.183:3000";
        address newApiServerAddress = deployer;

        registry.updateServerConfig(newMetadataBaseUrl, newApiServerAddress);

        console.log("Server configuration updated!");
        console.log("New API Server:", registry.apiServerAddress());
        console.log("New Metadata Base URL:", registry.metadataBaseUrl());

        vm.stopBroadcast();
    }
}

/**
 * @title TestProxyScript
 * @dev Script para testar o contrato proxy
 */
contract TestProxyScript is Script {
    function run() external {
        address payable proxyAddress = payable(vm.envAddress("CONTRACT_ADDRESS"));

        console.log("=== TEST PROXY SCRIPT ===");
        console.log("Testing contract at:", proxyAddress);

        VFinanceRegistry registry = VFinanceRegistry(proxyAddress);

        // Testes básicos
        console.log("Name:", registry.name());
        console.log("Symbol:", registry.symbol());
        console.log("Version:", registry.getVersion());
        console.log("Total Supply:", registry.totalSupply());
        console.log("Owner:", registry.owner());
        console.log("API Server:", registry.apiServerAddress());
        console.log("Metadata Base URL:", registry.metadataBaseUrl());
        console.log("Max Supply:", registry.MAX_SUPPLY());

        // Testar algumas marcas pré-registradas
        console.log("\n=== PRE-REGISTERED BRANDS ===");
        for (uint64 i = 1; i <= 10; i++) {
            try registry.getBrandName(i) returns (string memory brandName) {
                if (bytes(brandName).length > 0) {
                    console.log("Brand ID", i, ":", brandName);
                }
            } catch {
                break;
            }
        }

        // Testar se contrato existe
        console.log("\n=== CONTRACT EXISTENCE TESTS ===");
        bool exists1 = registry.doesContractExist("test-registry-id");
        console.log("Test registry ID exists:", exists1);

        bool exists2 = registry.doesHashExist(keccak256("test-hash"));
        console.log("Test hash exists:", exists2);
    }
}
