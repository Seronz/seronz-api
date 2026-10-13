package email

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

type Request struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}

func Send(c *gin.Context) {
	var req Request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user := os.Getenv("EMAIL_USER")
	pass := os.Getenv("EMAIL_PASS")

	m := gomail.NewMessage()
	m.SetHeader("From", user)
	m.SetHeader("To", user)
	m.SetHeader("Replay-To", req.Email)
	m.SetHeader("Subject", "Portofolio : "+req.Name)
	m.SetBody("text/html", "<h3>Daeri : "+req.Name+" ("+req.Email+")</h3><p>"+req.Message+"</p>")

	d := gomail.NewDialer("smtp.gmail.com", 587, user, pass)

	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "gagal mengirim pesan"})
		return
	}

	c.JSON(200, gin.H{"success": true})

}
