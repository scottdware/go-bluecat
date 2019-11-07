package bluecat

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/resty.v1"
)

// GetEntitiesByName returns an array of entities that match the specified parent, name, and object type.
//
// Parameter `name` is the name of the entity. Parameter `parentid` is the object ID of the parent object of the
// entities to be returned. Parameter `objecttype` is the type of object to be returned. This value must be one of the
// object types constants. Parameter `count` is the maximum number of objects to return. The default value is 10.
// This value cannot be null or empty. Parameter `start` indicates where in the list of returned objects to start
// returning objects. The list begins at an index of 0. This value cannot be null or empty.
//
// Returns an array of type APIEntity. The array is empty if there are no matching entities.
func (b *Bluecat) GetEntitiesByName(name string, parentid int64, objecttype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getEntitiesByName?name=%s&parentId=%d&type=%s&count=%d&start=%d",
		b.Server, b.URI, name, parentid, objecttype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetEntitesByName request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetEntitiesByName JSON parse", err)
	}

	return results, nil
}

// GetEntities returns multiple entities for the specified parent ID.
//
// Parameter `parentid` is the object ID of the parent object of the entities. Parameter `objecttype` is the type of
// object to be returned. This value must be one of the object types constants. Parameter `count` indicates the maximum
// number of child objects to return. Parameter `start` indicates where in the list of child objects to start returning
// entities. The list begins at an index of 0.
//
// Returns an array of type APIEntity. The array is empty if there are no matching entities.
func (b *Bluecat) GetEntities(parentid int64, objecttype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getEntities?parentId=%d&type=%s&count=%d&start=%d",
		b.Server, b.URI, parentid, objecttype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetEntities request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetEntities JSON parse", err)
	}

	return results, nil
}

// GetEntityByCIDR returns an IPv4 Network object from the database by calling it using CIDR notation.
//
// Parameter `cidr` is the CIDR notation of the IP4Network object type. Parameter `parentid` is the object ID of the
// network’s parent object. Parameter `objecttype` is the type of object returned: IP4Network. This must be one of the
// constants types constants.
//
// Returns the specified IPv4 block object from the database. Return type is APIEntity.
func (b *Bluecat) GetEntityByCIDR(cidr string, parentid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByCIDR?cidr=%s&parentId=%d&type=%s",
		b.Server, b.URI, cidr, parentid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetEntityByCIDR request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetEntityByCIDR JSON parse", err)
	}

	return results, nil
}

// GetEntityByID returns objects from the database referenced by their database ID and with its properties fields populated.
// For more information about the available options, refer to IPv4 objects in the Property Options Reference section of the API
// guide.
//
// Parameter `id` is the object ID of the target object.
//
// Returns the requested object from the database with its properties fields populated. Retury type is APIEntity.
func (b *Bluecat) GetEntityByID(id int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityById?id=%d",
		b.Server, b.URI, id)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetEntityByID request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetEntityByID JSON parse", err)
	}

	return results, nil
}

// GetEntityByName returns objects from the database referenced by their name field.
//
// Parameter `name` is the name of the entity. Parameter `parentid` is the ID of the target object’s parent object.
// Parameter `objecttype` is the type of object returned by the method. This string must be one of the object type constants.
//
// Returns an array of entities. The array is empty if there are no matching entities. Return type is APIEntity.
func (b *Bluecat) GetEntityByName(name string, parentid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByName?name=%s&parentId=%d&type=%s",
		b.Server, b.URI, name, parentid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetEntityByName request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetEntityByName JSON parse", err)
	}

	return results, nil
}

// GetEntityByPrefix returns an APIEntity for the specified IP block or network.
//
// Parameter `containerid` is the object ID of higher-level parent object IP block or configuration
// in which the IP block or network is located. Parameter `prefix` is the prefix value for the IP block or network.
// This value cannot be empty. Parameter `objecttype` is the type of object to be returned. This string must be one of
// the object type constants.
//
// Returns an APIEntity for the specified IPv6 block or network. The APIEntity is empty if the block or network does not exist.
func (b *Bluecat) GetEntityByPrefix(containerid int64, prefix, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByPrefix?containerId=%d&prefix=%s&type=%s",
		b.Server, b.URI, containerid, prefix, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetEntityByPrefix request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetEntityByPrefix JSON parse", err)
	}

	return results, nil
}

// GetEntityByRange returns an IPv4 DHCP range object by calling it using its range.
//
// Parameter `address1` is an IP address defining the lowest address or start of the range. Parameter `address2` is an
// IP address defining the highest address or end of the range. Parameter `parentid` is the object ID of the parent object
// of the DHCP range. Parameter `objecttype` is the type of object returned: DHCP4Range. This must be one of the object type constants.
//
// Returns the requested IPv4 block object from the database. Return type is APIEntity.
func (b *Bluecat) GetEntityByRange(address1, address2 string, parentid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getEntityByRange?address1=%s&address2=%s&parentId=%d&type=%s",
		b.Server, b.URI, address1, address2, parentid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetEntityByRange request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetEntityByRange JSON parse", err)
	}

	return results, nil
}

// CustomSearch searches for an array of entities by specifying object properties.
//
// Parameter `filters` is the list of properties on which the search will be based. The valid format is Field name=value.
// Refer to Reference: Supported search object types and fields for details. The field name is case-sensitive. In addition
// to the fields that are specified in the table, any user-defined fields will also be supported. The valid format for the
// Date type user-defined field value is DD-MMM-YYYY. You can also use partial formatting. For example, 10-Jan-2016, 10-Jan, Jan-2016 or 2016.
//
// Parameter `objecttype` must be one of the following object types: IP4Block, IP4Network, IP4Addr, GenericRecord, HostRecord,
// Any other objects with user-defined fields. Parameter `count` is the maximum number of objects to return. The value must be a
// positive value between 1 and 1000. This value cannot be null or empty. Parameter `start` indicates where in the list of
// returned objects to start returning objects. The value must be a non-negative value and cannot be null or empty.
//
// Returns an array of type APIEntity matching the specified object properties or returns an empty array. The APIEntity will
// at least contain Object Type, Object ID, Object Name, and Object Properties.
func (b *Bluecat) CustomSearch(filters, objecttype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/customSearch?filters=%s&type=%s&count=%d&start=%d",
		b.Server, b.URI, filters, objecttype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - CustomSearch request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - CustomSearch JSON parse", err)
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
//
// Returns an array of type APIEntity matching the keyword text and the category type, or returns an empty array.
func (b *Bluecat) SearchByCategory(keyword, category string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/searchByCategory?keyword=%s&category=%s&count=%d&start=%d",
		b.Server, b.URI, keyword, category, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - SearchByCategory request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - SearchByCategory JSON parse", err)
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
//
// Returns an array of type APIEntity matching the keyword text and the category type, or returns an empty array.
func (b *Bluecat) SearchByObjectTypes(keyword, objecttypes string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/searchByObjectTypes?keyword=%s&types=%s&count=%d&start=%d",
		b.Server, b.URI, keyword, objecttypes, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - SearchByObjectTypes request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - SearchByObjectTypes JSON parse", err)
	}

	return results, nil
}

// SearchResponsePolicyItem searches Response Policy items configured in local Response Policies or predefined BlueCat
// Security feed data. The search will return a list of all matching items in Address Manager across all configurations.
//
// Parameter `keyword` is the search string for which you wish to search.
//
// ^ — matches the beginning of a string. For example: ^ex matches ex ample but not t ex t.
//
// $ — atches the end of string. For example: ple$ matches exam ple but not ple ase.
//
// * — matches zero or more characters within a string. For example: ex*t matches exit and ex cellen t.
//
// Parameter `scope` is the scope in which the search is to be performed. The possible values are:
//
// RPItemSearchScope.LOCAL — to search policy items configured in local Response Policies.
//
// RPItemSearchScope.FEED — to search policy items configured in predefined BlueCat Security Feed data.
//
// RPItemSearchScope.ALL — to search policy items configured in both local Response Policies and predefined BlueCat
// Security Feed data.
//
// Parameter `count` is the total number of results to be returned. The possible value is a positive integer ranging
// from 1 to 1000. Parameter `start` is a starting number from where the search result will be returned. The possible
// value is a positive integer ranging from 0 to 999. For example, specifying 99 will return the search result from the
// 100th result to the maximum number that you specify with the count option.
//
// Returns an array of ResponsePolicySearchResult objects. Each object contains information of one Response Policy item
// found either in local Response Policies or BlueCat Security feed data.
func (b *Bluecat) SearchResponsePolicyItem(keyword, scope string, count, start int32) ([]ResponsePolicySearchResult, error) {
	var results []ResponsePolicySearchResult
	req := fmt.Sprintf("https://%s%s/searchResponsePolicyItem?keyword=%s&scope=%s&count=%d&start=%d",
		b.Server, b.URI, keyword, scope, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - SearchResponsePolicyItem request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - SearchResponsePolicyItem JSON parse", err)
	}

	return results, nil
}

// FindResponsePoliciesWithItem finds local DNS response policies with their associated response policy items.
//
// Parameter `configid` is the object ID of the configuration to which the local response policies are located.
// To view a complete list of all local response policies under all configurations that have an associated response
// policy item, set the value of this parameter to 0. Parameter `itemname` is the Fully Qualified Domain Name
// (FQDN) of the response policy item. The exact FQDN of the response policy item must be used when conducting a search.
//
// Returns a list (array) of type APIEntity, of local response policies along with the associated response policy item under
// a specific configuration or all configurations. This is determined by the input provided for the configurationId parameter.
func (b *Bluecat) FindResponsePoliciesWithItem(configid int64, itemname string) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/findResponsePoliciesWithItem?configurationId=%d&itemName=%s",
		b.Server, b.URI, configid, itemname)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - FindResponsePoliciesWithItem request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - FindResponsePoliciesWithItem JSON parse", err)
	}

	return results, nil
}

