package billing

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/alpacahq/alpacadecimal"
	"github.com/invopop/gobl/currency"
	"github.com/samber/lo"
	"github.com/samber/mo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	appsandbox "github.com/openmeterio/openmeter/openmeter/app/sandbox"
	"github.com/openmeterio/openmeter/openmeter/billing"
	billingadapter "github.com/openmeterio/openmeter/openmeter/billing/adapter"
	billingservice "github.com/openmeterio/openmeter/openmeter/billing/service"
	"github.com/openmeterio/openmeter/openmeter/billing/service/lineservice"
	"github.com/openmeterio/openmeter/openmeter/customer"
	"github.com/openmeterio/openmeter/openmeter/meter"
	"github.com/openmeterio/openmeter/openmeter/productcatalog"
	"github.com/openmeterio/openmeter/openmeter/productcatalog/feature"
	"github.com/openmeterio/openmeter/pkg/clock"
	"github.com/openmeterio/openmeter/pkg/currencyx"
	"github.com/openmeterio/openmeter/pkg/isodate"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
)

type InvoicingTestSuite struct {
	BaseSuite
}

func TestInvoicing(t *testing.T) {
	suite.Run(t, new(InvoicingTestSuite))
}

func (s *InvoicingTestSuite) TestPendingLineCreation() {
	namespace := "ns-create-invoice-workflow"
	now := time.Now().Truncate(time.Microsecond).In(time.UTC)
	periodEnd := now.Add(-time.Hour)
	periodStart := periodEnd.Add(-time.Hour * 24 * 30)
	issueAt := now.Add(-time.Minute)

	_ = s.InstallSandboxApp(s.T(), namespace)

	ctx := context.Background()

	// Given we have a test customer

	customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
		Namespace: namespace,

		CustomerMutate: customer.CustomerMutate{
			Name:         "Test Customer",
			PrimaryEmail: lo.ToPtr("test@test.com"),
			BillingAddress: &models.Address{
				Country:     lo.ToPtr(models.CountryCode("US")),
				PostalCode:  lo.ToPtr("12345"),
				State:       lo.ToPtr("NY"),
				City:        lo.ToPtr("New York"),
				Line1:       lo.ToPtr("1234 Test St"),
				Line2:       lo.ToPtr("Apt 1"),
				PhoneNumber: lo.ToPtr("1234567890"),
			},
			Currency: lo.ToPtr(currencyx.Code(currency.USD)),
			UsageAttribution: customer.CustomerUsageAttribution{
				SubjectKeys: []string{"test"},
			},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), customerEntity)
	require.NotEmpty(s.T(), customerEntity.ID)

	err = s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{
		{
			Namespace:     namespace,
			Slug:          "test",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
	})
	require.NoError(s.T(), err, "meter adapter should be able to replace meters")

	defer func() {
		err = s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{})
		require.NoError(s.T(), err, "meter adapter should be able to replace meters")
	}()

	_, err = s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
		Namespace: namespace,
		Name:      "test",
		Key:       "test",
		MeterSlug: lo.ToPtr("test"),
	})
	require.NoError(s.T(), err)

	// Given we have a default profile for the namespace

	var billingProfile billing.Profile
	s.T().Run("create default profile", func(t *testing.T) {
		minimalCreateProfileInput := MinimalCreateProfileInputTemplate
		minimalCreateProfileInput.Namespace = namespace

		profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

		require.NoError(t, err)
		require.NotNil(t, profile)
		billingProfile = *profile
	})

	var items []*billing.Line
	var HUFItem *billing.Line

	s.T().Run("CreateInvoiceItems", func(t *testing.T) {
		// When we create invoice items

		res, err := s.BillingService.CreatePendingInvoiceLines(ctx,
			billing.CreateInvoiceLinesInput{
				Namespace: namespace,
				Lines: []billing.LineWithCustomer{
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Namespace: namespace,
								Period:    billing.Period{Start: periodStart, End: periodEnd},

								InvoiceAt: issueAt,
								ManagedBy: billing.ManuallyManagedLine,

								Type: billing.InvoiceLineTypeFee,

								Name:     "Test item - USD",
								Currency: currencyx.Code(currency.USD),

								Metadata: map[string]string{
									"key": "value",
								},
							},
							FlatFee: &billing.FlatFeeLine{
								PerUnitAmount: alpacadecimal.NewFromFloat(100),
								Quantity:      alpacadecimal.NewFromFloat(1),
								Category:      billing.FlatFeeCategoryRegular,
								PaymentTerm:   productcatalog.InAdvancePaymentTerm,
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period: billing.Period{Start: periodStart, End: periodEnd},

								InvoiceAt: issueAt,
								ManagedBy: billing.ManuallyManagedLine,

								Type: billing.InvoiceLineTypeFee,

								Name:     "Test item - HUF",
								Currency: currencyx.Code(currency.HUF),
							},
							FlatFee: &billing.FlatFeeLine{
								PerUnitAmount: alpacadecimal.NewFromFloat(200),
								Quantity:      alpacadecimal.NewFromFloat(3),
								Category:      billing.FlatFeeCategoryRegular,
								PaymentTerm:   productcatalog.InAdvancePaymentTerm,
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period: billing.Period{Start: periodStart, End: periodEnd},

								InvoiceAt: issueAt,
								ManagedBy: billing.ManuallyManagedLine,

								Type: billing.InvoiceLineTypeUsageBased,

								Name:     "Test item - HUF",
								Currency: currencyx.Code(currency.HUF),
							},
							UsageBased: &billing.UsageBasedLine{
								Price: productcatalog.NewPriceFrom(productcatalog.TieredPrice{
									Mode: productcatalog.GraduatedTieredPrice,
									Tiers: []productcatalog.PriceTier{
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(100)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(10),
											},
										},
										{
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(100),
											},
										},
									},
								}),
								FeatureKey: "test",
							},
						},
						CustomerID: customerEntity.ID,
					},
				},
			})

		// Then we should have the items created
		require.NoError(s.T(), err)
		items = res

		// Then we should have an usd invoice automatically created
		usdInvoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Page: pagination.Page{
				PageNumber: 1,
				PageSize:   10,
			},

			Namespaces:       []string{namespace},
			Customers:        []string{customerEntity.ID},
			Expand:           billing.InvoiceExpandAll,
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Currencies:       []currencyx.Code{currencyx.Code(currency.USD)},
		})
		require.NoError(s.T(), err)
		require.Len(s.T(), usdInvoices.Items, 1)
		usdInvoice := usdInvoices.Items[0]

		usdInvoiceLine := usdInvoice.Lines.MustGet()[0]
		expectedUSDLine := &billing.Line{
			LineBase: billing.LineBase{
				ID:        items[0].ID,
				Namespace: namespace,

				Period: billing.Period{Start: periodStart.Truncate(time.Microsecond), End: periodEnd.Truncate(time.Microsecond)},

				InvoiceID: usdInvoice.ID,
				InvoiceAt: issueAt.In(time.UTC),
				ManagedBy: billing.ManuallyManagedLine,

				Type: billing.InvoiceLineTypeFee,

				Name:     "Test item - USD",
				Currency: currencyx.Code(currency.USD),

				Status: billing.InvoiceLineStatusValid,

				CreatedAt: usdInvoiceLine.CreatedAt.In(time.UTC),
				UpdatedAt: usdInvoiceLine.UpdatedAt.In(time.UTC),

				Metadata: map[string]string{
					"key": "value",
				},
			},
			FlatFee: &billing.FlatFeeLine{
				ConfigID:      usdInvoiceLine.FlatFee.ConfigID,
				PerUnitAmount: alpacadecimal.NewFromFloat(100),
				Quantity:      alpacadecimal.NewFromFloat(1),
				Category:      billing.FlatFeeCategoryRegular,
				PaymentTerm:   productcatalog.InAdvancePaymentTerm,
			},
		}
		// Let's make sure that the workflow config is cloned
		expectedInvoice := billing.Invoice{
			InvoiceBase: billing.InvoiceBase{
				Namespace: namespace,
				ID:        usdInvoice.ID,

				Type:     billing.InvoiceTypeStandard,
				Number:   "GATHER-TECU-USD-1",
				Currency: currencyx.Code(currency.USD),
				Status:   billing.InvoiceStatusGathering,

				CreatedAt: usdInvoice.CreatedAt,
				UpdatedAt: usdInvoice.UpdatedAt,

				Workflow: billing.InvoiceWorkflow{
					Config: billing.WorkflowConfig{
						Collection: billingProfile.WorkflowConfig.Collection,
						Invoicing:  billingProfile.WorkflowConfig.Invoicing,
						Payment:    billingProfile.WorkflowConfig.Payment,
					},
					SourceBillingProfileID: billingProfile.ID,
					AppReferences:          *billingProfile.AppReferences,
					Apps:                   billingProfile.Apps,
				},

				Customer: billing.InvoiceCustomer{
					CustomerID: customerEntity.ID,

					Name:           customerEntity.Name,
					BillingAddress: customerEntity.BillingAddress,
					UsageAttribution: billing.CustomerUsageAttribution{
						SubjectKeys: []string{"test"},
					},
				},
				Supplier: billingProfile.Supplier,
			},

			Lines:     billing.NewLineChildren([]*billing.Line{expectedUSDLine}),
			Discounts: billing.NewInvoiceDiscounts(nil),

			ExpandedFields: billing.InvoiceExpandAll,
		}
		_ = billingservice.UpdateInvoiceCollectionAt(&expectedInvoice, billingProfile.WorkflowConfig.Collection)

		require.Equal(s.T(),
			expectedInvoice.RemoveMetaForCompare(),
			usdInvoice.RemoveMetaForCompare())

		require.Len(s.T(), items, 3)
		// Validate that the create returns the expected items
		items[0].CreatedAt = expectedUSDLine.CreatedAt
		items[0].UpdatedAt = expectedUSDLine.UpdatedAt
		require.Equal(s.T(), items[0].RemoveMetaForCompare(), expectedUSDLine.RemoveMetaForCompare())
		require.NotEmpty(s.T(), items[1].ID)

		HUFItem = items[1]
	})

	s.T().Run("CreateInvoiceItems - HUF", func(t *testing.T) {
		// Then a HUF item is also created
		require.NotNil(s.T(), HUFItem.ID)

		// Then we have a different invoice for HUF
		hufInvoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Page: pagination.Page{
				PageNumber: 1,
				PageSize:   10,
			},

			Namespaces:       []string{namespace},
			Customers:        []string{customerEntity.ID},
			Expand:           billing.InvoiceExpandAll,
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Currencies:       []currencyx.Code{currencyx.Code(currency.HUF)},
		})
		require.NoError(s.T(), err)
		require.Len(s.T(), hufInvoices.Items, 1)

		hufInvoiceLines := hufInvoices.Items[0].Lines.MustGet()

		// Then we have two line items for the invoice
		require.Len(s.T(), hufInvoiceLines, 2)

		_, found := lo.Find(hufInvoiceLines, func(l *billing.Line) bool {
			return l.Type == billing.InvoiceLineTypeFee
		})
		require.True(s.T(), found, "manual fee item is present")

		// Then we should have the tiered price present
		tieredLine, found := lo.Find(hufInvoiceLines, func(l *billing.Line) bool {
			return l.Type == billing.InvoiceLineTypeUsageBased
		})

		require.True(s.T(), found, "tiered price item is present")
		require.Equal(s.T(), tieredLine.UsageBased.Price.Type(), productcatalog.TieredPriceType)
		tieredPrice, err := tieredLine.UsageBased.Price.AsTiered()
		require.NoError(s.T(), err)

		require.Equal(s.T(),
			tieredPrice,
			productcatalog.TieredPrice{
				Mode: productcatalog.GraduatedTieredPrice,
				Tiers: []productcatalog.PriceTier{
					{
						UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(100)),
						UnitPrice: &productcatalog.PriceTierUnitPrice{
							Amount: alpacadecimal.NewFromFloat(10),
						},
					},
					{
						UnitPrice: &productcatalog.PriceTierUnitPrice{
							Amount: alpacadecimal.NewFromFloat(100),
						},
					},
				},
			},
		)
	})

	s.T().Run("Expand scenarios - no  expand", func(t *testing.T) {
		invoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Page: pagination.Page{
				PageNumber: 1,
				PageSize:   10,
			},

			Namespaces:       []string{namespace},
			Customers:        []string{customerEntity.ID},
			Expand:           billing.InvoiceExpand{},
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Currencies:       []currencyx.Code{currencyx.Code(currency.USD)},
		})
		require.NoError(s.T(), err)
		require.Len(s.T(), invoices.Items, 1)
		invoice := invoices.Items[0]

		require.False(s.T(), invoice.Lines.IsPresent(), "no lines should be returned")
		require.NotNil(s.T(), invoice.Workflow, "workflow should be returned")
		require.Nil(s.T(), invoice.Workflow.Apps, "apps should not be resolved")
	})

	s.T().Run("Expand scenarios - app expand", func(t *testing.T) {
		invoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Page: pagination.Page{
				PageNumber: 1,
				PageSize:   10,
			},

			Namespaces: []string{namespace},
			Customers:  []string{customerEntity.ID},
			Expand: billing.InvoiceExpand{
				WorkflowApps: true,
			},
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Currencies:       []currencyx.Code{currencyx.Code(currency.USD)},
		})
		require.NoError(s.T(), err)
		require.Len(s.T(), invoices.Items, 1)
		invoice := invoices.Items[0]

		require.False(s.T(), invoice.Lines.IsPresent(), "no lines should be returned")
		require.NotNil(s.T(), invoice.Workflow, "workflow should be returned")
		require.NotNil(s.T(), invoice.Workflow.Apps, "apps should  be resolved")
		require.NotNil(s.T(), invoice.Workflow.Apps.Tax, "apps should be resolved")
		require.NotNil(s.T(), invoice.Workflow.Apps.Invoicing, "apps should be resolved")
		require.NotNil(s.T(), invoice.Workflow.Apps.Payment, "apps should be resolved")
	})
}

