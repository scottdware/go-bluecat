package bluecat

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

// GetEntities returns multiple entities for the specified parent ID.
//
// Parameter `parentid` is the object ID of the parent object of the entities.
// Parameter `objecttype` is the type of object to be returned. This value must be one of the object types constants.
// Parameter `count` indicates the maximum number of child objects to return.
// Parameter `start` indicates where in the list of child objects to start returning entities. The list begins at an index of 0.
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

// CustomSearch searches for an array of entities by specifying object properties.
//
// Parameter `objecttype` must be one of the following object types: IP4Block, IP4Network, IP4Addr, GenericRecord, HostRecord,
// Any other objects with user-defined fields.
//
// Parameter `filters` is the list of properties on which the search will be based. The valid format is Field name=value.
// Parameter `count` is the maximum number of objects to return. The value must be a positive value between 1 and 1000.
// This value cannot be null or empty.
//
// Parameter `start` indicates where in the list of returned objects to start returning objects.
// The value must be a non-negative value and cannot be null or empty.
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
// Parameter `keyword` is the search keyword string. This value cannot be null or empty.
// Parameter `category` is the entity category to be searched.
// Parameter `count` is the maximum number of objects to return. The default value is 10. This value cannot be null or empty.
// Parameter `start` indicates where in the list of returned objects to start returning objects.
// The list begins at an index of 0. This value cannot be null or empty.
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
// Parameter `keyword` is the search keyword string. This value cannot be null or empty.
// Parameter `objecttypes` is the object types for which to search, specified in the format: "type1[,type2...]"
// Parameter `count` is the maximum number of objects to return. The default value is 10. This value cannot be null or empty.
// Parameter `start` indicates where in the list of returned objects to start returning objects.
// The list begins at an index of 0. This value cannot be null or empty.
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
// policy item, set the value of this parameter to 0.
//
// Parameter `itemname` is the Fully Qualified Domain Name FQDN of the response policy item. The exact FQDN of the
// response policy item must be used when conducting a search.
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
// Parameter `count ` is the maximum number of access right child objects to return.
// Parameter `start` indicates where in the list of child access right objects to start returning objects.
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