// GetAccessRight retrieves an access right for a specified object.
//
// Parameter `entityid` is the object ID of the entity to which the access right is assigned.
// Parameter `userid` is the object ID of the user to whom the access right is applied.
//
// Returns the access right for the specified object. Return type is APIAccessRight.
func (b *Bluecat) GetAccessRight(entityid, userid int64) (APIAccessRight, error) {
	var results APIAccessRight
	req := fmt.Sprintf("https://%s%s/getAccessRight?entityId=%d&userId=%d",
		b.Server, b.URI, entityid, userid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetAccessRight request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetAccessRight JSON parse", err)
	}

	return results, nil
}

// GetAccessRightsForEntity retrieves an access right for a specified object.
//
// Parameter `entityid` is the object ID of the entity whose access rights are returned.
// Parameter `count ` is the maximum number of access right child objects to return. Parameter `start`
// indicates where in the list of child access right objects to start returning objects.
// The list begins at an index of 0.
//
// Returns an array of type APIAccessRight objects.
func (b *Bluecat) GetAccessRightsForEntity(entityid int64, count, start int32) ([]APIAccessRight, error) {
	var results []APIAccessRight
	req := fmt.Sprintf("https://%s%s/getAccessRightsForEntity?entityId=%d&count=%d&start=%d",
		b.Server, b.URI, entityid, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetAccessRightsForEntity request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetAccessRightsForEntity JSON parse", err)
	}

	return results, nil
}

// GetAccessRightsForUser returns an array of access rights for a specified user.
//
// Parameter `userid` is the object ID of the user whose access rights are returned. Parameter
// `count` is the maximum number of access right child objects to return. Parameter `start`
// indicates where in the list of child access right objects to start returning objects.
// The list begins at an index of 0.
//
// Returns an array of type APIAccessRight objects.
func (b *Bluecat) GetAccessRightsForUser(userid int64, count, start int32) ([]APIAccessRight, error) {
	var results []APIAccessRight
	req := fmt.Sprintf("https://%s%s/getAccessRightsForUser?userId=%d&count=%d&start=%d",
		b.Server, b.URI, userid, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetAccessRightsForUser request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetAccessRightsForUser JSON parse", err)
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
//
// Returns the list of additional IP addresses configured on the server in the format: [IP,serviceType|IP,serviceType].
// For example, 10.0.0.10/32,loopback|11.0.0.3/24,service|12.0.0.3/32,loopback. Return type is a string.
func (b *Bluecat) GetAdditionalIPAddresses(adonisid int64, properties string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getAdditionalIPAddresses?adonisId=%d&properties=%s",
		b.Server, b.URI, adonisid, properties)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetAdditionalIPAddresses request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
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
//
// Returns an array of Alias APIEntity objects.
func (b *Bluecat) GetAliasesByHint(options string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getAliasesByHint?options=%s&count=%d&start=%d",
		b.Server, b.URI, options, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetAliasesByHint request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetAliasesByHint JSON parse", err)
	}

	return results, nil
}

// GetAllUsedLocations returns a list of location objects that are used to annotate other objects.
//
// Returns an array of location APIEntity objects.
func (b *Bluecat) GetAllUsedLocations() ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getAllUsedLocations",
		b.Server, b.URI)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetAllUsedLocations request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetAllUsedLocations JSON parse", err)
	}

	return results, nil
}

// GetConfigurationGroups gets a list of all configuration groups in Address Manager.
//
// Returns a list of configuration groups. Return type is a string.
func (b *Bluecat) GetConfigurationGroups() (string, error) {
	req := fmt.Sprintf("https://%s%s/getConfigurationGroups",
		b.Server, b.URI)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetConfigurationGroups request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetConfigurationSetting returns the configuration setting.
//
// Parameter `configurationid` is the object ID of the configuration in which the setting is to be located. Parameter
// `setting` is the name of the specific setting to be read. Only the option inheritance OPTION_INHERITANCE setting is supported.
//
// Returns the properties of the setting of the configuration. Return type is a string.
func (b *Bluecat) GetConfigurationSetting(configurationid int64, setting string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getConfigurationSetting?configurationId=%d&settingName=%s",
		b.Server, b.URI, configurationid, setting)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetConfigurationGroups request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetConfigurationsByGroup gets a list of configurations in Address Manager based on the name of a configuration group.
//
// Parameter `group` is the name of the configuration group in which the configurations are to be located. Configuration
// group names are case sensitive.
//
// Returns a list/array of type APIEntity, of configurations based on the specified group.
func (b *Bluecat) GetConfigurationsByGroup(group string) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getConfigurationsByGroup?groupName=%s",
		b.Server, b.URI, group)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return nil, fmt.Errorf("%s - GetConfigurationsByGroup request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return nil, fmt.Errorf("%s - GetConfigurationsByGroup JSON parse", err)
	}

	return results, nil
}

// GetDHCP6ClientDeploymentOption returns DHCPv6 client options assigned for the object specified excluding the options
// inherited from the higher-level parent object.
//
// Parameter `entityid` is the object ID for the entity. Parameter `name` is the name of the DHCPv6 client option being added.
// This name must be one of the constants listed in DHCPv6 client options. Parameter `serverid` is the specific server or server
// group to which this option is deployed. To return an option that has not been assigned to a server role, set this value
// to zero. Omitting this parameter from the method call will result in an error.
//
// Returns the specified DHCPv6 client option object from the database. Return type is APIDeploymentOption.
func (b *Bluecat) GetDHCP6ClientDeploymentOption(entityid int64, name string, serverid int64) (APIDeploymentOption, error) {
	var results APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDHCP6ClientDeploymentOption?entityId=%d&name=%s&serverId=%d",
		b.Server, b.URI, entityid, name, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDHCP6ClientDeploymentOption request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDHCP6ClientDeploymentOption JSON parse", err)
	}

	return results, nil
}

