// +build !debug

/*
Copyright 2020 The kconnect Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package aws

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"

	"github.com/fidelity/kconnect/pkg/config"
)

func (p *ServiceProvider) resolveProfile(name string, cfg config.ConfigurationSet) error {
	if cfg.ExistsWithValue(name) {
		return nil
	}

	profile := ""
	prompt := &survey.Input{
		Message: "Enter the name of AWS profile",
	}
	if err := survey.AskOne(prompt, &profile, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("asking for profile name: %w", err)
	}

	if err := cfg.SetValue(name, profile); err != nil {
		p.logger.Errorf("failed setting profile flag to %s: %s", profile, err.Error())
		return fmt.Errorf("setting profile flag: %w", err)
	}

	return nil
}

func (p *ServiceProvider) resolveRegion(name string, cfg config.ConfigurationSet) error {
	if cfg.ExistsWithValue(name) {
		return nil
	}

	//TODO: this needs to be changed to present a list
	region := ""
	prompt := &survey.Input{
		Message: "Enter the name of AWS region",
	}
	if err := survey.AskOne(prompt, &region, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("asking for region name: %w", err)
	}

	if err := cfg.SetValue(name, region); err != nil {
		p.logger.Errorf("failed setting region flag to %s: %s", region, err.Error())
		return fmt.Errorf("setting region flag: %w", err)
	}

	return nil
}

func (p *ServiceProvider) resolveUsername(name string, cfg config.ConfigurationSet) error {
	if cfg.ExistsWithValue(name) {
		return nil
	}

	username := ""
	prompt := &survey.Input{
		Message: "Enter your username",
	}
	if err := survey.AskOne(prompt, &username, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("asking for username name: %w", err)
	}

	if err := cfg.SetValue(name, username); err != nil {
		p.logger.Errorf("failed setting username flag to %s: %s", username, err.Error())
		return fmt.Errorf("setting username flag: %w", err)
	}

	return nil
}

func (p *ServiceProvider) resolvePassword(name string, cfg config.ConfigurationSet) error {
	if cfg.ExistsWithValue(name) {
		return nil
	}

	//TODO: change to blank out the characters
	password := ""
	prompt := &survey.Password{
		Message: "Enter your password",
	}
	if err := survey.AskOne(prompt, &password, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("asking for password name: %w", err)
	}

	if err := cfg.SetValue(name, password); err != nil {
		p.logger.Errorf("failed setting password flag to %s: %s", password, err.Error())
		return fmt.Errorf("setting password flag: %w", err)
	}

	return nil
}

func (p *ServiceProvider) resolveIdpEndpoint(name string, cfg config.ConfigurationSet) error {
	if cfg.ExistsWithValue(name) {
		return nil
	}

	endpoint := ""
	prompt := &survey.Input{
		Message: "Enter the endpoint for the IdP",
	}
	if err := survey.AskOne(prompt, &endpoint, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("asking for idp-endpoint name: %w", err)
	}

	if err := cfg.SetValue(name, endpoint); err != nil {
		p.logger.Errorf("failed setting idp-endpoint flag to %s: %s", endpoint, err.Error())
		return fmt.Errorf("setting idp-endpoint flag: %w", err)
	}

	return nil
}

func (p *ServiceProvider) resolveIdpProvider(name string, cfg config.ConfigurationSet) error {
	if cfg.ExistsWithValue(name) {
		return nil
	}
	//TODO: get this from saml2aws????
	options := []string{"Akamai", "AzureAD", "ADFS", "ADFS2", "GoogleApps", "Ping", "JumpCloud", "Okta", "OneLogin", "PSU", "KeyCloak", "F5APM", "Shibboleth", "ShibbolethECP", "NetIQ"}

	idpProvider := ""
	prompt := &survey.Select{
		Message: "Select your identity provider",
		Options: options,
	}
	if err := survey.AskOne(prompt, &idpProvider, survey.WithValidator(survey.Required)); err != nil {
		return fmt.Errorf("asking for idp-provider: %w", err)
	}

	if err := cfg.SetValue(name, idpProvider); err != nil {
		p.logger.Errorf("failed setting idp-provider flag to %s: %s", idpProvider, err.Error())
		return fmt.Errorf("setting idp-provider flag: %w", err)
	}

	return nil
}
