package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	xendit "github.com/xendit/xendit-go/v5"
	"github.com/xendit/xendit-go/v5/invoice"
	"github.com/xendit/xendit-go/v5/payout"
)

var (
	ApiKey             = "XENDIT_API_KEY"
	SuccessRedirectURL = "https://www.google.com"
	client             = xendit.NewClient(ApiKey)
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

	resp, r, err := client.InvoiceApi.CreateInvoice(context.Background()).
		CreateInvoiceRequest(*createInvoiceRequest).
		Execute()

	if err != nil {
		return nil, r, err
	}

	return resp, r, nil
}

func GetInvoices() ([]invoice.Invoice, *http.Response, error) {
	resp, r, err := client.InvoiceApi.GetInvoices(context.Background()).Execute()
	if err != nil {
		return nil, r, err
	}

	return resp, r, nil
}

func GetInvoiceByID(invoiceID string) (*invoice.Invoice, *http.Response, error) {
	resp, r, err := client.InvoiceApi.GetInvoiceById(context.Background(), invoiceID).Execute()
	if err != nil {
		return nil, r, err
	}

	return resp, r, nil
}

func CreatePayout() (*payout.GetPayouts200ResponseDataInner, *http.Response, error) {
	idempotencyKey := "f819758c-8be6-4e13-ac6d-1b48a5184b26"

	createPayoutRequest := *payout.NewCreatePayoutRequest(
		"DISB-001",
		"ID_BCA",
		*payout.NewDigitalPayoutChannelProperties("1231314342"),
		float32(15000),
		"IDR",
	)

	createPayoutRequest.ChannelProperties.SetAccountHolderName("Ruben Adi")

	resp, r, err := client.PayoutApi.CreatePayout(context.Background()).
		IdempotencyKey(idempotencyKey).
		CreatePayoutRequest(createPayoutRequest).
		Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.CreatePayout``: %v\n", err.Error())

		b, _ := json.Marshal(err.FullError())
		fmt.Fprintf(os.Stderr, "Full Error Struct: %v\n", string(b))

		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return nil, r, err
	}

	fmt.Println(resp)
	return resp, r, nil
}