// GetDHCP6ServiceDeploymentOption returns DHCPv6 service options assigned for the object specified excluding the options
// inherited from the higher-level parent object.
//
// Parameter `entityid` is the database object ID for the entity to which the deployment option is assigned. Parameter `name`
// is the name of the DHCPv6 service option being added. This name must be one of the constants listed in DHCPv6 service options.
// Parameter `serverid` specifies the server or server group to which the option is deployed for the specified entity.
// To retrieve an option that has not been assigned to a server role, set this value to zero. Omitting this parameter from
// the method call will result in an error.
//
// Returns the requested DHCPv6 service option object from the database. Return type is APIDeploymentOption.
func (b *Bluecat) GetDHCP6ServiceDeploymentOption(entityid int64, name string, serverid int64) (APIDeploymentOption, error) {
	var results APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDHCP6ServiceDeploymentOption?entityId=%d&name=%s&serverId=%d",
		b.Server, b.URI, entityid, name, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDHCP6ServiceDeploymentOption request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDHCP6ServiceDeploymentOption JSON parse", err)
	}

	return results, nil
}

// GetDHCPClientDeploymentOption returns DHCPv4 client options assigned for the object specified excluding the options
// inherited from the higher-level parent object.
//
// Parameter `entityid` is the object ID for the entity to which the deployment option has been applied. Parameter `name`
// is the name of the DHCPv4 client option being added. This name must be one of the constants listed in DHCP client options.
// Parameter `serverid` is the specific server or server group to which this option is deployed. To return an option that
// has not been assigned to a server, set this value to 0 (zero). Omitting this parameter from the method call will
// result in an error.
//
// Returns the specified DHCPv4 client option object from the database. Return type is APIDeploymentOption.
func (b *Bluecat) GetDHCPClientDeploymentOption(entityid int64, name string, serverid int64) (APIDeploymentOption, error) {
	var results APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDHCPClientDeploymentOption?entityId=%d&name=%s&serverId=%d",
		b.Server, b.URI, entityid, name, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDHCPClientDeploymentOption request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDHCPClientDeploymentOption JSON parse", err)
	}

	return results, nil
}

// GetDHCPDeploymentRole retrieves the DHCP deployment role assigned to a specified object.
//
// Parameter `entityid` is the object ID for the object to which the deployment role is assigned. Parameter `serverinterfaceid`
// is the object ID of the server interface to which the role is assigned.
//
// Returns the DHCP deployment role assigned to the specified object, or returns an empty APIDeploymentRol if no role
// is defined. Return type is APIDeploymentRole.
func (b *Bluecat) GetDHCPDeploymentRole(entityid, serverinterfaceid int64) (APIDeploymentRole, error) {
	var results APIDeploymentRole
	req := fmt.Sprintf("https://%s%s/getDHCPDeploymentRole?entityId=%d&serverInterfaceId=%d",
		b.Server, b.URI, entityid, serverinterfaceid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDHCPDeploymentRole request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDHCPDeploymentRole JSON parse", err)
	}

	return results, nil
}

// GetDHCPServiceDeploymentOption returns DHCP service options assigned for the object specified excluding the options
// inherited from the higher-level parent object.
//
// Parameter `entityid` is the object ID for the entity to which the deployment option is assigned. Parameter `name`
// is the name of the DHCPv4 service option being retrieved. This name must be one of the constants listed in DHCP service options.
// Parameter `serverid` specifies the server or server group to which the option is deployed for the specified entity.
// To retrieve an option that has not been assigned to a server role, specify 0 as a value. Omitting this parameter from
// the method call will result in an error.
//
// Returns the requested DHCPv4 service option object from the database. Return type is APIDeploymentOption.
func (b *Bluecat) GetDHCPServiceDeploymentOption(entityid int64, name string, serverid int64) (APIDeploymentOption, error) {
	var results APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDHCPServiceDeploymentOption?entityId=%d&name=%s&serverId=%d",
		b.Server, b.URI, entityid, name, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDHCPServiceDeploymentOption request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDHCPServiceDeploymentOption JSON parse", err)
	}

	return results, nil
}

// GetDHCPVendorDeploymentOption retrieves a DHCP vendor deployment option assigned for the object specified excluding the
// options inherited from the higher-level parent object.
//
// Parameter `entityid` is the object ID of the entity to which the DHCP vendor deployment option is assigned. This must
// be the ID of a configuration, IPv4 block, IPv4 network, IPv4 address, IPv4 DHCP rage, server, MAC address, or MAC Pool.
// Parameter `optionid` is the object ID of the DHCP vendor option definition. Parameter `serverid` is the specific
// server or server group to which this option is deployed for the specified entity. To return an option that has not been
// assigned to a server, set this value to 0 (zero). Omitting this parameter from the method call will result in an error.
//
// Returns an APIDeploymentOption for the DHCP vendor client deployment option. Return type is APIDeploymentOption.
func (b *Bluecat) GetDHCPVendorDeploymentOption(entityid, optionid, serverid int64) (APIDeploymentOption, error) {
	var results APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDHCPVendorDeploymentOption?entityId=%d&optionId=%d&serverId=%d",
		b.Server, b.URI, entityid, optionid, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDHCPVendorDeploymentOption request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDHCPVendorDeploymentOption JSON parse", err)
	}

	return results, nil
}

// GetDNSDeploymentOption retrieves all DNS options assigned for the object specified excluding the options inherited
// from the higher-level parent object.
//
// Parameter `entityid` is the object ID for the entity to which this deployment option is assigned. Parameter `name` is the
// name of the DNS option. This name must be one of the constants listed in DNS options. Parameter `serverid` specifies
// the server or server group to which this option is assigned. To retrieve an option that has not been assigned to a
// server role, set this value to 0 (zero). Omitting this parameter from the method call will result in an error.
//
// Returns an instance of the type APIDeploymentOption that represents the DNS deployment option or empty if none were found.
// Return type is APIDeploymentOption.
func (b *Bluecat) GetDNSDeploymentOption(entityid int64, name string, serverid int64) (APIDeploymentOption, error) {
	var results APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDNSDeploymentOption?entityId=%d&name=%s&serverId=%d",
		b.Server, b.URI, entityid, name, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDNSDeploymentOption request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDNSDeploymentOption JSON parse", err)
	}

	return results, nil
}

// GetDNSDeploymentRoleForView retrieves the DNS deployment role assigned to a view-level objects in the IP space for ARPA zones.
//
// Parameter `entityid` is the object ID for the object to which the DNS deployment role is assigned. Parameter `serverinterfaceid`
// is the object ID of the server interface to which the DNS deployment role is assigned. Parameter `viewid` is the
// view in which the DNS deployment role is assigned.
//
// Returns the requested APIDeploymentRole object. Return type is APIDeploymentRole.
func (b *Bluecat) GetDNSDeploymentRoleForView(entityid, serverinterfaceid, viewid int64) (APIDeploymentRole, error) {
	var results APIDeploymentRole
	req := fmt.Sprintf("https://%s%s/getDNSDeploymentRoleForView?entityId=%d&serverInterfaceId=%d&viewId=%d",
		b.Server, b.URI, entityid, serverinterfaceid, viewid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDNSDeploymentRoleForView request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDNSDeploymentRoleForView JSON parse", err)
	}

	return results, nil
}