func (s *InvoicingTestSuite) TestCreateInvoice() {
	namespace := "ns-create-invoice-gathering-to-draft"
	now := time.Now().Truncate(time.Microsecond)
	periodEnd := now.Add(-time.Hour)
	periodStart := periodEnd.Add(-time.Hour * 24 * 30)
	line1IssueAt := now.Add(-2 * time.Hour)
	line2IssueAt := now.Add(-time.Hour)

	_ = s.InstallSandboxApp(s.T(), namespace)

	ctx := context.Background()

	// Given we have a test customer

	customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
		Namespace: namespace,

		CustomerMutate: customer.CustomerMutate{
			Name:         "Test Customer",
			PrimaryEmail: lo.ToPtr("test@test.com"),
			BillingAddress: &models.Address{
				Country: lo.ToPtr(models.CountryCode("US")),
			},
			Currency: lo.ToPtr(currencyx.Code(currency.USD)),
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), customerEntity)
	require.NotEmpty(s.T(), customerEntity.ID)

	// Given we have a default profile for the namespace

	minimalCreateProfileInput := MinimalCreateProfileInputTemplate
	minimalCreateProfileInput.Namespace = namespace

	profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	res, err := s.BillingService.CreatePendingInvoiceLines(ctx,
		billing.CreateInvoiceLinesInput{
			Namespace: namespace,
			Lines: []billing.LineWithCustomer{
				{
					Line: billing.Line{
						LineBase: billing.LineBase{
							Namespace: namespace,
							Period:    billing.Period{Start: periodStart, End: periodEnd},

							InvoiceAt: line1IssueAt,

							Type:      billing.InvoiceLineTypeFee,
							ManagedBy: billing.ManuallyManagedLine,

							Name:     "Test item1",
							Currency: currencyx.Code(currency.USD),

							Metadata: map[string]string{
								"key": "value",
							},
						},
						FlatFee: &billing.FlatFeeLine{
							PerUnitAmount: alpacadecimal.NewFromFloat(100),
							Quantity:      alpacadecimal.NewFromFloat(1),
							Category:      billing.FlatFeeCategoryRegular,
							PaymentTerm:   productcatalog.InAdvancePaymentTerm,
						},
					},
					CustomerID: customerEntity.ID,
				},
				{
					Line: billing.Line{
						LineBase: billing.LineBase{
							Namespace: namespace,
							Period:    billing.Period{Start: periodStart, End: periodEnd},

							InvoiceAt: line2IssueAt,

							Type:      billing.InvoiceLineTypeFee,
							ManagedBy: billing.ManuallyManagedLine,

							Name:     "Test item2",
							Currency: currencyx.Code(currency.USD),
						},
						FlatFee: &billing.FlatFeeLine{
							PerUnitAmount: alpacadecimal.NewFromFloat(200),
							Quantity:      alpacadecimal.NewFromFloat(3),
							Category:      billing.FlatFeeCategoryRegular,
							PaymentTerm:   productcatalog.InAdvancePaymentTerm,
						},
					},
					CustomerID: customerEntity.ID,
				},
			},
		})

	// Then we should have the items created
	require.NoError(s.T(), err)
	require.Len(s.T(), res, 2)
	line1ID := res[0].ID
	line2ID := res[1].ID
	require.NotEmpty(s.T(), line1ID)
	require.NotEmpty(s.T(), line2ID)

	// Expect that a single gathering invoice has been created
	require.Equal(s.T(), res[0].InvoiceID, res[1].InvoiceID)
	gatheringInvoiceID := billing.InvoiceID{
		Namespace: namespace,
		ID:        res[0].InvoiceID,
	}

	s.Run("Creating invoice in the future fails", func() {
		_, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customer.CustomerID{
				ID:        customerEntity.ID,
				Namespace: customerEntity.Namespace,
			},
			AsOf: lo.ToPtr(now.Add(time.Hour)),
		})

		require.Error(s.T(), err)
		require.ErrorAs(s.T(), err, &billing.ValidationError{})
	})

	s.Run("Creating invoice without any pending lines being available fails", func() {
		_, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customer.CustomerID{
				ID:        customerEntity.ID,
				Namespace: customerEntity.Namespace,
			},

			AsOf: lo.ToPtr(line1IssueAt.Add(-time.Minute)),
		})

		require.Error(s.T(), err)
		require.ErrorAs(s.T(), err, &billing.ValidationError{})
	})

	s.Run("Number of pending invoice lines is reported correctly by the adapter", func() {
		res, err := s.BillingAdapter.AssociatedLineCounts(ctx, billing.AssociatedLineCountsAdapterInput{
			Namespace:  namespace,
			InvoiceIDs: []string{gatheringInvoiceID.ID},
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), res.Counts, 1)
		require.Equal(s.T(), int64(2), res.Counts[gatheringInvoiceID])
	})

	s.Run("When creating an invoice with only item1 included", func() {
		invoice, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customer.CustomerID{
				ID:        customerEntity.ID,
				Namespace: customerEntity.Namespace,
			},
			AsOf: lo.ToPtr(line1IssueAt.Add(time.Minute)),
		})

		// Then we should have the invoice created
		require.NoError(s.T(), err)
		require.Len(s.T(), invoice, 1)

		// Then we should have item1 added to the invoice
		require.Len(s.T(), invoice[0].Lines.MustGet(), 1)
		require.Equal(s.T(), line1ID, invoice[0].Lines.MustGet()[0].ID)

		// Then we expect that the gathering invoice is still present, with item2
		gatheringInvoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
			Invoice: gatheringInvoiceID,
			Expand:  billing.InvoiceExpandAll,
		})
		require.NoError(s.T(), err)
		require.Nil(s.T(), gatheringInvoice.DeletedAt, "gathering invoice should be present")
		require.Len(s.T(), gatheringInvoice.Lines.MustGet(), 1)
		require.Equal(s.T(), line2ID, gatheringInvoice.Lines.MustGet()[0].ID)

		// We expect the freshly generated invoice to be in waiting for auto approval state
		require.Equal(s.T(), billing.InvoiceStatusDraftWaitingAutoApproval, invoice[0].Status)

		// We expect that the invoice can be listed by filtering to it's status_details_cache field
		invoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Namespaces:         []string{namespace},
			HasAvailableAction: []billing.InvoiceAvailableActionsFilter{billing.InvoiceAvailableActionsFilterApprove},
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), invoices.Items, 1)
		require.Equal(s.T(), invoice[0].ID, invoices.Items[0].ID)
	})

	s.Run("When creating an invoice with only item2 included, but bad asof", func() {
		_, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customer.CustomerID{
				ID:        customerEntity.ID,
				Namespace: customerEntity.Namespace,
			},
			IncludePendingLines: mo.Some([]string{line2ID}),
			AsOf:                lo.ToPtr(line1IssueAt.Add(time.Minute)),
		})

		// Then we should receive a validation error
		require.Error(s.T(), err)
		require.ErrorAs(s.T(), err, &billing.ValidationError{})
	})

	s.Run("When creating an invoice with only item2 included", func() {
		invoice, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customer.CustomerID{
				ID:        customerEntity.ID,
				Namespace: customerEntity.Namespace,
			},
			IncludePendingLines: mo.Some([]string{line2ID}),
			AsOf:                lo.ToPtr(now),
		})

		// Then we should have the invoice created
		require.NoError(s.T(), err)
		require.Len(s.T(), invoice, 1)

		// Then we should have item2 added to the invoice
		require.Len(s.T(), invoice[0].Lines.MustGet(), 1)
		require.Equal(s.T(), line2ID, invoice[0].Lines.MustGet()[0].ID)

		// Then we expect that the gathering invoice is deleted and empty
		gatheringInvoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
			Invoice: gatheringInvoiceID,
			Expand:  billing.InvoiceExpandAll,
		})
		require.NoError(s.T(), err)
		require.NotNil(s.T(), gatheringInvoice.DeletedAt, "gathering invoice should be present")
		require.Len(s.T(), gatheringInvoice.Lines.MustGet(), 0, "deleted gathering invoice is empty")
	})

	s.Run("When staging more lines the old gathering invoice gets reused", func() {
		res, err := s.BillingService.CreatePendingInvoiceLines(ctx,
			billing.CreateInvoiceLinesInput{
				Namespace: namespace,
				Lines: []billing.LineWithCustomer{
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Namespace: namespace,
								Period:    billing.Period{Start: periodStart, End: periodEnd},

								InvoiceAt: line1IssueAt,

								Type:      billing.InvoiceLineTypeFee,
								ManagedBy: billing.ManuallyManagedLine,

								Name:     "Test item1",
								Currency: currencyx.Code(currency.USD),

								Metadata: map[string]string{
									"key": "value",
								},
							},
							FlatFee: &billing.FlatFeeLine{
								PerUnitAmount: alpacadecimal.NewFromFloat(100),
								Quantity:      alpacadecimal.NewFromFloat(1),
								Category:      billing.FlatFeeCategoryRegular,
								PaymentTerm:   productcatalog.InAdvancePaymentTerm,
							},
						},
						CustomerID: customerEntity.ID,
					},
				},
			})

		s.NoError(err)
		s.Len(res, 1)

		newPendingLine := res[0]
		s.Equal(gatheringInvoiceID.ID, newPendingLine.InvoiceID)

		// The gathering invoice is undeleted
		gatheringInvoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
			Invoice: gatheringInvoiceID,
			Expand:  billing.InvoiceExpandAll,
		})
		s.NoError(err)
		s.Nil(gatheringInvoice.DeletedAt)
	})
}

type draftInvoiceInput struct {
	Namespace string
	Customer  *customer.Customer
}

func (i draftInvoiceInput) Validate() error {
	if i.Namespace == "" {
		return errors.New("namespace is required")
	}

	if err := i.Customer.Validate(); err != nil {
		return err
	}

	return nil
}

func (s *InvoicingTestSuite) createDraftInvoice(t *testing.T, ctx context.Context, in draftInvoiceInput) billing.Invoice {
	namespace := in.Customer.Namespace

	now := time.Now()
	invoiceAt := now.Add(-time.Second)
	periodEnd := now.Add(-24 * time.Hour)
	periodStart := periodEnd.Add(-24 * 30 * time.Hour)
	// Given we have a default profile for the namespace

	res, err := s.BillingService.CreatePendingInvoiceLines(ctx,
		billing.CreateInvoiceLinesInput{
			Namespace: in.Customer.Namespace,
			Lines: []billing.LineWithCustomer{
				{
					Line: billing.Line{
						LineBase: billing.LineBase{
							Namespace: namespace,
							Period:    billing.Period{Start: periodStart, End: periodEnd},

							InvoiceAt: invoiceAt,

							Type:      billing.InvoiceLineTypeFee,
							ManagedBy: billing.ManuallyManagedLine,

							Name:     "Test item1",
							Currency: currencyx.Code(currency.USD),

							Metadata: map[string]string{
								"key": "value",
							},
						},
						FlatFee: &billing.FlatFeeLine{
							PerUnitAmount: alpacadecimal.NewFromFloat(100),
							Quantity:      alpacadecimal.NewFromFloat(1),
							Category:      billing.FlatFeeCategoryRegular,
							PaymentTerm:   productcatalog.InAdvancePaymentTerm,
						},
					},
					CustomerID: in.Customer.ID,
				},
				{
					Line: billing.Line{
						LineBase: billing.LineBase{
							Namespace: namespace,
							Period:    billing.Period{Start: periodStart, End: periodEnd},

							InvoiceAt: invoiceAt,

							Type:      billing.InvoiceLineTypeFee,
							ManagedBy: billing.ManuallyManagedLine,

							Name:     "Test item2",
							Currency: currencyx.Code(currency.USD),
						},
						FlatFee: &billing.FlatFeeLine{
							PerUnitAmount: alpacadecimal.NewFromFloat(200),
							Quantity:      alpacadecimal.NewFromFloat(3),
							Category:      billing.FlatFeeCategoryRegular,
							PaymentTerm:   productcatalog.InAdvancePaymentTerm,
						},
					},
					CustomerID: in.Customer.ID,
				},
			},
		})

	require.NoError(s.T(), err)
	require.Len(s.T(), res, 2)
	line1ID := res[0].ID
	line2ID := res[1].ID
	require.NotEmpty(s.T(), line1ID)
	require.NotEmpty(s.T(), line2ID)

	invoice, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
		Customer: customer.CustomerID{
			ID:        in.Customer.ID,
			Namespace: in.Customer.Namespace,
		},
		AsOf: lo.ToPtr(now),
	})

	require.NoError(t, err)
	require.Len(t, invoice, 1)
	require.Len(t, invoice[0].Lines.MustGet(), 2)

	return invoice[0]
}

