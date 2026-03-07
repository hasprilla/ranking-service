package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hasprilla/ranking-service/models"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetArtistRanking(c *fiber.Ctx) error {
	// In a real scenario, we'd call the artist-service
	// For this migration, we'll implement the aggregation logic
	resp, err := httpClient.Get("http://artist-service:8080/api/v1/artists/public")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Artist service unavailable"})
	}
	defer resp.Body.Close()

	var artists []models.ArtistRanking
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode artists"})
	}

	// Sort by stars descending
	sort.Slice(artists, func(i, j int) bool {
		return artists[i].CurrentStars > artists[j].CurrentStars
	})

	// Take top 10
	limit := 10
	if len(artists) < 10 {
		limit = len(artists)
	}

	return c.JSON(artists[:limit])
}

func GetFanRanking(c *fiber.Ctx) error {
	resp, err := httpClient.Get("http://wallet-service:8080/api/v1/wallet/top-donors")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Wallet service unavailable"})
	}
	defer resp.Body.Close()

	var fans []models.FanRanking
	if err := json.NewDecoder(resp.Body).Decode(&fans); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode fans"})
	}

	return c.JSON(fans)
}