// GetDNSDeploymentRole retrieves a DNS deployment role from a specified object.
//
// Parameter `entityid` is the object ID for the object to which the DNS deployment role is assigned. Parameter
// `serverinterfaceid` is the object ID of the server interface to which the DNS deployment role is assigned.
//
// Returns a DNS deployment role from the specified object, or returns an empty APIDeploymentRole if no role is defined.
// Return type is APIDeploymentRole.
func (b *Bluecat) GetDNSDeploymentRole(entityid, serverinterfaceid int64) (APIDeploymentRole, error) {
	var results APIDeploymentRole
	req := fmt.Sprintf("https://%s%s/getDNSDeploymentRole?entityId=%d&serverInterfaceId=%d",
		b.Server, b.URI, entityid, serverinterfaceid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDNSDeploymentRole request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDNSDeploymentRole JSON parse", err)
	}

	return results, nil
}

// GetDeploymentOptions retrieves deployment options for Address Manager DNS and DHCP services.
//
// Parameter `entityid` is the object ID of the entity to which the DNS or DHCP deployment option is assigned. Parameter
// `optiontypes` is the type of deployment options. Multiple options can be separated by a | (pipe) character.
// This value must be one of the following items:
//
// DNSOption
//
// DNSRawOption
//
// DHCPRawOption
//
// DHCPV6RawOption
//
// DHCPV4ClientOption
//
// DHCPV6ClientOption
//
// DHCPServiceOption
//
// DHCPV6ServiceOption
//
// VendorClientOption
//
// StartOfAuthority
//
// For complete list of Option Types and Object Types constants, refer to Option types and Option types.
//
// If Invalid deployment option types or invalid strings are specified, the API execution will fail and return the
// error message: " Invalid deployment option found ". For example, if the user passes DHCPv6ClientOption for IPv4 networks,
// it will return this error message as DHCPv6 client options are not a valid for IPv4 networks.
//
// If specified as an empty string "", all deployment options for the specified entity will be returned. Depending on the type
// of DNS deployment option being retrieved, the format of the value might differ. For more information, refer to
// Reference: Deployment option value formats.
//
// Parameter `serverid` is the specific server or server group to which options are deployed. The valid values are as follows:
//
// >0 — returns only the options that are linked to the specified server ID.
//
// <0 — returns all options regardless of the server ID specified.
//
// =0 — returns only the options that are linked to all servers.
//
// Returns all deployment options, array of type APIDeploymentOption, assigned to the specified object including inherited
// options from higher level parent objects. If an option is inherited and overridden, then only the overriding option will be returned.
func (b *Bluecat) GetDeploymentOptions(entityid int64, optiontypes string, serverid int64) ([]APIDeploymentOption, error) {
	var results []APIDeploymentOption
	req := fmt.Sprintf("https://%s%s/getDeploymentOptions?entityId=%d&optionTypes=%s&serverId=%d",
		b.Server, b.URI, entityid, optiontypes, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDeploymentOptions request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDeploymentOptions JSON parse", err)
	}

	return results, nil
}

// GetDeploymentRoles returns the DNS and DHCP deployment roles associated with the specified object. For DNS views and zones,
// GetDeploymentRoles returns DNS deployment roles. For IP address space objects, such as IPv4 blocks and networks, IPv6 blocks
// and networks, DHCP classes, and MAC pools, GetDeploymentRoles returns DNS and DHCP deployment roles.
//
// Returns an array of APIDeploymentRole objects representing the deployment roles associated with the specified object.
func (b *Bluecat) GetDeploymentRoles(entityid int64) ([]APIDeploymentRole, error) {
	var results []APIDeploymentRole
	req := fmt.Sprintf("https://%s%s/getDeploymentRoles?entityId=%d",
		b.Server, b.URI, entityid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDeploymentRoles request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDeploymentRoles JSON parse", err)
	}

	return results, nil
}

// GetDeploymentTaskStatus returns the deployment status of the deployment task that was created using the selectiveDeploy API method.
//
// Parameter `deploymenttasktoken` is the string token value that is returned from the selectiveDeploy} API method.
//
// Returns a string value of the overall deployment status and the deployment status of individual entities. Return type is a string.
func (b *Bluecat) GetDeploymentTaskStatus(deploymenttasktoken string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getDeploymentRoles?entityId=%s",
		b.Server, b.URI, deploymenttasktoken)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetDeploymentTaskStatus request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetDiscoveredDeviceArpEntries returns all ARP entries of a specific device discovered by running an IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4 reconciliation policy.
//
// Returns all ARP entries of a specific device. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDeviceArpEntries(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDeviceArpEntries?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceArpEntries request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceArpEntries JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDeviceHosts returns all hosts of a specific device discovered by running an IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4
// reconciliation policy.
//
// Returns all hosts of a specific device. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDeviceHosts(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDeviceHosts?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceHosts request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceHosts JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDeviceInterfaces returns all interfaces of a specific device discovered by running an IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4
// reconciliation policy.
//
// Returns all interfaces of a specific device. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDeviceInterfaces(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDeviceInterfaces?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceInterfaces request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceInterfaces JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDeviceMacAddressEntries returns all MAC address entries of a specific device discovered by running an
// IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4
// reconciliation policy.
//
// Returns all MAC address entries of a specific device. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDeviceMacAddressEntries(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDeviceMacAddressEntries?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceMacAddressEntries request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceMacAddressEntries JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDeviceNetworks returns all networks of a specific device discovered by running an IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4
// reconciliation policy.
//
// Returns all networks of a specific device. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDeviceNetworks(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDeviceNetworks?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceNetworks request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceNetworks JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDeviceVlans returns all VLANs of a specific device discovered by running an IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4
// reconciliation policy.
//
// Returns all Vlans of a specific device. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDeviceVlans(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDeviceVlans?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceVlans request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDeviceVlans JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDevice returns the object ID of the discovered device by running an IPv4 reconciliation policy.
//
// Parameter `deviceid` is the object ID of the discovered device. Parameter `policyid` is the object ID for the IPv4
// reconciliation policy.
//
// Returns the object ID of the discovered device. Return type is APIEntity.
func (b *Bluecat) GetDiscoveredDevice(deviceid, policyid int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDevice?deviceId=%d&policyId=%d",
		b.Server, b.URI, deviceid, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDevice request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDevice JSON parse", err)
	}

	return results, nil
}

// GetDiscoveredDevices returns a list of discovered Layer 2 or Layer 3 devices by running an IPv4 reconciliation policy specified.
//
// Parameter `policyid` is the object ID for the IPv4 reconciliation policy.
//
// Returns an array of discovered Layer 2 or Layer 3 devices. Return type is an array of APIEntity.
func (b *Bluecat) GetDiscoveredDevices(deviceid, policyid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getDiscoveredDevices?policyId=%d",
		b.Server, b.URI, policyid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDevices request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetDiscoveredDevices JSON parse", err)
	}

	return results, nil
}

// GetEntitiesByNameUsingOptions returns an array of entities that match the specified name and object type. Searching
// behavior can be changed by using the options.
//
// Parameter `name` is the name of the entity. Parameter `options` is a tring containing options. Currently the only available
// option is ObjectProperties.ignoreCase. By default, the value is set to false. Setting this option to true will ignore
// the case-sensitivity used while searching entities by name.
//
// ObjectProperties.ignoreCase = [true | false]
//
// Parameter `parentid` is the object ID of the parent object of the entities to be returned. Parameter `objecttype` is the
// type of object to be returned. This value must be one of the object types listed in Object types. Parameter `count`
// is the maximum number of objects to return. The default value is 10. This value cannot be null or empty. Parameter `start`
// indicates where in the list of returned objects to start returning objects. The list begins at an index of 0. This
// value cannot be null or empty.
//
// Returns an array of type APIEntity. The array is empty if there are no matching entities.
func (b *Bluecat) GetEntitiesByNameUsingOptions(name, options string, parentid int64, objecttype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getEntitiesByNameUsingOptions?name=%s&options=%s&parentId=%d&type=%s&count=%d&start=%d",
		b.Server, b.URI, name, options, parentid, objecttype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetEntitiesByNameUsingOptions request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetEntitiesByNameUsingOptions JSON parse", err)
	}

	return results, nil
}

