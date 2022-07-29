package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPAddress struct {
	DnsNameLabel *string                     `json:"dnsNameLabel,omitempty"`
	Fqdn         *string                     `json:"fqdn,omitempty"`
	IP           *string                     `json:"ip,omitempty"`
	Ports        []Port                      `json:"ports"`
	Type         ContainerGroupIPAddressType `json:"type"`
}