package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/layer5io/meshkit/logger"
	"github.com/meshery/helm-kanvas-snapshot/internal/errors"
	"github.com/meshery/helm-kanvas-snapshot/internal/log"
	"github.com/spf13/cobra"
)

var (
	ProviderToken          string
	MesheryApiBaseUrl      string
	MesheryCloudApiBaseUrl string
	Log                    logger.Handler
)

var (
	chartURI   string
	email      string
	designName string
)

var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

var generateKanvasSnapshotCmd = &cobra.Command{
	Use:   "kanvas",
	Short: "Generate a Kanvas snapshot using a Helm chart",
	Long: `Generate a Kanvas snapshot by providing a Helm chart URI.

		This command allows you to generate a snapshot in Meshery using a Helm chart.

		Example usage:

		helm kanvas-snapshot -n nginx-helm -f https://meshery.github.io/meshery.io/charts/meshery-v0.7.109.tgz -e your-email@example.com

		Flags:
		-f, --file string	URI to Helm chart (required)
		-n, --name string	(optional name for the Meshery design
		-e, --email string	email address to notify when snapshot is ready (required)
		-h			Help for Helm Kanvas Snapshot plugin`,

	RunE: func(cmd *cobra.Command, args []string) error {
		// Use the extracted name from URI if not provided
		if designName == "" {
			designName = ExtractNameFromURI(chartURI)
			Log.Warnf("No design name provided. Using extracted name: %s", designName)
		}
		if email != "" && !isValidEmail(email) {
			handleError(errors.ErrInvalidEmailFormat(email))
		}
		// Create Meshery Snapshot
		designID, err := CreateMesheryDesign(chartURI, designName, email)
		if err != nil {
			handleError(errors.ErrCreatingMesheryDesign(err))
		}

		assetLocation := fmt.Sprintf("https://raw.githubusercontent.com/layer5labs/meshery-extensions-packages/master/action-assets/%s.png", designID)

		// Generate Snapshot
		err = GenerateSnapshot(designID, chartURI, email, assetLocation)
		if err != nil {
			handleError(errors.ErrGeneratingSnapshot(err))
		}

		if email == "" {
			loader(2*time.Minute + 40*time.Second) // Loader running for 2 minutes and 40 seconds
			Log.Infof("\nSnapshot generated. Snapshot URL: %s\n", assetLocation)
		} else {
			Log.Info("You will be notified via email when your snapshot is ready.")
		}
		return nil
	},
}

type MesheryDesignPayload struct {
	Save  bool   `json:"save"`
	URL   string `json:"url"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func loader(duration time.Duration) {
	total := int(duration.Seconds()) // Total time in seconds
	progress := 0

	for progress <= total {
		printProgressBar(progress, total)
		time.Sleep(1 * time.Second) // Sleep for 1 second to update progress
		progress++
	}
	fmt.Println() // Print a new line at the end for better output formatting
}

func printProgressBar(progress, total int) {
	barWidth := 25

	percentage := float64(progress) / float64(total)
	barProgress := int(percentage * float64(barWidth))

	bar := "[" + fmt.Sprintf("%s%s", repeat("=", barProgress), repeat("-", barWidth-barProgress)) + "]"
	fmt.Printf("\rProgress %s %.2f%% Complete", bar, percentage*100)
}

// Helper function to repeat a character n times
func repeat(char string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += char
	}
	return result
}

// ExtractNameFromURI extracts the name from the URI
func ExtractNameFromURI(uri string) string {
	filename := filepath.Base(uri)
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func handleError(err error) {
	if err == nil {
		return
	}
	if Log != nil {
		Log.Error(err)
	} else {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	}
	os.Exit(1)
}

func CreateMesheryDesign(uri, name, email string) (string, error) {
	payload := MesheryDesignPayload{
		Save:  true,
		URL:   uri,
		Name:  name,
		Email: email,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		Log.Info("Failed to marshal payload:", err)
		return "", errors.ErrDecodingAPI(err)
	}

	sourceType := "Helm Chart"
	encodedChartType := url.PathEscape(sourceType)
	fullURL := fmt.Sprintf("%s/api/pattern/%s", MesheryApiBaseUrl, encodedChartType)

	// Create the request
	req, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		Log.Info("Failed to create new request:", err)
		return "", errors.ErrHTTPPostRequest(err)
	}

	// Set headers and log them
	req.Header.Set("Cookie", ProviderToken)
	req.Header.Set("Origin", MesheryApiBaseUrl)
	req.Header.Set("Host", MesheryApiBaseUrl)
	req.Header.Set("Content-Type", "text/plain;charset=UTF-8")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", errors.ErrHTTPPostRequest(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", errors.ErrUnexpectedResponseCode(resp.StatusCode, string(body))
	}

	// Decode response
	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		return "", errors.ErrDecodingAPI(fmt.Errorf("failed to decode json. body: %s, error: %w", body, err))
	}

	// Extracting the design ID from the result
	if len(result) > 0 {
		if id, ok := result[0]["id"].(string); ok {
			Log.Infof("Successfully created Meshery design. ID: %s", id)
			return id, nil
		}
	}

	return "", errors.ErrCreatingMesheryDesign(fmt.Errorf("failed to extract design ID from response"))
}

func GenerateSnapshot(designID, chartURI, email, assetLocation string) error {
	payload := map[string]interface{}{
		"Payload": map[string]string{
			"application_type": "Helm Chart",
			"designID":         designID,
			"email":            email,
			"assetLocation":    assetLocation,
		},
	}

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Create the POST request
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/api/integrations/trigger/workflow", MesheryCloudApiBaseUrl),
		bytes.NewBuffer(payloadBytes),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Cookie", ProviderToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", fmt.Sprintf("%s/dashboard", MesheryCloudApiBaseUrl))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return errors.ErrHTTPPostRequest(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return errors.ErrUnexpectedResponseCode(resp.StatusCode, string(body))
	}

	// Decode response
	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		body, _ := io.ReadAll(resp.Body)
		return errors.ErrDecodingAPI(fmt.Errorf("failed to decode json. body: %s, error: %w", body, err))
	}

	return nil
}

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func Main() {
	generateKanvasSnapshotCmd.Flags().StringVarP(&chartURI, "file", "f", "", "URI to Helm chart (required)")
	generateKanvasSnapshotCmd.Flags().StringVarP(&designName, "design-name", "n", "", "Optional name for the Meshery design")
	generateKanvasSnapshotCmd.Flags().StringVarP(&email, "email", "e", "", "Optional email to associate with the Meshery design")

	_ = generateKanvasSnapshotCmd.MarkFlagRequired("file")
	_ = generateKanvasSnapshotCmd.MarkFlagRequired("email")

	if err := generateKanvasSnapshotCmd.Execute(); err != nil {
		Log.Error(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		Log = log.SetupMeshkitLogger("helm-kanvas-snapshot", false, os.Stdout)
	})
}
