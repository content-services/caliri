/*
Candlepin

Testing EntitlementsAPIService

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

func Test_caliri_EntitlementsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EntitlementsAPIService GetEntitlement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementsAPI.GetEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService GetUpstreamCert", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dbid string

		resp, httpRes, err := apiClient.EntitlementsAPI.GetUpstreamCert(context.Background(), dbid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService HasEntitlement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var consumerUuid string
		var productId string

		resp, httpRes, err := apiClient.EntitlementsAPI.HasEntitlement(context.Background(), consumerUuid, productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService ListAllForConsumer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EntitlementsAPI.ListAllForConsumer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService MigrateEntitlement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var entitlementId string

		resp, httpRes, err := apiClient.EntitlementsAPI.MigrateEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService RegenerateEntitlementCertificatesForProduct", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var productId string

		resp, httpRes, err := apiClient.EntitlementsAPI.RegenerateEntitlementCertificatesForProduct(context.Background(), productId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService Unbind", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var dbid string

		httpRes, err := apiClient.EntitlementsAPI.Unbind(context.Background(), dbid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EntitlementsAPIService UpdateEntitlement", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var entitlementId string

		httpRes, err := apiClient.EntitlementsAPI.UpdateEntitlement(context.Background(), entitlementId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