func (s *InvoicingTestSuite) TestInvoicingFlow() {
	cases := []struct {
		name           string
		workflowConfig billing.WorkflowConfig
		advance        func(t *testing.T, ctx context.Context, invoice billing.Invoice)
		expectedState  billing.InvoiceStatus
	}{
		{
			name: "instant issue",
			workflowConfig: billing.WorkflowConfig{
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: true,
					DraftPeriod: lo.Must(isodate.String("PT0S").Parse()),
					DueAfter:    lo.Must(isodate.String("P1W").Parse()),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},
			advance: func(t *testing.T, ctx context.Context, invoice billing.Invoice) {
				// When trying to advance an issued invoice, we get an error
				_, err := s.BillingService.AdvanceInvoice(ctx, billing.AdvanceInvoiceInput{
					ID:        invoice.ID,
					Namespace: invoice.Namespace,
				})

				require.ErrorIs(t, err, billing.ErrInvoiceCannotAdvance)
				require.ErrorAs(t, err, &billing.ValidationError{})
			},
			expectedState: billing.InvoiceStatusPaid,
		},
		{
			name: "draft period bypass with manual approve",
			workflowConfig: billing.WorkflowConfig{
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: true,
					DraftPeriod: lo.Must(isodate.String("PT1H").Parse()),
					DueAfter:    lo.Must(isodate.String("P1W").Parse()),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},
			advance: func(t *testing.T, ctx context.Context, invoice billing.Invoice) {
				require.Equal(s.T(), billing.InvoiceStatusDraftWaitingAutoApproval, invoice.Status)

				// Approve the invoice, should become DraftReadyToIssue
				invoice, err := s.BillingService.ApproveInvoice(ctx, billing.ApproveInvoiceInput{
					ID:        invoice.ID,
					Namespace: invoice.Namespace,
				})

				require.NoError(s.T(), err)
				require.Equal(s.T(), billing.InvoiceStatusPaid, invoice.Status)
			},
			expectedState: billing.InvoiceStatusPaid,
		},
		{
			name: "manual approvement flow",
			workflowConfig: billing.WorkflowConfig{
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: false,
					DraftPeriod: lo.Must(isodate.String("PT0H").Parse()),
					DueAfter:    lo.Must(isodate.String("P1W").Parse()),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},
			advance: func(t *testing.T, ctx context.Context, invoice billing.Invoice) {
				require.Equal(s.T(), billing.InvoiceStatusDraftManualApprovalNeeded, invoice.Status)
				require.Equal(s.T(), billing.InvoiceStatusDetails{
					AvailableActions: billing.InvoiceAvailableActions{
						Approve: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusPaymentProcessingPending,
						},
					},
				}, invoice.StatusDetails)

				// Approve the invoice, should become Issued
				invoice, err := s.BillingService.ApproveInvoice(ctx, billing.ApproveInvoiceInput{
					ID:        invoice.ID,
					Namespace: invoice.Namespace,
				})

				require.NoError(s.T(), err)
				require.Equal(s.T(), billing.InvoiceStatusPaid, invoice.Status)
			},
			expectedState: billing.InvoiceStatusPaid,
		},
		// sandbox payment status override metadata
		{
			name: "app sandbox failed payment simulation",
			workflowConfig: billing.WorkflowConfig{
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: false,
					DraftPeriod: lo.Must(isodate.String("PT0H").Parse()),
					DueAfter:    lo.Must(isodate.String("P1W").Parse()),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},
			advance: func(t *testing.T, ctx context.Context, invoice billing.Invoice) {
				require.Equal(s.T(), billing.InvoiceStatusDraftManualApprovalNeeded, invoice.Status)

				// Let's instruct the sandbox to fail the invoice
				_, err := s.BillingService.UpdateInvoice(ctx, billing.UpdateInvoiceInput{
					Invoice: invoice.InvoiceID(),
					EditFn: func(invoice *billing.Invoice) error {
						invoice.Metadata = map[string]string{
							appsandbox.TargetPaymentStatusMetadataKey: appsandbox.TargetPaymentStatusFailed,
						}

						return nil
					},
				})
				s.NoError(err)

				// Approve the invoice, should become InvoiceStatusPaymentProcessingFailed
				invoice, err = s.BillingService.ApproveInvoice(ctx, billing.ApproveInvoiceInput{
					ID:        invoice.ID,
					Namespace: invoice.Namespace,
				})

				require.NoError(s.T(), err)
				require.Equal(s.T(), billing.InvoiceStatusPaymentProcessingFailed, invoice.Status)
				require.Len(s.T(), invoice.ValidationIssues, 1)

				validationIssue := invoice.ValidationIssues[0]
				require.ElementsMatch(s.T(), billing.ValidationIssues{
					{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      validationIssue.Code,
						Message:   validationIssue.Message,
						Component: "app.sandbox.invoiceCustomers.initiate_payment",
					},
				}, invoice.ValidationIssues.RemoveMetaForCompare())
			},
			expectedState: billing.InvoiceStatusPaymentProcessingFailed,
		},
	}

	ctx := context.Background()

	for i, tc := range cases {
		s.T().Run(tc.name, func(t *testing.T) {
			namespace := fmt.Sprintf("ns-invoicing-flow-happy-path-%d", i)

			_ = s.InstallSandboxApp(s.T(), namespace)

			// Given we have a test customer
			customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
				Namespace: namespace,

				CustomerMutate: customer.CustomerMutate{
					Name:         "Test Customer",
					PrimaryEmail: lo.ToPtr("test@test.com"),
					BillingAddress: &models.Address{
						Country: lo.ToPtr(models.CountryCode("US")),
					},
					Currency: lo.ToPtr(currencyx.Code(currency.USD)),
				},
			})
			require.NoError(s.T(), err)
			require.NotNil(s.T(), customerEntity)
			require.NotEmpty(s.T(), customerEntity.ID)

			// Given we have a billing profile
			minimalCreateProfileInput := MinimalCreateProfileInputTemplate
			minimalCreateProfileInput.Namespace = namespace
			minimalCreateProfileInput.WorkflowConfig = tc.workflowConfig

			profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

			require.NoError(s.T(), err)
			require.NotNil(s.T(), profile)

			invoice := s.createDraftInvoice(s.T(), ctx, draftInvoiceInput{
				Namespace: namespace,
				Customer:  customerEntity,
			})
			require.NotNil(s.T(), invoice)

			// When we advance the invoice
			tc.advance(t, ctx, invoice)

			resultingInvoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
				Invoice: billing.InvoiceID{
					Namespace: namespace,
					ID:        invoice.ID,
				},
				Expand: billing.InvoiceExpandAll,
			})

			require.NoError(s.T(), err)
			require.NotNil(s.T(), resultingInvoice)
			require.Equal(s.T(), tc.expectedState, resultingInvoice.Status)
		})
	}
}

type ValidationIssueIntrospector interface {
	IntrospectValidationIssues(ctx context.Context, invoice billing.InvoiceID) ([]billingadapter.ValidationIssueWithDBMeta, error)
}

