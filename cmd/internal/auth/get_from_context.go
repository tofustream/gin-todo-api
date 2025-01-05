package auth

import (
    "github.com/gin-gonic/gin"
)

const accountIDKey = "accountID"

// context から account id を取得する
func GetAccountIDFromContext(ctx *gin.Context) (string, bool) {
    maybeAccountID, exists := ctx.Get(accountIDKey)
    if !exists {
        return "", false
    }
    return maybeAccountID.(string), true
}
