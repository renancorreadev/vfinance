package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	jwtSecret string
}

func NewAuthHandler(jwtSecret string) *AuthHandler {
	return &AuthHandler{jwtSecret: jwtSecret}
}

// @Summary Generate JWT Token
// @Description Gera um token JWT válido para autenticação
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} object{token=string,success=boolean,message=string} "Token gerado com sucesso"
// @Failure 500 {object} object{valid=boolean,success=boolean,message=string,error=string} "Erro interno do servidor"
// @Router /api/auth/token [post]
func (h *AuthHandler) GenerateToken(c *gin.Context) {
	// Para simplificar, vamos gerar um token válido por 24 horas
	claims := jwt.MapClaims{
		"user_id": "admin",
		"role":    "admin",
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"valid":   false,
			"success": false,
			"message": "Erro ao gerar token",
			"error":   "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      tokenString,
		"success":    true,
		"message":    "Token gerado com sucesso",
	})
}

// @Summary Validate JWT Token
// @Description Valida um token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} object{valid=boolean,success=boolean,message=string,user_id=string,role=string,exp=integer} "Token válido"
// @Failure 401 {object} object{valid=boolean,success=boolean,message=string,error=string} "Token inválido ou ausente"
// @Router /api/auth/validate [get]
func (h *AuthHandler) ValidateToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"valid":   false,
			"success": false,
			"message": "Token de autorização necessário",
			"error":   "Unauthorized",
		})
		return
	}

	// Verifica se o header começa com "Bearer " e extrai o token
	if len(authHeader) < 7 || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"valid":   false,
			"success": false,
			"message": "Formato de autorização inválido. Use 'Bearer <token>'",
			"error":   "Invalid Authorization Format",
		})
		return
	}

	tokenString := authHeader[7:] // Remove "Bearer "

	// Verifica se o token não está vazio
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"valid":   false,
			"success": false,
			"message": "Token não pode estar vazio",
			"error":   "Empty Token",
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica se o método de assinatura é o esperado
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(h.jwtSecret), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"valid":   false,
			"success": false,
			"message": "Token inválido: " + err.Error(),
			"error":   "Invalid Token",
		})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"valid":   false,
			"success": false,
			"message": "Token expirado ou inválido",
			"error":   "Token Expired or Invalid",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"valid":   false,
			"success": false,
			"message": "Token inválido",
			"error":   "Invalid Claims",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"valid":   true,
		"success": true,
		"message": "Token válido",
		"user_id": claims["user_id"],
		"role":    claims["role"],
		"exp":     claims["exp"],
	})
}
