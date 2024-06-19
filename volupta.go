import (
	"context"
	"fmt"
	"io"

	workflows "cloud.google.com/go/workflows/apiv1"
	"cloud.google.com/go/workflows/apiv1/workflowspb"
)

// waitForLocalExecution waits for a local execution to complete.
func waitForLocalExecution(w io.Writer, projectID, location, executionID string) error {
	// projectID := "my-project-id"
	// location := "us-central1"
	// executionID := "my-execution-id"
	ctx := context.Background()
	client, err := workflows.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("workflows.NewClient: %v", err)
	}
	defer client.Close()

	req := &workflowspb.GetExecutionRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/executions/%s", projectID, location, executionID),
	}

	resp, err := client.GetExecution(ctx, req)
	if err != nil {
		return fmt.Errorf("GetExecution: %v", err)
	}

	fmt.Fprintf(w, "Execution state: %v\n", resp.State)

	return nil
}
  
