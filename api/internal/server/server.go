package server

import (
	"vfinance-api/docs"
	"vfinance-api/internal/blockchain"
	"vfinance-api/internal/config"
	"vfinance-api/internal/handlers"
	"vfinance-api/internal/middleware"
	"vfinance-api/internal/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	config *config.Config
	db     *gorm.DB
	router *gin.Engine
}

func New(cfg *config.Config, db *gorm.DB) *Server {
	return &Server{
		config: cfg,
		db:     db,
		router: gin.Default(),
	}
}

func (s *Server) setupRoutes() error {
	// Inicializar cliente blockchain
	blockchainClient, err := blockchain.NewClient(
		s.config.EthereumRPC,
		s.config.ContractAddress,
		s.config.PrivateKey,
		1337, // ChainID padrão para Besu local
	)
	if err != nil {
		return err
	}

	// Inicializar serviços
	metadataService := services.NewMetadataService(s.db)
	contractService := services.NewContractService(s.db, blockchainClient, metadataService)

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler(s.config.JWTSecret)
	metadataHandler := handlers.NewMetadataHandler(metadataService)
	contractHandler := handlers.NewContractHandler(contractService)

	// Middleware global
	s.router.Use(middleware.RateLimit())
	s.router.Use(gin.Recovery())

	// Rotas da API
	api := s.router.Group("/api")
	{
		// Rotas de autenticação
		auth := api.Group("/auth")
		{
			auth.POST("/token", authHandler.GenerateToken)
			auth.GET("/validate", authHandler.ValidateToken)
		}

		// Rotas de metadados
		metadata := api.Group("/metadata")
		{
			metadata.POST("/:hash", middleware.JWTAuth(s.config.JWTSecret), metadataHandler.StoreMetadata)
			metadata.GET("/:hash", metadataHandler.GetMetadata)
			metadata.PUT("/:hash", middleware.JWTAuth(s.config.JWTSecret), metadataHandler.UpdateMetadata)
			metadata.DELETE("/:hash", middleware.JWTAuth(s.config.JWTSecret), metadataHandler.DeleteMetadata)
		}

		// Rotas de contratos
		contracts := api.Group("/contracts")
		{
			contracts.POST("/", middleware.JWTAuth(s.config.JWTSecret), contractHandler.RegisterContract)
			contracts.GET("/:regConId", contractHandler.GetContract)
			contracts.GET("/active", contractHandler.GetActiveContracts)
			contracts.GET("/hash/:hash", contractHandler.GetContractByHash)
			contracts.GET("/stats", contractHandler.GetStats)
			contracts.GET("/token/:tokenId", contractHandler.GetContractByTokenId)
			contracts.GET("/chassis/:chassis", contractHandler.GetContractByChassis)
			contracts.GET("/metadata-url/:hash", contractHandler.GetMetadataUrl)
			contracts.GET("/metadata-url/registry/:registryId", contractHandler.GetMetadataUrlByRegistryId)
		}
	}

	// Rota de health check
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Configurar Swagger
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = "localhost:" + s.config.Port
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Rotas do Swagger
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	s.router.GET("/docs", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	return nil
}

func (s *Server) Start() error {
	if err := s.setupRoutes(); err != nil {
		return err
	}

	return s.router.Run(":" + s.config.Port)
}
