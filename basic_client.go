package ofxgo

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

// A function that is called before performing a request against an OFX server.
// Returning a non-nil error will prevent the request from being sent.
type PreRequestFunc func(*http.Request) error

// BasicClient provides a standard Client implementation suitable for most
// financial institutions. BasicClient uses default, non-zero settings, even if
// its fields are not initialized.
type BasicClient struct {
	// Request fields to overwrite with the client's values. If nonempty,
	// defaults are used
	SpecVersion ofxVersion // VERSION in header
	AppID       string     // SONRQ>APPID
	AppVer      string     // SONRQ>APPVER

	// Don't insert newlines or indentation when marshalling to SGML/XML
	NoIndent bool

	// Use carriage returns on new lines
	CarriageReturn bool

	// allows overriding the HTTP client used to perform the request
	Client *http.Client

	// allows modification of the raw request before sending
	PreRequestFuncs []PreRequestFunc

	// allow insecure (non-HTTPS) requests
	Insecure bool
}

// Version returns the OFX specification version this BasicClient will marshal
// Requests as. Defaults to "203" if the client's SpecVersion field is empty.
func (c *BasicClient) OfxVersion() ofxVersion {
	if c.SpecVersion.Valid() {
		return c.SpecVersion
	}

	return OfxVersion203
}

func (c *BasicClient) httpClient() *http.Client {
	if c.Client == nil {
		return http.DefaultClient
	} else {
		return c.Client
	}
}

func (c *BasicClient) prepRequest(req *http.Request) error {
	for _, fn := range c.PreRequestFuncs {
		if fn != nil {
			if err := fn(req); err != nil {
				return err
			}
		}
	}

	return nil
}

// ID returns this BasicClient's OFX AppID field, defaulting to "OFXGO" if
// unspecified.
func (c *BasicClient) ID() String {
	if len(c.AppID) > 0 {
		return String(c.AppID)
	}
	return String("OFXGO")
}

// Version returns this BasicClient's version number as a string, defaulting to
// "0001" if unspecified.
func (c *BasicClient) Version() String {
	if len(c.AppVer) > 0 {
		return String(c.AppVer)
	}
	return String("0001")
}

// IndentRequests returns true if the marshaled XML should be indented (and
// contain newlines, since the two are linked in the current implementation)
func (c *BasicClient) IndentRequests() bool {
	return !c.NoIndent
}

// CarriageReturnNewLines returns true if carriage returns should be used on new lines, false otherwise
func (c *BasicClient) CarriageReturnNewLines() bool {
	return c.CarriageReturn
}

func (c *BasicClient) RawRequest(URL string, r io.Reader) (*http.Response, error) {
	if !c.Insecure {
		if !strings.HasPrefix(URL, "https://") {
			return nil, errors.New("Refusing to send OFX request with possible plain-text password over non-https protocol")
		}
	}

	if req, err := http.NewRequest(`POST`, URL, r); err == nil {
		req.Header.Set(`Content-Type`, `application/x-ofx`)

		if err := c.prepRequest(req); err != nil {
			return nil, err
		}

		response, err := c.httpClient().Do(req)

		if err != nil {
			return nil, err
		}

		if response.StatusCode != 200 {
			return nil, errors.New("OFXQuery request status: " + response.Status)
		}

		return response, nil
	} else {
		return nil, err
	}
}

func (c *BasicClient) RequestNoParse(r *Request) (*http.Response, error) {
	return clientRequestNoParse(c, r)
}

func (c *BasicClient) Request(r *Request) (*Response, error) {
	return clientRequest(c, r)
}