// GetHostRecordsByHint returns an array of objects with host record type.
//
// Parameter `options` is a string containing options. The supported options are hint and retrieveFields. Separate multiple
// options with a pipe character. For example: hint=^abc|retrieveFields=false
//
// If the hint option is not specified in the string, searching criteria will be based on the same as zone host record.
// The following wildcards are supported in the hint option.
//
// ^ — matches the beginning of a string. For example: ^ex matches ex ample but not t ex t.
//
// $ — matches the end of a string. For example: ple$ matches exam ple but not ple ase.
//
// ^ $ — matches the exact characters between the two wildcards. For example: ^example$ only matches example.
//
// ? — matches any one character. For example: ex?t matches exit.
//
// * — matches one or more characters within a string. For example: ex*t matches exit and ex cellen t.
//
// The default value for the retrieveFields option is set to false. If the option is set to true, user-defined field will
// be returned. If the options string does not contain retrieveFields, user-defined field will not be returned.
//
// Parameter `count` indicates the maximum of child objects that this method will return. The value must be less
// than or equal to 10. Parameter `start` indicates where in the list of objects to start returning objects.
// The list begins at an index of 0.
//
// Returns an array of host record APIEntity objects.
func (b *Bluecat) GetHostRecordsByHint(options string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getHostRecordsByHint?options=%s&count=%d&start=%d",
		b.Server, b.URI, options, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetHostRecordsByHint request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetHostRecordsByHint JSON parse", err)
	}

	return results, nil
}

// GetIP4Address returns the details for the requested IPv4 address object.
//
// Parameter `address` is the IPv4 address. Parameter `containerid` is the object ID for the configuration, block, network,
// or DHCP range in which this address is located.
//
// Returns the requested IPv4 Address object from the database. Return type is APIEntity.
func (b *Bluecat) GetIP4Address(address string, containerid int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getIP4Address?address=%s&containerId=%d",
		b.Server, b.URI, address, containerid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetIP4Address request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetIP4Address JSON parse", err)
	}

	return results, nil
}

// GetIP4NetworksByHint returns an array of IPv4 networks found under a given container object. You can filter the networks
// can using the ObjectProperties.hint, ObjectProperties.accessRight, and ObjectProperties.overrideType options.
//
// Parameter `containerid` is the object ID for the container object. It can be the object ID of any object in the parent
// object hierarchy. The highest parent object is the configuration level.
//
// Parameter `options` is string containing options. The option names available in the ObjectProperties are
// ObjectProperties.hint, ObjectProperties.accessRight, and ObjectProperties.overrideType. Separate multiple options
// with a pipe character. For example: hint=ab|overrideType=HostRecord|accessRight=ADD
//
// The values for the ObjectProperties.hint option can be the prefix of the IP address for a network or the name of a network.
// The following example will match networks that have the network ID starting with 192.168. For example, 192.168.0.0/24 or 192.168.1.0/24.
//
// Example 1: String options = ObjectProperties.hint + “=198.168|”
//
// The following example will match networks that have a name starting with “abc”. For example, “abc”, “abc123” or “abcdef”.
//
// Example 2: String options = ObjectProperties.hint + “=abc|”
//
// Matching networks to a network ID Example 1 will take precedence over matching networks to a name Example 2.
//
// The values for the ObjectProperties.accessRight and ObjectProperties.overrideType options must be one of the constants
// listed in Access right values and Object types. For example:
//
// String options = ObjectProperties.accessRight + "=" + AccessRightValues.AddAccess + "|"+ ObjectProperties.overrideType + "=" + ObjectTypes.HostRecord;
//
// Parameter `count` indicates the maximum number of child objects that this method will return. The maximum number
// of child objects that can be returned is 10. Parameter `start` indicates where in the list of objects to start returning
// objects. The list begins at an index of 0.
//
// Returns an array, type APIEntity, of IPv4 networks based on the input argument without their properties fields populated, or returns
// an empty array if containerId is invalid. If no access right option is specified, the View access level will be used by default.
func (b *Bluecat) GetIP4NetworksByHint(containerid int64, options string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getIP4NetworksByHint?containerId=%d&options=%s&count=%d&start=%d",
		b.Server, b.URI, containerid, options, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetIP4NetworksByHint request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetIP4NetworksByHint JSON parse", err)
	}

	return results, nil
}

// GetIP6Address returns an APIEntity for the specified IPv6 address.
//
// Parameter `address` is the IPv6 address. Parameter `containerid` is the object ID of the container in which the IPv6
// address is located. The container can be a configuration, an IPv6 block, or an IPv6 network.
//
// Returns an APIEntity for the specified IPv6 address. The APIEntity is empty of the IPv6 address does not exist.
func (b *Bluecat) GetIP6Address(address string, containerid int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getIP6Address?address=%s&containerId=%d",
		b.Server, b.URI, address, containerid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetIP6Address request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetIP6Address JSON parse", err)
	}

	return results, nil
}

// GetIP6ObjectsByHint returns an array of IPv6 objects found under a given container object. The networks can be filtered
// by using ObjectProperties.hint and ObjectProperties.accessRight options. Currently, it only supports IPv6 networks.
//
// Parameter `containerid` is the object ID for the container object. It can be the object ID of any object in the parent
// object hierarchy. The highest parent object is the configuration level. Parameter `objecttype` is the type of object
// containing the IPv6 network. Currently, it only supports ObjectTypes.IP6Network.
//
// Parameter `options` is a string containing options. The Option names available in ObjectProperties are ObjectProperties.hint
// and ObjectProperties.accessRight. Multiple options can be separated by a pipe | character. For example:
//
// hint=ab|
//
// accessRight=ADD
//
// The values for the ObjectProperties.hint option can be the prefix of the IP address for a network or the name of a network.
//
// The following example will match networks that have the network ID starting with 2000::. For example, 2000::/64.
//
// Example 1: String options = ObjectProperties.hint + “=2000::|”
//
// The following example will match networks that have a name starting with “abc”. For example, “abc”, “abc123” or “abcdef”.
//
// Example 2: String options = ObjectProperties.hint + “=abc|”
//
// Matching networks to a network ID Example 1 will take precedence over matching networks to a name Example 2.
//
// The values for the ObjectProperties.accessRight option must be one of the constants listed in Access right values and
// Object types. For example:
//
// String options = ObjectProperties.accessRight + "=" + AccessRightValues.AddAccess;
//
// Parameter `count` indicates the maximum number of child objects that this method will return. Parameter `start` indicates
// where in the list of objects to start returning objects. The list begins at an index of 0.
//
// Returns an array, type APIEntity, of IPv6 objects based on the input argument without their properties fields populated, or
// returns an empty array if containerId is invalid. If no access right option is specified, the View access level will be used by default.
func (b *Bluecat) GetIP6ObjectsByHint(containerid int64, objecttype, options string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getIP6ObjectsByHint?containerId=%d&objectType=%s&options=%s&count=%d&start=%d",
		b.Server, b.URI, containerid, objecttype, options, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetIP6ObjectsByHint request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetIP6ObjectsByHint JSON parse", err)
	}

	return results, nil
}

