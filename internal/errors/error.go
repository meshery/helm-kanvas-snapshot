package errors

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	ErrInvalidChartURICode       = "kanvas-snapshot-900"
	ErrCreatingLogFileCode       = "kanvas-snapshot-901"
	ErrCreatingMesheryDesignCode = "kanvas-snapshot-902"
	ErrGeneratingSnapshotCode    = "kanvas-snapshot-903"
	ErrHTTPPostRequestCode       = "kanvas-snapshot-904"
	ErrDecodingAPICode           = "kanvas-snapshot-905"
	ErrRemovingFileCode          = "kanvas-snapshot-906"
	ErrGettingCWDCode            = "kanvas-snapshot-907"
	ErrCreatingLogDirCode        = "kanvas-snapshot-908"
)

func ErrInvalidChartURI(err error) error {
	return errors.New(ErrInvalidChartURICode, errors.Alert,
		[]string{"Invalid or missing Helm chart URI."},
		[]string{err.Error()},
		[]string{"The provided URI for the Helm chart is either missing or invalid."},
		[]string{"Ensure the Helm chart URI is correctly provided."},
	)
}

func ErrCreatingLogFile(err error, path string) error {
	return errors.New(ErrCreatingLogFileCode, errors.Alert,
		[]string{fmt.Sprintf("Failed to create log file at path: %s", path)},
		[]string{err.Error()},
		[]string{"An error occurred while trying to create the log file."},
		[]string{"Check file permissions or the file path."},
	)
}

func ErrCreatingLogDir(err error) error {
	return errors.New(ErrCreatingLogDirCode, errors.Alert,
		[]string{"Failed to create log directory."},
		[]string{err.Error()},
		[]string{"The log directory could not be created, possibly due to insufficient file system permissions or directory conflicts."},
		[]string{"Check the application's file system permissions, ensure that the target directory path is correct, and verify that there are no conflicts (e.g., existing files with the same name)."},
	)
}

func ErrGettingCWD(err error) error {
	return errors.New(ErrGettingCWDCode, errors.Alert,
		[]string{"Failed to get current working directory."},
		[]string{err.Error()},
		[]string{"The function was unable to determine the current working directory, possibly due to insufficient permissions or a system error."},
		[]string{"Verify that the application has permission to access the file system and that the environment is properly configured."},
	)
}

func ErrCreatingMesheryDesign(err error) error {
	return errors.New(ErrCreatingMesheryDesignCode, errors.Alert,
		[]string{"Failed to create Meshery design."},
		[]string{err.Error()},
		[]string{"Meshery Design creation failed due to an error."},
		[]string{"Check Meshery API connection and ensure the payload is correct."},
	)
}

func ErrGeneratingSnapshot(err error) error {
	return errors.New(ErrGeneratingSnapshotCode, errors.Alert,
		[]string{"Failed to generate snapshot."},
		[]string{err.Error()},
		[]string{"Snapshot generation failed due to an error."},
		[]string{"Check Meshery Cloud API connection and payload."},
	)
}

func ErrHTTPPostRequest(err error) error {
	return errors.New(ErrHTTPPostRequestCode, errors.Alert,
		[]string{"Failed to perform HTTP POST request."},
		[]string{err.Error()},
		[]string{"HTTP POST request failed during interaction with Meshery API."},
		[]string{"Check Meshery API endpoint and ensure valid request payload."},
	)
}

func ErrDecodingAPI(err error) error {
	return errors.New(ErrDecodingAPICode, errors.Alert,
		[]string{"Failed to decode API response."},
		[]string{err.Error()},
		[]string{"API response could not be decoded into the expected format."},
		[]string{"Ensure the Meshery API response format is correct."},
	)
}

func ErrRemovingFile(err error, path string) error {
	return errors.New(ErrRemovingFileCode, errors.Alert,
		[]string{fmt.Sprintf("Failed to remove file at path: %s", path)},
		[]string{err.Error()},
		[]string{"File deletion failed due to an error."},
		[]string{"Check file permissions and file existence."},
	)
}
