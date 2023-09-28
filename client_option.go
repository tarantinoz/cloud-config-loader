package cloudconfigloader

import "github.com/pkg/errors"

type LoaderOption func(c *defaultClient) error

func WithBranch(branch string) LoaderOption {
	return func(c *defaultClient) error {
		if branch == "" {
			return errors.New("branch must not be empty")
		}

		c.branch = branch
		return nil
	}
}

func WithFormat(format Format) LoaderOption {
	return func(c *defaultClient) error {
		if !format.Valid() {
			return errors.Errorf("[%s] is not an acceptable format", format)
		}

		c.format = format
		return nil
	}
}

func WithScheme(scheme string) LoaderOption {
	return func(c *defaultClient) error {
		if scheme == "" {
			return errors.New("scheme must not be empty")
		}

		c.scheme = scheme
		return nil
	}
}

func WithBasicAuth(username, password string) LoaderOption {
	return func(c *defaultClient) error {
		if username == "" {
			return errors.New("username must not be empty")
		}

		c.basicAuth = &basicAuthInfo{
			username: username,
			password: password,
		}
		return nil
	}
}