func (s *InvoicingTestSuite) TestInvoicingFlowErrorHandling() {
	cases := []struct {
		name           string
		workflowConfig billing.WorkflowConfig
		advance        func(t *testing.T, ctx context.Context, ns string, customer *customer.Customer, mockApp *appsandbox.MockApp) *billing.Invoice
		expectedState  billing.InvoiceStatus
	}{
		{
			name: "validation issue - different sources",
			workflowConfig: billing.WorkflowConfig{
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: true,
					DraftPeriod: lo.Must(isodate.String("PT0S").Parse()),
					DueAfter:    lo.Must(isodate.String("P1W").Parse()),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},
			advance: func(t *testing.T, ctx context.Context, ns string, customer *customer.Customer, mockApp *appsandbox.MockApp) *billing.Invoice {
				calcMock := s.InvoiceCalculator.EnableMock()
				defer s.InvoiceCalculator.DisableMock(t)

				validationIssueGetter, ok := s.BillingAdapter.(ValidationIssueIntrospector)
				require.True(t, ok)

				// Given that the app will return a validation error
				mockApp.OnValidateInvoice(billing.NewValidationError("test1", "validation error"))
				calcMock.OnCalculate(nil)

				// When we create a draft invoice
				invoice := s.createDraftInvoice(s.T(), ctx, draftInvoiceInput{
					Namespace: ns,
					Customer:  customer,
				})
				require.NotNil(s.T(), invoice)

				// Then we should end up in draft_invalid state
				require.Equal(s.T(), billing.InvoiceStatusDraftInvalid, invoice.Status)
				require.Equal(s.T(), billing.InvoiceStatusDetails{
					AvailableActions: billing.InvoiceAvailableActions{
						Retry: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusPaymentProcessingPending,
						},
						Delete: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusDeleted,
						},
					},
					Immutable: false,
				}, invoice.StatusDetails)
				require.Equal(s.T(), billing.ValidationIssues{
					{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test1",
						Message:   "validation error",
						Component: "app.sandbox.invoiceCustomers.validate",
					},
				}, invoice.ValidationIssues.RemoveMetaForCompare())

				// Then we have the issues captured in the database
				issues, err := validationIssueGetter.IntrospectValidationIssues(ctx, billing.InvoiceID{
					Namespace: ns,
					ID:        invoice.ID,
				})
				require.NoError(t, err)
				require.Len(t, issues, 1)
				require.Equal(t,
					billing.ValidationIssue{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test1",
						Message:   "validation error",
						Component: "app.sandbox.invoiceCustomers.validate",
					},
					issues[0].ValidationIssue,
				)
				require.Nil(t, issues[0].DeletedAt)
				customerValidationIssueID := issues[0].ID
				require.NotEmpty(t, customerValidationIssueID)

				calcMock.AssertExpectations(t)
				mockApp.Reset(t)

				// Given that the issue is fixed, but a new one is introduced by editing the invoice
				mockApp.OnValidateInvoice(nil)
				calcMock.OnCalculate(billing.NewValidationError("test2", "validation error"))

				// regardless the state transition will be the same for now.
				invoice, err = s.BillingService.RetryInvoice(ctx, billing.RetryInvoiceInput{
					ID:        invoice.ID,
					Namespace: invoice.Namespace,
				})
				require.NoError(s.T(), err)
				require.NotNil(s.T(), invoice)

				// Then we should end up in draft_invalid state
				require.Equal(s.T(), billing.InvoiceStatusDraftInvalid, invoice.Status)
				require.Equal(s.T(), billing.InvoiceStatusDetails{
					AvailableActions: billing.InvoiceAvailableActions{
						Retry: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusPaymentProcessingPending,
						},
						Delete: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusDeleted,
						},
					},
					Immutable: false,
				}, invoice.StatusDetails)
				require.Equal(s.T(), billing.ValidationIssues{
					{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test2",
						Message:   "validation error",
						Component: billing.ValidationComponentOpenMeter,
					},
				}, invoice.ValidationIssues.RemoveMetaForCompare())

				// Then we have the new issues captured in the database, the old one deleted, as Retry changes the severity
				// we will have a new validation issue
				issues, err = validationIssueGetter.IntrospectValidationIssues(ctx, billing.InvoiceID{
					Namespace: ns,
					ID:        invoice.ID,
				})
				require.NoError(t, err)
				require.Len(t, issues, 3)

				// The old issue should be deleted
				invoiceIssue, ok := lo.Find(issues, func(i billingadapter.ValidationIssueWithDBMeta) bool {
					return i.ID == customerValidationIssueID
				})
				require.True(t, ok, "old issue should be present")
				require.NotNil(t, invoiceIssue.DeletedAt)
				require.Equal(t,
					billing.ValidationIssue{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test1",
						Message:   "validation error",
						Component: "app.sandbox.invoiceCustomers.validate",
					},
					invoiceIssue.ValidationIssue,
				)

				// A new version of the issue is present with downgraded severity, to facilitate the retry
				downgradedIssue, ok := lo.Find(issues, func(i billingadapter.ValidationIssueWithDBMeta) bool {
					return i.Code == "test1" && i.Severity == billing.ValidationIssueSeverityWarning
				})
				require.True(t, ok, "the issue should be present")
				require.NotNil(t, downgradedIssue.DeletedAt)
				require.Equal(t,
					billing.ValidationIssue{
						Severity:  billing.ValidationIssueSeverityWarning,
						Code:      "test1",
						Message:   "validation error",
						Component: "app.sandbox.invoiceCustomers.validate",
					},
					downgradedIssue.ValidationIssue,
				)

				// The new issue should not be deleted
				calculationErrorIssue, ok := lo.Find(issues, func(i billingadapter.ValidationIssueWithDBMeta) bool {
					return i.Code == "test2"
				})
				require.True(t, ok, "new issue should be present")
				require.Equal(t,
					billing.ValidationIssue{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test2",
						Message:   "validation error",
						Component: "openmeter",
					},
					calculationErrorIssue.ValidationIssue,
				)

				mockApp.Reset(t)
				calcMock.Reset(t)

				// Given that both issues are present, both will be reported
				mockApp.OnValidateInvoice(billing.NewValidationError("test1", "validation error"))
				calcMock.OnCalculate(billing.NewValidationError("test2", "validation error"))

				// regardless the state transition will be the same for now.
				invoice, err = s.BillingService.RetryInvoice(ctx, billing.RetryInvoiceInput{
					ID:        invoice.ID,
					Namespace: invoice.Namespace,
				})
				require.NoError(s.T(), err)
				require.NotNil(s.T(), invoice)

				// Then we should end up in draft_invalid state
				require.Equal(s.T(), billing.InvoiceStatusDraftInvalid, invoice.Status)
				require.Equal(s.T(), billing.InvoiceStatusDetails{
					AvailableActions: billing.InvoiceAvailableActions{
						Retry: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusPaymentProcessingPending,
						},
						Delete: &billing.InvoiceAvailableActionDetails{
							ResultingState: billing.InvoiceStatusDeleted,
						},
					},
					Immutable: false,
				}, invoice.StatusDetails)
				require.ElementsMatch(s.T(), billing.ValidationIssues{
					{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test1",
						Message:   "validation error",
						Component: "app.sandbox.invoiceCustomers.validate",
					},
					{
						Severity:  billing.ValidationIssueSeverityCritical,
						Code:      "test2",
						Message:   "validation error",
						Component: billing.ValidationComponentOpenMeter,
					},
				}, invoice.ValidationIssues.RemoveMetaForCompare())

				// The database now has both  critical issues active (but no new ones are created)
				issues, err = validationIssueGetter.IntrospectValidationIssues(ctx, billing.InvoiceID{
					Namespace: ns,
					ID:        invoice.ID,
				})
				require.NoError(t, err)
				criticalIssues := lo.Filter(issues, func(i billingadapter.ValidationIssueWithDBMeta, _ int) bool {
					return i.Severity == billing.ValidationIssueSeverityCritical
				})
				require.Len(t, criticalIssues, 2)

				_, deletedIssueFound := lo.Find(criticalIssues, func(i billingadapter.ValidationIssueWithDBMeta) bool {
					return i.DeletedAt != nil
				})
				require.False(t, deletedIssueFound, "no issues should be deleted")

				return &invoice
			},
			expectedState: billing.InvoiceStatusDraftInvalid,
		},
		{
			name: "validation issue - warnings allow state transitions",
			workflowConfig: billing.WorkflowConfig{
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: true,
					DraftPeriod: lo.Must(isodate.String("PT0S").Parse()),
					DueAfter:    lo.Must(isodate.String("P1W").Parse()),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},
			advance: func(t *testing.T, ctx context.Context, ns string, customer *customer.Customer, mockApp *appsandbox.MockApp) *billing.Invoice {
				calcMock := s.InvoiceCalculator.EnableMock()
				defer s.InvoiceCalculator.DisableMock(t)

				// Given that the app will return a validation error
				mockApp.OnValidateInvoice(billing.NewValidationWarning("test1", "validation warning"))
				mockApp.OnUpsertInvoice(nil)
				mockApp.OnFinalizeInvoice(nil)
				calcMock.OnCalculate(nil)

				// When we create a draft invoice
				invoice := s.createDraftInvoice(s.T(), ctx, draftInvoiceInput{
					Namespace: ns,
					Customer:  customer,
				})
				require.NotNil(s.T(), invoice)

				// We are using the mock app factory, so we won't have automatic payment handling provided by the sandbox app
				require.Equal(s.T(), billing.InvoiceStatusPaymentProcessingPending, invoice.Status)
				require.Equal(s.T(), billing.InvoiceStatusDetails{
					AvailableActions: billing.InvoiceAvailableActions{},
					Immutable:        true,
				}, invoice.StatusDetails)
				require.Equal(s.T(), billing.ValidationIssues{
					{
						Severity:  billing.ValidationIssueSeverityWarning,
						Code:      "test1",
						Message:   "validation warning",
						Component: "app.sandbox.invoiceCustomers.validate",
					},
				}, invoice.ValidationIssues.RemoveMetaForCompare())

				return &invoice
			},
			expectedState: billing.InvoiceStatusPaymentProcessingPending,
		},
	}

	ctx := context.Background()

	for i, tc := range cases {
		s.T().Run(tc.name, func(t *testing.T) {
			namespace := fmt.Sprintf("ns-invoicing-flow-valid-%d", i)

			_ = s.InstallSandboxApp(s.T(), namespace)

			mockApp := s.SandboxApp.EnableMock(t)
			defer s.SandboxApp.DisableMock()

			// Given we have a test customer
			customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
				Namespace: namespace,

				CustomerMutate: customer.CustomerMutate{
					Name:         "Test Customer",
					PrimaryEmail: lo.ToPtr("test@test.com"),
					BillingAddress: &models.Address{
						Country: lo.ToPtr(models.CountryCode("US")),
					},
					Currency: lo.ToPtr(currencyx.Code(currency.USD)),
				},
			})
			require.NoError(s.T(), err)
			require.NotNil(s.T(), customerEntity)
			require.NotEmpty(s.T(), customerEntity.ID)

			// Given we have a billing profile
			minimalCreateProfileInput := MinimalCreateProfileInputTemplate
			minimalCreateProfileInput.Namespace = namespace
			minimalCreateProfileInput.WorkflowConfig = tc.workflowConfig

			profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

			require.NoError(s.T(), err)
			require.NotNil(s.T(), profile)

			// When we advance the invoice
			invoice := tc.advance(t, ctx, namespace, customerEntity, mockApp)

			mockApp.AssertExpectations(t)

			resultingInvoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
				Invoice: billing.InvoiceID{
					Namespace: namespace,
					ID:        invoice.ID,
				},
				Expand: billing.InvoiceExpandAll,
			})

			require.NoError(s.T(), err)
			require.NotNil(s.T(), resultingInvoice)
			require.Equal(s.T(), tc.expectedState, resultingInvoice.Status)
		})
	}
}

func (s *InvoicingTestSuite) TestBillingProfileChange() {
	namespace := "ns-billing-profile-default-change"
	ctx := context.Background()

	_ = s.InstallSandboxApp(s.T(), namespace)

	oldCreateProfileInput := MinimalCreateProfileInputTemplate
	oldCreateProfileInput.Namespace = namespace
	oldCreateProfileInput.WorkflowConfig.Invoicing.ProgressiveBilling = true

	oldBillingProfile, err := s.BillingService.CreateProfile(ctx, oldCreateProfileInput)
	s.NoError(err)
	s.NotNil(oldBillingProfile)

	newCreateProfileInput := MinimalCreateProfileInputTemplate
	newCreateProfileInput.Namespace = namespace
	newCreateProfileInput.WorkflowConfig.Invoicing.ProgressiveBilling = true

	newBillingProfile, err := s.BillingService.CreateProfile(ctx, newCreateProfileInput)
	s.NoError(err)
	s.NotNil(newBillingProfile)

	defaultProfile, err := s.BillingService.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
		Namespace: namespace,
	})
	s.NoError(err)
	s.NotNil(defaultProfile)

	s.Equal(newBillingProfile.ID, defaultProfile.ID)
	s.NotEqual(newBillingProfile.ID, oldBillingProfile.ID)

	// Changing the old profile to default works

	oldBillingProfile.AppReferences = nil

	oldProfileAsDefault, err := s.BillingService.UpdateProfile(ctx, billing.UpdateProfileInput(oldBillingProfile.BaseProfile))
	s.NoError(err)
	s.NotNil(oldProfileAsDefault)
	s.True(oldProfileAsDefault.Default)

	defaultProfile, err = s.BillingService.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
		Namespace: namespace,
	})
	s.NoError(err)
	s.NotNil(defaultProfile)

	s.Equal(oldProfileAsDefault.ID, defaultProfile.ID)
}

