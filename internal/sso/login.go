// Package sso handles AWS SSO authentication via the aws cli.
package sso

import "github.com/krause-timo/awsctx/internal/awscli"

// Login ensures the aws cli is available and performs an SSO login for the given profile.
func Login(profile string) error {
	if err := awscli.Require(); err != nil {
		return err
	}

	return awscli.SSOLogin(profile)
}