// GetIPRangeByIP returns the DHCP range containing the specified IPv4 or IPv6 address. Use this method to find the configuration,
// IPv4 or IPv6 block, network, or DHCP range containing a specified address. You can specify the type of object to be returned,
// or you can leave the type of object empty to find the most direct container for the object.
//
// Parameter `address` is an IPv4 or IPv6 address. Parameter `containerid` is the object ID of the container in which the
// IPv4 or IPv6 address is located. This can be a configuration, IPv4 or IPv6 block, network, or DHCP range. When you do
// not know the block, network, or range in which the address is located, specify the configuration.
//
// Parameter `objecttype` is the type of object containing the IP address. Specify ObjectTypes.IP4Block or ObjectTypes.IP6Block,
// ObjectTypes.IP4Network or ObjectTypes.IP6Network, or ObjectTypes.DHCP4Range or ObjectTypes.DHCP6Range to find the block,
// network, or range containing the IPv4 or IPv6 address. Specify an empty string "" to return the most direct container
// for the IPv4 or IPv6 address.
//
// Returns an APIEntity for the object containing the specified address.
func (b *Bluecat) GetIPRangeByIP(address string, containerid int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getIPRangeByIP?address=%s&containerId=%d&type=%s",
		b.Server, b.URI, address, containerid, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetIPRangeByIP request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetIPRangeByIP JSON parse", err)
	}

	return results, nil
}

// GetKSK returns a string containing all active Key Signing Keys (KSK) for a given entityId value in a specified output
// format with its start time and expire time, divided by a delimiter (|). The list of returned KSKs is sorted in
// descending order by expiry date.
//
// Parameter `entityid` is the object ID of the entity associated with the KSK. The only supported entity types are zone,
// IPv4 block, and IPv4 network. Parameter `format` is the output format of the KSK of an entity. The value must be one
// of the constants listed in DNSSEC key format.
//
// Returns a string containing up to two active KSK(s) of an entity.
func (b *Bluecat) GetKSK(entityid int64, format string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getKSK?entityId=%d&format=%s",
		b.Server, b.URI, entityid, format)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetKSK request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetLinkedEntities returns an array of entities containing the entities linked to a specified entity. The array is
// empty if there are no linked entities.
//
// Parameter `entityid` is the object ID of the entity for which to return linked entities. Parameter `linkedtype` is the
// type of linked entities which need to be returned. This value must be one of the types listed in Object types.
//
// While specifying a resource record as the entityId, if you want to find all the records CNAME, MX, or SRV records having
// links to this record, you can use RecordWithLink for the type parameter.
//
// When specifying a MAC address as the entityId, this method returns the IPv4 address associated with the MAC address.
// When appropriate, leaseTimeand expiryTimeinformation also appears in the returned properties string.
//
// Parameter `count` is the maximum number of objects to return. Parameter `start` indicates where in the list of returned
// objects to start returning objects. The list begins at an index of 0. This value cannot be null or empty.
//
// Returns a string containing up to two active KSK(s) of an entity. Return type is an array of APIEntity.
func (b *Bluecat) GetLinkedEntities(entityid int64, linkedtype string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getLinkedEntities?entityId=%d&type=%s&count=%d&start=%d",
		b.Server, b.URI, entityid, linkedtype, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetLinkedEntities request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetLinkedEntities JSON parse", err)
	}

	return results, nil
}

// GetLocationByCode returns the location object with the specified hierarchical location code.
//
// Parameter `code` is the hierarchical location code consists of a set of 1 to 3 alpha-numeric strings separated by a
// space. The first two characters indicate a country, followed by next three characters which indicate a city in UN/LOCODE.
// New custom locations created under a UN/LOCODE city are appended to the end of the hierarchy.
// For example, CA TOR OF1 indicates:
//
// CA - Canada
//
// TOR - Toronto
//
// OF1 - Office 1
//
// The code is case-sensitive. It must be all UPPER CASE letters. The county code and child location code should be
// alphanumeric strings.
//
// Returns the APIEntity that matches the specified hierarchical location code. If no entity is found, returns an empty APIEntity.
func (b *Bluecat) GetLocationByCode(code string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getLocationByCode?code=%s",
		b.Server, b.URI, code)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetLocationByCode request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetLocationByCode JSON parse", err)
	}

	return results, nil
}

// GetMACAddress returns an APIEntity for a MAC address.
//
// Parameter `configid` is the object ID of the configuration in which the MAC address is located. Parameter `macaddress`
// is the MAC address in the format nnnnnnnnnnnn, nn-nn-nn-nn-nn-nn or nn:nn:nn:nn:nn:nn, where nn is a hexadecimal value.
//
// Returns an APIEntity for the MAC address. Returns an empty APIEntity if the MAC address does not exist.
func (b *Bluecat) GetMACAddress(configid int64, macaddress string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getMACAddress?configurationId=%d&macAddress=%s",
		b.Server, b.URI, configid, macaddress)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetMACAddress request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetMACAddress JSON parse", err)
	}

	return results, nil
}