func (s *InvoicingTestSuite) TestUBPProgressiveInvoicing() {
	namespace := "ns-ubp-invoicing-progressive"
	ctx := context.Background()

	periodStart := lo.Must(time.Parse(time.RFC3339, "2024-09-02T12:13:14Z"))
	periodEnd := lo.Must(time.Parse(time.RFC3339, "2024-09-03T12:13:14Z"))

	_ = s.InstallSandboxApp(s.T(), namespace)

	err := s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{
		{
			Namespace:     namespace,
			Slug:          "flat-per-unit",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
		{
			Namespace:     namespace,
			Slug:          "flat-per-usage",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
		{
			Namespace:     namespace,
			Slug:          "tiered-graduated",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
		{
			Namespace:     namespace,
			Slug:          "tiered-volume",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
	})
	require.NoError(s.T(), err, "meter adapter replace meters")

	defer func() {
		err = s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{})
		require.NoError(s.T(), err, "meter adapter replace meters")
	}()

	// Let's initialize the mock streaming connector with data that is out of the period so that we
	// can start with empty values
	for _, slug := range []string{"flat-per-unit", "flat-per-usage", "tiered-graduated", "tiered-volume"} {
		s.MockStreamingConnector.AddSimpleEvent(slug, 0, periodStart.Add(-time.Minute))
	}

	defer s.MockStreamingConnector.Reset()

	// Let's create the features
	// TODO[later]: we need to handle archived features, do we want to issue a warning? Can features be archived when used
	// by a draft invoice?
	features := ubpFeatures{
		flatPerUnit: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "flat-per-unit",
			Key:       "flat-per-unit",
			MeterSlug: lo.ToPtr("flat-per-unit"),
		})),
		flatPerUsage: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "flat-per-usage",
			Key:       "flat-per-usage",
			MeterSlug: lo.ToPtr("flat-per-usage"),
		})),
		tieredGraduated: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "tiered-graduated",
			Key:       "tiered-graduated",
			MeterSlug: lo.ToPtr("tiered-graduated"),
		})),
		tieredVolume: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "tiered-volume",
			Key:       "tiered-volume",
			MeterSlug: lo.ToPtr("tiered-volume"),
		})),
	}

	// Given we have a test customer

	customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
		Namespace: namespace,

		CustomerMutate: customer.CustomerMutate{
			Name:         "Test Customer",
			PrimaryEmail: lo.ToPtr("test@test.com"),
			BillingAddress: &models.Address{
				Country:     lo.ToPtr(models.CountryCode("US")),
				PostalCode:  lo.ToPtr("12345"),
				State:       lo.ToPtr("NY"),
				City:        lo.ToPtr("New York"),
				Line1:       lo.ToPtr("1234 Test St"),
				Line2:       lo.ToPtr("Apt 1"),
				PhoneNumber: lo.ToPtr("1234567890"),
			},
			Currency: lo.ToPtr(currencyx.Code(currency.USD)),
			UsageAttribution: customer.CustomerUsageAttribution{
				SubjectKeys: []string{"test"},
			},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), customerEntity)
	require.NotEmpty(s.T(), customerEntity.ID)

	// Given we have a default profile for the namespace
	minimalCreateProfileInput := MinimalCreateProfileInputTemplate
	minimalCreateProfileInput.Namespace = namespace
	minimalCreateProfileInput.WorkflowConfig.Invoicing.ProgressiveBilling = true

	profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	lines := ubpPendingLines{}
	s.Run("create pending invoice items", func() {
		// When we create pending invoice items
		pendingLines, err := s.BillingService.CreatePendingInvoiceLines(ctx,
			billing.CreateInvoiceLinesInput{
				Namespace: namespace,
				Lines: []billing.LineWithCustomer{
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - FLAT per unit",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.flatPerUnit.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.UnitPrice{
									Amount:        alpacadecimal.NewFromFloat(100),
									MaximumAmount: lo.ToPtr(alpacadecimal.NewFromFloat(2000)),
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - FLAT per any usage",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.flatPerUsage.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.FlatPrice{
									Amount:      alpacadecimal.NewFromFloat(100),
									PaymentTerm: productcatalog.InArrearsPaymentTerm,
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - Tiered graduated",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.tieredGraduated.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.TieredPrice{
									Mode: productcatalog.GraduatedTieredPrice,
									Tiers: []productcatalog.PriceTier{
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(10)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(100),
											},
										},
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(20)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(90),
											},
										},
										{
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(80),
											},
										},
									},
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - Tiered volume",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.tieredVolume.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.TieredPrice{
									Mode: productcatalog.VolumeTieredPrice,
									Tiers: []productcatalog.PriceTier{
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(10)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(100),
											},
										},
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(20)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(90),
											},
										},
										{
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(80),
											},
										},
									},
									MinimumAmount: lo.ToPtr(alpacadecimal.NewFromFloat(3000)),
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
				},
			},
		)
		require.NoError(s.T(), err)
		require.Len(s.T(), pendingLines, 4)

		// The pending invoice items should be truncated to 1 min resolution (start => up to next, end down to previous)
		for _, line := range pendingLines {
			require.Equal(s.T(),
				billing.Period{
					Start: lo.Must(time.Parse(time.RFC3339, "2024-09-02T12:13:00Z")),
					End:   lo.Must(time.Parse(time.RFC3339, "2024-09-03T12:13:00Z")),
				},
				line.Period,
				"period should be truncated to 1 min resolution",
			)

			require.Equal(s.T(),
				line.InvoiceAt,
				periodEnd,
				"invoice at should be unchanged",
			)
		}

		lines = ubpPendingLines{
			flatPerUnit:     pendingLines[0],
			flatPerUsage:    pendingLines[1],
			tieredGraduated: pendingLines[2],
			tieredVolume:    pendingLines[3],
		}
	})

	s.Run("create invoice with empty truncated periods", func() {
		asOf := periodStart.Add(time.Second)
		_, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.ErrorIs(s.T(), err, billing.ErrInvoiceCreateNoLines)
		require.ErrorAs(s.T(), err, &billing.ValidationError{})
	})

	s.Run("create mid period invoice", func() {
		// Usage
		s.MockStreamingConnector.AddSimpleEvent("flat-per-unit", 10, periodStart)

		// Period
		asOf := periodStart.Add(time.Hour)
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)

		invoice := out[0]

		s.DebugDumpInvoice("mid period ubp progressive invoice", invoice)

		require.Len(s.T(), invoice.ValidationIssues, 0)

		invoiceLines := invoice.Lines.MustGet()
		require.Len(s.T(), invoiceLines, 3)

		// Let's resolve the lines by parent
		flatPerUnit := s.lineWithParent(invoiceLines, lines.flatPerUnit.ID)
		flatPerUsage := s.lineWithParent(invoiceLines, lines.flatPerUsage.ID)
		tieredGraduated := s.lineWithParent(invoiceLines, lines.tieredGraduated.ID)

		// The invoice should not have:
		// - the volume item as that must be invoiced in arreas
		require.NotContains(s.T(), lo.Map(invoiceLines, func(l *billing.Line, _ int) string {
			return l.ID
		}), []string{
			flatPerUnit.ID,
			flatPerUsage.ID,
			tieredGraduated.ID,
		})

		expectedPeriod := billing.Period{
			Start: periodStart.Truncate(time.Minute),
			End:   periodStart.Add(time.Hour).Truncate(time.Minute),
		}
		for _, line := range invoiceLines {
			require.True(s.T(), expectedPeriod.Equal(line.Period), "period should be changed for the line items")
		}

		// Let's validate the output of the split itself
		tieredGraduatedChildren := s.getLineChildLines(ctx, namespace, lines.tieredGraduated.ID)
		require.True(s.T(), tieredGraduatedChildren.ParentLine.Period.Equal(lines.tieredGraduated.Period))
		require.Equal(s.T(), flatPerUnit.UsageBased.Quantity.InexactFloat64(), float64(10), "flat per unit should have 10 units")
		require.Equal(s.T(), billing.InvoiceLineStatusSplit, tieredGraduatedChildren.ParentLine.Status, "parent should be split [id=%s]", tieredGraduatedChildren.ParentLine.ID)
		require.Len(s.T(), tieredGraduatedChildren.ChildLines, 2, "there should be to child lines [id=%s]", tieredGraduatedChildren.ParentLine.ID)
		require.True(s.T(), tieredGraduatedChildren.ChildLines[0].Period.Equal(billing.Period{
			Start: periodStart.Truncate(time.Minute),
			End:   periodStart.Add(time.Hour).Truncate(time.Minute),
		}), "first child period should be truncated")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[0].InvoiceAt.Equal(periodStart.Add(time.Hour).Truncate(time.Minute)), "first child should be issued at the end of parent's period")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[1].Period.Equal(billing.Period{
			Start: periodStart.Add(time.Hour).Truncate(time.Minute),
			End:   periodEnd.Truncate(time.Minute),
		}), "second child period should be until the end of parent's period")

		// Let's validate detailed line items
		requireDetailedLines(s.T(), flatPerUnit, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.UnitPriceUsageChildUniqueReferenceID: {
					Quantity:      10,
					PerUnitAmount: 100,
				},
			},
		})

		// Let's validate the totals
		requireTotals(s.T(), expectedTotals{
			Amount: 1000,
			Total:  1000,
		}, flatPerUnit.Children.MustGet()[0].Totals)

		requireTotals(s.T(), expectedTotals{
			Amount: 1000,
			Total:  1000,
		}, flatPerUnit.Totals)

		requireTotals(s.T(), expectedTotals{
			Amount: 1000,
			Total:  1000,
		}, out[0].Totals)

		s.Run("update line item", func() {
			updatedInvoice, err := s.BillingService.UpdateInvoice(ctx, billing.UpdateInvoiceInput{
				Invoice: invoice.InvoiceID(),
				EditFn: func(invoice *billing.Invoice) error {
					line := invoice.Lines.GetByID(flatPerUnit.ID)
					if line == nil {
						return fmt.Errorf("line not found")
					}

					line.UsageBased.Price = productcatalog.NewPriceFrom(productcatalog.UnitPrice{
						Amount: alpacadecimal.NewFromFloat(250),
					})
					return nil
				},
			})
			require.NoError(s.T(), err)
			require.NotNil(s.T(), updatedInvoice)

			line := updatedInvoice.Lines.GetByID(flatPerUnit.ID)
			s.NotNil(line)

			// TODO[later]: we need to decide how to handle the situation where the line is updated, but there are split
			// lines

			require.Equal(s.T(), float64(250), lo.Must(line.UsageBased.Price.AsUnit()).Amount.InexactFloat64())
			require.True(s.T(), flatPerUnit.UpdatedAt.Before(line.UpdatedAt), "updated at should be changed")
			require.True(s.T(), flatPerUnit.CreatedAt.Equal(line.CreatedAt), "created at should not be changed")

			requireTotals(s.T(), expectedTotals{
				Amount: 2500,
				Total:  2500,
			}, line.Totals)

			invoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
				Invoice: billing.InvoiceID{
					Namespace: namespace,
					ID:        out[0].ID,
				},
				Expand: billing.InvoiceExpand{},
			})
			require.NoError(s.T(), err)

			requireTotals(s.T(), expectedTotals{
				Amount: 2500,
				Total:  2500,
			}, invoice.Totals)
		})

		s.Run("invalid update of a line item", func() {
			_, err := s.BillingService.UpdateInvoice(ctx, billing.UpdateInvoiceInput{
				Invoice: invoice.InvoiceID(),
				EditFn: func(invoice *billing.Invoice) error {
					line := invoice.Lines.GetByID(flatPerUnit.ID)
					if line == nil {
						return fmt.Errorf("line not found")
					}

					line.UsageBased.Price = productcatalog.NewPriceFrom(productcatalog.TieredPrice{
						Mode: productcatalog.VolumeTieredPrice,
						Tiers: []productcatalog.PriceTier{
							{
								UnitPrice: &productcatalog.PriceTierUnitPrice{
									Amount: alpacadecimal.NewFromFloat(250),
								},
							},
						},
					})

					return nil
				},
			})

			require.Error(s.T(), err)
			require.ErrorAs(s.T(), err, &billing.ValidationError{})
			require.ErrorIs(s.T(), err, billing.ErrInvoiceLinesNotBillable)
		})

		s.Run("deleting a valid line item worked", func() {
			updatedInvoice, err := s.BillingService.UpdateInvoice(ctx, billing.UpdateInvoiceInput{
				Invoice: invoice.InvoiceID(),
				EditFn: func(invoice *billing.Invoice) error {
					line := invoice.Lines.GetByID(flatPerUnit.ID)
					if line == nil {
						return fmt.Errorf("line not found")
					}

					line.DeletedAt = lo.ToPtr(clock.Now())
					return nil
				},
				IncludeDeletedLines: true,
			})
			require.NoError(s.T(), err)

			require.Len(s.T(), updatedInvoice.Lines.MustGet(), 3)

			deletedLine := updatedInvoice.Lines.GetByID(flatPerUnit.ID)
			require.NotNil(s.T(), deletedLine)
			require.NotNil(s.T(), deletedLine.DeletedAt)

			requireTotals(s.T(), expectedTotals{
				Amount: 0,
				Total:  0,
			}, updatedInvoice.Totals)

			// Let's validate without deleted line fetching
			updatedInvoice, err = s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
				Invoice: out[0].InvoiceID(),
				Expand:  billing.InvoiceExpandAll.SetDeletedLines(false),
			})
			require.NoError(s.T(), err)

			require.NotContains(s.T(), lo.Map(updatedInvoice.Lines.MustGet(), func(l *billing.Line, _ int) string {
				return l.ID
			}), []string{flatPerUnit.ID})

			requireTotals(s.T(), expectedTotals{
				Amount: 0,
				Total:  0,
			}, updatedInvoice.Totals)
		})

		s.Run("invoice deletion works", func() {
			// Mock invoicing app
			mockApp := s.SandboxApp.EnableMock(s.T())
			defer s.SandboxApp.DisableMock()

			s.Run("when a validation error occurs, the error is returned", func() {
				// InvoiceDeletion fails
				validationError := billing.NewValidationError("delete-failed", "invoice cannot be deleted")
				mockApp.OnDeleteInvoice(validationError)

				err := s.BillingService.DeleteInvoice(ctx, out[0].InvoiceID())
				require.Error(s.T(), err)
				require.ErrorAs(s.T(), err, &billing.ValidationError{})

				validationIssue := billing.ValidationIssue{}
				require.True(s.T(), errors.As(err, &validationIssue))
				require.Equal(s.T(), validationIssue.Code, validationError.Code)
				require.Equal(s.T(), validationIssue.Message, validationError.Message)

				deletedInvoice, err := s.BillingService.GetInvoiceByID(ctx, billing.GetInvoiceByIdInput{
					Invoice: out[0].InvoiceID(),
					Expand:  billing.InvoiceExpandAll,
				})
				require.NoError(s.T(), err)
				require.NotNil(s.T(), deletedInvoice.DeletedAt)
				require.Equal(s.T(), billing.InvoiceStatusDeleteFailed, deletedInvoice.Status)

				mockApp.AssertExpectations(s.T())
			})

			s.Run("when a generic error occurs, the error is added to the validation errors", func() {
				mockApp.Reset(s.T())

				// InvoiceDeletion fails
				mockApp.OnDeleteInvoice(errors.New("generic error"))

				invoice, err := s.BillingService.RetryInvoice(ctx, out[0].InvoiceID())
				require.NotNil(s.T(), invoice)
				require.NoError(s.T(), err)
				require.Len(s.T(), invoice.ValidationIssues, 1)
				require.Equal(s.T(), billing.InvoiceStatusDeleteFailed, invoice.Status)

				validationIssue := invoice.ValidationIssues[0]
				require.Empty(s.T(), validationIssue.Code)
				require.Equal(s.T(), "generic error", validationIssue.Message)
				require.Equal(s.T(), billing.ValidationIssueSeverityCritical, validationIssue.Severity)

				mockApp.AssertExpectations(s.T())
			})

			s.Run("when the sync passes, the invoice is deleted", func() {
				mockApp.Reset(s.T())

				mockApp.OnDeleteInvoice(nil)

				invoice, err := s.BillingService.RetryInvoice(ctx, out[0].InvoiceID())
				require.NotNil(s.T(), invoice)
				require.NoError(s.T(), err)
				require.Len(s.T(), invoice.ValidationIssues, 0)
				require.Equal(s.T(), billing.InvoiceStatusDeleted, invoice.Status)

				mockApp.AssertExpectations(s.T())
			})
		})
	})

	s.Run("create mid period invoice - pt2", func() {
		// Mock invoicing app
		mockApp := s.SandboxApp.EnableMock(s.T())
		defer s.SandboxApp.DisableMock()

		mockApp.OnValidateInvoice(nil)
		mockApp.OnUpsertInvoice(func(i billing.Invoice) *billing.UpsertInvoiceResult {
			lines := i.FlattenLinesByID()

			out := billing.NewUpsertInvoiceResult()

			for _, line := range lines {
				if line.Type == billing.InvoiceLineTypeFee {
					// We set the external id the same as the line id to make it easier to test the output.
					out.AddLineExternalID(line.ID, line.ID)
				}

				// We set the external id the same as the discount id to make it easier to test the output.
				for discountId := range line.FlattenDiscountsByID() {
					out.AddLineDiscountExternalID(discountId, discountId)
				}
			}

			out.SetInvoiceNumber("INV-123")

			return out
		})

		// Usage
		s.MockStreamingConnector.AddSimpleEvent("flat-per-unit", 20, periodStart.Add(time.Minute*100))
		s.MockStreamingConnector.AddSimpleEvent("tiered-graduated", 15, periodStart.Add(time.Minute*100))

		asOf := periodStart.Add(2 * time.Hour)
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)
		require.Len(s.T(), out[0].ValidationIssues, 0)

		invoiceLines := out[0].Lines.MustGet()

		require.Len(s.T(), invoiceLines, 3)

		// Let's resolve the lines by parent
		flatPerUnit := s.lineWithParent(invoiceLines, lines.flatPerUnit.ID)
		flatPerUsage := s.lineWithParent(invoiceLines, lines.flatPerUsage.ID)
		tieredGraduated := s.lineWithParent(invoiceLines, lines.tieredGraduated.ID)

		// The invoice should not have:
		// - the volume item as that must be invoiced in arreas
		require.NotContains(s.T(), lo.Map(invoiceLines, func(l *billing.Line, _ int) string {
			return l.ID
		}), []string{
			flatPerUnit.ID,
			flatPerUsage.ID,
			tieredGraduated.ID,
		})

		expectedPeriod := billing.Period{
			Start: periodStart.Add(time.Hour).Truncate(time.Minute),
			End:   periodStart.Add(2 * time.Hour).Truncate(time.Minute),
		}
		for _, line := range invoiceLines {
			require.True(s.T(), expectedPeriod.Equal(line.Period), "period should be changed for the line items")
		}

		// Let's validate the output of the split itself
		tieredGraduatedChildren := s.getLineChildLines(ctx, namespace, lines.tieredGraduated.ID)
		require.True(s.T(), tieredGraduatedChildren.ParentLine.Period.Equal(lines.tieredGraduated.Period))
		require.Equal(s.T(), billing.InvoiceLineStatusSplit, tieredGraduatedChildren.ParentLine.Status, "parent should be split [id=%s]", tieredGraduatedChildren.ParentLine.ID)
		require.Len(s.T(), tieredGraduatedChildren.ChildLines, 3, "there should be to child lines [id=%s]", tieredGraduatedChildren.ParentLine.ID)
		require.True(s.T(), tieredGraduatedChildren.ChildLines[0].Period.Equal(billing.Period{
			Start: periodStart.Truncate(time.Minute),
			End:   periodStart.Add(time.Hour).Truncate(time.Minute),
		}), "first child period should be truncated")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[1].Period.Equal(billing.Period{
			Start: periodStart.Add(time.Hour).Truncate(time.Minute),
			End:   periodStart.Add(2 * time.Hour).Truncate(time.Minute),
		}), "second child period should be between the first and the third child's period")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[1].InvoiceAt.Equal(periodStart.Add(2*time.Hour).Truncate(time.Minute)), "second child should be issued at the end of parent's period")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[2].Period.Equal(billing.Period{
			Start: periodStart.Add(2 * time.Hour).Truncate(time.Minute),
			End:   periodEnd.Truncate(time.Minute),
		}), "third child period should be until the end of parent's period")

		// Detailed lines
		requireDetailedLines(s.T(), flatPerUnit, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.UnitPriceUsageChildUniqueReferenceID: {
					Quantity:      20,
					PerUnitAmount: 100,
					Discounts: map[string]float64{
						billing.LineMaximumSpendReferenceID: 1000,
					},
				},
			},
		})

		requireDetailedLines(s.T(), tieredGraduated, lineExpectations{
			Details: map[string]feeLineExpect{
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 1): {
					Quantity:      10,
					PerUnitAmount: 100,
				},
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 2): {
					Quantity:      5,
					PerUnitAmount: 90,
				},
			},
		})

		// Let's validate the totals
		requireTotals(s.T(), expectedTotals{
			Amount:         2000,
			DiscountsTotal: 1000,
			Total:          1000,
		}, flatPerUnit.Totals)

		requireTotals(s.T(), expectedTotals{
			Amount: 1450,
			Total:  1450,
		}, tieredGraduated.Totals)

		requireTotals(s.T(), expectedTotals{
			Amount:         3450,
			DiscountsTotal: 1000,
			Total:          2450,
		}, out[0].Totals)

		// Invoice app testing

		require.Equal(s.T(), "INV-123", out[0].Number)

		for _, line := range out[0].FlattenLinesByID() {
			switch {
			case line.Type == billing.InvoiceLineTypeFee:
				require.Equal(s.T(), line.ID, line.ExternalIDs.Invoicing)
			case line.Type == billing.InvoiceLineTypeUsageBased:
				require.Empty(s.T(), line.ExternalIDs.Invoicing)
			default:
				s.T().Errorf("unexpected line type: %s", line.Type)
			}

			// Test discounts
			for _, discount := range line.FlattenDiscountsByID() {
				require.Equal(s.T(), discount.ID, discount.ExternalIDs.Invoicing)
			}
		}

		mockApp.AssertExpectations(s.T())

		s.Run("validate invoice finalization", func() {
			mockApp.OnUpsertInvoice(func(i billing.Invoice) *billing.UpsertInvoiceResult {
				lines := i.FlattenLinesByID()

				out := billing.NewUpsertInvoiceResult()

				for _, line := range lines {
					if line.Type == billing.InvoiceLineTypeFee {
						out.AddLineExternalID(line.ID, "final_upsert_"+line.ID)
					}

					for discountId := range line.FlattenDiscountsByID() {
						out.AddLineDiscountExternalID(discountId, "final_upsert_"+discountId)
					}
				}

				return out
			})

			finalizedInvoiceResult := billing.NewFinalizeInvoiceResult()
			finalizedInvoiceResult.SetPaymentExternalID("payment_external_id")
			mockApp.OnFinalizeInvoice(finalizedInvoiceResult)

			// Let's finalize the invoice
			finalizedInvoice, err := s.BillingService.ApproveInvoice(ctx, out[0].InvoiceID())
			require.NoError(s.T(), err)
			require.NotNil(s.T(), finalizedInvoice)

			require.Equal(s.T(), "payment_external_id", finalizedInvoice.ExternalIDs.Payment)
			// Invoice app testing
			for _, line := range finalizedInvoice.FlattenLinesByID() {
				switch {
				case line.Type == billing.InvoiceLineTypeFee:
					require.Equal(s.T(), "final_upsert_"+line.ID, line.ExternalIDs.Invoicing)
				case line.Type == billing.InvoiceLineTypeUsageBased:
					require.Empty(s.T(), line.ExternalIDs.Invoicing)
				default:
					s.T().Errorf("unexpected line type: %s", line.Type)
				}

				// Test discounts
				for _, discount := range line.FlattenDiscountsByID() {
					require.Equal(s.T(), "final_upsert_"+discount.ID, discount.ExternalIDs.Invoicing)
				}
			}

			mockApp.AssertExpectations(s.T())
		})
	})

	s.Run("create end of period invoice", func() {
		// Usage
		afterPreviousTest := periodStart.Add(3 * time.Hour)
		s.MockStreamingConnector.AddSimpleEvent("tiered-volume", 25, afterPreviousTest)
		s.MockStreamingConnector.AddSimpleEvent("tiered-graduated", 15, afterPreviousTest)

		asOf := periodEnd
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)

		invoiceLines := out[0].Lines.MustGet()

		require.Len(s.T(), invoiceLines, 4)

		// Let's resolve the lines by parent
		flatPerUnit := s.lineWithParent(invoiceLines, lines.flatPerUnit.ID)
		flatPerUsage := s.lineWithParent(invoiceLines, lines.flatPerUsage.ID)
		tieredGraduated := s.lineWithParent(invoiceLines, lines.tieredGraduated.ID)
		tieredVolume, tieredVolumeFound := lo.Find(invoiceLines, func(l *billing.Line) bool {
			return l.ID == lines.tieredVolume.ID
		})
		require.True(s.T(), tieredVolumeFound, "tiered volume line should be present")
		require.Equal(s.T(), tieredVolume.ID, lines.tieredVolume.ID, "tiered volume line should be the same (no split occurred)")

		require.NotContains(s.T(), lo.Map(invoiceLines, func(l *billing.Line, _ int) string {
			return l.ID
		}), []string{
			flatPerUnit.ID,
			flatPerUsage.ID,
			tieredGraduated.ID,
			lines.tieredVolume.ID,
		})

		expectedPeriod := billing.Period{
			Start: periodStart.Add(2 * time.Hour).Truncate(time.Minute),
			End:   periodEnd.Truncate(time.Minute),
		}
		for _, line := range []*billing.Line{flatPerUnit, flatPerUsage, tieredGraduated} {
			require.True(s.T(), expectedPeriod.Equal(line.Period), "period should be changed for the line items")
		}
		require.True(s.T(), tieredVolume.Period.Equal(lines.tieredVolume.Period), "period should be unchanged for the tiered volume line")

		// Let's validate the output of the split itself: no new split should have occurred
		tieredGraduatedChildren := s.getLineChildLines(ctx, namespace, lines.tieredGraduated.ID)
		require.True(s.T(), tieredGraduatedChildren.ParentLine.Period.Equal(lines.tieredGraduated.Period))
		require.Equal(s.T(), billing.InvoiceLineStatusSplit, tieredGraduatedChildren.ParentLine.Status, "parent should be split [id=%s]", tieredGraduatedChildren.ParentLine.ID)
		require.Len(s.T(), tieredGraduatedChildren.ChildLines, 3, "there should be to child lines [id=%s]", tieredGraduatedChildren.ParentLine.ID)
		require.True(s.T(), tieredGraduatedChildren.ChildLines[0].Period.Equal(billing.Period{
			Start: periodStart.Truncate(time.Minute),
			End:   periodStart.Add(time.Hour).Truncate(time.Minute),
		}), "first child period should be truncated")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[1].Period.Equal(billing.Period{
			Start: periodStart.Add(time.Hour).Truncate(time.Minute),
			End:   periodStart.Add(2 * time.Hour).Truncate(time.Minute),
		}), "second child period should be between the first and the third child's period")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[1].InvoiceAt.Equal(periodStart.Add(2*time.Hour).Truncate(time.Minute)), "second child should be issued at the end of parent's period")
		require.True(s.T(), tieredGraduatedChildren.ChildLines[2].Period.Equal(billing.Period{
			Start: periodStart.Add(2 * time.Hour).Truncate(time.Minute),
			End:   periodEnd.Truncate(time.Minute),
		}), "third child period should be until the end of parent's period")

		// Details
		requireDetailedLines(s.T(), flatPerUsage, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.FlatPriceChildUniqueReferenceID: {
					Quantity:      1,
					PerUnitAmount: 100,
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount: 100,
			Total:  100,
		}, flatPerUsage.Totals)

		requireDetailedLines(s.T(), tieredVolume, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.VolumeUnitPriceChildUniqueReferenceID: {
					Quantity:      25,
					PerUnitAmount: 80,
				},
				lineservice.VolumeMinSpendChildUniqueReferenceID: {
					Quantity:      1,
					PerUnitAmount: 1000,
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount:       2000,
			ChargesTotal: 1000,
			Total:        3000,
		}, tieredVolume.Totals)

		requireDetailedLines(s.T(), tieredGraduated, lineExpectations{
			Details: map[string]feeLineExpect{
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 2): {
					Quantity:      5,
					PerUnitAmount: 90,
				},
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 3): {
					Quantity:      10,
					PerUnitAmount: 80,
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount: 1250,
			Total:  1250,
		}, tieredGraduated.Totals)

		// invoice totals
		requireTotals(s.T(), expectedTotals{
			Amount:       3350,
			ChargesTotal: 1000,
			Total:        4350,
		}, out[0].Totals)
	})
}

