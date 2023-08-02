// Copyright (c) 2023 Cloudnatively Services Pvt Ltd
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"io"
	"net/http"
	"net/url"
	"pb/pkg/config"
	"time"
)

type HttpClient struct {
	client  http.Client
	profile *config.Profile
}

func DefaultClient() HttpClient {
	return HttpClient{
		client: http.Client{
			Timeout: 60 * time.Second,
		},
		profile: &DefaultProfile,
	}
}

func (client *HttpClient) baseApiUrl(path string) (x string) {
	x, _ = url.JoinPath(client.profile.Url, "api/v1/", path)
	return
}

func (client *HttpClient) NewRequest(method string, path string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, client.baseApiUrl(path), body)
	if err != nil {
		return
	}
	req.SetBasicAuth(client.profile.Username, client.profile.Password)
	req.Header.Add("Content-Type", "application/json")
	return
}