// GetMaxAllowedRange finds the maximum possible address range to which the existing IPv4 DHCP range can be extended.
// This method only supports the IPv4 DHCP range.
//
// Parameter `rangeid` is the object ID of the IPv4 DHCP range.
//
// Returns the possible start address and end address for the specified IPv4 DHCP range object in the form of array of length 2.
func (b *Bluecat) GetMaxAllowedRange(rangeid int64) (string, error) {
	req := fmt.Sprintf("https://%s%s/getMaxAllowedRange?rangeId=%d",
		b.Server, b.URI, rangeid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetMaxAllowedRange request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetNetworkLinkedProperties returns an array of IP addresses with linked records and the IP addresses that are
// assigned as DHCP Reserved, Static, or Gateway.
//
// Returns an array of IP address APIEntity objects with their linked host records and the IP addresses that are assigned
// as DHCP Reserved, Static or Gateway. The output has the following format: hostId : hostName : zoneId : zoneName : viewId : viewName : hasAlias;.
func (b *Bluecat) GetNetworkLinkedProperties(networkid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getNetworkLinkedProperties?networkId=%d",
		b.Server, b.URI, networkid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetNetworkLinkedProperties request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetNetworkLinkedProperties JSON parse", err)
	}

	return results, nil
}

// GetNextAvailableIP4Address returns the IPv4 address for the next available (unallocated) address within a
// configuration, block, or network.
//
// Parameter `parentid` is the object ID for configuration, block, or network in which to look for the next available address.
//
// Returns the next available IPv4 address in an existing network as a string.
func (b *Bluecat) GetNextAvailableIP4Address(parentid int64) (string, error) {
	req := fmt.Sprintf("https://%s%s/getNextAvailableIP4Address?parentId=%d",
		b.Server, b.URI, parentid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetNextAvailableIP4Address request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetNextAvailableIP4Network returns the object ID for the next available (unused) network within a configuration or block.
//
// Parameter `autocreate` is a boolean value indicates whether the next available network should be created if it does not exist.
// Parameter `islargerallowed` is a boolean value indicates whether to return larger networks than those specified
// with the size parameter. Parameter `parentid` is the object ID of the network’s parent object.
//
// Parameter `size` is the size of the network, expressed as a power of 2. The size represents the number of hosts on the
// network. For example, if you are to find or create a /24 network, the size would be 256.
//
// Returns the object ID for the existing next available IPv4 network or, if the next available network did not exist and
// autoCreate was set to true, the newly created IPv4 network.
func (b *Bluecat) GetNextAvailableIP4Network(autocreate, islargerallowed bool, parentid, size int64) (string, error) {
	req := fmt.Sprintf("https://%s%s/getNextAvailableIP4Network?autoCreate=%t&isLargerAllowed=%t&parentId=%d&size=%d",
		b.Server, b.URI, autocreate, islargerallowed, parentid, size)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetNextAvailableIP4Network request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetNextAvailableIPRange returns the object ID for the next available (unused) block or network within a configuration or block.
//
// Parameter `parentid` is the object ID of the parent object under which the next available range resides
// configuration or block.
//
// Parameter `properties` is a string containing the following properties and values:
//
// reuseExisting — True or False. This Boolean value indicates whether to search existing empty networks to find the
// available IP range of specified size.
//
// isLargerAllowed — True or False. This Boolean value indicates whether to return larger networks than those
// specified with the sizeparameter.
//
// autoCreate — True or False. This Boolean value indicates whether the next available IP range should be created in the
// parent object if it does not exist.
//
// traversalMethod — This parameter identifies the appropriate search algorithm to find the suitable object.
// The possible values are:
//
// TraversalMethodology.NO_TRAVERSAL NO_TRAVERSAL — will attempt to find the next range directly under the specified
// parent object. It will not search through to the lower level objects.
//
// TraversalMethodology.DEPTH_FIRST DEPTH_FIRST — will attempt to find the next range under the specified object by iteratively
// through its children one by one. After exploring the object recursively for its child ranges, it will move to the next child object.
//
// TraversalMethodology.BREADTH_FIRST BREADTH_FIRST — will attempt to find the next range under the specified object by
// iterative levels. It will first find the range immediately below the specified parent object. If not found, then it
// will attempt to find the range under all the first child objects.
//
// Parameter `size` is the size of the range, expressed as a power of 2. Parameter `objecttype` is the type of the range
// object to be fetched. Currently IPv4 block and network are supported.
//
// Returns the object ID, type APIEntity, for the existing next available IPv4 range or, if the next available IP range does not exist and
// autoCreate was set to true, the newly created IPv4 range.
func (b *Bluecat) GetNextAvailableIPRange(parentid int64, properties string, size int64, objecttype string) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getNextAvailableIPRange?parentId=%d&properties=%s&size=%d&type=%s",
		b.Server, b.URI, parentid, properties, size, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetNextAvailableIPRange request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetNextAvailableIPRange JSON parse", err)
	}

	return results, nil
}

// GetNextAvailableIPRanges returns the object ID for the next available (unused) block or network within a configuration
// or block.
//
// Parameter `parentid` is the object ID of the parent object under which the next available range resides configuration
// or block.
//
// Parameter `properties` is a string containing the following properties and values:
//
// reuseExisting — True or False. This Boolean value indicates whether to search existing empty networks to find the
// available IP range of specified size.
//
// isLargerAllowed — True or False. This Boolean value indicates whether to return larger networks than those
// specified with the sizeparameter.
//
// autoCreate — True or False. This Boolean value indicates whether the next available IP range should be created in the
// parent object if it does not exist.
//
// traversalMethod — This parameter identifies the appropriate search algorithm to find the suitable object.
// The possible values are:
//
// TraversalMethodology.DEPTH_FIRST DEPTH_FIRST — will attempt to find the next range under the specified object by iteratively
// through its children one by one. After exploring the object recursively for its child ranges, it will move to the next child object.
//
// If a value is not provided for the following parameters: reuseExisting and autoCreate, the default values will be as
// follows: reuseExisting = false and autoCreate = true.
//
// Parameter `size` is the size of the range, expressed as a power of 2. Parameter `objecttype` is the type of the range
// object to be fetched. Currently IPv4 block and network are supported. Parameter `count`  is the number of networks to
// be found. If the number of networks count is greater than 1: isLargerAllowed and traversalMethod properties will not
// be applicable. The DEPTH_FIRST methodology will be used to search objects.
//
// Returns consecutive matching IPv4 range object IDs. If the next available ranges do not exist and you have set the
// autoCreate property to true, new IPv4 ranges will be created and their object IDs will be returned. Return type is
// an array of APIEntity.
func (b *Bluecat) GetNextAvailableIPRanges(parentid int64, properties string, size int64, objecttype string, count int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getNextAvailableIPRanges?parentId=%d&properties=%s&size=%d&type=%s&count=%d",
		b.Server, b.URI, parentid, properties, size, objecttype, count)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetNextAvailableIPRanges request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetNextAvailableIPRanges JSON parse", err)
	}

	return results, nil
}

// GetNextIP4Address returns the next available IP addresses in octet notation under specified circumstances.
//
// Parameter `parentid` is the network or configuration Id. Parameter `properties` is a string that contains three properties;
// skip, offset and excludeDHCPRange. The values for skip and offset must be IPv4 addresses and must appear in dotted octet notation.
//
// skip - This is optional. It is used to specify the IP address ranges or IP addresses to skip, separated by comma.
// A hyphen (-), not a dash is used to separate the start and end addresses. (Do not use the skip property if the parentId
// is a configuration Id. If you do, an error message appears, ‘Skip is not allowed for configuration level’.)
//
// offset - This is optional. This is to specify from which address to start to assign IPv4 Address.
//
// excludeDHCPRange - This specifies whether IP addresses in DHCP ranges should be excluded from assignment. The value
// is either true or false, default value is false. For example:
//
// skip=10.10.10.128-10.10.11.200,10.10.11.210|
//
// offset=10.10.10.100|excludeDHCPRange=true|
//
// Returns the IPv4 address in octet notation. Return type is a string.
func (b *Bluecat) GetNextIP4Address(parentid int64, properties string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getNextIP4Address?parentId=%d&properties=%s",
		b.Server, b.URI, parentid, properties)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetNextIP4Address request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetParent returns the parent entity of a given entity.
//
// Parameter `entityid` is the entity ID of the parent object.
//
// Returns the APIEntity for the parent entity with its properties fields populated.
func (b *Bluecat) GetParent(entityid int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getParent?entityId=%d",
		b.Server, b.URI, entityid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetParent request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetParent JSON parse", err)
	}

	return results, nil
}

// GetProbeData returns the JSON response from the properties field of the APIData object.
//
// Parameter `definedprobe` are pre-defined SQL queries that will be triggered to collect data. The available values
// are LEASE_COUNT_PER_DATE and NETWORK_BLOOM.
//
// Returns the JSON response from the properties field of the APIData object. Return type is APIData.
func (b *Bluecat) GetProbeData(definedprobe string) (APIData, error) {
	var results APIData
	req := fmt.Sprintf("https://%s%s/getProbeData?definedProbe=%s",
		b.Server, b.URI, definedprobe)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetProbeData request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetProbeData JSON parse", err)
	}

	return results, nil
}

