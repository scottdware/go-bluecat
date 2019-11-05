package bluecat

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

// GetEntitiesByName returns an array of entities that match the specified parent, name, and object type.
//
// Parameter `name` is the name of the entity. Parameter `parentid` is the object ID of the parent object of the
// entities to be returned. Parameter `objecttype` is the type of object to be returned. This value must be one of the
// object types constants. Parameter `count` is the maximum number of objects to return. The default value is 10.
// This value cannot be null or empty. Parameter `start` indicates where in the list of returned objects to start
// returning objects. The list begins at an index of 0. This value cannot be null or empty.
func (b *Bluecat) GetEntitiesByName(name string, parentid int64, objecttype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getEntitiesByName?name=%s&parentId=%d&type=%s&count=%d&start=%d",
		b.Server, b.URI, name, parentid, objecttype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetEntitesByName request error", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetEntitiesByName JSON error", err)
	}

	return results, nil
}

// GetEntities returns multiple entities for the specified parent ID.
//
// Parameter `parentid` is the object ID of the parent object of the entities. Parameter `objecttype` is the type of
// object to be returned. This value must be one of the object types constants. Parameter `count` indicates the maximum
// number of child objects to return. Parameter `start` indicates where in the list of child objects to start returning
// entities. The list begins at an index of 0.
func (b *Bluecat) GetEntities(parentid int64, objecttype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getEntities?parentId=%d&type=%s&count=%d&start=%d",
		b.Server, b.URI, parentid, objecttype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("GetEntities request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("GetEntities JSON parse error: %s", err)
	}

	return results, nil
}

// GetEntityByCIDR returns an IPv4 Network object from the database by calling it using CIDR notation.
//
// Parameter `cidr` is the CIDR notation of the IP4Network object type. Parameter `parentid` is the object ID of the
// network’s parent object. Parameter `objecttype` is the type of object returned: IP4Network. This must be one of the
// constants types constants.
func (b *Bluecat) GetEntityByCIDR(cidr string, parentid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByCIDR?cidr=%s&parentId=%d&type=%s",
		b.Server, b.URI, cidr, parentid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("GetEntityByCIDR request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("GetEntityByCIDR JSON parse error: %s", err)
	}

	return results, nil
}

// GetEntityByID returns objects from the database referenced by their database ID and with its properties fields populated.
// For more information about the available options, refer to IPv4 objects in the Property Options Reference section of the API
// guide.
//
// Parameter `id` is the object ID of the target object.
func (b *Bluecat) GetEntityByID(id int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityById?id=%d",
		b.Server, b.URI, id)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("GetEntityByID request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("GetEntityByID JSON parse error: %s", err)
	}

	return results, nil
}

// GetEntityByName returns objects from the database referenced by their name field.
//
// Parameter `name` is the name of the entity. Parameter `parentid` is the ID of the target object’s parent object.
// Parameter `objecttype` is the type of object returned by the method. This string must be one of the object type constants.
func (b *Bluecat) GetEntityByName(name string, parentid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByName?name=%s&parentId=%d&type=%s",
		b.Server, b.URI, name, parentid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("GetEntityByName request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("GetEntityByName JSON parse error: %s", err)
	}

	return results, nil
}

// GetEntityByPrefix returns an APIEntity for the specified IP block or network.
//
// Parameter `containerid` is the object ID of higher-level parent object IP block or configuration
// in which the IP block or network is located. Parameter `prefix` is the prefix value for the IP block or network.
// This value cannot be empty. Parameter `objecttype` is the type of object to be returned. This string must be one of
// the object type constants.
func (b *Bluecat) GetEntityByPrefix(containerid int64, prefix, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByPrefix?containerId=%d&prefix=%s&type=%s",
		b.Server, b.URI, containerid, prefix, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("GetEntityByPrefix request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("GetEntityByPrefix JSON parse error: %s", err)
	}

	return results, nil
}

// GetEntityByRange returns an IPv4 DHCP range object by calling it using its range.
//
// Parameter `address1` is an IP address defining the lowest address or start of the range. Parameter `address2` is an
// IP address defining the highest address or end of the range. Parameter `parentid` is the object ID of the parent object
// of the DHCP range. Parameter `objecttype` is the type of object returned: DHCP4Range. This must be one of the object type constants.
func (b *Bluecat) GetEntityByRange(address1, address2 string, parentid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByRange?address1=%s&address2=%s&parentId=%d&type=%s",
		b.Server, b.URI, address1, address2, parentid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("GetEntityByRange request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("GetEntityByRange JSON parse error: %s", err)
	}

	return results, nil
}

