// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import "forge-std/Script.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "../src/VFinanceRegistry.sol";

/**
 * @title DeployVFinanceRegistryV2
 * @dev Script para deploy do contrato VFinanceRegistry com proxy UUPS
 */
contract DeployVFinanceRegistry is Script {
    // Default configurations for development
    string constant NAME = "VFinance Registry";
    string constant SYMBOL = "VREG";
    string constant METADATA_BASE_URL = "http://localhost:8080";
    // No registration fee for development

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address deployer = vm.addr(deployerPrivateKey);

        console.log("Deploying contracts with the account:", deployer);
        console.log("Account balance:", deployer.balance);

        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy da implementação
        VFinanceRegistry implementation = new VFinanceRegistry();
        console.log("Implementation deployed to:", address(implementation));

        // 2. Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            VFinanceRegistry.initialize.selector,
            NAME,
            SYMBOL,
            METADATA_BASE_URL,
            deployer // API server address (dev)
        );

        // 3. Deploy do proxy
        ERC1967Proxy proxy = new ERC1967Proxy(
            address(implementation),
            initData
        );
        console.log("Proxy deployed to:", address(proxy));

        // 4. Verify initialization worked
        VFinanceRegistry registry = VFinanceRegistry(payable(address(proxy)));
        console.log("Contract name:", registry.name());
        console.log("Contract symbol:", registry.symbol());
        console.log("Contract version:", registry.getVersion());
        console.log("Owner:", registry.owner());
        console.log("API Server:", registry.apiServerAddress());
        console.log("Metadata Base URL:", registry.metadataBaseUrl());

        vm.stopBroadcast();

        // 5. Salvar endereços para uso posterior
        console.log("\n=== DEPLOYMENT SUMMARY ===");
        console.log("Implementation:", address(implementation));
        console.log("Proxy (Main Contract):", address(proxy));
        console.log("Owner/Admin:", deployer);
        console.log("Metadata Base URL:", METADATA_BASE_URL);

        // Salvar em arquivo .env local
        string memory envContent = string(
            abi.encodePacked(
                "CONTRACT_ADDRESS=",
                vm.toString(address(proxy)),
                "\nIMPLEMENTATION_ADDRESS=",
                vm.toString(address(implementation))
            )
        );

        vm.writeFile("./deployment.env", envContent);
        console.log("\nDeployment addresses saved to ./deployment.env");
    }
}

/**
 * @title UpgradeVFinanceRegistryV2
 * @dev Script para upgrade do contrato existente
 */
contract UpgradeVFinanceRegistryV2 is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        address payable proxyAddress = payable(
            vm.envAddress("CONTRACT_ADDRESS")
        );

        console.log("Upgrading contract at:", proxyAddress);

        vm.startBroadcast(deployerPrivateKey);

        // 1. Deploy nova implementação
        VFinanceRegistry newImplementation = new VFinanceRegistry();
        console.log(
            "New implementation deployed to:",
            address(newImplementation)
        );

        // 2. Executar upgrade
        VFinanceRegistry registry = VFinanceRegistry(proxyAddress);
        registry.upgradeToAndCall(
            address(newImplementation),
            "" // Sem dados de inicialização adicional
        );

        // 3. Verificar upgrade
        console.log("Upgrade completed!");
        console.log("New version:", registry.getVersion());

        vm.stopBroadcast();
    }
}

/**
 * @title TestDeployment
 * @dev Script para testar o deployment local
 */
contract TestDeployment is Script {
    function run() external {
        address payable contractAddress = payable(
            vm.envAddress("CONTRACT_ADDRESS")
        );

        console.log("Testing contract at:", contractAddress);

        VFinanceRegistry registry = VFinanceRegistry(contractAddress);

        // Testes básicos
        console.log("Name:", registry.name());
        console.log("Symbol:", registry.symbol());
        console.log("Version:", registry.getVersion());
        console.log("Total Supply:", registry.totalSupply());
        console.log("Owner:", registry.owner());
        console.log("API Server:", registry.apiServerAddress());
        console.log("Metadata Base URL:", registry.metadataBaseUrl());
        console.log("Max Supply:", registry.MAX_SUPPLY());
    }
}
