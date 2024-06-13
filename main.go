package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		res, r, err := CreateInvoice()
		if err != nil {
			return c.Status(r.StatusCode).JSON(fiber.Map{
				"http_response": r.StatusCode,
				"error":         err.Error(),
			})
		}

		return c.Status(r.StatusCode).JSON(fiber.Map{
			"http_response": r.StatusCode,
			"response":      res,
		})
	})

	app.Get("/invoices", func(c *fiber.Ctx) error {
		res, r, err := GetInvoices()
		if err != nil {
			return c.Status(r.StatusCode).JSON(fiber.Map{
				"http_response": r.StatusCode,
				"error":         err.Error(),
			})
		}

		return c.Status(r.StatusCode).JSON(fiber.Map{
			"http_response": r.StatusCode,
			"response":      res,
		})
	})

	app.Get("/invoices/:id", func(c *fiber.Ctx) error {
		invoiceID := c.Params("id")
		resp, r, err := GetInvoiceByID(invoiceID)
		if err != nil {
			return c.Status(r.StatusCode).JSON(fiber.Map{
				"http_response": r.StatusCode,
				"error":         err.Error(),
			})
		}

		return c.Status(r.StatusCode).JSON(fiber.Map{
			"http_response": r.StatusCode,
			"response":      resp,
		})
	})

	app.Post("/payout", func(c *fiber.Ctx) error {
		resp, r, err := CreatePayout()
		if err != nil {
			return c.Status(r.StatusCode).JSON(fiber.Map{
				"http_response": r.StatusCode,
				"error":         err.Error(),
			})
		}

		return c.Status(r.StatusCode).JSON(fiber.Map{
			"http_response": r.StatusCode,
			"response":      resp,
		})
	})

	app.Listen(":8002")
}
