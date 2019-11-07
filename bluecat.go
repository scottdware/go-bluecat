// Package bluecat is Go client to interact with the Bluecat API.
package bluecat

import (
	"crypto/tls"
	"fmt"
	"regexp"

	"gopkg.in/resty.v1"
)

// Bluecat contains all of the server information that is used when calling the package functions.
type Bluecat struct {
	Server    string
	URI       string
	AuthToken string
}

// APIAccessRight class controls access right objects.
type APIAccessRight struct {
	EntityID   int64  `json:"entityId"`
	UserID     int64  `json:"userId"`
	Value      string `json:"value"`
	Overrides  string `json:"overrides"`
	Properties string `json:"properties"`
}

// APIData contains API Data with a timestamp
type APIData struct {
	Name       string `json:"name"`
	Properties string `json:"properties"`
}

// APIDeploymentOption configure both DHCP and DNS services on the network. They are available as
// DHCP client and service options, as well as standard DNS options. Deployment options support the
// standard object functions.
type APIDeploymentOption struct {
	ID         int64  `json:"id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Value      string `json:"value"`
	Properties string `json:"properties"`
}

// APIDeploymentRole manages the deployment roles that control the services provided by
// Address Manager-managed servers. These objects support the standard object functions.
type APIDeploymentRole struct {
	ID                int64  `json:"id"`
	Type              string `json:"type"`
	Service           string `json:"service"`
	EntityID          int64  `json:"entityId"`
	ServerInterfaceID int64  `json:"serverInterfaceId"`
	Properties        string `json:"properties"`
}

// APIEntity represents all entities except options, roles, and access rights. It manages
// all other types by passing the values for the object as a delimited properties string of name–value pairs.
type APIEntity struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Properties string `json:"properties"`
}

// APIUserDefinedField enables you to add user-defined fields to each of the Address Manager object types.
// This class allows API users to query and gather user-defined fields information for a specified object type.
type APIUserDefinedField struct {
	Name                string `json:"name"`
	DisplayName         string `json:"displayName"`
	Type                string `json:"type"`
	DefaultValue        string `json:"defaultValue"`
	ValidatorProperties string `json:"validatorProperties"`
	Properties          string `json:"properties"`
	PredefinedValues    string `json:"predefinedValues"`
	Required            bool   `json:"required"`
	HideFromSearch      bool   `json:"hideFromSearch"`
}

// ResponsePolicySearchResult represents the response policy items that are configured either in local response
// policies or BlueCat Security feed data.
type ResponsePolicySearchResult struct {
	Name       string `json:"name"`
	PolicyType string `json:"policyType"`
	ParentIDs  string `json:"parentIds"`
	Category   string `json:"category"`
	ConfigID   int64  `json:"configId"`
}

// getAuthToken returns the Bluecat session authentication token which is used to authenticate
// all of the API calls to the BLuecat server.
func getAuthToken(server, user, pass string) (string, error) {
	sessionToken := regexp.MustCompile(`^.*(BAMAuthToken:\s+[\w=]+)\s+.*$`)

	loginReq := fmt.Sprintf("https://%s/Services/REST/v1/login?username=%s&password=%s", server, user, pass)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(loginReq)

	if err != nil {
		return "", fmt.Errorf("%s - getAuthToken login", err)
	}

	token := sessionToken.FindStringSubmatch(resp.String())
	if len(token) <= 0 {
		return "", fmt.Errorf("%s - getAuthToken token parse", string(resp.Body()))
	}

	return token[1], nil
}

func init() {
	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
}

// NewSession initializes a session against the specificed Bluecat server.
func NewSession(server, user, pass string) (*Bluecat, error) {
	token, err := getAuthToken(server, user, pass)
	if err != nil {
		return nil, fmt.Errorf("%s - NewSession initialization", err)
	}

	bc := &Bluecat{
		Server:    server,
		URI:       "/Services/REST/v1",
		AuthToken: token,
	}

	return bc, nil
}
