// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import "../src/VFinanceRegistry.sol";

/**
 * @title VFinanceRegistryV2Test
 * @dev Comprehensive test suite for VFinanceRegistry contract
 * @notice Tests registry system that uses ERC721 infrastructure (not traditional NFTs)
 */
contract VFinanceRegistryV2Test is Test {
    VFinanceRegistry public registry;
    VFinanceRegistry public implementation;
    ERC1967Proxy public proxy;

    address public owner;
    address public apiServer;
    address public user1;
    address public user2;

    string constant NAME = "VFinance Registry";
    string constant SYMBOL = "VFIN";
    string constant METADATA_BASE_URL = "https://api.vfinance.com.br";

    // Test data
    string constant REGISTRY_ID = "6193a9b1-38a8-4d4b-a21a-8a8ca9aef9c9";
    string constant CONTRACT_NUMBER = "250";
    string constant CHASSIS = "8A1CB8205DL478385";
    string constant LICENSE_PLATE = "PGA4J42";
    string constant BRAND_NAME = "RENAULT";
    string constant MODEL_NAME = "CLIO AUTHENTIQUE 1.0";
    uint32 constant CONTRACT_DATE = 1646697600; // 2022-03-08
    uint128 constant TOTAL_VALUE = 23900 ether;

    event ContractRegistered(
        uint256 indexed tokenId,
        bytes32 indexed registryIdHash,
        bytes32 indexed chassisHash,
        bytes32 metadataHash,
        uint32 timestamp
    );

    event VehicleTracked(
        uint256 indexed tokenId,
        bytes32 indexed chassis,
        bytes32 indexed licensePlate,
        uint128 totalValue,
        uint64 brandId,
        uint64 modelId
    );

    event MetadataUpdated(
        uint256 indexed tokenId,
        bytes32 indexed oldHash,
        bytes32 indexed newHash,
        uint32 timestamp
    );

    event StatusChanged(
        uint256 indexed tokenId,
        bool indexed active,
        uint32 timestamp
    );

    event SystemConfigured(address indexed apiServer, string metadataUrl);

    function setUp() public {
        // Setup accounts
        owner = makeAddr("owner");
        apiServer = makeAddr("apiServer");
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");

        vm.startPrank(owner);

        // Deploy implementation
        implementation = new VFinanceRegistry();

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            VFinanceRegistry.initialize.selector,
            NAME,
            SYMBOL,
            METADATA_BASE_URL,
            apiServer
        );

        // Deploy proxy
        proxy = new ERC1967Proxy(address(implementation), initData);
        registry = VFinanceRegistry(payable(address(proxy)));

        vm.stopPrank();
    }

    // =============================================================
    //                      INITIALIZATION TESTS
    // =============================================================

    function testInitialization() public {
        assertEq(registry.name(), NAME);
        assertEq(registry.symbol(), SYMBOL);
        assertEq(registry.metadataBaseUrl(), METADATA_BASE_URL);
        assertEq(registry.apiServerAddress(), apiServer);
        assertEq(registry.owner(), owner);
        assertEq(registry.getVersion(), "2.0.0");
        assertEq(registry.totalSupply(), 0);
    }

    function testInitialBrandsRegistered() public {
        // Check that common brands were registered during initialization
        assertEq(registry.getBrandName(1), "TOYOTA");
        assertEq(registry.getBrandName(2), "HONDA");
        assertEq(registry.getBrandName(3), "VOLKSWAGEN");
        assertEq(registry.getBrandName(4), "FIAT");
        assertEq(registry.getBrandName(5), "CHEVROLET");
        assertEq(registry.getBrandName(6), "FORD");
        assertEq(registry.getBrandName(7), "NISSAN");
        assertEq(registry.getBrandName(8), "HYUNDAI");
        assertEq(registry.getBrandName(9), "RENAULT");
        assertEq(registry.getBrandName(10), "PEUGEOT");
    }

    function testCannotInitializeTwice() public {
        vm.expectRevert(); // InvalidInitialization error in OpenZeppelin v5
        registry.initialize(NAME, SYMBOL, METADATA_BASE_URL, apiServer);
    }

    // =============================================================
    //                   CONTRACT REGISTRATION TESTS
    // =============================================================

    function testRegisterContract() public {
        vm.startPrank(apiServer);

        bytes32 registryIdHash = keccak256(bytes(REGISTRY_ID));
        bytes32 chassisHash = keccak256(bytes(CHASSIS));

        // We expect the events but don't check exact values for hash and timestamp (they're generated)
        vm.expectEmit(true, true, false, false);
        emit ContractRegistered(1, registryIdHash, chassisHash, bytes32(0), 0);

        // For VehicleTracked, chassis and licensePlate are converted to bytes32
        bytes32 chassisBytes = bytes32(bytes(CHASSIS));
        bytes32 licensePlateBytes = bytes32(bytes(LICENSE_PLATE));

        vm.expectEmit(true, true, true, false);
        emit VehicleTracked(
            1,
            chassisBytes,
            licensePlateBytes,
            TOTAL_VALUE,
            9,
            1
        ); // RENAULT is brand ID 9, model ID will be 1

        (uint256 tokenId, bytes32 metadataHash) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        assertEq(tokenId, 1);
        assertNotEq(metadataHash, bytes32(0));
        assertEq(registry.totalSupply(), 1);
        assertEq(registry.ownerOf(tokenId), address(registry)); // Token belongs to contract (non-transferable)

        vm.stopPrank();
    }

    function testRegisterContractCreatesNewModel() public {
        vm.startPrank(apiServer);

        string memory newModel = "NEW MODEL 2024";

        (uint256 tokenId, ) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            newModel
        );

        // Get vehicle data
        (, VFinanceRegistry.VehicleCore memory vehicleCore) = registry
            .getContract(tokenId);

        // Check that new model was created
        string memory storedModel = registry.getModelName(vehicleCore.modelId);
        assertEq(storedModel, newModel);

        vm.stopPrank();
    }

    function testCannotRegisterWithEmptyRegistryId() public {
        vm.startPrank(apiServer);

        vm.expectRevert(VFinanceRegistry.InvalidInput.selector);
        registry.registerContract(
            "",
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();
    }

    function testCannotRegisterWithEmptyChassis() public {
        vm.startPrank(apiServer);

        vm.expectRevert(VFinanceRegistry.InvalidInput.selector);
        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            "",
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();
    }

    function testCannotRegisterDuplicateRegistryId() public {
        vm.startPrank(apiServer);

        // Register first contract
        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        // Try to register with same registry ID
        vm.expectRevert(VFinanceRegistry.ContractAlreadyExists.selector);
        registry.registerContract(
            REGISTRY_ID,
            "251",
            CONTRACT_DATE,
            "DIFFERENT_CHASSIS",
            "DIF-PLATE",
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();
    }

    function testCannotRegisterDuplicateChassis() public {
        vm.startPrank(apiServer);

        // Register first contract
        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        // Try to register with same chassis
        vm.expectRevert(VFinanceRegistry.ContractAlreadyExists.selector);
        registry.registerContract(
            "different-registry-id",
            "251",
            CONTRACT_DATE,
            CHASSIS,
            "DIF-PLATE",
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();
    }

    function testCannotRegisterUnauthorized() public {
        vm.startPrank(user1);

        vm.expectRevert(VFinanceRegistry.UnauthorizedAccess.selector);
        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();
    }

    // =============================================================
    //                        VIEW FUNCTION TESTS
    // =============================================================

    function testGetContract() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, ) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        (
            VFinanceRegistry.ContractRecord memory contractRecord,
            VFinanceRegistry.VehicleCore memory vehicleCore
        ) = registry.getContract(tokenId);

        assertEq(contractRecord.registryId, keccak256(bytes(REGISTRY_ID)));
        assertEq(contractRecord.contractDate, CONTRACT_DATE);
        assertTrue(contractRecord.active);
        assertEq(contractRecord.registeredBy, apiServer);

        assertEq(vehicleCore.totalValue, TOTAL_VALUE);
        assertEq(vehicleCore.brandId, 9); // RENAULT

        vm.stopPrank();
    }

    function testGetContractByRegistryId() public {
        vm.startPrank(apiServer);

        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        (
            VFinanceRegistry.ContractRecord memory contractRecord,
            VFinanceRegistry.VehicleCore memory vehicleCore
        ) = registry.getContractByRegistryId(REGISTRY_ID);

        assertEq(contractRecord.registryId, keccak256(bytes(REGISTRY_ID)));
        assertEq(vehicleCore.totalValue, TOTAL_VALUE);

        vm.stopPrank();
    }

    function testGetContractByChassis() public {
        vm.startPrank(apiServer);

        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        (
            VFinanceRegistry.ContractRecord memory contractRecord,
            VFinanceRegistry.VehicleCore memory vehicleCore
        ) = registry.getContractByChassis(CHASSIS);

        assertEq(contractRecord.registryId, keccak256(bytes(REGISTRY_ID)));
        assertEq(vehicleCore.totalValue, TOTAL_VALUE);

        vm.stopPrank();
    }

    function testDoesContractExist() public {
        vm.startPrank(apiServer);

        assertFalse(registry.doesContractExist(REGISTRY_ID));

        registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        assertTrue(registry.doesContractExist(REGISTRY_ID));

        vm.stopPrank();
    }

    function testGetActiveContracts() public {
        vm.startPrank(apiServer);

        // Register multiple contracts
        for (uint i = 0; i < 5; i++) {
            registry.registerContract(
                string(abi.encodePacked(REGISTRY_ID, vm.toString(i))),
                string(abi.encodePacked(CONTRACT_NUMBER, vm.toString(i))),
                CONTRACT_DATE,
                string(abi.encodePacked(CHASSIS, vm.toString(i))),
                string(abi.encodePacked(LICENSE_PLATE, vm.toString(i))),
                TOTAL_VALUE,
                BRAND_NAME,
                MODEL_NAME
            );
        }

        uint256[] memory activeContracts = registry.getActiveContracts(0, 10);
        assertEq(activeContracts.length, 5);

        // Test pagination
        uint256[] memory paginatedContracts = registry.getActiveContracts(2, 2);
        assertEq(paginatedContracts.length, 2);
        assertEq(paginatedContracts[0], 3);
        assertEq(paginatedContracts[1], 4);

        vm.stopPrank();
    }

    function testGetActiveContractsInvalidLimit() public {
        vm.expectRevert(VFinanceRegistry.InvalidInput.selector);
        registry.getActiveContracts(0, 0);

        vm.expectRevert(VFinanceRegistry.InvalidInput.selector);
        registry.getActiveContracts(0, 101);
    }

    // =============================================================
    //                    UPDATE FUNCTION TESTS
    // =============================================================

    function testUpdateMetadataHash() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, bytes32 oldHash) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        bytes32 newHash = keccak256("new metadata hash");

        vm.expectEmit(true, true, true, false);
        emit MetadataUpdated(tokenId, oldHash, newHash, 0);

        registry.updateMetadataHash(tokenId, newHash);

        (VFinanceRegistry.ContractRecord memory contractRecord, ) = registry
            .getContract(tokenId);
        assertEq(contractRecord.metadataHash, newHash);

        vm.stopPrank();
    }

    function testUpdateStatus() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, ) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.expectEmit(true, true, true, false);
        emit StatusChanged(tokenId, false, 0);

        registry.updateStatus(tokenId, false);

        (VFinanceRegistry.ContractRecord memory contractRecord, ) = registry
            .getContract(tokenId);
        assertFalse(contractRecord.active);

        vm.stopPrank();
    }

    function testCannotUpdateMetadataHashUnauthorized() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, ) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();

        vm.startPrank(user1);

        vm.expectRevert(VFinanceRegistry.UnauthorizedAccess.selector);
        registry.updateMetadataHash(tokenId, keccak256("new hash"));

        vm.stopPrank();
    }

    function testCannotUpdateInvalidToken() public {
        vm.startPrank(apiServer);

        vm.expectRevert(VFinanceRegistry.TokenNotFound.selector);
        registry.updateMetadataHash(999, keccak256("new hash"));

        vm.stopPrank();
    }

    // =============================================================
    //                      NFT FUNCTION TESTS
    // =============================================================

    function testTokenURI() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, bytes32 metadataHash) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        string memory expectedURI = string(
            abi.encodePacked(
                METADATA_BASE_URL,
                "/api/metadata/0x",
                _toHexString(metadataHash)
            )
        );

        string memory actualURI = registry.tokenURI(tokenId);
        assertEq(actualURI, expectedURI);

        vm.stopPrank();
    }

    function testTokenURIInvalidToken() public {
        vm.expectRevert(VFinanceRegistry.TokenNotFound.selector);
        registry.tokenURI(999);
    }

    function testGetMetadataUrl() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, bytes32 metadataHash) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        string memory expectedUrl = string(
            abi.encodePacked(
                METADATA_BASE_URL,
                "/api/metadata/0x",
                _toHexString(metadataHash)
            )
        );

        string memory actualUrl = registry.getMetadataUrl(metadataHash);
        assertEq(actualUrl, expectedUrl);

        vm.stopPrank();
    }

    function testGetMetadataUrlByRegistryId() public {
        vm.startPrank(apiServer);

        (, bytes32 metadataHash) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        string memory expectedUrl = string(
            abi.encodePacked(
                METADATA_BASE_URL,
                "/api/metadata/0x",
                _toHexString(metadataHash)
            )
        );

        string memory actualUrl = registry.getMetadataUrlByRegistryId(
            REGISTRY_ID
        );
        assertEq(actualUrl, expectedUrl);

        vm.stopPrank();
    }

    // =============================================================
    //                   NON-TRANSFERABLE TESTS
    // =============================================================

    function testTokensAreNonTransferable() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, ) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();

        // Test that transfers are blocked (they will fail at _update level)
        vm.expectRevert("Registry tokens are non-transferable");
        registry.transferFrom(address(registry), user1, tokenId);

        vm.expectRevert("Registry tokens are non-transferable");
        registry.safeTransferFrom(address(registry), user1, tokenId);

        vm.expectRevert("Registry tokens are non-transferable");
        registry.safeTransferFrom(address(registry), user1, tokenId, "");
    }

    function testApprovalsAreBlocked() public {
        vm.startPrank(apiServer);

        (uint256 tokenId, ) = registry.registerContract(
            REGISTRY_ID,
            CONTRACT_NUMBER,
            CONTRACT_DATE,
            CHASSIS,
            LICENSE_PLATE,
            TOTAL_VALUE,
            BRAND_NAME,
            MODEL_NAME
        );

        vm.stopPrank();

        // Test that approvals are blocked
        vm.expectRevert("Registry tokens are non-transferable");
        registry.approve(user1, tokenId);

        vm.expectRevert("Registry tokens are non-transferable");
        registry.setApprovalForAll(user1, true);
    }

    // =============================================================
    //                     ADMIN FUNCTION TESTS
    // =============================================================

    function testUpdateServerConfig() public {
        vm.startPrank(owner);

        string memory newUrl = "https://new-api.vfinance.com.br";
        address newApiServer = makeAddr("newApiServer");

        vm.expectEmit(true, false, false, true);
        emit SystemConfigured(newApiServer, newUrl);

        registry.updateServerConfig(newUrl, newApiServer);

        assertEq(registry.metadataBaseUrl(), newUrl);
        assertEq(registry.apiServerAddress(), newApiServer);

        vm.stopPrank();
    }

    function testRegisterBrandManually() public {
        vm.startPrank(owner);

        string memory newBrand = "TESLA";
        uint64 brandId = registry.registerBrand(newBrand);

        assertEq(registry.getBrandName(brandId), newBrand);
        assertEq(brandId, 11); // After the 10 initial brands

        vm.stopPrank();
    }

    function testRegisterModelManually() public {
        vm.startPrank(owner);

        string memory newModel = "MODEL S";
        uint64 modelId = registry.registerModel(newModel);

        assertEq(registry.getModelName(modelId), newModel);
        assertEq(modelId, 1); // First model

        vm.stopPrank();
    }

    function testCannotUpdateServerConfigUnauthorized() public {
        vm.startPrank(user1);

        vm.expectRevert(); // OwnableUnauthorizedAccount error in OpenZeppelin v5
        registry.updateServerConfig("new-url", makeAddr("newApiServer"));

        vm.stopPrank();
    }

    // =============================================================
    //                      UPGRADE TESTS
    // =============================================================

    function testUpgrade() public {
        vm.startPrank(owner);

        // Deploy new implementation
        VFinanceRegistry newImplementation = new VFinanceRegistry();

        // Upgrade
        registry.upgradeToAndCall(address(newImplementation), "");

        // Verify upgrade worked
        assertEq(registry.getVersion(), "2.0.0");

        vm.stopPrank();
    }

    function testCannotUpgradeUnauthorized() public {
        vm.startPrank(user1);

        VFinanceRegistry newImplementation = new VFinanceRegistry();

        vm.expectRevert(); // OwnableUnauthorizedAccount error in OpenZeppelin v5
        registry.upgradeToAndCall(address(newImplementation), "");

        vm.stopPrank();
    }

    // =============================================================
    //                      EMERGENCY TESTS
    // =============================================================

    function testReceiveETH() public {
        uint256 amount = 1 ether;

        vm.deal(user1, amount);
        vm.startPrank(user1);

        (bool success, ) = address(registry).call{value: amount}("");
        assertTrue(success);

        assertEq(address(registry).balance, amount);

        vm.stopPrank();
    }

    // =============================================================
    //                     HELPER FUNCTIONS
    // =============================================================

    function _toHexString(bytes32 hash) internal pure returns (string memory) {
        bytes memory buffer = new bytes(64);
        for (uint256 i = 0; i < 32; i++) {
            buffer[i * 2] = _toHexChar(uint8(hash[i]) / 16);
            buffer[i * 2 + 1] = _toHexChar(uint8(hash[i]) % 16);
        }
        return string(buffer);
    }

    function _toHexChar(uint8 digit) internal pure returns (bytes1) {
        return
            bytes1(
                digit < 10
                    ? uint8(bytes1("0")) + digit
                    : uint8(bytes1("a")) + digit - 10
            );
    }

    // =============================================================
    //                      FUZZ TESTS
    // =============================================================

    function testFuzzRegisterContract(
        string calldata registryId,
        string calldata contractNumber,
        uint32 contractDate,
        string calldata chassis,
        string calldata licensePlate,
        uint128 totalValue,
        string calldata brandName,
        string calldata modelName
    ) public {
        vm.assume(
            bytes(registryId).length > 0 && bytes(registryId).length <= 32
        );
        vm.assume(bytes(chassis).length > 0 && bytes(chassis).length <= 32);
        vm.assume(bytes(brandName).length > 0 && bytes(brandName).length <= 32);
        vm.assume(bytes(modelName).length > 0 && bytes(modelName).length <= 32);
        vm.assume(contractDate > 0);

        vm.startPrank(apiServer);

        (uint256 tokenId, bytes32 metadataHash) = registry.registerContract(
            registryId,
            contractNumber,
            contractDate,
            chassis,
            licensePlate,
            totalValue,
            brandName,
            modelName
        );

        assertGt(tokenId, 0);
        assertNotEq(metadataHash, bytes32(0));
        assertTrue(registry.doesContractExist(registryId));

        vm.stopPrank();
    }

    function testFuzzGetActiveContracts(uint256 offset, uint256 limit) public {
        vm.assume(limit > 0 && limit <= 100);

        // Register some contracts first
        vm.startPrank(apiServer);
        for (uint i = 0; i < 5; i++) {
            registry.registerContract(
                string(abi.encodePacked("registry-", vm.toString(i))),
                string(abi.encodePacked("contract-", vm.toString(i))),
                CONTRACT_DATE,
                string(abi.encodePacked("chassis-", vm.toString(i))),
                string(abi.encodePacked("plate-", vm.toString(i))),
                TOTAL_VALUE,
                BRAND_NAME,
                MODEL_NAME
            );
        }
        vm.stopPrank();

        uint256[] memory activeContracts = registry.getActiveContracts(
            offset,
            limit
        );

        if (offset >= 5) {
            assertEq(activeContracts.length, 0);
        } else {
            assertLe(activeContracts.length, limit);
            assertLe(activeContracts.length, 5 - offset);
        }
    }
}
