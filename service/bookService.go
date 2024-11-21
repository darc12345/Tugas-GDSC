package service

import (
	"main/model"

	"github.com/gin-gonic/gin"
)

func (h *ServiceDB) GetBooksService(c *gin.Context) ([]model.Book, error) {
	return h.RepoDB.GetBooksRepo(c)

}
func (h *ServiceDB) GetBookByIDService(c *gin.Context) (model.Book, error) {
	return h.RepoDB.GetBookByIDRepo(c)

}
func (h *ServiceDB) PostBooksService(c *gin.Context) (model.Book, error) {
	return h.RepoDB.PostBooksRepo(c)
}
func (h *ServiceDB) DeleteBooksService(c *gin.Context) error {
	return h.RepoDB.DeleteBooksRepo(c)
}
func (h *ServiceDB) PutBooksService(c *gin.Context) (model.Book, error) {
	return h.RepoDB.PutBooksRepo(c)

}
