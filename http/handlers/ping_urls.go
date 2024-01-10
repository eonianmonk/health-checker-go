package handlers

import (
	"healthchecker/http"
	"healthchecker/http/requests"
	"healthchecker/http/responses"

	"github.com/gofiber/fiber/v2"
)

// GET
func PingURLs(c *fiber.Ctx) error {
	log := http.GetLogger(c)
	var urls requests.UrlsRequest
	err := c.BodyParser(urls)
	if err != nil {
		errr := responses.ErrorResponse{Error: "failed to read urls from request body"}
		log.Errorf("failed to parse request body: %s", err.Error())
		c.Status(fiber.StatusBadRequest).JSON(errr)
		return nil
	}
	err = urls.VerifyURLs() 
	if err != nil {
		errr := responses.ErrorResponse{Error: err.Error()}
		log.Errorf("%s", err.Error())
		c.Status(fiber.StatusBadRequest).JSON(errr)
		return nil
	} 
	//pinger := http.GetPinger(c)
	//statuses, err := pinger.PingEm(urls.Urls)
	return nil
}