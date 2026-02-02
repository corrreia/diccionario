package diccionario

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExistsResponse is the response sent back for the exists endpoint.
type ExistsResponse struct {
	// Exists is true if the word exists; otherwise, false.
	Exists bool `json:"exists"`
}

// WordExists returns true if the word exists in the word list.
// It performs case insensitive matching to the words in the wordlist.
func (s *Server) WordExists(c *gin.Context) {
	word := c.Param("word")

	valid, err := ValidateWord(word)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if !valid {
		c.String(http.StatusBadRequest, "invalid word")
		return
	}

	exists, err := s.w.WordExists(word)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ExistsResponse{Exists: exists})
}
