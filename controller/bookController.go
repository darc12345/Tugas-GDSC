package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *ControllerDB) GetBooksHandler(c *gin.Context) {
	books, err := h.serviceDB.GetBooksService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getBooksStructure{books})

}
func (h *ControllerDB) GetBookByIDHandler(c *gin.Context) {
	book, err := h.serviceDB.GetBookByIDService(c)
	if err != nil {
		c.JSON(http.StatusNotFound, messageStructure{err.Error()})
		return
	}
	c.JSON(http.StatusOK, getBookStructure{book})
}
func (h *ControllerDB) PostBooksHandler(c *gin.Context) {
	book, err := h.serviceDB.PostBooksService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusCreated, postBookStructure{"books inserted", book})
}
func (h *ControllerDB) DeleteBooksHandler(c *gin.Context) {
	err := h.serviceDB.DeleteBooksService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, deleteBookStructure{"Book deleted"})
}
func (h *ControllerDB) PutBooksHandler(c *gin.Context) {
	book, err := h.serviceDB.PutBooksService(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, messageStructure{err.Error()})
		return
	}
	c.JSON(http.StatusOK, putBookStructure{"data updated", book})
}
