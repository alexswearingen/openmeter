// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"github.com/openmeterio/openmeter/pkg/pagination"
)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bs *BalanceSnapshotQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BalanceSnapshot], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bs.ctx.Offset = &zero
	bs.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bs.Clone()
	pagedQuery := bs

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BalanceSnapshot]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BalanceSnapshot] = (*BalanceSnapshotQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bi *BillingInvoiceQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BillingInvoice], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bi.ctx.Offset = &zero
	bi.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bi.Clone()
	pagedQuery := bi

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BillingInvoice]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BillingInvoice] = (*BillingInvoiceQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bii *BillingInvoiceItemQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BillingInvoiceItem], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bii.ctx.Offset = &zero
	bii.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bii.Clone()
	pagedQuery := bii

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BillingInvoiceItem]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BillingInvoiceItem] = (*BillingInvoiceItemQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bp *BillingProfileQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BillingProfile], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bp.ctx.Offset = &zero
	bp.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bp.Clone()
	pagedQuery := bp

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BillingProfile]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BillingProfile] = (*BillingProfileQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bwc *BillingWorkflowConfigQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BillingWorkflowConfig], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bwc.ctx.Offset = &zero
	bwc.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bwc.Clone()
	pagedQuery := bwc

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BillingWorkflowConfig]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BillingWorkflowConfig] = (*BillingWorkflowConfigQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bwco *BillingWorkflowConfigOverrideQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BillingWorkflowConfigOverride], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bwco.ctx.Offset = &zero
	bwco.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bwco.Clone()
	pagedQuery := bwco

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BillingWorkflowConfigOverride]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BillingWorkflowConfigOverride] = (*BillingWorkflowConfigOverrideQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (c *CustomerQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Customer], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	c.ctx.Offset = &zero
	c.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := c.Clone()
	pagedQuery := c

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Customer]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Customer] = (*CustomerQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (cs *CustomerSubjectsQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*CustomerSubjects], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	cs.ctx.Offset = &zero
	cs.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := cs.Clone()
	pagedQuery := cs

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*CustomerSubjects]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*CustomerSubjects] = (*CustomerSubjectsQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (e *EntitlementQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Entitlement], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	e.ctx.Offset = &zero
	e.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := e.Clone()
	pagedQuery := e

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Entitlement]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Entitlement] = (*EntitlementQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (f *FeatureQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Feature], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	f.ctx.Offset = &zero
	f.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := f.Clone()
	pagedQuery := f

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Feature]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Feature] = (*FeatureQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (gr *GrantQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Grant], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	gr.ctx.Offset = &zero
	gr.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := gr.Clone()
	pagedQuery := gr

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Grant]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Grant] = (*GrantQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (nc *NotificationChannelQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*NotificationChannel], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	nc.ctx.Offset = &zero
	nc.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := nc.Clone()
	pagedQuery := nc

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*NotificationChannel]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*NotificationChannel] = (*NotificationChannelQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (ne *NotificationEventQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*NotificationEvent], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	ne.ctx.Offset = &zero
	ne.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := ne.Clone()
	pagedQuery := ne

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*NotificationEvent]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*NotificationEvent] = (*NotificationEventQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (neds *NotificationEventDeliveryStatusQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*NotificationEventDeliveryStatus], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	neds.ctx.Offset = &zero
	neds.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := neds.Clone()
	pagedQuery := neds

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*NotificationEventDeliveryStatus]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*NotificationEventDeliveryStatus] = (*NotificationEventDeliveryStatusQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (nr *NotificationRuleQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*NotificationRule], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	nr.ctx.Offset = &zero
	nr.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := nr.Clone()
	pagedQuery := nr

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*NotificationRule]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*NotificationRule] = (*NotificationRuleQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (ur *UsageResetQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*UsageReset], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	ur.ctx.Offset = &zero
	ur.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := ur.Clone()
	pagedQuery := ur

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*UsageReset]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*UsageReset] = (*UsageResetQuery)(nil)
