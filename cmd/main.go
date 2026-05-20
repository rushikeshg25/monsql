package cmd

import (
	"errors"
	"net/url"
)

type monsql struct {
	mongoURL string
}

func main() {
	m := monsql{
		mongoURL: "",
	}
	_ = m // prevent unused variable error if main is run
}

func (m *monsql) create(rawURL string) error {
	sanitized, err := validateAndSanitizeURL(rawURL)
	if err != nil {
		return err
	}
	m.mongoURL = sanitized
	return nil
}

func validateAndSanitizeURL(rawURL string) (string, error) {
	sanitized, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return sanitized.String(), nil
}

func (m *monsql) run() error {
	//parse the SQL Query

	err := errors.New("Dummy")
	return err
}
