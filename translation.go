package main

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/translate"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func main() {
	r := gin.Default()
	r.GET("/translation", func(c *gin.Context) {
		english, _ := translateText("en", c.Query("term"))
		c.JSON(http.StatusOK, gin.H{
			"text": english,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}
