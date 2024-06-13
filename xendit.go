package main

import (
	"context"
	"net/http"

	xendit "github.com/xendit/xendit-go/v5"
	"github.com/xendit/xendit-go/v5/invoice"
)

var (
	ApiKey             = "XENDIT_API_KEY"
	SuccessRedirectURL = "https://www.google.com"
)

func CreateInvoice() (*invoice.Invoice, *http.Response, error) {
	// Initialize a new item
	item := invoice.InvoiceItem{
		Name:     "Activity Futsal",
		Price:    10000,
		Quantity: 1,
	}

	// Create a slice of InvoiceItem
	items := []invoice.InvoiceItem{item}

	// Create a new invoice request
	createInvoiceRequest := invoice.NewCreateInvoiceRequest("jsfsdfsdfs_ID", 10000)
	createInvoiceRequest.SuccessRedirectUrl = &SuccessRedirectURL
	createInvoiceRequest.Items = items
	createInvoiceRequest.Customer = invoice.NewCustomerObject()
	createInvoiceRequest.Customer.SetEmail("adisuryo22@gmail.com")
	shouldSendEmail := true
	createInvoiceRequest.ShouldSendEmail = &shouldSendEmail

	xenditClient := xendit.NewClient(ApiKey)

	resp, r, err := xenditClient.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(*createInvoiceRequest).
		Execute()

	if err != nil {
		return nil, r, err
	}

	return resp, r, nil
}

func GetInvoices() ([]invoice.Invoice, *http.Response, error) {
	client := xendit.NewClient(ApiKey)

	resp, r, err := client.InvoiceApi.GetInvoices(context.Background()).Execute()
	if err != nil {
		return nil, r, err
	}

	return resp, r, nil
}

func GetInvoiceByID(invoiceID string) (*invoice.Invoice, *http.Response, error) {
	client := xendit.NewClient(ApiKey)

	resp, r, err := client.InvoiceApi.GetInvoiceById(context.Background(), invoiceID).Execute()
	if err != nil {
		return nil, r, err
	}

	return resp, r, nil
}
