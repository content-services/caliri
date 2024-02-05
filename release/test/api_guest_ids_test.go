/*
Candlepin

Testing GuestIdsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package caliri

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/content-services/caliri/release/v4"
)

func Test_caliri_GuestIdsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GuestIdsAPIService DeleteGuest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumerUuid string
		var guestId string

		httpRes, err := apiClient.GuestIdsAPI.DeleteGuest(context.Background(), consumerUuid, guestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestIdsAPIService GetGuestId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumerUuid string
		var guestId string

		resp, httpRes, err := apiClient.GuestIdsAPI.GetGuestId(context.Background(), consumerUuid, guestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestIdsAPIService GetGuestIds", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumerUuid string

		resp, httpRes, err := apiClient.GuestIdsAPI.GetGuestIds(context.Background(), consumerUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestIdsAPIService UpdateGuest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumerUuid string
		var guestId string

		httpRes, err := apiClient.GuestIdsAPI.UpdateGuest(context.Background(), consumerUuid, guestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestIdsAPIService UpdateGuests", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumerUuid string

		httpRes, err := apiClient.GuestIdsAPI.UpdateGuests(context.Background(), consumerUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}