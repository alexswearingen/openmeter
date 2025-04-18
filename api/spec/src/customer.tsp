import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";

import "./rest.tsp";
import "./errors.tsp";
import "./types.tsp";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

namespace OpenMeter;

/**
 * Customer API.
 */
@route("/api/v1/customers")
@tag("Customers")
interface Customers {
  /**
   * Create a new customer.
   */
  @post
  @operationId("createCustomer")
  @summary("Create customer")
  create(@body customer: Rest.Resource.ResourceCreateModel<Customer>): {
    @statusCode _: 201;
    @body body: Customer;
  } | CommonErrors;

  /**
   * List customers.
   */
  @get
  @operationId("listCustomers")
  @summary("List customers")
  list(...ListCustomersParams): PaginatedResponse<Customer> | CommonErrors;

  /**
   * Get a customer by ID or key.
   */
  @get
  @route("/{customerIdOrKey}")
  @operationId("getCustomer")
  @summary("Get customer")
  get(
    @path customerIdOrKey: ULIDOrKey,
  ): Customer | NotFoundError | CommonErrors;

  /**
   * Update a customer by ID.
   */
  @put
  @route("/{customerIdOrKey}")
  @operationId("updateCustomer")
  @summary("Update customer")
  update(
    @path customerIdOrKey: ULIDOrKey,
    @body customer: Rest.Resource.ResourceReplaceModel<Customer>,
  ): Customer | NotFoundError | CommonErrors;

  /**
   * Delete a customer by ID.
   */
  @delete
  @route("/{customerIdOrKey}")
  @operationId("deleteCustomer")
  @summary("Delete customer")
  delete(@path customerIdOrKey: ULIDOrKey): {
    @statusCode _: 204;
  } | NotFoundError | CommonErrors;

  /**
   * List customers app data.
   */
  @get
  @route("/{customerIdOrKey}/apps")
  @operationId("listCustomerAppData")
  @summary("List customer app data")
  listAppData(
    @path customerIdOrKey: ULIDOrKey,
    ...ListCustomerAppDataParams,
  ): PaginatedResponse<App.CustomerAppData> | NotFoundError | CommonErrors;

  /**
   * Upsert customer app data.
   */
  @put
  @route("/{customerIdOrKey}/apps")
  @operationId("upsertCustomerAppData")
  @summary("Upsert customer app data")
  upsertAppData(
    @path customerIdOrKey: ULIDOrKey,
    @body appData: App.CustomerAppData[],
  ): App.CustomerAppData[] | NotFoundError | CommonErrors;

  /**
   * Delete customer app data.
   */
  @delete
  @route("/{customerIdOrKey}/apps/{appId}")
  @operationId("deleteCustomerAppData")
  @summary("Delete customer app data")
  deleteAppData(@path customerIdOrKey: ULIDOrKey, @path appId: ULID): {
    @statusCode _: 204;
  } | NotFoundError | CommonErrors;

  /**
   * Lists all subscriptions for a customer.
   */
  @get
  @route("/{customerIdOrKey}/subscriptions")
  @operationId("listCustomerSubscriptions")
  @summary("List customer subscriptions")
  listCustomerSubscriptions(
    @path customerIdOrKey: ULIDOrKey,
    ...QueryPagination,
  ): PaginatedResponse<ProductCatalog.Subscription> | NotFoundError | CommonErrors;

  /**
   * Checks customer access to a given feature (by key). All entitlement types share the hasAccess property in their value response, but multiple other properties are returned based on the entitlement type.
   */
  @get
  @operationId("getCustomerEntitlementValue")
  @route("/{customerIdOrKey}/entitlements/{featureKey}/value")
  @summary("Get entitlement value")
  getCustomerEntitlementValue(
    @path customerIdOrKey: ULIDOrKey,
    @path featureKey: Key,
    @query(#{ explode: true }) time?: DateTime,
  ): Entitlements.EntitlementValue | OpenMeter.CommonErrors | OpenMeter.NotFoundError;
}

/**
 * Query params for listing customers.
 */
@friendlyName("queryCustomerList")
model ListCustomersParams {
  ...QueryPagination;

  /**
   * Order customers by.
   * @TODO: set default value to CustomerOrderBy.name
   */
  ...OpenMeter.QueryOrdering<CustomerOrderBy>;

  /**
   * Include deleted customers.
   */
  @query
  includeDeleted?: boolean = false;

  /**
   * Filter customers by key.
   * Case-sensitive exact match.
   */
  @query
  key?: string;

  /**
   * Filter customers by name.
   * Case-insensitive partial match.
   */
  @query
  @example("ACME")
  name?: string;

  /**
   * Filter customers by primary email.
   * Case-insensitive partial match.
   */
  @query
  @example("acme@test.com")
  primaryEmail?: string;

  /**
   * Filter customers by usage attribution subject.
   * Case-insensitive partial match.
   */
  @query
  @example("my_subject_key")
  subject?: string;

  /**
   * Filter customers by the plan key of their susbcription.
   */
  @query
  @example("pro_plan")
  planKey?: string;
}

/**
 * Query params for listing customers.
 */
@friendlyName("queryCustomerList")
model ListCustomerAppDataParams {
  ...QueryPagination;

  /**
   * Filter customer data by app type.
   */
  @query
  @example(App.Type.Stripe)
  type?: App.Type;
}

/**
 * Order by options for customers.
 */
@friendlyName("CustomerOrderBy")
enum CustomerOrderBy {
  id: "id",
  name: "name",
  createdAt: "createdAt",
}

/**
 * A customer object.
 */
@friendlyName("Customer")
@example(#{
  id: "01G65Z755AFWAKHE12NY0CQ9FH",
  name: "ACME Inc.",
  usageAttribution: #{ subjectKeys: #["my_subject_key"] },
  createdAt: DateTime.fromISO("2024-01-01T01:01:01.001Z"),
  updatedAt: DateTime.fromISO("2024-01-01T01:01:01.001Z"),
})
model Customer {
  ...Resource;

  /**
   * An optional unique key of the customer.
   * Useful to reference the customer in external systems.
   * For example, your database ID.
   */
  @summary("Key")
  key?: ExternalKey;

  /**
   * Mapping to attribute metered usage to the customer
   */
  @summary("Usage Attribution")
  usageAttribution: CustomerUsageAttribution;

  /**
   * The primary email address of the customer.
   */
  @summary("Primary Email")
  primaryEmail?: string;

  /**
   * Currency of the customer.
   * Used for billing, tax and invoicing.
   */
  @summary("Currency")
  currency?: CurrencyCode;

  /**
   * The billing address of the customer.
   * Used for tax and invoicing.
   */
  @summary("Billing Address")
  billingAddress?: Address;

  /**
   * The ID of the Subscription if the customer has one.
   */
  @summary("Current Subscription ID")
  @visibility(Lifecycle.Read)
  currentSubscriptionId?: ULID;
}

/**
 * Mapping to attribute metered usage to the customer.
 * One customer can have multiple subjects,
 * but one subject can only belong to one customer.
 */
@friendlyName("CustomerUsageAttribution")
model CustomerUsageAttribution {
  /**
   * The subjects that are attributed to the customer.
   * @TODO allow multiple subjects
   */
  @summary("SubjectKeys")
  @minItems(1)
  @maxItems(1)
  subjectKeys: string[];
}
