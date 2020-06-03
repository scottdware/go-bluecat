package bluecat

import (
	"fmt"
	"strings"

	"gopkg.in/resty.v1"
)

// addACL

// addAccessRight

// addAdditionalIPAddresses

// addAliasRecord

// addBulkHostRecord

// addCustomOptionDefinition

// addDHCP4RangeBySize

// addDHCP4Range

// addDHCP6ClientDeploymentOption

// addDHCP6RangeBySize

// addDHCP6Range

// addDHCP6ServiceDeploymentOption

// addDHCPClientDeploymentOption

// addDHCPDeploymentRole

// addDHCPServiceDeploymentOption

// addDHCPSubClass

// addDHCPVendorDeploymentOption

// addDNSDeploymentOption

// addDNSDeploymentRole

// addDeviceInstance

// addDeviceSubtype

// addDeviceType

// addDevice

// addEntity

// addEnumNumber

// addEnumZone

// addExternalHostRecord

// AddGenericRecord Adds generic records.
//
// Parameter `absolutename` is the FQDN of the record. If you are adding a record in a zone that is linked to a incremental
// naming policy, you must add a single hash sign (#) at the appropriate location in the FQDN. Depending on the policy order value,
// the location of the single hash sign varies.
//
// Parameter `properties` adds object properties, including comments and user-defined fields. Parameter `rdata` is the data for the
// resource record, in BIND format. For example, A records follow the format 10.0.0.4. Parameter `ttl` is the time-to-live (TTL) value
// for the record. To ignore the TTL, set this value to -1.
//
// Parameter `objecttype` is the type of record. Valid settings for this parameter are the generic resource record types supported in
// Address Manager: A6, AAAA, AFSDB, APL, CAA, CERT, DHCID, DNAME, DNSKEY, DS, ISDN, KEY, KX, LOC, MB, MG, MINFO, MR, NS, NSAP, PX, RP,
// RT, SINK, SSHFP, TLSA, WKS, and X25.
//
// Parameter `viewid` is the object ID for the parent view to which you are adding the record.
//
// Returns the object ID for the new generic resource record.
func (b *Bluecat) AddGenericRecord(absolutename, properties, rdata string, ttl int, objecttype string, viewid int) (string, error) {
	req := fmt.Sprintf("https://%s%s/addGenericRecord?absoluteName=%s&rdata=%s&ttl=%d&type=%s&viewId=%d&properties=%s",
		b.Server, b.URI, absolutename, rdata, ttl, objecttype, viewid, properties)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Post(req)

	if err != nil {
		return "", fmt.Errorf("%s - addGenericRecord request", err)
	}

	if strings.Contains(resp.String(), "Invalid") {
		return "", fmt.Errorf("%s - addGenericRecord response", resp.String())
	}

	return resp.String(), nil
}
