package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) Register(router *gin.Engine) {
	router.POST("/user", h.CreateUser)
	router.GET("/user", h.GetUser)
	router.DELETE("/user", h.DeleteUser)
	router.POST("/book", h.CreateBook)
	router.GET("/book", h.SearchBook)
	router.POST("/book/buy", h.TransactionUserToUser)
	router.GET("/book/buy", h.GetUsersTransactions)
}
