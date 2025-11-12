package diccionario

import (
	"net/http"

	"github.com/for-the-kidz/diccionario/wordlist"
	"github.com/gin-gonic/gin"
)

// Server abstraction.
type Server struct {
	r *gin.Engine
	w wordlist.WordList
}

// NewServer instantiates a new HTTP Server.
func NewServer() (s *Server) {
	s = &Server{
		r: gin.Default(),
		w: wordlist.New("/words.txt"),
	}

	s.r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	s.r.GET("/exists/:word", s.WordExists)
	s.r.POST("/add", s.Add)
	s.r.GET("/matches/:prefix", s.Matches)

	return
}

// Run the HTTP server. It will block until a fatal error is encountered.
func (s *Server) Run(addr string) error {
	return s.r.Run(addr)
}
