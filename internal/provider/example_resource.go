package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &exampleResource{}

func NewExampleResource() resource.Resource {
	return &exampleResource{}
}

type exampleResource struct{}

type exampleResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

func (r *exampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (r *exampleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages an example thing resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Example identifier",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the thing",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of the thing",
				Optional:    true,
			},
			"last_updated": schema.StringAttribute{
				Description: "Timestamp of last update",
				Computed:    true,
			},
		},
	}
}

// Create creates the resource and sets initial state
func (r *exampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan exampleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate an ID for demonstration
	plan.ID = types.StringValue(fmt.Sprintf("example-%d", time.Now().Unix()))
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	// In real provider: make API call to create resource

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read refreshes the Terraform state with latest data
func (r *exampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state exampleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// In real provider: make API call to read resource
	// If resource not found: call resp.State.RemoveResource()

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets updated state
func (r *exampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan exampleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	// In real provider: make API call to update resource

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes state
func (r *exampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state exampleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// In real provider: make API call to delete resource

	// State automatically removed if Delete returns without error
}
