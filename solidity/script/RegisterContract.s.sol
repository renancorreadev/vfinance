// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/VFinanceRegistry.sol";

/**
 * @title RegisterContract
 * @dev Script para registrar contratos de exemplo no VFinanceRegistry
 * @notice Execute: forge script script/RegisterContract.s.sol --broadcast --rpc-url $RPC_URL
 */
contract RegisterContract is Script {
    VFinanceRegistry public registry;

    // Dados de exemplo para registro
    struct ExampleContract {
        string registryId;
        string contractNumber;
        uint32 contractDate;
        string chassis;
        string licensePlate;
        uint128 totalValue;
        string brandName;
        string modelName;
    }

    function run() external {
        // Configuração do ambiente
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address contractAddress = vm.envAddress("CONTRACT_ADDRESS");

        vm.startBroadcast(deployerPrivateKey);

        // Conectar ao contrato já deployado
        registry = VFinanceRegistry(payable(contractAddress));

        console.log("=== REGISTRANDO CONTRATOS DE EXEMPLO ===");
        console.log("Contrato:", address(registry));
        console.log("API Server:", registry.apiServerAddress());

        // Registrar multiplos contratos de exemplo
        registerExampleContracts();

        vm.stopBroadcast();

        console.log("=== REGISTRO CONCLUIDO ===");
        console.log("Total de contratos:", registry.totalSupply());
    }

    function registerExampleContracts() internal {
        ExampleContract[5] memory examples = [
            ExampleContract({
                registryId: "VFIN-2024-001",
                contractNumber: "FIN123456789",
                contractDate: uint32(block.timestamp),
                chassis: "9BWZZZ377VT004251",
                licensePlate: "BRA2E19",
                totalValue: 75000 ether, // R$ 75.000
                brandName: "VOLKSWAGEN",
                modelName: "GOLF TSI HIGHLINE 1.4"
            }),
            ExampleContract({
                registryId: "VFIN-2024-002",
                contractNumber: "FIN987654321",
                contractDate: uint32(block.timestamp),
                chassis: "9BFZZZ377VT004252",
                licensePlate: "BRA3F28",
                totalValue: 45000 ether, // R$ 45.000
                brandName: "TOYOTA",
                modelName: "COROLLA GLI 1.8"
            }),
            ExampleContract({
                registryId: "VFIN-2024-003",
                contractNumber: "FIN456789123",
                contractDate: uint32(block.timestamp),
                chassis: "9CGZZZ377VT004253",
                licensePlate: "BRA4G37",
                totalValue: 120000 ether, // R$ 120.000
                brandName: "BMW",
                modelName: "320I SPORT 2.0"
            }),
            ExampleContract({
                registryId: "VFIN-2024-004",
                contractNumber: "FIN789123456",
                contractDate: uint32(block.timestamp),
                chassis: "9AHZZZ377VT004254",
                licensePlate: "BRA5H46",
                totalValue: 35000 ether, // R$ 35.000
                brandName: "HONDA",
                modelName: "CIVIC LX 1.6"
            }),
            ExampleContract({
                registryId: "VFIN-2024-005",
                contractNumber: "FIN654321987",
                contractDate: uint32(block.timestamp),
                chassis: "9FJZZZ377VT004255",
                licensePlate: "BRA6I55",
                totalValue: 95000 ether, // R$ 95.000
                brandName: "AUDI",
                modelName: "A3 SEDAN ATTRACTION 1.4"
            })
        ];

        for (uint i = 0; i < examples.length; i++) {
            registerSingleContract(examples[i], i + 1);
        }
    }

    function registerSingleContract(
        ExampleContract memory example,
        uint256 index
    ) internal {
        console.log("--- Registrando Contrato", index, "---");
        console.log("Registry ID:", example.registryId);
        console.log("Chassi:", example.chassis);
        console.log("Marca:", example.brandName);
        console.log("Modelo:", example.modelName);

        (uint256 tokenId, bytes32 metadataHash) = registry.registerContract(
            example.registryId,
            example.contractNumber,
            example.contractDate,
            example.chassis,
            example.licensePlate,
            example.totalValue,
            example.brandName,
            example.modelName
        );

        console.log("Sucesso!");
        console.log("Token ID:", tokenId);
        console.log("Metadata Hash:", vm.toString(metadataHash));
        console.log("Token URI:", registry.tokenURI(tokenId));

        // Verificar dados salvos
        verifyRegisteredContract(tokenId, example);
    }

    function verifyRegisteredContract(
        uint256 tokenId,
        ExampleContract memory example
    ) internal view {
        // Buscar dados do contrato
        (
            VFinanceRegistry.ContractRecord memory contractData,
            VFinanceRegistry.VehicleCore memory vehicleData
        ) = registry.getContract(tokenId);

        // Verificações básicas
        require(contractData.active, "Contrato deve estar ativo");
        require(
            vehicleData.totalValue == example.totalValue,
            "Valor total incorreto"
        );

        console.log("Verificacao OK - Dados salvos corretamente");
    }

    // Função para buscar e exibir contratos existentes
    function listExistingContracts() external view {
        uint256 totalSupply = registry.totalSupply();
        console.log("=== CONTRATOS EXISTENTES ===");
        console.log("Total:", totalSupply);

        if (totalSupply > 0) {
            // Listar primeiros 10 contratos
            uint256 limit = totalSupply > 10 ? 10 : totalSupply;

            for (uint256 i = 1; i <= limit; i++) {
                (
                    VFinanceRegistry.ContractRecord memory contractData,
                    VFinanceRegistry.VehicleCore memory vehicleData
                ) = registry.getContract(i);

                console.log("Token ID:", i);
                console.log("Ativo:", contractData.active);
                console.log("Valor:", vehicleData.totalValue / 1 ether, "ETH");
                console.log("Brand ID:", vehicleData.brandId);
                console.log("Model ID:", vehicleData.modelId);
            }
        }
    }

    // Função para testar consultas básicas (sem dependências)
    function testQueries() external view {
        console.log("=== TESTANDO CONSULTAS BASICAS ===");

        uint256 totalSupply = registry.totalSupply();
        console.log("Total de contratos:", totalSupply);

        // Testar configurações básicas
        console.log("Nome:", registry.name());
        console.log("Simbolo:", registry.symbol());
        console.log("Versao:", registry.getVersion());
        console.log("Owner:", registry.owner());
        console.log("API Server:", registry.apiServerAddress());
        console.log("Metadata Base URL:", registry.metadataBaseUrl());
    }

    // Função para demonstrar sistema de metadados
    function demonstrateMetadataSystem() external view {
        console.log("=== SISTEMA DE METADADOS ===");

        uint256 totalSupply = registry.totalSupply();
        console.log("Total de contratos:", totalSupply);

        if (totalSupply > 0) {
            // Pegar primeiro contrato
            (
                VFinanceRegistry.ContractRecord memory contractData, // VFinanceRegistry.VehicleCore memory vehicleData

            ) = registry.getContract(1);

            console.log(
                "Metadata Hash:",
                vm.toString(contractData.metadataHash)
            );
            console.log("Token URI:", registry.tokenURI(1));
            console.log(
                "Direct Metadata URL:",
                registry.getMetadataUrl(contractData.metadataHash)
            );
        } else {
            console.log("Nenhum contrato encontrado para demonstrar metadados");
        }
    }

    // Funcao auxiliar para converter bytes32 para string (para debugging)
    function bytes32ToString(
        bytes32 data
    ) internal pure returns (string memory) {
        uint256 length = 0;
        while (length < 32 && data[length] != 0) {
            length++;
        }
        bytes memory result = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            result[i] = data[i];
        }
        return string(result);
    }
}
