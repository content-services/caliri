/*
Candlepin

Testing OwnerAPIService

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

func Test_caliri_OwnerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OwnerAPIService Claim", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var anonymousOwnerKey string

		resp, httpRes, err := apiClient.OwnerAPI.Claim(context.Background(), anonymousOwnerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService CountConsumers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.CountConsumers(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService CreateActivationKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.CreateActivationKey(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService CreateEnvironment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.CreateEnvironment(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService CreateOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OwnerAPI.CreateOwner(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService CreatePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.CreatePool(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService CreateUeberCertificate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.CreateUeberCertificate(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService DeleteLogLevel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		httpRes, err := apiClient.OwnerAPI.DeleteLogLevel(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService DeleteOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		httpRes, err := apiClient.OwnerAPI.DeleteOwner(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetConsumersSyspurpose", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetConsumersSyspurpose(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetHypervisors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetHypervisors(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetImports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetImports(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetOwner(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetOwnerContentAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetOwnerContentAccess(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetOwnerInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetOwnerInfo(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetOwnerSubscriptions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetOwnerSubscriptions(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetSyspurpose", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetSyspurpose(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetUeberCertificate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetUeberCertificate(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService GetUpstreamConsumers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.GetUpstreamConsumers(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService HealEntire", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.HealEntire(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService ImportManifestAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.ImportManifestAsync(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService ListConsumers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.ListConsumers(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService ListEnvironments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.ListEnvironments(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService ListOwnerPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.ListOwnerPools(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService ListOwners", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.OwnerAPI.ListOwners(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService OwnerActivationKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.OwnerActivationKeys(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService OwnerEntitlements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.OwnerEntitlements(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService OwnerServiceLevels", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.OwnerServiceLevels(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService RefreshPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.RefreshPools(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService SetLogLevel", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.SetLogLevel(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService UndoImports", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.UndoImports(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService UpdateOwner", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		resp, httpRes, err := apiClient.OwnerAPI.UpdateOwner(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OwnerAPIService UpdatePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ownerKey string

		httpRes, err := apiClient.OwnerAPI.UpdatePool(context.Background(), ownerKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