func (s *InvoicingTestSuite) TestUBPGraduatingFlatFeeTier1() {
	namespace := "ns-ubp-invoicing-graduated-flat-fee-tier-1"
	ctx := context.Background()

	periodStart := lo.Must(time.Parse(time.RFC3339, "2024-09-02T12:13:14Z"))
	periodEnd := lo.Must(time.Parse(time.RFC3339, "2024-09-03T12:13:14Z"))

	_ = s.InstallSandboxApp(s.T(), namespace)

	err := s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{
		{
			Namespace:     namespace,
			Slug:          "tiered-graduated",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
	})
	require.NoError(s.T(), err, "failed to replace meters")

	defer func() {
		err = s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{})
		require.NoError(s.T(), err, "failed to replace meters")
	}()

	// Let's initialize the mock streaming connector with data that is out of the period so that we
	// can start with empty values
	for _, slug := range []string{"flat-per-unit", "flat-per-usage", "tiered-graduated", "tiered-volume"} {
		s.MockStreamingConnector.AddSimpleEvent(slug, 0, periodStart.Add(-time.Minute))
	}

	defer s.MockStreamingConnector.Reset()

	// Let's create the features
	features := ubpFeatures{
		tieredGraduated: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "tiered-graduated",
			Key:       "tiered-graduated",
			MeterSlug: lo.ToPtr("tiered-graduated"),
		})),
	}

	// Given we have a test customer

	customerEntity := s.CreateTestCustomer(namespace, "test")
	require.NotNil(s.T(), customerEntity)
	require.NotEmpty(s.T(), customerEntity.ID)

	// Given we have a default profile for the namespace
	minimalCreateProfileInput := MinimalCreateProfileInputTemplate
	minimalCreateProfileInput.Namespace = namespace
	minimalCreateProfileInput.WorkflowConfig.Invoicing.ProgressiveBilling = true

	profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	var pendingLine *billing.Line
	s.Run("create pending invoice items", func() {
		// When we create pending invoice items
		pendingLines, err := s.BillingService.CreatePendingInvoiceLines(ctx,
			billing.CreateInvoiceLinesInput{
				Namespace: namespace,
				Lines: []billing.LineWithCustomer{
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - Tiered graduated",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.tieredGraduated.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.TieredPrice{
									Mode: productcatalog.GraduatedTieredPrice,
									Tiers: []productcatalog.PriceTier{
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(10)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(10),
											},
											FlatPrice: &productcatalog.PriceTierFlatPrice{
												Amount: alpacadecimal.NewFromFloat(100),
											},
										},
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(20)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(5),
											},
											FlatPrice: &productcatalog.PriceTierFlatPrice{
												Amount: alpacadecimal.NewFromFloat(200),
											},
										},
										{
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(80),
											},
										},
									},
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
				},
			},
		)
		require.NoError(s.T(), err)
		require.Len(s.T(), pendingLines, 1)

		pendingLine = pendingLines[0]
	})

	s.Run("create mid period invoice, no usage", func() {
		// Period
		asOf := periodStart.Add(time.Hour)
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)

		s.DebugDumpInvoice("mid period ubp progressive invoice, no usage", out[0])

		require.Len(s.T(), out[0].ValidationIssues, 0)

		invoiceLines := out[0].Lines.MustGet()
		require.Len(s.T(), invoiceLines, 1)

		// Let's resolve the lines by parent
		tieredGraduated := s.lineWithParent(invoiceLines, pendingLine.ID)

		requireTotals(s.T(), expectedTotals{
			Amount: 100,
			Total:  100,
		}, tieredGraduated.Totals)

		// Let's validate the output of the split itself
		tieredGraduatedChildren := s.getLineChildLines(ctx, namespace, tieredGraduated.ID)
		s.Len(tieredGraduatedChildren.ChildLines, 1)
		childLine := tieredGraduatedChildren.ChildLines[0]

		requireTotals(s.T(), expectedTotals{
			Amount: 100,
			Total:  100,
		}, childLine.Totals)
		s.Equal(*childLine.ChildUniqueReferenceID, "graduated-tiered-1-flat-price")
	})

	s.Run("create mid period invoice 2, no usage", func() {
		// Period
		asOf := periodStart.Add(2 * time.Hour)
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)

		s.DebugDumpInvoice("mid period ubp progressive 2nd invoice, no usage", out[0])

		require.Len(s.T(), out[0].ValidationIssues, 0)

		invoiceLines := out[0].Lines.MustGet()
		require.Len(s.T(), invoiceLines, 1)

		tieredGraduated := s.lineWithParent(invoiceLines, pendingLine.ID)

		requireTotals(s.T(), expectedTotals{
			Amount: 0,
			Total:  0,
		}, tieredGraduated.Totals)

		// Let's validate the output of the split itself
		tieredGraduatedChildren := s.getLineChildLines(ctx, namespace, tieredGraduated.ID)
		s.Len(tieredGraduatedChildren.ChildLines, 0)
	})

	s.Run("create new invoice, with usage", func() {
		// Period

		s.MockStreamingConnector.AddSimpleEvent("tiered-graduated", 15, periodStart.Add(time.Minute*130)) // 2h10m

		asOf := periodStart.Add(3 * time.Hour)
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)

		s.DebugDumpInvoice("mid period ubp progressive invoice, has usage", out[0])

		invoice := out[0]

		require.Len(s.T(), invoice.ValidationIssues, 0)

		invoiceLines := out[0].Lines.MustGet()
		require.Len(s.T(), invoiceLines, 1)

		expectedTotal := float64(10*10 /* usage for the first tier */ + 5*5 /* usage for the second tier */ + 200 /* flat price for the 2nd tier */)
		requireTotals(s.T(), expectedTotals{
			Amount: expectedTotal,
			Total:  expectedTotal,
		}, invoiceLines[0].Totals)

		requireDetailedLines(s.T(), invoiceLines[0], lineExpectations{
			Details: map[string]feeLineExpect{
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 1): {
					Quantity:      10,
					PerUnitAmount: 10,
				},
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 2): {
					Quantity:      5,
					PerUnitAmount: 5,
				},
				fmt.Sprintf(lineservice.GraduatedTieredFlatPriceChildUniqueReferenceID, 2): {
					Quantity:      1,
					PerUnitAmount: 200,
				},
			},
		})
	})
}

