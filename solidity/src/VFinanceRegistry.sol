// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

/**
 * @title VFinanceRegistry
 * @dev Optimized vehicle financing contract registry using NFT infrastructure
 * @notice Uses NFT tokenURI system to connect with metadata server, not as traditional NFTs
 * @author VFinance Team
 * @custom:security-contact security@vfinance.com.br
 * @custom:version 2.0.0
 * @custom:besu-optimized Optimized for Hyperledger Besu
 */
contract VFinanceRegistry is
    Initializable,
    ERC721Upgradeable,
    OwnableUpgradeable,
    UUPSUpgradeable
{
    // =============================================================
    //                          CONSTANTS
    // =============================================================

    /// @dev Contract version for upgrade control
    string public constant VERSION = "2.0.0";

    /// @dev Maximum contracts that can be registered (optimized for Besu)
    uint256 public constant MAX_SUPPLY = 10000000; // 10M contracts

    /// @dev Event prefix for traceability
    bytes32 private constant EVENT_PREFIX = keccak256("VFINANCE_REGISTRY_V2");

    // =============================================================
    //                           STRUCTS
    // =============================================================

    /**
     * @dev Main contract data (optimized for storage)
     */
    struct ContractRecord {
        bytes32 registryId; // Unique contract ID converted to bytes32 (gas savings)
        bytes32 contractNumber; // Contract number in bytes32
        uint32 contractDate; // Contract date timestamp (uint32 sufficient until 2106)
        bytes32 metadataHash; // SHA256 hash of metadata
        uint32 timestamp; // Registration timestamp (uint32)
        address registeredBy; // Address that registered
        bool active; // Active/inactive status
    }

    /**
     * @dev Critical vehicle data for traceability (packed for efficiency)
     */
    struct VehicleCore {
        bytes32 chassis; // Vehicle chassis in bytes32 (17 chars fit)
        bytes32 licensePlate; // License plate in bytes32 (7-8 chars)
        uint128 totalValue; // Total value in wei (uint128 supports up to ~340 trillion)
        uint64 brandId; // Brand ID (internal enum)
        uint64 modelId; // Model ID (internal enum)
    }

    // =============================================================
    //                        STATE VARIABLES
    // =============================================================

    /// @dev Mapping from tokenId to contract data (packed struct)
    mapping(uint256 => ContractRecord) public contracts;

    /// @dev Mapping from registryId hash to tokenId (bytes32 for efficiency)
    mapping(bytes32 => uint256) public registryIdHashToTokenId;

    /// @dev Mapping from metadata hash to tokenId
    mapping(bytes32 => uint256) public metadataHashToTokenId;

    /// @dev Mapping from tokenId to vehicle data (packed struct)
    mapping(uint256 => VehicleCore) public vehicleCores;

    /// @dev Mapping from chassis hash to tokenId (anti-duplication)
    mapping(bytes32 => uint256) public chassisHashToTokenId;

    /// @dev Current token counter (optimized)
    uint256 private _currentIndex;

    /// @dev Authorized API server address
    address public apiServerAddress;

    /// @dev Base URL for metadata (compact string)
    string public metadataBaseUrl;

    /// @dev Mapping for brands (ID -> name) - storage savings
    mapping(uint64 => string) public brands;
    uint64 private _brandIdCounter;

    /// @dev Mapping for models (ID -> name) - storage savings
    mapping(uint64 => string) public models;
    uint64 private _modelIdCounter;

    // =============================================================
    //                            EVENTS
    // =============================================================

    /**
     * @dev Main registration event (optimized for Besu indexing)
     */
    event ContractRegistered(
        uint256 indexed tokenId,
        bytes32 indexed registryIdHash,
        bytes32 indexed chassisHash,
        bytes32 metadataHash,
        uint32 timestamp
    );

    /**
     * @dev Vehicle traceability event
     */
    event VehicleTracked(
        uint256 indexed tokenId,
        bytes32 indexed chassis,
        bytes32 indexed licensePlate,
        uint128 totalValue,
        uint64 brandId,
        uint64 modelId
    );

    /**
     * @dev Metadata update event
     */
    event MetadataUpdated(
        uint256 indexed tokenId,
        bytes32 indexed oldHash,
        bytes32 indexed newHash,
        uint32 timestamp
    );

    /**
     * @dev Status change event
     */
    event StatusChanged(
        uint256 indexed tokenId,
        bool indexed active,
        uint32 timestamp
    );

    /**
     * @dev System configuration event
     */
    event SystemConfigured(address indexed apiServer, string metadataUrl);

    /**
     * @dev Brand registered event
     */
    event BrandRegistered(uint64 indexed brandId, string brandName);

    /**
     * @dev Model registered event
     */
    event ModelRegistered(uint64 indexed modelId, string modelName);

    // =============================================================
    //                            ERRORS
    // =============================================================

    error ContractAlreadyExists();
    error ContractNotFound();
    error InvalidInput();
    error UnauthorizedAccess();
    error DuplicateHash();
    error MaxSupplyExceeded();
    error TokenNotFound();

    // =============================================================
    //                           MODIFIERS
    // =============================================================

    /**
     * @dev Only API server or owner
     */
    modifier onlyAuthorized() {
        if (msg.sender != apiServerAddress && msg.sender != owner()) {
            revert UnauthorizedAccess();
        }
        _;
    }

    /**
     * @dev Check if token exists
     */
    modifier validToken(uint256 tokenId) {
        if (tokenId == 0 || tokenId >= _currentIndex) {
            revert TokenNotFound();
        }
        _;
    }

    // =============================================================
    //                        INITIALIZATION
    // =============================================================

    /**
     * @dev Initialize the contract (optimized for Besu)
     */
    function initialize(
        string memory name,
        string memory symbol,
        string memory _metadataBaseUrl,
        address _apiServerAddress
    ) public initializer {
        __ERC721_init(name, symbol);
        __Ownable_init(msg.sender);
        __UUPSUpgradeable_init();

        metadataBaseUrl = _metadataBaseUrl;
        apiServerAddress = _apiServerAddress;
        _currentIndex = 1; // Start from tokenId 1

        // Register common brands (future gas savings)
        _registerBrand("TOYOTA");
        _registerBrand("HONDA");
        _registerBrand("VOLKSWAGEN");
        _registerBrand("FIAT");
        _registerBrand("CHEVROLET");
        _registerBrand("FORD");
        _registerBrand("NISSAN");
        _registerBrand("HYUNDAI");
        _registerBrand("RENAULT");
        _registerBrand("PEUGEOT");

        emit SystemConfigured(_apiServerAddress, _metadataBaseUrl);
    }

    // =============================================================
    //                      CORE FUNCTIONALITY
    // =============================================================

    /**
     * @dev Register optimized contract (main function)
     * @param registryId Unique contract ID (string converted to bytes32)
     * @param contractNumber Contract number
     * @param contractDate Contract date timestamp
     * @param chassis Vehicle chassis
     * @param licensePlate Vehicle license plate
     * @param totalValue Contract value in wei
     * @param brandName Brand name (will be converted to ID)
     * @param modelName Model name (will be converted to ID)
     * @return tokenId Created registry token ID (for internal reference only)
     * @return metadataHash Generated metadata hash (connects to metadata server)
     */
    function registerContract(
        string calldata registryId,
        string calldata contractNumber,
        uint32 contractDate,
        string calldata chassis,
        string calldata licensePlate,
        uint128 totalValue,
        string calldata brandName,
        string calldata modelName
    ) external onlyAuthorized returns (uint256 tokenId, bytes32 metadataHash) {
        // Basic validations
        if (bytes(registryId).length == 0 || bytes(chassis).length == 0) {
            revert InvalidInput();
        }
        if (_currentIndex > MAX_SUPPLY) revert MaxSupplyExceeded();

        // Convert strings to bytes32 (gas savings)
        bytes32 registryIdHash = keccak256(bytes(registryId));
        bytes32 contractNumberBytes = _stringToBytes32(contractNumber);
        bytes32 chassisBytes = _stringToBytes32(chassis);
        bytes32 licensePlateBytes = _stringToBytes32(licensePlate);

        // Check for duplicates
        if (registryIdHashToTokenId[registryIdHash] != 0)
            revert ContractAlreadyExists();
        if (chassisHashToTokenId[keccak256(bytes(chassis))] != 0)
            revert ContractAlreadyExists();

        // Generate tokenId and metadata hash
        tokenId = _currentIndex++;
        metadataHash = keccak256(
            abi.encodePacked(
                registryIdHash,
                contractNumberBytes,
                contractDate,
                chassisBytes,
                tokenId,
                block.timestamp
            )
        );

        // Check unique hash
        if (metadataHashToTokenId[metadataHash] != 0) revert DuplicateHash();

        // Get brand and model IDs (register if doesn't exist)
        uint64 brandId = _getOrCreateBrandId(brandName);
        uint64 modelId = _getOrCreateModelId(modelName);

        // Mint registry token (non-transferable, used only for tokenURI system)
        _mint(address(this), tokenId);

        // Store data (packed structs for savings)
        contracts[tokenId] = ContractRecord({
            registryId: registryIdHash,
            contractNumber: contractNumberBytes,
            contractDate: contractDate,
            metadataHash: metadataHash,
            timestamp: uint32(block.timestamp),
            registeredBy: msg.sender,
            active: true
        });

        vehicleCores[tokenId] = VehicleCore({
            chassis: chassisBytes,
            licensePlate: licensePlateBytes,
            totalValue: totalValue,
            brandId: brandId,
            modelId: modelId
        });

        // Update mappings
        registryIdHashToTokenId[registryIdHash] = tokenId;
        metadataHashToTokenId[metadataHash] = tokenId;
        chassisHashToTokenId[keccak256(bytes(chassis))] = tokenId;

        // Emit optimized events
        emit ContractRegistered(
            tokenId,
            registryIdHash,
            keccak256(bytes(chassis)),
            metadataHash,
            uint32(block.timestamp)
        );

        emit VehicleTracked(
            tokenId,
            chassisBytes,
            licensePlateBytes,
            totalValue,
            brandId,
            modelId
        );

        return (tokenId, metadataHash);
    }

    /**
     * @dev Update metadata hash (optimized)
     */
    function updateMetadataHash(
        uint256 tokenId,
        bytes32 newMetadataHash
    ) external onlyAuthorized validToken(tokenId) {
        if (newMetadataHash == bytes32(0)) revert InvalidInput();
        if (metadataHashToTokenId[newMetadataHash] != 0) revert DuplicateHash();

        bytes32 oldHash = contracts[tokenId].metadataHash;

        // Update mappings
        delete metadataHashToTokenId[oldHash];
        contracts[tokenId].metadataHash = newMetadataHash;
        metadataHashToTokenId[newMetadataHash] = tokenId;

        emit MetadataUpdated(
            tokenId,
            oldHash,
            newMetadataHash,
            uint32(block.timestamp)
        );
    }

    /**
     * @dev Update contract status
     */
    function updateStatus(
        uint256 tokenId,
        bool active
    ) external onlyAuthorized validToken(tokenId) {
        contracts[tokenId].active = active;
        emit StatusChanged(tokenId, active, uint32(block.timestamp));
    }

    // =============================================================
    //                        VIEW FUNCTIONS
    // =============================================================

    /**
     * @dev Get complete contract data by tokenId
     */
    function getContract(
        uint256 tokenId
    )
        external
        view
        validToken(tokenId)
        returns (ContractRecord memory, VehicleCore memory)
    {
        return (contracts[tokenId], vehicleCores[tokenId]);
    }

    /**
     * @dev Get contract by registryId string
     */
    function getContractByRegistryId(
        string calldata registryId
    ) external view returns (ContractRecord memory, VehicleCore memory) {
        bytes32 registryIdHash = keccak256(bytes(registryId));
        uint256 tokenId = registryIdHashToTokenId[registryIdHash];
        if (tokenId == 0) revert ContractNotFound();
        return (contracts[tokenId], vehicleCores[tokenId]);
    }

    /**
     * @dev Get contract by metadata hash
     */
    function getContractByHash(
        bytes32 metadataHash
    ) external view returns (ContractRecord memory, VehicleCore memory) {
        uint256 tokenId = metadataHashToTokenId[metadataHash];
        if (tokenId == 0) revert ContractNotFound();
        return (contracts[tokenId], vehicleCores[tokenId]);
    }

    /**
     * @dev Get contract by chassis
     */
    function getContractByChassis(
        string calldata chassis
    ) external view returns (ContractRecord memory, VehicleCore memory) {
        bytes32 chassisHash = keccak256(bytes(chassis));
        uint256 tokenId = chassisHashToTokenId[chassisHash];
        if (tokenId == 0) revert ContractNotFound();
        return (contracts[tokenId], vehicleCores[tokenId]);
    }

    /**
     * @dev Get active contracts with pagination (optimized)
     */
    function getActiveContracts(
        uint256 offset,
        uint256 limit
    ) external view returns (uint256[] memory tokenIds) {
        if (limit == 0 || limit > 100) revert InvalidInput();

        uint256 totalTokens = _currentIndex - 1;
        uint256 activeCount = 0;

        // Count active contracts
        for (uint256 i = 1; i <= totalTokens; i++) {
            if (contracts[i].active) {
                activeCount++;
            }
        }

        if (offset >= activeCount) {
            return new uint256[](0);
        }

        uint256 resultSize = activeCount - offset;
        if (resultSize > limit) {
            resultSize = limit;
        }

        tokenIds = new uint256[](resultSize);
        uint256 currentIndex = 0;
        uint256 resultIndex = 0;

        for (uint256 i = 1; i <= totalTokens && resultIndex < resultSize; i++) {
            if (contracts[i].active) {
                if (currentIndex >= offset) {
                    tokenIds[resultIndex] = i;
                    resultIndex++;
                }
                currentIndex++;
            }
        }

        return tokenIds;
    }

    /**
     * @dev Get total supply
     */
    function totalSupply() public view returns (uint256) {
        return _currentIndex - 1;
    }

    /**
     * @dev Check if contract exists by registryId
     */
    function doesContractExist(
        string calldata registryId
    ) external view returns (bool) {
        bytes32 registryIdHash = keccak256(bytes(registryId));
        return registryIdHashToTokenId[registryIdHash] != 0;
    }

    /**
     * @dev Check if hash exists
     */
    function doesHashExist(bytes32 metadataHash) external view returns (bool) {
        return metadataHashToTokenId[metadataHash] != 0;
    }

    /**
     * @dev Get brand name by ID
     */
    function getBrandName(
        uint64 brandId
    ) external view returns (string memory) {
        return brands[brandId];
    }

    /**
     * @dev Get model name by ID
     */
    function getModelName(
        uint64 modelId
    ) external view returns (string memory) {
        return models[modelId];
    }

    // =============================================================
    //                    REGISTRY TOKEN FUNCTIONS
    // =============================================================

    /**
     * @dev Return token URI (metadata server endpoint)
     * @notice This connects to the metadata server, not for NFT marketplace use
     */
    function tokenURI(
        uint256 tokenId
    ) public view override validToken(tokenId) returns (string memory) {
        bytes32 metadataHash = contracts[tokenId].metadataHash;
        return
            string(
                abi.encodePacked(
                    metadataBaseUrl,
                    "/api/metadata/0x",
                    _toHexString(metadataHash)
                )
            );
    }

    /**
     * @dev Override update function to make tokens non-transferable
     * @notice These are registry entries, not tradeable NFTs
     */
    function _update(
        address to,
        uint256 tokenId,
        address auth
    ) internal virtual override returns (address) {
        address from = _ownerOf(tokenId);

        // Allow minting (from == address(0)) but block all transfers
        if (from != address(0) && to != address(0)) {
            revert("Registry tokens are non-transferable");
        }

        return super._update(to, tokenId, auth);
    }

    /**
     * @dev Override approve functions since tokens are non-transferable
     */
    function approve(address, uint256) public pure override {
        revert("Registry tokens are non-transferable");
    }

    function setApprovalForAll(address, bool) public pure override {
        revert("Registry tokens are non-transferable");
    }

    /**
     * @dev Get metadata URL directly by hash (more efficient than tokenURI)
     * @notice Primary method for accessing metadata server
     */
    function getMetadataUrl(
        bytes32 metadataHash
    ) external view returns (string memory) {
        return
            string(
                abi.encodePacked(
                    metadataBaseUrl,
                    "/api/metadata/0x",
                    _toHexString(metadataHash)
                )
            );
    }

    /**
     * @dev Get metadata URL by registry ID
     */
    function getMetadataUrlByRegistryId(
        string calldata registryId
    ) external view returns (string memory) {
        bytes32 registryIdHash = keccak256(bytes(registryId));
        uint256 tokenId = registryIdHashToTokenId[registryIdHash];
        if (tokenId == 0) revert ContractNotFound();

        bytes32 metadataHash = contracts[tokenId].metadataHash;
        return
            string(
                abi.encodePacked(
                    metadataBaseUrl,
                    "/api/metadata/0x",
                    _toHexString(metadataHash)
                )
            );
    }

    // =============================================================
    //                      ADMIN FUNCTIONS
    // =============================================================

    /**
     * @dev Update server configuration
     */
    function updateServerConfig(
        string memory newMetadataBaseUrl,
        address newApiServerAddress
    ) external onlyOwner {
        metadataBaseUrl = newMetadataBaseUrl;
        apiServerAddress = newApiServerAddress;

        emit SystemConfigured(newApiServerAddress, newMetadataBaseUrl);
    }

    /**
     * @dev Register new brand manually
     */
    function registerBrand(
        string calldata brandName
    ) external onlyOwner returns (uint64 brandId) {
        return _registerBrand(brandName);
    }

    /**
     * @dev Register new model manually
     */
    function registerModel(
        string calldata modelName
    ) external onlyOwner returns (uint64 modelId) {
        return _registerModel(modelName);
    }

    // =============================================================
    //                      INTERNAL FUNCTIONS
    // =============================================================

    /**
     * @dev Convert string to bytes32 (up to 32 chars)
     */
    function _stringToBytes32(
        string memory source
    ) internal pure returns (bytes32 result) {
        bytes memory tempEmptyStringTest = bytes(source);
        if (tempEmptyStringTest.length == 0) {
            return 0x0;
        }
        assembly {
            result := mload(add(source, 32))
        }
    }

    /**
     * @dev Convert bytes32 to hex string
     */
    function _toHexString(bytes32 hash) internal pure returns (string memory) {
        bytes memory buffer = new bytes(64);
        for (uint256 i = 0; i < 32; i++) {
            buffer[i * 2] = _toHexChar(uint8(hash[i]) / 16);
            buffer[i * 2 + 1] = _toHexChar(uint8(hash[i]) % 16);
        }
        return string(buffer);
    }

    /**
     * @dev Convert digit to hex char
     */
    function _toHexChar(uint8 digit) internal pure returns (bytes1) {
        return
            bytes1(
                digit < 10
                    ? uint8(bytes1("0")) + digit
                    : uint8(bytes1("a")) + digit - 10
            );
    }

    /**
     * @dev Register new brand internally
     */
    function _registerBrand(
        string memory brandName
    ) internal returns (uint64 brandId) {
        brandId = ++_brandIdCounter;
        brands[brandId] = brandName;
        emit BrandRegistered(brandId, brandName);
        return brandId;
    }

    /**
     * @dev Register new model internally
     */
    function _registerModel(
        string memory modelName
    ) internal returns (uint64 modelId) {
        modelId = ++_modelIdCounter;
        models[modelId] = modelName;
        emit ModelRegistered(modelId, modelName);
        return modelId;
    }

    /**
     * @dev Get or create brand ID
     */
    function _getOrCreateBrandId(
        string calldata brandName
    ) internal returns (uint64) {
        // Simple linear search for now (could be optimized with mapping)
        for (uint64 i = 1; i <= _brandIdCounter; i++) {
            if (keccak256(bytes(brands[i])) == keccak256(bytes(brandName))) {
                return i;
            }
        }
        return _registerBrand(brandName);
    }

    /**
     * @dev Get or create model ID
     */
    function _getOrCreateModelId(
        string calldata modelName
    ) internal returns (uint64) {
        // Simple linear search for now (could be optimized with mapping)
        for (uint64 i = 1; i <= _modelIdCounter; i++) {
            if (keccak256(bytes(models[i])) == keccak256(bytes(modelName))) {
                return i;
            }
        }
        return _registerModel(modelName);
    }

    // =============================================================
    //                        UUPS UPGRADE
    // =============================================================

    /**
     * @dev Authorize upgrade (owner only)
     */
    function _authorizeUpgrade(
        address newImplementation
    ) internal override onlyOwner {}

    /**
     * @dev Get current contract version
     */
    function getVersion() external pure returns (string memory) {
        return VERSION;
    }

    // =============================================================
    //                      EMERGENCY FUNCTIONS
    // =============================================================

    /**
     * @dev Emergency function to recover ERC20 tokens
     */
    function emergencyWithdrawToken(
        address token,
        uint256 amount
    ) external onlyOwner {
        // Implement if needed for accidentally sent ERC20 tokens
    }

    /**
     * @dev Receive ETH directly
     */
    receive() external payable {
        // Accept ETH
    }
}
