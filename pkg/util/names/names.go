package names

import (
	"fmt"
	"strings"
)

// executionSpacePrefix is the prefix of execution space
const executionSpacePrefix = "karmada-es-"

// GenerateExecutionSpaceName generates execution space name for the given member cluster
func GenerateExecutionSpaceName(memberClusterName string) (string, error) {
	if memberClusterName == "" {
		return "", fmt.Errorf("the member cluster name is empty")
	}
	executionSpace := executionSpacePrefix + memberClusterName
	return executionSpace, nil
}

// GetMemberClusterName returns member cluster name for the given execution space
func GetMemberClusterName(executionSpaceName string) (string, error) {
	if !strings.HasPrefix(executionSpaceName, executionSpacePrefix) {
		return "", fmt.Errorf("the execution space name is in wrong format")
	}
	return strings.TrimPrefix(executionSpaceName, executionSpacePrefix), nil
}

// GenerateBindingName will generate binding name by namespace, kind and name
func GenerateBindingName(namespace, kind, name string) string {
	return strings.ToLower(namespace + "-" + kind + "-" + name)
}

// GenerateOwnerLabelValue will get owner label value.
func GenerateOwnerLabelValue(namespace, name string) string {
	return namespace + "." + name
}