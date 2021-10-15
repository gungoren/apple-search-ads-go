/**
Copyright (C) 2021 Mehmet Gungoren.
This file is part of apple-search-ads-go, a package for working with Apple's
Search Ads API.
apple-search-ads-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
apple-search-ads-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with apple-search-ads-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asa

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserAcl(t *testing.T) {
	t.Parallel()

	testEndpointWithResponse(t, "{}", &UserACLListResponse{}, func(ctx context.Context, client *Client) (interface{}, *Response, error) {
		return client.AccessControlList.GetUserACL(ctx)
	})
}

func TestDeserialization(t *testing.T) {
	t.Parallel()

	sampleJSONResponse := "../test/response_body_json_files/user_acl_list_response.json"

	content, err := ioutil.ReadFile(sampleJSONResponse)
	assert.NoError(t, err)

	model := &UserACLListResponse{}

	err = json.Unmarshal(content, model)
	assert.NoError(t, err)

	assert.Equal(t, len(model.UserAcls), 1)
	assert.Nil(t, model.Pagination)

	userACL := model.UserAcls[0]

	assert.Equal(t, "USD", userACL.Currency)
	assert.Equal(t, 1, int(userACL.OrgID))
	assert.Equal(t, "orgName", userACL.OrgName)
	assert.Equal(t, PaymentModelLoc, userACL.PaymentModel)

	assert.Equal(t, []UserACLRoleName{UserACLRoleNameAPIAccountManager}, userACL.RoleNames)
	assert.Equal(t, "Asia/Hong_Kong", string(userACL.TimeZone))
}
