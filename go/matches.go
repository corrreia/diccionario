package diccionario

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// MatchesResponse is the response sent back for the prefix matches endpoint.
type MatchesResponse struct {
	// Exists is true if the word exists; otherwise, false.
	Matches []string `json:"matches"`
}

// Matches returns a list of words that matched the given prefix.
// It performs case insensitive matching to the words in the wordlist.
func (s *Server) Matches(c *gin.Context) {
	prefix := strings.ToLower(c.Param("prefix"))

	log.Println("checking for matches for prefix:", prefix)

	wordlist, err := s.w.GetWords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ApiError{Err: err, Desc: "unable to retrieve word list"})
		return
	}

	resp := MatchesResponse{Matches: make([]string, 0)}

	for _, w := range wordlist {
		if strings.HasPrefix(strings.ToLower(w), prefix) {
			resp.Matches = append(resp.Matches, w)
		}
	}

	c.JSON(http.StatusOK, resp)
}
