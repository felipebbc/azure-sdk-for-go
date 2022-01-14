//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
)

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListBySubscription.json
func ExampleLabsClient_ListBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	pager := client.ListBySubscription(&armdevtestlabs.LabsClientListBySubscriptionOptions{Expand: nil,
		Filter:  nil,
		Top:     nil,
		Orderby: nil,
	})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListByResourceGroup.json
func ExampleLabsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		&armdevtestlabs.LabsClientListByResourceGroupOptions{Expand: nil,
			Filter:  nil,
			Top:     nil,
			Orderby: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_Get.json
func ExampleLabsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<name>",
		&armdevtestlabs.LabsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LabsClientGetResult)
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateOrUpdate.json
func ExampleLabsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.Lab{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tagName1": to.StringPtr("tagValue1"),
			},
			Properties: &armdevtestlabs.LabProperties{
				LabStorageType: armdevtestlabs.StorageType("{Standard|Premium}").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LabsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_Delete.json
func ExampleLabsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_Update.json
func ExampleLabsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.LabFragment{},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LabsClientUpdateResult)
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_CreateEnvironment.json
func ExampleLabsClient_BeginCreateEnvironment() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateEnvironment(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.LabVirtualMachineCreationParameter{
			Name:     to.StringPtr("<name>"),
			Location: to.StringPtr("<location>"),
			Properties: &armdevtestlabs.LabVirtualMachineCreationParameterProperties{
				AllowClaim:              to.BoolPtr(true),
				DisallowPublicIPAddress: to.BoolPtr(true),
				GalleryImageReference: &armdevtestlabs.GalleryImageReference{
					Offer:     to.StringPtr("<offer>"),
					OSType:    to.StringPtr("<ostype>"),
					Publisher: to.StringPtr("<publisher>"),
					SKU:       to.StringPtr("<sku>"),
					Version:   to.StringPtr("<version>"),
				},
				LabSubnetName:       to.StringPtr("<lab-subnet-name>"),
				LabVirtualNetworkID: to.StringPtr("<lab-virtual-network-id>"),
				Password:            to.StringPtr("<password>"),
				Size:                to.StringPtr("<size>"),
				StorageType:         to.StringPtr("<storage-type>"),
				UserName:            to.StringPtr("<user-name>"),
			},
			Tags: map[string]*string{
				"tagName1": to.StringPtr("tagValue1"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ExportResourceUsage.json
func ExampleLabsClient_BeginExportResourceUsage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginExportResourceUsage(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.ExportResourceUsageParameters{
			BlobStorageAbsoluteSasURI: to.StringPtr("<blob-storage-absolute-sas-uri>"),
			UsageStartDate:            to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-12-01T00:00:00Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ImportVirtualMachine.json
func ExampleLabsClient_BeginImportVirtualMachine() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginImportVirtualMachine(ctx,
		"<resource-group-name>",
		"<name>",
		armdevtestlabs.ImportLabVirtualMachineRequest{
			DestinationVirtualMachineName:  to.StringPtr("<destination-virtual-machine-name>"),
			SourceVirtualMachineResourceID: to.StringPtr("<source-virtual-machine-resource-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ListVhds.json
func ExampleLabsClient_ListVhds() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdevtestlabs.NewLabsClient("<subscription-id>", cred, nil)
	pager := client.ListVhds("<resource-group-name>",
		"<name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}