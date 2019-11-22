package messagelayer

import (
	"fmt"
	"strings"

	"k8s.io/klog"

	"github.com/kubeedge/beehive/pkg/core/model"
	"github.com/kubeedge/kubeedge/cloud/pkg/edgecontroller/config"
	"github.com/kubeedge/kubeedge/common/constants"
)

// BuildResource return a string as "beehive/pkg/core/model".Message.Router.Resource
func BuildResource(nodeID, namespace, resourceType, resourceID string) (resource string, err error) {
	if namespace == "" || resourceType == "" {
		if !config.EdgeSiteEnabled && nodeID == "" {
			err = fmt.Errorf("required parameter are not set (node id, namespace or resource type)")
		} else {
			err = fmt.Errorf("required parameter are not set (namespace or resource type)")
		}
		return
	}

	resource = fmt.Sprintf("%s%s%s%s%s%s%s", constants.ResNode, constants.ResourceSep, nodeID, constants.ResourceSep, namespace, constants.ResourceSep, resourceType)
	if config.EdgeSiteEnabled {
		resource = fmt.Sprintf("%s%s%s", namespace, constants.ResourceSep, resourceType)
	}
	if resourceID != "" {
		resource += fmt.Sprintf("%s%s", constants.ResourceSep, resourceID)
	}
	return
}

// GetNodeID from "beehive/pkg/core/model".Message.Router.Resource
func GetNodeID(msg model.Message) (string, error) {
	sli := strings.Split(msg.GetResource(), constants.ResourceSep)
	if len(sli) <= constants.ResourceNodeIDIndex {
		return "", fmt.Errorf("node id not found")
	}
	return sli[constants.ResourceNodeIDIndex], nil
}

// GetNamespace from "beehive/pkg/core/model".Model.Router.Resource
func GetNamespace(msg model.Message) (string, error) {
	sli := strings.Split(msg.GetResource(), constants.ResourceSep)
	length := constants.ResourceNamespaceIndex
	if config.EdgeSiteEnabled {
		length = constants.EdgeSiteResourceNamespaceIndex
	}
	if len(sli) <= length {
		return "", fmt.Errorf("namespace not found")
	}
	var res string
	var index uint8
	if config.EdgeSiteEnabled {
		res = sli[constants.EdgeSiteResourceNamespaceIndex]
		index = constants.EdgeSiteResourceNamespaceIndex
	} else {
		res = sli[constants.ResourceNamespaceIndex]
		index = constants.ResourceNamespaceIndex
	}
	klog.V(4).Infof("The namespace is %s, %d", res, index)
	return res, nil
}

// GetResourceType from "beehive/pkg/core/model".Model.Router.Resource
func GetResourceType(msg model.Message) (string, error) {
	sli := strings.Split(msg.GetResource(), constants.ResourceSep)
	length := constants.ResourceResourceTypeIndex
	if config.EdgeSiteEnabled {
		length = constants.EdgeSiteResourceResourceTypeIndex
	}
	if len(sli) <= length {
		return "", fmt.Errorf("resource type not found")
	}

	var res string
	var index uint8
	if config.EdgeSiteEnabled {
		res = sli[constants.EdgeSiteResourceResourceTypeIndex]
		index = constants.EdgeSiteResourceResourceTypeIndex
	} else {
		res = sli[constants.ResourceResourceTypeIndex]
		index = constants.ResourceResourceTypeIndex
	}
	klog.Infof("The resource type is %s, %d", res, index)

	return res, nil
}

// GetResourceName from "beehive/pkg/core/model".Model.Router.Resource
func GetResourceName(msg model.Message) (string, error) {
	sli := strings.Split(msg.GetResource(), constants.ResourceSep)
	length := constants.ResourceResourceNameIndex
	if config.EdgeSiteEnabled {
		length = constants.EdgeSiteResourceResourceNameIndex
	}
	if len(sli) <= length {
		return "", fmt.Errorf("resource name not found")
	}

	var res string
	var index uint8
	if config.EdgeSiteEnabled {
		res = sli[constants.EdgeSiteResourceResourceNameIndex]
		index = constants.EdgeSiteResourceResourceNameIndex
	} else {
		res = sli[constants.ResourceResourceNameIndex]
		index = constants.ResourceResourceNameIndex
	}
	klog.Infof("The resource name is %s, %d", res, index)
	return res, nil
}
