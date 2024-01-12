package handlers

import (
	"encoding/json"
	"healthchecker/http/middleware"
	"healthchecker/http/requests"
	"healthchecker/http/responses"

	"github.com/gofiber/fiber/v2"
)

// GET
func PingURLs(c *fiber.Ctx) error {
	log := middleware.GetLogger(c)
	log.Printf("received ping request: %s", string(c.Body()))

	var urls requests.UrlsRequest
	err := json.Unmarshal(c.Body(),&urls.Urls)
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
	pinger := middleware.GetPinger(c)
	statuses, err := pinger.PingEm(urls.Urls)
	if err != nil {
		errr := responses.ErrorResponse{Error: err.Error()}
		log.Errorf("%s", err.Error())
		c.Status(fiber.StatusBadRequest).JSON(errr)
		return nil
	}
	c.Status(fiber.StatusOK).JSON(statuses)
	return nil
}