// CustomSearch searches for an array of entities by specifying object properties.
//
// Parameter `objecttype` must be one of the following object types: IP4Block, IP4Network, IP4Addr, GenericRecord, HostRecord,
// Any other objects with user-defined fields. Parameter `filters` is the list of properties on which the search will be based.
// The valid format is Field name=value. Parameter `count` is the maximum number of objects to return. The value must be a
// positive value between 1 and 1000. This value cannot be null or empty. Parameter `start` indicates where in the list of
// returned objects to start returning objects. The value must be a non-negative value and cannot be null or empty.
func (b *Bluecat) CustomSearch(objecttype, filters string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/customSearch?type=%s&filters=%s&count=%d&start=%d",
		b.Server, b.URI, objecttype, filters, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("CustomSearch request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("CustomSearch JSON parse error: %s", err)
	}

	return results, nil
}

// SearchByCategory returns an array of entities by searching for keywords associated with objects of
// a specified object category.
//
// Parameter `keyword` is the search keyword string. This value cannot be null or empty. Parameter `category`
// is the entity category to be searched. Parameter `count` is the maximum number of objects to return.
// The default value is 10. This value cannot be null or empty. Parameter `start` indicates where in the list
// of returned objects to start returning objects. The list begins at an index of 0. This value cannot be null or empty.
func (b *Bluecat) SearchByCategory(keyword, category string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/searchByCategory?keyword=%s&category=%s&count=%d&start=%d",
		b.Server, b.URI, keyword, category, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("SearchByCategory request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("SearchByCategory JSON parse error: %s", err)
	}

	return results, nil
}

// SearchByObjectTypes returns an array of entities by searching for keywords associated with objects of a
// specified object type. You can search for multiple object types with a single method call.
//
// Parameter `keyword` is the search keyword string. This value cannot be null or empty. Parameter `objecttypes`
// is the object types for which to search, specified in the format: "type1[,type2...]". Parameter `count` is
// the maximum number of objects to return. The default value is 10. This value cannot be null or empty.
// Parameter `start` indicates where in the list of returned objects to start returning objects. The list begins
// at an index of 0. This value cannot be null or empty.
func (b *Bluecat) SearchByObjectTypes(keyword, objecttypes string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/searchByObjectTypes?keyword=%s&types=%s&count=%d&start=%d",
		b.Server, b.URI, keyword, objecttypes, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("SearchByObjectTypes request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("SearchByObjectTypes JSON parse error: %s", err)
	}

	return results, nil
}

// FindResponsePoliciesWithItem finds local DNS response policies with their associated response policy items.
//
// Parameter `configid` is the object ID of the configuration to which the local response policies are located.
// To view a complete list of all local response policies under all configurations that have an associated response
// policy item, set the value of this parameter to 0. Parameter `itemname` is the Fully Qualified Domain Name
// (FQDN) of the response policy item. The exact FQDN of the response policy item must be used when conducting a search.
func (b *Bluecat) FindResponsePoliciesWithItem(configid int64, itemname string) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/findResponsePoliciesWithItem?configurationId=%d&itemName=%s",
		b.Server, b.URI, configid, itemname)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("FindResponsePoliciesWithItem request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("FindResponsePoliciesWithItem JSON parse error: %s", err)
	}

	return results, nil
}