func (s *InvoicingTestSuite) TestUBPNonProgressiveInvoicing() {
	namespace := "ns-ubp-invoicing-non-progressive"
	ctx := context.Background()

	periodStart := lo.Must(time.Parse(time.RFC3339, "2024-09-02T12:13:14Z"))
	periodEnd := lo.Must(time.Parse(time.RFC3339, "2024-09-03T12:13:14Z"))

	_ = s.InstallSandboxApp(s.T(), namespace)

	err := s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{
		{
			Namespace:     namespace,
			Slug:          "flat-per-unit",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
		{
			Namespace:     namespace,
			Slug:          "flat-per-usage",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
		{
			Namespace:     namespace,
			Slug:          "tiered-graduated",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
		{
			Namespace:     namespace,
			Slug:          "tiered-volume",
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
	})
	require.NoError(s.T(), err, "failed to replace meters")

	defer func() {
		err = s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{})
		require.NoError(s.T(), err, "failed to replace meters")
	}()

	// Let's initialize the mock streaming connector with data that is out of the period so that we
	// can start with empty values
	for _, slug := range []string{"flat-per-unit", "flat-per-usage", "tiered-graduated", "tiered-volume"} {
		s.MockStreamingConnector.AddSimpleEvent(slug, 0, periodStart.Add(-time.Minute))
	}

	defer s.MockStreamingConnector.Reset()

	// Let's create the features
	// TODO[later]: we need to handle archived features, do we want to issue a warning? Can features be archived when used
	// by a draft invoice?
	features := ubpFeatures{
		flatPerUnit: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "flat-per-unit",
			Key:       "flat-per-unit",
			MeterSlug: lo.ToPtr("flat-per-unit"),
		})),
		flatPerUsage: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "flat-per-usage",
			Key:       "flat-per-usage",
			MeterSlug: lo.ToPtr("flat-per-usage"),
		})),
		tieredGraduated: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "tiered-graduated",
			Key:       "tiered-graduated",
			MeterSlug: lo.ToPtr("tiered-graduated"),
		})),
		tieredVolume: lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
			Namespace: namespace,
			Name:      "tiered-volume",
			Key:       "tiered-volume",
			MeterSlug: lo.ToPtr("tiered-volume"),
		})),
	}

	// Given we have a test customer

	customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
		Namespace: namespace,

		CustomerMutate: customer.CustomerMutate{
			Name:         "Test Customer",
			PrimaryEmail: lo.ToPtr("test@test.com"),
			BillingAddress: &models.Address{
				Country:     lo.ToPtr(models.CountryCode("US")),
				PostalCode:  lo.ToPtr("12345"),
				State:       lo.ToPtr("NY"),
				City:        lo.ToPtr("New York"),
				Line1:       lo.ToPtr("1234 Test St"),
				Line2:       lo.ToPtr("Apt 1"),
				PhoneNumber: lo.ToPtr("1234567890"),
			},
			Currency: lo.ToPtr(currencyx.Code(currency.USD)),
			UsageAttribution: customer.CustomerUsageAttribution{
				SubjectKeys: []string{"test"},
			},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), customerEntity)
	require.NotEmpty(s.T(), customerEntity.ID)

	// Given we have a default profile for the namespace
	minimalCreateProfileInput := MinimalCreateProfileInputTemplate
	minimalCreateProfileInput.Namespace = namespace

	profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	lines := ubpPendingLines{}
	s.Run("create pending invoice items", func() {
		// When we create pending invoice items
		pendingLines, err := s.BillingService.CreatePendingInvoiceLines(ctx,
			billing.CreateInvoiceLinesInput{
				Namespace: namespace,
				Lines: []billing.LineWithCustomer{
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - FLAT per unit",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.flatPerUnit.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.UnitPrice{
									Amount:        alpacadecimal.NewFromFloat(100),
									MaximumAmount: lo.ToPtr(alpacadecimal.NewFromFloat(2000)),
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - FLAT per any usage",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.flatPerUsage.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.FlatPrice{
									Amount:      alpacadecimal.NewFromFloat(100),
									PaymentTerm: productcatalog.InArrearsPaymentTerm,
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								ManagedBy: billing.ManuallyManagedLine,
								Currency:  currencyx.Code(currency.USD),
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - Tiered graduated",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.tieredGraduated.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.TieredPrice{
									Mode: productcatalog.GraduatedTieredPrice,
									Tiers: []productcatalog.PriceTier{
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(10)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(100),
											},
										},
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(20)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(90),
											},
										},
										{
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(80),
											},
										},
									},
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								Currency:  currencyx.Code(currency.USD),
								ManagedBy: billing.ManuallyManagedLine,
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - Tiered volume",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: features.tieredVolume.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.TieredPrice{
									Mode: productcatalog.VolumeTieredPrice,
									Tiers: []productcatalog.PriceTier{
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(10)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(100),
											},
										},
										{
											UpToAmount: lo.ToPtr(alpacadecimal.NewFromFloat(20)),
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(90),
											},
										},
										{
											UnitPrice: &productcatalog.PriceTierUnitPrice{
												Amount: alpacadecimal.NewFromFloat(80),
											},
										},
									},
									MinimumAmount: lo.ToPtr(alpacadecimal.NewFromFloat(3000)),
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
				},
			},
		)
		require.NoError(s.T(), err)
		require.Len(s.T(), pendingLines, 4)

		// The pending invoice items should be truncated to 1 min resolution (start => up to next, end down to previous)
		for _, line := range pendingLines {
			require.Equal(s.T(),
				billing.Period{
					Start: lo.Must(time.Parse(time.RFC3339, "2024-09-02T12:13:00Z")),
					End:   lo.Must(time.Parse(time.RFC3339, "2024-09-03T12:13:00Z")),
				},
				line.Period,
				"period should be truncated to 1 min resolution",
			)

			require.Equal(s.T(),
				line.InvoiceAt,
				periodEnd,
				"invoice at should be unchanged",
			)
		}

		lines = ubpPendingLines{
			flatPerUnit:     pendingLines[0],
			flatPerUsage:    pendingLines[1],
			tieredGraduated: pendingLines[2],
			tieredVolume:    pendingLines[3],
		}
	})

	// Usage:
	s.MockStreamingConnector.AddSimpleEvent("flat-per-unit", 10, periodStart)
	s.MockStreamingConnector.AddSimpleEvent("flat-per-unit", 20, periodStart.Add(time.Minute*100))
	s.MockStreamingConnector.AddSimpleEvent("tiered-graduated", 15, periodStart.Add(time.Minute*100))
	s.MockStreamingConnector.AddSimpleEvent("tiered-volume", 25, periodStart.Add(3*time.Hour))
	s.MockStreamingConnector.AddSimpleEvent("tiered-graduated", 15, periodStart.Add(3*time.Hour))

	s.Run("create invoice with empty truncated periods", func() {
		asOf := periodStart.Add(time.Second)
		_, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.ErrorIs(s.T(), err, billing.ErrInvoiceCreateNoLines)
		require.ErrorAs(s.T(), err, &billing.ValidationError{})
	})

	s.Run("create mid period invoice", func() {
		// Period
		asOf := periodStart.Add(time.Hour)
		_, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.ErrorIs(s.T(), err, billing.ErrInvoiceCreateNoLines)
		require.ErrorAs(s.T(), err, &billing.ValidationError{})
	})

	s.Run("create end of period invoice", func() {
		asOf := periodEnd
		out, err := s.BillingService.InvoicePendingLines(ctx, billing.InvoicePendingLinesInput{
			Customer: customerEntity.GetID(),
			AsOf:     &asOf,
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), out, 1)

		invoiceLines := out[0].Lines.MustGet()

		require.Len(s.T(), invoiceLines, 4)

		// Given that we didn't have to do a split the line IDs should be the same as the original lines
		flatPerUnit := s.lineByID(invoiceLines, lines.flatPerUnit.ID)
		flatPerUsage := s.lineByID(invoiceLines, lines.flatPerUsage.ID)
		tieredGraduated := s.lineByID(invoiceLines, lines.tieredGraduated.ID)
		tieredVolume := s.lineByID(invoiceLines, lines.tieredVolume.ID)

		expectedPeriod := billing.Period{
			Start: periodStart.Truncate(time.Minute),
			End:   periodEnd.Truncate(time.Minute),
		}
		for _, line := range []*billing.Line{flatPerUnit, flatPerUsage, tieredGraduated, tieredVolume} {
			require.True(s.T(), expectedPeriod.Equal(line.Period), "period should not be changed for the line items")
		}

		// Details
		requireDetailedLines(s.T(), flatPerUsage, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.FlatPriceChildUniqueReferenceID: {
					Quantity:      1,
					PerUnitAmount: 100,
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount: 100,
			Total:  100,
		}, flatPerUsage.Totals)

		requireDetailedLines(s.T(), flatPerUnit, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.UnitPriceUsageChildUniqueReferenceID: {
					Quantity:      30,
					PerUnitAmount: 100,
					Discounts: map[string]float64{
						billing.LineMaximumSpendReferenceID: 1000,
					},
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount:         3000,
			DiscountsTotal: 1000,
			Total:          2000,
		}, flatPerUnit.Totals)

		requireDetailedLines(s.T(), tieredVolume, lineExpectations{
			Details: map[string]feeLineExpect{
				lineservice.VolumeUnitPriceChildUniqueReferenceID: {
					Quantity:      25,
					PerUnitAmount: 80,
				},
				lineservice.VolumeMinSpendChildUniqueReferenceID: {
					Quantity:      1,
					PerUnitAmount: 1000,
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount:       2000,
			ChargesTotal: 1000,
			Total:        3000,
		}, tieredVolume.Totals)

		requireDetailedLines(s.T(), tieredGraduated, lineExpectations{
			Details: map[string]feeLineExpect{
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 1): {
					Quantity:      10,
					PerUnitAmount: 100,
				},
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 2): {
					Quantity:      10,
					PerUnitAmount: 90,
				},
				fmt.Sprintf(lineservice.GraduatedTieredPriceUsageChildUniqueReferenceID, 3): {
					Quantity:      10,
					PerUnitAmount: 80,
				},
			},
		})

		requireTotals(s.T(), expectedTotals{
			Amount: 2700,
			Total:  2700,
		}, tieredGraduated.Totals)

		// invoice totals
		requireTotals(s.T(), expectedTotals{
			Amount:         7800,
			ChargesTotal:   1000,
			DiscountsTotal: 1000,
			Total:          7800,
		}, out[0].Totals)
	})
}

