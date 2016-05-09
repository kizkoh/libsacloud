package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testLoadBalancerJSON = `
{
        "ID": "123456789012",
        "Class": "loadbalancer",
        "Name": "\u308d\u304a\u3069\u3070\u3089\u3093\u3055",
        "Description": "\u30ed\u30aa\u30c9\u30d0\u30e9\u30f3\u3055\u306e\u8aac\u660e",
        "Plan": {
            "ID": 1
        },
        "Settings": {
            "LoadBalancer": [
                {
                    "VirtualIPAddress": "192.168.200.50",
                    "Port": "80",
                    "DelayLoop": "1000",
                    "SorryServer": "",
                    "Servers": [
                       {
                            "IPAddress": "192.168.200.51",
                            "Port": "80",
                            "HealthCheck": {
                                "Protocol": "http",
                                "Path": "\/index.html",
                                "Status": "200"
                            },
                            "Enabled": "True",
                            "Status": "DOWN",
                            "ActiveConn": "0"
                        },
                        {
                            "IPAddress": "192.168.200.52",
                            "Port": "80",
                            "HealthCheck": {
                                "Protocol": "ping"
                            },
                            "Enabled": "True"
                        }
                    ]
                }
            ]
        },
        "SettingsHash": "924521e812f96157a83d138c79d423fb",
        "Remark": {
            "Zone": {
                "ID": 31002
            },
            "Switch": {
                "ID": "123456789012"
            },
            "VRRP": {
                "VRID": 1
            },
            "Network": {
                "NetworkMaskLen": 24,
                "DefaultRoute": "192.168.200.1"
            },
            "Servers": [
                {
                    "IPAddress": "192.168.200.11"
                }
            ],
            "Plan": {
                "ID": 1
            }
        },
        "Availability": "available",
        "Instance": {
            "Status": "up",
            "StatusChangedAt": "2016-04-29T18:29:17+09:00"
        },
        "ServiceClass": "cloud\/appliance\/loadbalancer\/1",
        "CreatedAt": "2016-04-29T18:27:18+09:00",
        "Icon": {
            "ID": "112300511981",
            "URL": "https:\/\/secure.sakura.ad.jp\/cloud\/zone\/is1b\/api\/cloud\/1.1\/icon\/112300511981.png",
            "Name": "CentOS",
            "Scope": "shared"
        },
        "Switch": {
            "ID": "112800442260",
            "Name": "\u3059\u3046\u3043\u3063\u3061",
            "Internet": null,
            "Scope": "user",
            "Availability": "available",
            "Zone": {
                "ID": 31002,
                "Name": "is1b",
                "Region": {
                    "ID": 310,
                    "Name": "\u77f3\u72e9"
                }
            }
        },
        "Interfaces": [
            {
                "IPAddress": null,
                "UserIPAddress": "192.168.200.11",
                "HostName": null,
                "Switch": {
                    "ID": "112800442260",
                    "Name": "\u3059\u3046\u3043\u3063\u3061",
                    "Scope": "user",
                    "Subnet": null,
                    "UserSubnet": {
                        "DefaultRoute": "192.168.200.1",
                        "NetworkMaskLen": 24
                    }
                }
            }
        ],
        "Tags": [
            "\u3042\u3042",
            "\u3044\u3044",
            "\u3046\u3046"
        ]
    }
    	`
)

func TestMarshalLoadBalancerJSON(t *testing.T) {
	//standard plan
	var lb LoadBalancer
	err := json.Unmarshal([]byte(testLoadBalancerJSON), &lb)

	assert.NoError(t, err)
	assert.NotEmpty(t, lb)

	assert.NotEmpty(t, lb.ID)
	assert.NotEmpty(t, lb.Remark)

	assert.NotEmpty(t, lb.Remark.Servers)
	assert.NotEmpty(t, lb.Remark.Network)
	assert.NotEmpty(t, lb.Remark.Switch)
	assert.NotEmpty(t, lb.Remark.VRRP)
	assert.NotEmpty(t, lb.Remark.Zone)
	//assert.NotEmpty(t, lb.Remark.Plan)

	assert.NotEmpty(t, lb.Instance)
	assert.NotEmpty(t, lb.Interfaces)

	assert.NotEmpty(t, lb.Settings.LoadBalancer)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].VirtualIPAddress)

	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0])

	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].IPAddress)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].HealthCheck.Protocol)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].HealthCheck.Path)
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[0].HealthCheck.Status)

	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[1])
	assert.NotEmpty(t, lb.Settings.LoadBalancer[0].Servers[1].HealthCheck.Protocol)
	assert.Empty(t, lb.Settings.LoadBalancer[0].Servers[1].HealthCheck.Path)
	assert.Empty(t, lb.Settings.LoadBalancer[0].Servers[1].HealthCheck.Status)

}