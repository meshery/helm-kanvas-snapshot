package main

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	ErrInvalidChartURI        = "meshery_plugin_invalid_chart_uri"
	ErrCreateLogFileCode      = "meshery_plugin_create_log_file_failed"
	ErrCreateMesheryDesign    = "meshery_plugin_create_meshery_design_failed"
	ErrGenerateSnapshotFailed = "meshery_plugin_generate_snapshot_failed"
	ErrHTTPPostRequestFailed  = "meshery_plugin_http_post_failed"
	ErrAPIDecodeFailed        = "meshery_plugin_api_decode_failed"
	ErrFileRemoveFailed       = "meshery_plugin_file_remove_failed"
)

func ErrInvalidChartURIHandler(err error) error {
	return errors.New(ErrInvalidChartURI, errors.Alert,
		[]string{"Invalid or missing Helm chart URI."},
		[]string{err.Error()},
		[]string{"The provided URI for the Helm chart is either missing or invalid."},
		[]string{"Ensure the Helm chart URI is correctly provided."},
	)
}

func ErrCreateLogFileHandler(err error, path string) error {
	return errors.New(ErrCreateLogFileCode, errors.Alert,
		[]string{fmt.Sprintf("Failed to create log file at path: %s", path)},
		[]string{err.Error()},
		[]string{"An error occurred while trying to create the log file."},
		[]string{"Check file permissions or the file path."},
	)
}

func ErrCreateMesheryDesignHandler(err error) error {
	return errors.New(ErrCreateMesheryDesign, errors.Alert,
		[]string{"Failed to create Meshery design."},
		[]string{err.Error()},
		[]string{"Meshery Design creation failed due to an error."},
		[]string{"Check Meshery API connection and ensure the payload is correct."},
	)
}

func ErrGenerateSnapshotHandler(err error) error {
	return errors.New(ErrGenerateSnapshotFailed, errors.Alert,
		[]string{"Failed to generate snapshot."},
		[]string{err.Error()},
		[]string{"Snapshot generation failed due to an error."},
		[]string{"Check Meshery Cloud API connection and payload."},
	)
}

func ErrHTTPPostRequestHandler(err error) error {
	return errors.New(ErrHTTPPostRequestFailed, errors.Alert,
		[]string{"Failed to perform HTTP POST request."},
		[]string{err.Error()},
		[]string{"HTTP POST request failed during interaction with Meshery API."},
		[]string{"Check Meshery API endpoint and ensure valid request payload."},
	)
}

func ErrAPIDecodeHandler(err error) error {
	return errors.New(ErrAPIDecodeFailed, errors.Alert,
		[]string{"Failed to decode API response."},
		[]string{err.Error()},
		[]string{"API response could not be decoded into the expected format."},
		[]string{"Ensure the Meshery API response format is correct."},
	)
}

func ErrFileRemoveHandler(err error, path string) error {
	return errors.New(ErrFileRemoveFailed, errors.Alert,
		[]string{fmt.Sprintf("Failed to remove file at path: %s", path)},
		[]string{err.Error()},
		[]string{"File deletion failed due to an error."},
		[]string{"Check file permissions and file existence."},
	)
}
