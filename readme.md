# golang pustaka api

last in minute 2:19:17

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "sakura endo",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Name":    "sakura endo",
		"Age":     20,
		"Address": "Tokyo",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(200, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(200, gin.H{
		"title": title,
		"price": price,
	})

}