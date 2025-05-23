package httpdriver

import (
	"github.com/samber/lo"

	"github.com/openmeterio/openmeter/api"
	"github.com/openmeterio/openmeter/openmeter/customer"
	"github.com/openmeterio/openmeter/pkg/currencyx"
	"github.com/openmeterio/openmeter/pkg/models"
)

func MapCustomerCreate(body api.CustomerCreate) customer.CustomerMutate {
	return customer.CustomerMutate{
		Key:              body.Key,
		Name:             body.Name,
		Description:      body.Description,
		UsageAttribution: customer.CustomerUsageAttribution(body.UsageAttribution),
		PrimaryEmail:     body.PrimaryEmail,
		BillingAddress:   MapAddress(body.BillingAddress),
		Currency:         mapCurrency(body.Currency),
	}
}

func mapCurrency(apiCurrency *string) *currencyx.Code {
	if apiCurrency == nil {
		return nil
	}

	return lo.ToPtr(currencyx.Code(*apiCurrency))
}

func MapAddress(apiAddress *api.Address) *models.Address {
	if apiAddress == nil {
		return nil
	}

	address := models.Address{
		City:        apiAddress.City,
		State:       apiAddress.State,
		PostalCode:  apiAddress.PostalCode,
		Line1:       apiAddress.Line1,
		Line2:       apiAddress.Line2,
		PhoneNumber: apiAddress.PhoneNumber,
	}

	if apiAddress.Country != nil {
		address.Country = lo.ToPtr(models.CountryCode(*apiAddress.Country))
	}

	return &address
}

// CustomerToAPI converts a Customer to an API Customer
func CustomerToAPI(c customer.Customer) (api.Customer, error) {
	apiCustomer := api.Customer{
		Id:                    c.ManagedResource.ID,
		Key:                   c.Key,
		Name:                  c.Name,
		UsageAttribution:      api.CustomerUsageAttribution{SubjectKeys: c.UsageAttribution.SubjectKeys},
		PrimaryEmail:          c.PrimaryEmail,
		Description:           c.Description,
		CreatedAt:             c.CreatedAt,
		UpdatedAt:             c.UpdatedAt,
		DeletedAt:             c.DeletedAt,
		CurrentSubscriptionId: c.CurrentSubscriptionID,
	}

	if c.BillingAddress != nil {
		address := api.Address{
			City:        c.BillingAddress.City,
			State:       c.BillingAddress.State,
			PostalCode:  c.BillingAddress.PostalCode,
			Line1:       c.BillingAddress.Line1,
			Line2:       c.BillingAddress.Line2,
			PhoneNumber: c.BillingAddress.PhoneNumber,
		}

		if c.BillingAddress.Country != nil {
			address.Country = lo.ToPtr(string(*c.BillingAddress.Country))
		}

		apiCustomer.BillingAddress = &address
	}

	if c.Currency != nil {
		apiCustomer.Currency = lo.ToPtr(string(*c.Currency))
	}

	return apiCustomer, nil
}
