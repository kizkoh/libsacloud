package sacloud

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testServerJSON = `
{
    "ID": 999999999999,
    "Name": "lisacloud-test-server-name",
    "HostName": "libsacloud-test-server-name.com",
    "Description": "Description",
    "Availability": "available",
    "ServiceClass": "cloud\/plan\/1core-1gb",
    "CreatedAt": "2016-04-25T21:36:47+09:00",
    "Icon": null,
    "ServerPlan": ` + testServerPlanJSON + `,
    "Zone": ` + testZoneJSON + `,
    "Instance": ` + testInstanceJSON + `,
    "Disks": [ ` + testDiskJSON + `
    ],
    "Interfaces": [ ` + testPublicInterfaceJSON + `
    	,` + testPrivateInterfaceJSON + `
    ],
    "Appliance": null,
    "Tags": [
	"@virtio-net-pci"
    ]
}

`

var testVNCProxyJSON = `
{
    "Success": true,
    "Status": "OK",
    "Host": "sac-is1b-ssl.sakura.ad.jp",
    "Port": "9999",
    "Password": "xxxxxxxx",
    "VNCFile": "[connection]\nhost=sac-is1b-ssl.sakura.ad.jp\nport=9999\npassword=xxxxxxxxxxxxx",
    "is_ok": true
}
`

var testVNCSizeJSON = `
{
    "Width": "800",
    "Height": "599",
    "is_ok": true
}
`

var testVNCSnapshotJSON = `
{
    "Image": "iVBORw0KGgoAAAANSUhEUgAAAyAAAAJXCAIAAADkQqfyAAANp0lEQVR4nO3dUZKaWABAUUn1lnr\/CzCLmg9nKAYQgb7GbnPOl7Hx8USrvMWDyjAMwwUAgM6vV08AAODdCCwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCICSwAgJjAAgCIDb9\/\/75cLp+fn7d\/X6\/X2z9vD6Z\/+qLr9ZoMNc7w0J9Wt\/z6OPu3H\/e43HL6p+kGs+erD2K203bY7zOBe9+3o9\/Do18GALhsnMF60s\/588bZv4vpT+byVUenunP78XjuPLDjJG9OTGznfF7lqRO4d6ye9OECwMxPWiLM62oZLq\/6AT6aX\/wB6gqA0z4ebjFbIlme\/pk9ni39LPNl47XbM9k+5zTd9TlHx6n2O3vjY\/xdr9dxPevQuuds2I3Vydlrx\/ey\/\/PdOf72VPfMf+dXZfv79nDA7f4GgJ0en8Ha+TMzbjZb0hr\/OXtcOX2JzOwdPeO6qy8aD+Oe3\/iNJcWH64x73su95Noz\/p7dVUui975vz\/4eAsDU4zNYl8XZlG9odn7i9DyPjlPt955pYz0MoHtnvLZL5WgpLrfcczpqOsmN3TldBMB72BVY39myJHa2wurC3P5xTu\/3hJ2BO242ncz2Ff1fmdL2+OeOxrcteAA45Cdd5P5mHp5YOtFDJ9bXXGkEALk\/HVjnuiHc47QnDp15ameeH4Tp9N6gsVbfzokzhfc+tT\/\/PQTgrzIMw7Dn7rPV5\/fcdDa+fOPus4eX5qzO5974q8+vXlK9usdD42xs\/3DylzuH+uHx37mL7eO8WmPLq9eX2y9ncvpzXP3+bO\/l6F0I27OdPTPb9fb3DQA2DMMwvHoO8F08+3I6AP4SrsGCf6krACo\/\/i5C+Iqj67AAsIclQgCAmCVCAICYwAIAiAksAICYwAIAiAksAICYwAIAiAksAICYwAIAiAksAICYwAIAiAksAICYwAIAiAksAICYwAIAiAksAICYwAIAiAksAICYwAIAiH0cfcH1er09+Pz8rCfz+gmEg9+GetVRAgBe6PAZrJcXw1Mn8PJ3BwC8gcNnsNjpaKs54wUAb+PjMlkXuxl\/46fP3\/vhH7Ng2gerj5erb3vGX93d9pOz8ZdTOurh8Zm9\/cvaUuPGcQAA3syv8Yf\/5vL\/OFg+ObWnWu4l157x9+xuY\/6Je+NPn19ObHlMxtfem+f4PADw033c+uD2w3\/5fxlsl8rRc0LLLfecjhpfu727NqoAAL7i4\/JfvqwucuXnhB6Of+4sjnM\/AMD38e9dhCfW156xHvc3GxsXAPjpfk1\/19+gsVbfjhv0AIA\/aRiG4dBdcqs1trx6fbn9ZZE4G3fhrbq3\/b35T1914i7F1bsd9zy\/Pc7X72oEAL65YRiGV8\/hx1NLAMCU\/4vwJJdMAQD3OIN13tdvfgQA3pLAAgCIWSIEAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIgJLACAmMACAIj9A4hgQGHrk9lhAAAAAElFTkSuQmCC",
    "is_ok": true
}
`

func TestMarshalServerJSON(t *testing.T) {
	var server Server
	err := json.Unmarshal([]byte(testServerJSON), &server)

	assert.NoError(t, err)
	assert.NotEmpty(t, server)

	assert.NotEmpty(t, server.ID)
	assert.NotEmpty(t, server.ServerPlan)
	assert.NotEmpty(t, server.Zone)
	assert.NotEmpty(t, server.Disks)
	assert.NotEmpty(t, server.Interfaces)
	assert.NotEmpty(t, server.Instance)
	assert.NotEmpty(t, server.propTags)

	assert.True(t, server.MaintenanceScheduled())
}

func TestMarshalVNCProxyJSON(t *testing.T) {
	var vnc VNCProxyResponse
	err := json.Unmarshal([]byte(testVNCProxyJSON), &vnc)

	assert.NoError(t, err)
	assert.NotEmpty(t, vnc)
	assert.NotEmpty(t, vnc.VNCFile)
}

func TestMarshalVNCSizeJSON(t *testing.T) {
	var vnc VNCSizeResponse
	err := json.Unmarshal([]byte(testVNCSizeJSON), &vnc)

	assert.NoError(t, err)
	assert.NotEmpty(t, vnc)
	assert.NotEmpty(t, vnc.Width)
	assert.NotEmpty(t, vnc.Height)

}

func TestMarshalVNCSnapshotJSON(t *testing.T) {
	var vnc VNCSnapshotResponse
	err := json.Unmarshal([]byte(testVNCSnapshotJSON), &vnc)

	assert.NoError(t, err)
	assert.NotEmpty(t, vnc)
	assert.NotEmpty(t, vnc.Image)

}
