package sidecars

import (
	"errors"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

// ErrNoSupportedSidecar occurs when we can't find any premade sidecar for the given Istio version
var ErrNoSupportedSidecar = errors.New("no valid istio sidecar found for this istio version")

// GetIstioSidecar will return an Istio sidecar for the given
// version of Istio, with the given jwtPolicy, to run
// in the gateway-proxy pod
func GetIstioSidecar(istioVersion, jwtPolicy string, istioMetaMeshID string, istioMetaClusterID string) (*corev1.Container, error) {

	istio17to119Match, _ := regexp.MatchString("1.(7|8|9|1[0-9]).[0-9]*", istioVersion)

	if istio17to119Match {
		return generateIstio17to119Sidecar(istioVersion, jwtPolicy, istioMetaMeshID, istioMetaClusterID), nil
	} else if strings.HasPrefix(istioVersion, "1.6.") {
		return generateIstio16Sidecar(istioVersion, jwtPolicy, istioMetaMeshID, istioMetaClusterID), nil
	}
	return nil, ErrNoSupportedSidecar
}
