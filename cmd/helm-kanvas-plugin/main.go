package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/layer5io/meshkit/logger"
	"github.com/meshery/helm-kanvas-snapshot/internal/errors"
	"github.com/meshery/helm-kanvas-snapshot/internal/log"
)

var (
	GithubToken            string
	MesheryToken           string
	MesheryCloudApiCookie  string
	MesheryApiCookie       string
	Owner                  string
	Repo                   string
	Workflow               string
	Branch                 string
	MesheryApiBaseUrl      string
	MesheryCloudApiBaseUrl string
	SystemID               string
	HelmPluginDir          string
	logFilePath            string
	Log                    logger.Handler
	LogError               logger.Handler
)

type Config struct {
	GithubToken            string
	MesheryToken           string
	MesheryCloudApiCookie  string
	MesheryApiCookie       string
	HelmPluginDir          string
	Owner                  string
	Repo                   string
	Workflow               string
	Branch                 string
	MesheryApiBaseUrl      string
	MesheryCloudApiBaseUrl string
	SystemID               string
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

}

func printProgressBar(progress, total int) {
	barWidth := 25

	percentage := float64(progress) / float64(total)
	barProgress := int(percentage * float64(barWidth))

	bar := "[" + fmt.Sprintf("%s%s", repeat("=", barProgress), repeat("-", barWidth-barProgress)) + "]"
	Log.Infof("\rProgress %s %.2f%% Complete", bar, percentage*100)
}

// Helper function to repeat a character n times
func repeat(char string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += char
	}
	return result
}

func CreateLogFile() error {
	cwd, err := os.Getwd()
	if err != nil {
		return errors.ErrGettingCWD(err)
	}

	logDir := filepath.Join(fmt.Sprintf("%s/snapshot", cwd), "log")

	// Create the /log directory if it doesn't exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err = os.Mkdir(logDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	logFilePath = filepath.Join(logDir, "snapshot")

	// Create the log file if it doesn't exist
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		file, err := os.Create(logFilePath)
		if err != nil {
			return errors.ErrCreatingLogFile(err, logFilePath)
		}
		file.Close()
	}

	Log.Infof("Log file created successfully at: %s\n", logFilePath)
	return nil
}

// ExtractNameFromURI extracts the name from the URI
func ExtractNameFromURI(uri string) string {
	filename := filepath.Base(uri)
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func handleError(err error) {
	if err != nil {
		LogError.Error(err)

		// Check if the log file exists before writing to it
		if _, fileErr := os.Stat(logFilePath); !os.IsNotExist(fileErr) {
			os.WriteFile(logFilePath, []byte(fmt.Sprintf("%s - %s", time.Now().Format("2006-01-02 15:04:05"), err.Error())), os.ModeAppend)
		}
		os.Exit(1)
	}
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
		return "", err
	}
	sourceType := "Helm Chart"
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/pattern/%s", MesheryApiBaseUrl, sourceType), bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", MesheryApiCookie)
	req.Header.Set("Origin", MesheryApiBaseUrl)
	req.Header.Set("Host", MesheryApiBaseUrl)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.ReadAll(resp.Body)
		return "", err
	}
	// Expecting a JSON array in the response
	var result []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if len(result) > 0 {
		if id, ok := result[0]["id"].(string); ok {
			return id, nil
		}
	}

	return "", errors.ErrHTTPPostRequest(err)
}

func GenerateSnapshot(designID, chartURI, email, assetLocation string) error {

	payload := map[string]interface{}{
		"Owner":        Owner,
		"Repo":         Repo,
		"Workflow":     Workflow,
		"Branch":       Branch,
		"github_token": GithubToken,
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

	req.Header.Set("Cookie", MesheryCloudApiCookie)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("SystemID", SystemID)
	req.Header.Set("Referer", fmt.Sprintf("%s/dashboard", MesheryCloudApiBaseUrl))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		_, err := io.ReadAll(resp.Body)
		return err
	}

	return nil
}

func main() {
	Log = log.SetupMeshkitLogger("kanvas-snapshot", false, os.Stdout)

	chartURI := flag.String("f", "", "URI to Helm chart")
	name := flag.String("n", "", "Optional name for the Meshery design")
	email := flag.String("e", "", "Optional email to associate with the Meshery design")
	flag.Parse()

	if chartURI == nil || *chartURI == "" {
		Log.Info("url to helm chart is required")
		os.Exit(1)
	}

	CreateLogFile()

	// Use the extracted name from URI if not provided
	if name == nil || *name == "" {
		*name = ExtractNameFromURI(*chartURI)
		Log.Warnf("No name provided. Using extracted name: %s", *name)
	}

	// Create Meshery Snapshot
	designID, err := CreateMesheryDesign(*chartURI, *name, *email)
	if err != nil {
		handleError(errors.ErrCreatingMesheryDesign(err))
	}

	assetLocation := fmt.Sprintf("https://raw.githubusercontent.com/layer5labs/meshery-extensions-packages/master/action-assets/%s.png", designID)

	// Generate Snapshot
	err = GenerateSnapshot(designID, *chartURI, *email, assetLocation)
	if err != nil {
		handleError(errors.ErrGeneratingSnapshot(err))
	}

	if *email == "" {
		loader(2*time.Minute + 40*time.Second) // Loader running for 2 minutes and 40 seconds
		Log.Infof("\nSnapshot generated successfully. Snapshot URL: %s\n", assetLocation)
	} else {
		Log.Info("An email will be sent to the provided email containing the snapshot in a few minutes.")
	}

	err = os.Remove(logFilePath)
	if err != nil {
		handleError(errors.ErrRemovingFile(err, logFilePath))
	}

}