// GetAccessRight retrieves an access right for a specified object.
//
// Parameter `entityid` is the object ID of the entity to which the access right is assigned.
// Parameter `userid` is the object ID of the user to whom the access right is applied.
func (b *Bluecat) GetAccessRight(entityid, userid int64) (APIAccessRight, error) {
	var results APIAccessRight
	req := fmt.Sprintf("https://%s%s/getAccessRight?entityId=%d&userId=%d",
		b.Server, b.URI, entityid, userid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("GetAccessRight request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("GetAccessRight JSON parse error: %s", err)
	}

	return results, nil
}

// GetAccessRightsForEntity retrieves an access right for a specified object.
//
// Parameter `entityid` is the object ID of the entity whose access rights are returned.
// Parameter `count ` is the maximum number of access right child objects to return. Parameter `start`
// indicates where in the list of child access right objects to start returning objects.
// The list begins at an index of 0.
func (b *Bluecat) GetAccessRightsForEntity(entityid int64, count, start int32) ([]APIAccessRight, error) {
	var results []APIAccessRight
	req := fmt.Sprintf("https://%s%s/getAccessRightsForEntity?entityId=%d&count=%d&start=%d",
		b.Server, b.URI, entityid, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("GetAccessRightsForEntity request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("GetAccessRightsForEntity JSON parse error: %s", err)
	}

	return results, nil
}

// GetAccessRightsForUser returns an array of access rights for a specified user.
//
// Parameter `userid` is the object ID of the user whose access rights are returned. Parameter
// `count` is the maximum number of access right child objects to return. Parameter `start`
// indicates where in the list of child access right objects to start returning objects.
// The list begins at an index of 0.
func (b *Bluecat) GetAccessRightsForUser(userid int64, count, start int32) ([]APIAccessRight, error) {
	var results []APIAccessRight
	req := fmt.Sprintf("https://%s%s/getAccessRightsForUser?userId=%d&count=%d&start=%d",
		b.Server, b.URI, userid, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("GetAccessRightsForUser request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("GetAccessRightsForUser JSON parse error: %s", err)
	}

	return results, nil
}

// GetAdditionalIPAddresses returns IPv4 addresses and loopback addresses added to the Service interface for DNS services.
//
// Parameter `adonisid` is the database object ID of the server on which additional services IP address have been added.
// The supported `properties` parameter is: serviceType — type of service for which a list of IP addresses will be retrieved.
// Available types are AdditionalIPServiceType.SERVICE and AdditionalIPServiceType.LOOPBACK. If serviceType is not provided,
// all additional IP addresses of the services interface will be returned.
//
// The return `string` returns the list of additional IP addresses configured on the server in the format: [IP,serviceType|IP,serviceType].
// For example, 10.0.0.10/32,loopback|11.0.0.3/24,service|12.0.0.3/32,loopback.
func (b *Bluecat) GetAdditionalIPAddresses(adonisid int64, properties string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getAdditionalIPAddresses?adonisId=%d&properties=%s",
		b.Server, b.URI, adonisid, properties)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("GetAdditionalIPAddresses request error: %s", err)
	}

	return resp.String(), nil
}

// GetAliasesByHint returns an array of CNAMEs with linked record name.
//
// Parameter `options` is a string containing options. The supported options are hint and retrieveFields. Separate multiple
// options with a pipe character.
//
// For example: hint=^abc|retrieveFields=false
//
// If the hint option is not specified in the string, searching criteria will be based on the same as zone alias.
// The following wildcards are supported in the hint option.
//
// ^ — matches the beginning of a string. For example: ^ex matches ex ample but not t ex t.
//
// $ — matches the end of a string. For example: ple$ matches exam ple but not ple ase.
//
// ^ $ — matches the exact characters between the two wildcards. For example: ^example$ only matches example.
//
// ? — matches any one character. For example: ex?t matches ex i t.
//
// * — matches one or more characters within a string. For example: ex*t matches exit and ex cellen t.
//
// The default value for the retrieveFields option is set to false. If the option is set to true, user-defined field will
// be returned. If the options string does not contain retrieveFields, user-defined field will not be returned. Parameter
// `count` indicates the maximum of child objects that this method will return. The value must be less than or equal to 10.
// Parameter `start` indicates where in the list of objects to start returning objects. The list begins at an index of 0.
func (b *Bluecat) GetAliasesByHint(options string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getAliasesByHint?options=%s&count=%d&start=%d",
		b.Server, b.URI, options, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("GetAliasesByHint request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("GetAliasesByHint JSON parse error: %s", err)
	}

	return results, nil
}

// GetAllUsedLocations returns a list of location objects that are used to annotate other objects.
func (b *Bluecat) GetAllUsedLocations() ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getAllUsedLocations",
		b.Server, b.URI)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("GetAllUsedLocations request error: %s", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("GetAllUsedLocations JSON parse error: %s", err)
	}

	return results, nil
}

// GetConfigurationGroups gets a list of all configuration groups in Address Manager.
func (b *Bluecat) GetConfigurationGroups() (string, error) {
	//var results []string
	req := fmt.Sprintf("https://%s%s/getConfigurationGroups",
		b.Server, b.URI)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("GetConfigurationGroups request error: %s", err)
	}

	//if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
	//	return nil, fmt.Errorf("GetConfigurationGroups JSON parse error: %s", err)
	//}

	return resp.String(), nil
}