// GetProbeStatus returns the status of the triggered data collection process.
//
// Parameter `definedprobe` are pre-defined SQL queries that will be triggered to collect data. The available values
// are LEASE_COUNT_PER_DATE and NETWORK_BLOOM.
//
// Returns a pre-defined value from 0 to 3, depending on the status of the data collection process. Return type is a string.
func (b *Bluecat) GetProbeStatus(definedprobe string) (string, error) {
	req := fmt.Sprintf("https://%s%s/getProbeStatus?definedProbe=%s",
		b.Server, b.URI, definedprobe)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetProbeStatus request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetReplicationInfo retrieves information regarding the status of replication in Address Manager through the API.
//
// Returns a JSON string containing the hostname, status of replication, latency, the IP address of the Primary and
// standby servers, and cluster information.
func (b *Bluecat) GetReplicationInfo() (string, error) {
	req := fmt.Sprintf("https://%s%s/getReplicationInfo",
		b.Server, b.URI)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetReplicationInfo request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetServerDeploymentRoles returns a list of all deployment roles associated with the server.
//
// Parameter `serverid` is the object ID of the server with which deployment roles are associated.
//
// Returns a list of all deployment roles associated with the server. Return type is an array of APIDeploymentRole.
func (b *Bluecat) GetServerDeploymentRoles(serverid int64) ([]APIDeploymentRole, error) {
	var results []APIDeploymentRole
	req := fmt.Sprintf("https://%s%s/getServerDeploymentRoles?serverId=%d",
		b.Server, b.URI, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetServerDeploymentRoles request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetServerDeploymentRoles JSON parse", err)
	}

	return results, nil
}

// GetServerDeploymentStatus returns the deployment status of the server.
//
// For the parameter `properties` the valid value is empty. Parameter `serverid` is the object ID of the server whose
// deployment status needs to be checked.
//
// Returns status code for deployment of a particular server. Return type is a string.
func (b *Bluecat) GetServerDeploymentStatus(properties string, serverid int64) (string, error) {
	req := fmt.Sprintf("https://%s%s/getServerDeploymentStatus?properties=%s&serverId=%d",
		b.Server, b.URI, properties, serverid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetServerDeploymentStatus request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetServerForRole returns a list of all servers associated with the specified deployment role.
//
// Parameter `roleid` is the object ID for the deployment role whose servers are to be returned.
//
// Returns an APIEntity object representing the servers associated with the specified deployment role.
func (b *Bluecat) GetServerForRole(roleid int64) (APIEntity, error) {
	var results APIEntity
	req := fmt.Sprintf("https://%s%s/getServerForRole?roleId=%d",
		b.Server, b.URI, roleid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetServerForRole request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetServerForRole JSON parse", err)
	}

	return results, nil
}

// GetSharedNetworks returns multiple IPv4 networks linked to the given shared network tag.
//
// Parameter `tagid` is the object ID of the tag that is linked with shared IPv4 networks. If tagId is not valid,
// an error will be returned.
//
// Returns an array of type APIEntity, of all the IPv4 networks linked to the given shared network tag. If no networks
// are found, returns an empty array.
func (b *Bluecat) GetSharedNetworks(tagid int64) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getSharedNetworks?tagId=%d",
		b.Server, b.URI, tagid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetSharedNetworks request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetSharedNetworks JSON parse", err)
	}

	return results, nil
}

// GetSystemInfo gets Address Manager system information.
//
// Returns Address Manager system information. Return type is a string.
func (b *Bluecat) GetSystemInfo() (string, error) {
	req := fmt.Sprintf("https://%s%s/getSystemInfo",
		b.Server, b.URI)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetSystemInfo request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetTemplateTaskStatus gets the IPv4 template task status when the template is applied.
//
// Parameter `taskid` is the task ID of the IPv4 network template.
//
// Returns a JSON string that contains the template status.
func (b *Bluecat) GetTemplateTaskStatus(taskid int64) (string, error) {
	req := fmt.Sprintf("https://%s%s/getTemplateTaskStatus?taskId=%d",
		b.Server, b.URI, taskid)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - GetTemplateTaskStatus request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// GetUserDefinedFields returns the user-defined fields information.
//
// Parameter `requiredfieldsonly` specifies whether all user-defined fields of the object type will be returned or not.
// If set to true, only required fields will be returned. Parameter `objecttype` is the type of object the user-defined
// field belongs to. This must be one of the constants listed in Object types.
//
// Returns the user-defined fields information. Return type is an array of APIUserDefinedField.
func (b *Bluecat) GetUserDefinedFields(requiredfieldsonly bool, objecttype string) ([]APIUserDefinedField, error) {
	var results []APIUserDefinedField
	req := fmt.Sprintf("https://%s%s/getUserDefinedFields?requiredFieldsOnly=%t&type=%s",
		b.Server, b.URI, requiredfieldsonly, objecttype)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetUserDefinedFields request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetUserDefinedFields JSON parse", err)
	}

	return results, nil
}

// GetZonesByHint returns an array of accessible zones of child objects for a given containerId value.
//
// Parameter `containerid` is the object ID for the container object. It can be the object ID of any object in the parent
// object hierarchy. The highest parent object can be the configuration level.
//
// Parameter `options` is a string containing options. The Option names available in the ObjectProperties are
// ObjectProperties.hint, ObjectProperties.accessRight, and ObjectProperties.overrideType. Multiple options can be
// separated by a pipe character. For example:
//
// hint=ab|overrideType=HostRecord|accessRight=ADD
//
// The values for ObjectProperties.hint option can be the prefix of a zone name. For example:
//
// String options = ObjectProperties.hint + "=abc|"
//
// The values for the ObjectProperties.accessRight and ObjectProperties.overrideType options must be one of the constants
// listed in Access right values and Object types. For example:
//
// String options = ObjectProperties.accessRight + "=" + AccessRightValues.AddAccess + "|"+ ObjectProperties.overrideType + "=" + ObjectTypes.HostRecord;
//
// Parameter `count` indicates the maximum number of child objects that this method will return. The maximum number of
// child objects cannot exceed more than 10. Parameter `start` indicates where in the list of objects to start returning
// objects. The list begins at an index of 0.
//
// Returns an array, of type APIEntity, of zones based on the input argument without their properties fields populated, or returns
// an empty array if containerId is invalid. If no access right option is specified, the View access level will be used by default.
func (b *Bluecat) GetZonesByHint(containerid int64, options string, count, start int32) ([]APIEntity, error) {
	var results []APIEntity
	req := fmt.Sprintf("https://%s%s/getZonesByHint?containerId=%d&options=%s&count=%d&start=%d",
		b.Server, b.URI, containerid, options, count, start)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return results, fmt.Errorf("%s - GetZonesByHint request", err)
	}

	if err := json.Unmarshal([]byte(resp.String()), &results); err != nil {
		return results, fmt.Errorf("%s - GetZonesByHint JSON parse", err)
	}

	return results, nil
}

// IsAddressAllocated queries a MAC address to determine if the address has been allocated to an IP address.
//
// Parameter `configid` is the object ID of the parent configuration in which the MAC address resides. Parameter `ipaddress`
// is the IPv4 DHCP allocated address to be checked against the MAC address. Parameter `macaddress` is the
// MAC address in the format nnnnnnnnnnnn, nn-nn-nn-nn-nn-nn or nn:nn:nn:nn:nn:nn, where nn is a hexadecimal value.
//
// Returns a Boolean value indicating whether the address is allocated.
func (b *Bluecat) IsAddressAllocated(configid int64, ipaddress, macaddress string) (string, error) {
	req := fmt.Sprintf("https://%s%s/isAddressAllocated?configurationId=%d&ipAddress=%s&macAddress=%s",
		b.Server, b.URI, configid, ipaddress, macaddress)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - IsAddressAllocated request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// IsMigrationRunning returns true or false to indicate if the migration service is running. Specify a filename to
// determine if the specified file is migrating. Specify an empty string ("") to determine if any migration files are
// migrating or queued for migration.
//
// Parameter `filename` is the filename of the XML file in the data/migration/incoming directory. Do not include a path
// in the filename. This value can be empty.
//
// Returns a Boolean value indicating if the specified file is currently migrating. When an empty string is specified
// for the filename, returns a true if there are any migration files queued for migration or currently migrating.
func (b *Bluecat) IsMigrationRunning(filename string) (string, error) {
	req := fmt.Sprintf("https://%s%s/isMigrationRunning?filename=%s",
		b.Server, b.URI, filename)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return "", fmt.Errorf("%s - IsMigrationRunning request", err)
	}

	formatted := strings.TrimLeft(strings.TrimRight(resp.String(), "\""), "\"")
	return formatted, nil
}

// LinkEntities establishes a link between two specified Address Manager entities.
//
// Parameter `entity1id` is the object ID of the first entity in the pair of linked entities. Parameter `entity2id` is
// the object ID of the second entity in the pair of linked entities. Parameter `properties` Adds object properties,
// including user-defined fields.
func (b *Bluecat) LinkEntities(entity1id, entity2id int64, properties string) error {
	req := fmt.Sprintf("https://%s%s/linkEntities?entity1Id=%d&entity2Id=%d&properties=%s",
		b.Server, b.URI, entity1id, entity2id, properties)
	_, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", fmt.Sprintf("%s", b.AuthToken)).
		Get(req)

	if err != nil {
		return fmt.Errorf("%s - LinkEntities request", err)
	}

	return nil
}
