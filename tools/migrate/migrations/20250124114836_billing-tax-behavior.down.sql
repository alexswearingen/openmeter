-- reverse: modify "billing_workflow_configs" table
ALTER TABLE "billing_workflow_configs" DROP COLUMN "invoice_tax_behavior";
-- reverse: modify "billing_customer_overrides" table
ALTER TABLE "billing_customer_overrides" DROP COLUMN "invoice_tax_behavior";