func (s *InvoicingTestSuite) lineWithParent(lines []*billing.Line, parentID string) *billing.Line {
	s.T().Helper()
	for _, line := range lines {
		if line.ParentLineID != nil && *line.ParentLineID == parentID {
			return line
		}
	}

	require.Fail(s.T(), "line with parent not found")
	return nil
}

func (s *InvoicingTestSuite) lineByID(lines []*billing.Line, id string) *billing.Line {
	s.T().Helper()
	for _, line := range lines {
		if line.ID == id {
			return line
		}
	}

	require.Fail(s.T(), "line not found")
	return nil
}

type getChlildLinesResponse struct {
	ParentLine *billing.Line
	ChildLines []*billing.Line
}

func (s *InvoicingTestSuite) getLineChildLines(ctx context.Context, ns string, parentID string) getChlildLinesResponse {
	res, err := s.BillingAdapter.ListInvoiceLines(ctx, billing.ListInvoiceLinesAdapterInput{
		Namespace:                  ns,
		ParentLineIDs:              []string{parentID},
		ParentLineIDsIncludeParent: true,
	})
	require.NoError(s.T(), err)

	if len(res) == 0 {
		require.Fail(s.T(), "no child lines found")
	}

	response := getChlildLinesResponse{}

	for _, line := range res {
		if line.ID == parentID {
			response.ParentLine = line
		} else {
			response.ChildLines = append(response.ChildLines, line)
		}
	}

	slices.SortFunc(response.ChildLines, func(a, b *billing.Line) int {
		switch {
		case a.Period.Start.Equal(b.Period.Start):
			return 0
		case a.Period.Start.Before(b.Period.Start):
			return -1
		default:
			return 1
		}
	})

	require.NotEmpty(s.T(), response.ParentLine.ID)
	return response
}

type ubpPendingLines struct {
	flatPerUnit     *billing.Line
	flatPerUsage    *billing.Line
	tieredGraduated *billing.Line
	tieredVolume    *billing.Line
}

type ubpFeatures struct {
	flatPerUnit     feature.Feature
	flatPerUsage    feature.Feature
	tieredGraduated feature.Feature
	tieredVolume    feature.Feature
}

type lineExpectations struct {
	Details map[string]feeLineExpect
}

type feeLineExpect struct {
	Quantity      float64
	PerUnitAmount float64
	Discounts     map[string]float64
}

func requireDetailedLines(t *testing.T, line *billing.Line, expectations lineExpectations) {
	t.Helper()
	require.NotNil(t, line)
	children := line.Children.MustGet()

	require.Len(t, children, len(expectations.Details))

	detailsById := lo.GroupBy(children, func(l *billing.Line) string {
		return *l.ChildUniqueReferenceID
	})

	for key, expect := range expectations.Details {
		require.Contains(t, detailsById, key, "detail %s should be present", key)
		detail := detailsById[key][0]

		require.Equal(t, detail.Type, billing.InvoiceLineTypeFee, "line type should be fee")
		require.Equal(t, expect.Quantity, detail.FlatFee.Quantity.InexactFloat64(), "quantity should match")
		require.Equal(t, expect.PerUnitAmount, detail.FlatFee.PerUnitAmount.InexactFloat64(), "per unit amount should match")

		discounts := detail.Discounts.MustGet()
		require.Len(t, discounts, len(expect.Discounts), "discounts should match")

		discountsById := lo.GroupBy(discounts, func(d billing.LineDiscount) string {
			return *d.ChildUniqueReferenceID
		})

		for discountType, discountExpect := range expect.Discounts {
			require.Contains(t, discountsById, discountType, "discount %s should be present", discountType)
			discount := discountsById[discountType][0]

			require.Equal(t, discountExpect, discount.Amount.InexactFloat64(), "discount amount should match")
		}
	}
}

type expectedTotals struct {
	// Amount is the total amount value of the line before taxes, discounts and commitments
	Amount float64 `json:"amount"`
	// ChargesTotal is the amount of value of the line that are due to additional charges
	ChargesTotal float64 `json:"chargesTotal"`
	// DiscountsTotal is the amount of value of the line that are due to discounts
	DiscountsTotal float64 `json:"discountsTotal"`

	// TaxesInclusiveTotal is the total amount of taxes that are included in the line
	TaxesInclusiveTotal float64 `json:"taxesInclusiveTotal"`
	// TaxesExclusiveTotal is the total amount of taxes that are excluded from the line
	TaxesExclusiveTotal float64 `json:"taxesExclusiveTotal"`
	// TaxesTotal is the total amount of taxes that are included in the line
	TaxesTotal float64 `json:"taxesTotal"`

	// Total is the total amount value of the line after taxes, discounts and commitments
	Total float64 `json:"total"`
}

func requireTotals(t *testing.T, expected expectedTotals, totals billing.Totals) {
	t.Helper()
	totalsFloat := expectedTotals{
		Amount:              totals.Amount.InexactFloat64(),
		ChargesTotal:        totals.ChargesTotal.InexactFloat64(),
		DiscountsTotal:      totals.DiscountsTotal.InexactFloat64(),
		TaxesInclusiveTotal: totals.TaxesInclusiveTotal.InexactFloat64(),
		TaxesExclusiveTotal: totals.TaxesExclusiveTotal.InexactFloat64(),
		TaxesTotal:          totals.TaxesTotal.InexactFloat64(),
		Total:               totals.Total.InexactFloat64(),
	}

	require.Equal(t, expected, totalsFloat)
}

func (s *InvoicingTestSuite) TestGatheringInvoiceRecalculation() {
	namespace := "ns-gathering-invoice-calc"
	ctx := context.Background()

	periodStart := lo.Must(time.Parse(time.RFC3339, "2024-09-02T12:13:14Z"))
	periodEnd := lo.Must(time.Parse(time.RFC3339, "2024-09-03T12:13:14Z"))
	clock.SetTime(periodStart)
	defer clock.ResetTime()

	_ = s.InstallSandboxApp(s.T(), namespace)

	meterSlug := "flat-per-unit"

	err := s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{
		{
			Namespace:     namespace,
			Slug:          meterSlug,
			WindowSize:    meter.WindowSizeMinute,
			Aggregation:   meter.MeterAggregationSum,
			EventType:     "test",
			ValueProperty: "$.value",
		},
	})
	require.NoError(s.T(), err, "failed to replace meters")

	defer func() {
		err = s.MeterAdapter.ReplaceMeters(ctx, []meter.Meter{})
		require.NoError(s.T(), err, "failed to replace meters")
	}()

	// Let's initialize the mock streaming connector with data that is out of the period so that we
	// can start with empty values
	s.MockStreamingConnector.AddSimpleEvent(meterSlug, 0, periodStart.Add(-time.Minute))

	defer s.MockStreamingConnector.Reset()

	flatPerUnitFeature := lo.Must(s.FeatureService.CreateFeature(ctx, feature.CreateFeatureInputs{
		Namespace: namespace,
		Name:      "flat-per-unit",
		Key:       "flat-per-unit",
		MeterSlug: lo.ToPtr("flat-per-unit"),
	}))

	// Given we have a test customer

	customerEntity, err := s.CustomerService.CreateCustomer(ctx, customer.CreateCustomerInput{
		Namespace: namespace,

		CustomerMutate: customer.CustomerMutate{
			Name:         "Test Customer",
			PrimaryEmail: lo.ToPtr("test@test.com"),
			BillingAddress: &models.Address{
				Country: lo.ToPtr(models.CountryCode("US")),
			},
			Currency: lo.ToPtr(currencyx.Code(currency.USD)),
			UsageAttribution: customer.CustomerUsageAttribution{
				SubjectKeys: []string{"test"},
			},
		},
	})
	require.NoError(s.T(), err)
	require.NotNil(s.T(), customerEntity)
	require.NotEmpty(s.T(), customerEntity.ID)

	// Given we have a default profile for the namespace
	minimalCreateProfileInput := MinimalCreateProfileInputTemplate
	minimalCreateProfileInput.Namespace = namespace

	profile, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	s.Run("create pending invoice items", func() {
		// When we create pending invoice items
		pendingLines, err := s.BillingService.CreatePendingInvoiceLines(ctx,
			billing.CreateInvoiceLinesInput{
				Namespace: namespace,
				Lines: []billing.LineWithCustomer{
					{
						Line: billing.Line{
							LineBase: billing.LineBase{
								Period:    billing.Period{Start: periodStart, End: periodEnd},
								InvoiceAt: periodEnd,
								Currency:  currencyx.Code(currency.USD),
								ManagedBy: billing.ManuallyManagedLine,
								Type:      billing.InvoiceLineTypeUsageBased,
								Name:      "UBP - FLAT per unit",
							},
							UsageBased: &billing.UsageBasedLine{
								FeatureKey: flatPerUnitFeature.Key,
								Price: productcatalog.NewPriceFrom(productcatalog.UnitPrice{
									Amount:        alpacadecimal.NewFromFloat(100),
									MaximumAmount: lo.ToPtr(alpacadecimal.NewFromFloat(2000)),
								}),
							},
						},
						CustomerID: customerEntity.ID,
					},
				},
			},
		)
		require.NoError(s.T(), err)
		require.Len(s.T(), pendingLines, 1)
	})

	s.Run("fetch gathering invoice", func() {
		invoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Namespaces:       []string{namespace},
			Customers:        []string{customerEntity.ID},
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Expand: billing.InvoiceExpand{
				RecalculateGatheringInvoice: true,
			},
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), invoices.Items, 1)

		gatheringInvoice := invoices.Items[0]
		require.Equal(s.T(), float64(0), gatheringInvoice.Totals.Total.InexactFloat64())
	})

	// when we have some traffic on the meter, the invoice should be recalculated
	s.Run("invoice recalculation", func() {
		s.MockStreamingConnector.AddSimpleEvent(meterSlug, 10, periodStart.Add(time.Minute))

		invoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Namespaces:       []string{namespace},
			Customers:        []string{customerEntity.ID},
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Expand: billing.InvoiceExpand{
				RecalculateGatheringInvoice: true,
			},
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), invoices.Items, 1)

		gatheringInvoice := invoices.Items[0]
		require.Equal(s.T(), float64(1000), gatheringInvoice.Totals.Total.InexactFloat64())
	})

	// Max spend is reached
	s.Run("invoice recalculation - max spend", func() {
		s.MockStreamingConnector.AddSimpleEvent(meterSlug, 30, periodStart.Add(2*time.Minute))

		invoices, err := s.BillingService.ListInvoices(ctx, billing.ListInvoicesInput{
			Namespaces:       []string{namespace},
			Customers:        []string{customerEntity.ID},
			ExtendedStatuses: []billing.InvoiceStatus{billing.InvoiceStatusGathering},
			Expand: billing.InvoiceExpand{
				RecalculateGatheringInvoice: true,
			},
		})

		require.NoError(s.T(), err)
		require.Len(s.T(), invoices.Items, 1)

		gatheringInvoice := invoices.Items[0]
		requireTotals(s.T(), expectedTotals{
			Amount:         4000,
			Total:          2000,
			DiscountsTotal: 2000,
		}, gatheringInvoice.Totals)
	})
}
