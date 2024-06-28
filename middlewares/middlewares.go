package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

var client *auth.Client

func InitFirebase(jsonFilePath string) {

	opt := option.WithCredentialsFile(jsonFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
}

func verifyToken(idToken string) (*auth.Token, error) {
	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Authorizationヘッダーからトークンを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			log.Println("error: Unauthorized because of authHeader None")
			c.Abort()
			return
		}

		idToken := authHeader[len("Bearer "):]

		// トークンを検証してユーザーを認証
		token, err := verifyToken(idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			errMsg := fmt.Sprintf("error: Unauthorized : %v", err)
			log.Println(errMsg)
			c.Abort()
			return
		}

		// コンテキストにトークンを保存して後続の処理で利用可能にする
		c.Set("token", token)

		// 次のミドルウェアまたはハンドラー関数を呼び出す
		c.Next()
	}
}
