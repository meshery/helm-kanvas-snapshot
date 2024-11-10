package errors

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	ErrInvalidChartURICode          = "kanvas-snapshot-900"
	ErrCreatingMesheryDesignCode    = "kanvas-snapshot-901"
	ErrGeneratingSnapshotCode       = "kanvas-snapshot-902"
	ErrHTTPPostRequestCode          = "kanvas-snapshot-903"
	ErrDecodingAPICode              = "kanvas-snapshot-905"
	ErrUnexpectedResponseCodeCode   = "kanvas-snapshot-906"
	ErrRequiredFieldNotProvidedCode = "kanvas-snapshot-907"
	ErrInvalidEmailFormatCode       = "kanvas-snapshot-908"
)

func ErrInvalidChartURI(err error) error {
	return errors.New(ErrInvalidChartURICode, errors.Alert,
		[]string{"Invalid or missing Helm chart URI."},
		[]string{err.Error()},
		[]string{"The provided URI for the Helm chart is either missing or invalid."},
		[]string{"Ensure the Helm chart URI is correctly provided."},
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

func ErrUnexpectedResponseCode(statusCode int, body string) error {
	return errors.New(ErrUnexpectedResponseCodeCode, errors.Alert,
		[]string{"Received unexpected response code from Meshery API."},
		[]string{fmt.Sprintf("Status Code: %d, Body: %s", statusCode, body)},
		[]string{"The API returned an unexpected status code."},
		[]string{"Check the request details and ensure the Meshery API is functioning correctly."},
	)
}

func ErrRequiredFieldNotProvided(err error, field string) error {
	return errors.New(ErrRequiredFieldNotProvidedCode, errors.Alert,
		[]string{"All required flags are not passed."},
		[]string{err.Error()},
		[]string{fmt.Sprintf("Required flag \"%s\" is not passed.", field)},
		[]string{fmt.Sprintf("Ensure value for flag \"%s\" is correctly provided.", field)},
	)
}

func ErrInvalidEmailFormat(email string) error {
	return errors.New(ErrInvalidEmailFormatCode, errors.Alert,
		[]string{"Invalid email format provided."},
		[]string{fmt.Sprintf("The provided email '%s' is not a valid email format.", email)},
		[]string{"The email provided for the Kanvas snapshot request is not in the correct format."},
		[]string{"Ensure the email address follows the correct format (e.g., user@example.com)."},
	)
}
