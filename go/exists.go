package diccionario

import (
	"log"
	"net/http"
	"strings"

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

	log.Println("checking if word exists:", word)

	wordlist, err := s.w.GetWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{Err: err, Desc: "unable to retrieve word list"})
		return
	}

	resp := ExistsResponse{Exists: false}

	for _, w := range wordlist {
		if strings.HasPrefix(w, word) {
			resp.Exists = true
		}
	}

	c.JSON(http.StatusOK, resp)
}
