// Code generated for package generated by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/controller.yaml
// assets/controller_sa.yaml
// assets/csidriver.yaml
// assets/namespace.yaml
// assets/node.yaml
// assets/node_nfs.yaml
// assets/node_sa.yaml
// assets/rbac/controller_privileged_binding.yaml
// assets/rbac/node_privileged_binding.yaml
// assets/rbac/privileged_role.yaml
// assets/rbac/provisioner_binding.yaml
// assets/rbac/provisioner_role.yaml
// assets/rbac/snapshotter_binding.yaml
// assets/rbac/snapshotter_role.yaml
package generated

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _controllerYaml = []byte(`kind: Deployment
apiVersion: apps/v1
metadata:
  name: openstack-manila-csi-controllerplugin
  namespace: openshift-manila-csi-driver
  annotations:
    config.openshift.io/inject-proxy: csi-driver
spec:
  selector:
    matchLabels:
      app: openstack-manila-csi
      component: controllerplugin
  serviceName: manila-csi-driver-controller
  replicas: 1
  template:
    metadata:
      labels:
        app: openstack-manila-csi
        component: controllerplugin
    spec:
      nodeSelector:
        node-role.kubernetes.io/master: ""
      serviceAccount: manila-csi-driver-controller-sa
      priorityClassName: system-cluster-critical
      tolerations:
        - key: CriticalAddonsOnly
          operator: Exists
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: "NoSchedule"
      containers:
        - name: csi-driver
          image: ${DRIVER_IMAGE}
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
          args:
            - --v=${LOG_LEVEL}
            - --nodeid=$(NODE_ID)
            - --endpoint=$(CSI_ENDPOINT)
            - --drivername=$(DRIVER_NAME)
            - --share-protocol-selector=$(MANILA_SHARE_PROTO)
            - --fwdendpoint=$(FWD_CSI_ENDPOINT)
          env:
            - name: DRIVER_NAME
              value: manila.csi.openstack.org
            - name: NODE_ID
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
            - name: CSI_ENDPOINT
              value: unix:///plugin/csi.sock
            - name: MANILA_SHARE_PROTO
              value: NFS
            - name: FWD_CSI_ENDPOINT
              value: unix:///plugin/csi-nfs.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /plugin
            - name: cacert
              mountPath: /etc/kubernetes/static-pod-resources/configmaps/cloud-config
          resources:
            requests:
              cpu: 10m
              memory: 50Mi
        # TODO: fix manila CSI driver not to require NFS driver socket!
        - name: csi-driver-nfs
          image: ${NFS_DRIVER_IMAGE}
          resources:
            requests:
              memory: 20Mi
              cpu: 5m
          args:
            - "--nodeid=$(NODE_ID)"
            - "--endpoint=unix://plugin/csi-nfs.sock"
            - "--mount-permissions=0777"
          env:
            - name: NODE_ID
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: socket-dir
              mountPath: /plugin
          resources:
            requests:
              cpu: 10m
              memory: 50Mi
        - name: csi-provisioner
          image: ${PROVISIONER_IMAGE}
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
          args:
            - --csi-address=$(ADDRESS)
            - --feature-gates=Topology=true
            - --v=${LOG_LEVEL}
            - --timeout=120s
          env:
            - name: ADDRESS
              value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
            - name: socket-dir
              mountPath: /var/lib/csi/sockets/pluginproxy/
          resources:
            requests:
              cpu: 10m
              memory: 50Mi
        - name: csi-snapshotter
          image: ${SNAPSHOTTER_IMAGE}
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
          args:
            - --csi-address=$(ADDRESS)
            - --v=${LOG_LEVEL}
          env:
          - name: ADDRESS
            value: /var/lib/csi/sockets/pluginproxy/csi.sock
          volumeMounts:
          - mountPath: /var/lib/csi/sockets/pluginproxy/
            name: socket-dir
          resources:
            requests:
              cpu: 10m
              memory: 50Mi
      volumes:
        - name: socket-dir
          emptyDir: {}
        - name: cacert
          # If present, extract ca-bundle.pem to
          # /etc/kubernetes/static-pod-resources/configmaps/cloud-config
          # Let the pod start when the ConfigMap does not exist or the certificate
          # is not preset there. The certificate file will be created once the
          # ConfigMap is created / the cerificate is added to it.
          configMap:
            name: cloud-provider-config
            items:
            - key: ca-bundle.pem
              path: ca-bundle.pem
            optional: true
`)

func controllerYamlBytes() ([]byte, error) {
	return _controllerYaml, nil
}

func controllerYaml() (*asset, error) {
	bytes, err := controllerYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _controller_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: manila-csi-driver-controller-sa
  namespace: openshift-manila-csi-driver
`)

func controller_saYamlBytes() ([]byte, error) {
	return _controller_saYaml, nil
}

func controller_saYaml() (*asset, error) {
	bytes, err := controller_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "controller_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _csidriverYaml = []byte(`apiVersion: storage.k8s.io/v1
kind: CSIDriver
metadata:
  name: manila.csi.openstack.org
  annotations:
      # This CSIDriver is managed by an OCP CSI operator
      csi.openshift.io/managed: "true"
spec:
  attachRequired: false
  podInfoOnMount: false
`)

func csidriverYamlBytes() ([]byte, error) {
	return _csidriverYaml, nil
}

func csidriverYaml() (*asset, error) {
	bytes, err := csidriverYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "csidriver.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _namespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-manila-csi-driver
  annotations:
    include.release.openshift.io/self-managed-high-availability: "true"
    openshift.io/node-selector: ""
`)

func namespaceYamlBytes() ([]byte, error) {
	return _namespaceYaml, nil
}

func namespaceYaml() (*asset, error) {
	bytes, err := namespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _nodeYaml = []byte(`kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: openstack-manila-csi-nodeplugin
  namespace: openshift-manila-csi-driver
  annotations:
    config.openshift.io/inject-proxy: csi-driver
spec:
  selector:
    matchLabels:
      app: openstack-manila-csi
      component: nodeplugin
  template:
    metadata:
      labels:
        app: openstack-manila-csi
        component: nodeplugin
    spec:
      hostNetwork: true
      dnsPolicy: ClusterFirstWithHostNet
      serviceAccount: manila-csi-driver-node-sa
      priorityClassName: system-node-critical
      tolerations:
        - operator: Exists
      containers:
        - name: csi-driver
          securityContext:
            privileged: true
          image: ${DRIVER_IMAGE}
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
          args:
            - --v=${LOG_LEVEL}
            - "--nodeid=$(NODE_ID)"
            - "--endpoint=$(CSI_ENDPOINT)"
            - "--drivername=$(DRIVER_NAME)"
            - "--share-protocol-selector=$(MANILA_SHARE_PROTO)"
            - "--fwdendpoint=$(FWD_CSI_ENDPOINT)"
          env:
            - name: DRIVER_NAME
              value: manila.csi.openstack.org
            - name: NODE_ID
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
            - name: CSI_ENDPOINT
              value: unix:///var/lib/kubelet/plugins/manila.csi.openstack.org/csi.sock
            - name: FWD_CSI_ENDPOINT
              value: unix:///var/lib/kubelet/plugins/csi-nfsplugin/csi.sock
            - name: MANILA_SHARE_PROTO
              value: NFS
          volumeMounts:
            - name: plugin-dir
              mountPath: /var/lib/kubelet/plugins/manila.csi.openstack.org
            - name: fwd-plugin-dir
              mountPath: /var/lib/kubelet/plugins/csi-nfsplugin
            - name: cacert
              mountPath: /etc/kubernetes/static-pod-resources/configmaps/cloud-config
          resources:
            requests:
              cpu: 10m
              memory: 50Mi
        - name: csi-node-driver-registrar
          securityContext:
            privileged: true
          image: ${NODE_DRIVER_REGISTRAR_IMAGE}
          resources:
            requests:
              memory: 20Mi
              cpu: 5m
          args:
            - --v=${LOG_LEVEL}
            - --csi-address=/csi/csi.sock
            - --kubelet-registration-path=/var/lib/kubelet/plugins/manila.csi.openstack.org/csi.sock
          lifecycle:
            preStop:
              exec:
                command: ["/bin/sh", "-c", "rm -rf /var/lib/kubelet/plugins/manila.csi.openstack.org/csi.sock"]
          env:
            - name: KUBE_NODE_NAME
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: plugin-dir
              mountPath: /csi
            - name: registration-dir
              mountPath: /registration
          resources:
            requests:
              cpu: 5m
              memory: 20Mi
      volumes:
        - name: registration-dir
          hostPath:
            path: /var/lib/kubelet/plugins_registry/
            type: Directory
        - name: plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/manila.csi.openstack.org
            type: DirectoryOrCreate
        - name: fwd-plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/csi-nfsplugin
            type: DirectoryOrCreate
        - name: cacert
          # Extract ca-bundle.pem to /etc/kubernetes/static-pod-resources/configmaps/cloud-config if present.
          # Let the pod start when the ConfigMap does not exist or the certificate
          # is not preset there. The certificate file will be created once the
          # ConfigMap is created / the cerificate is added to it.
          configMap:
            name: cloud-provider-config
            items:
            - key: ca-bundle.pem
              path: ca-bundle.pem
            optional: true
`)

func nodeYamlBytes() ([]byte, error) {
	return _nodeYaml, nil
}

func nodeYaml() (*asset, error) {
	bytes, err := nodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node_nfsYaml = []byte(`kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: csi-nodeplugin-nfsplugin
  namespace: openshift-manila-csi-driver
spec:
  selector:
    matchLabels:
      app: openstack-manila-csi
      component: nfs-nodeplugin
  template:
    metadata:
      labels:
        app: openstack-manila-csi
        component: nfs-nodeplugin
    spec:
      hostNetwork: true
      dnsPolicy: ClusterFirstWithHostNet
      serviceAccount: manila-csi-driver-node-sa
      priorityClassName: system-node-critical
      tolerations:
        - operator: Exists
      containers:
        - name: csi-driver
          securityContext:
            privileged: true
          image: ${NFS_DRIVER_IMAGE}
          resources:
            requests:
              memory: 50Mi
              cpu: 10m
          args:
            - --v=${LOG_LEVEL}
            - "--nodeid=$(NODE_ID)"
            - "--endpoint=unix://plugin/csi.sock"
            - "--mount-permissions=0777"
          env:
            - name: NODE_ID
              valueFrom:
                fieldRef:
                  fieldPath: spec.nodeName
          volumeMounts:
            - name: plugin-dir
              mountPath: /plugin
            - name: pods-mount-dir
              mountPath: /var/lib/kubelet/pods
              mountPropagation: Bidirectional
      volumes:
        - name: plugin-dir
          hostPath:
            path: /var/lib/kubelet/plugins/csi-nfsplugin
            type: DirectoryOrCreate
        - name: pods-mount-dir
          hostPath:
            path: /var/lib/kubelet/pods
            type: DirectoryOrCreate
`)

func node_nfsYamlBytes() ([]byte, error) {
	return _node_nfsYaml, nil
}

func node_nfsYaml() (*asset, error) {
	bytes, err := node_nfsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node_nfs.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _node_saYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: manila-csi-driver-node-sa
  namespace: openshift-manila-csi-driver`)

func node_saYamlBytes() ([]byte, error) {
	return _node_saYaml, nil
}

func node_saYaml() (*asset, error) {
	bytes, err := node_saYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "node_sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacController_privileged_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-controller-privileged-binding
subjects:
  - kind: ServiceAccount
    name: manila-csi-driver-controller-sa
    namespace: openshift-manila-csi-driver
roleRef:
  kind: ClusterRole
  name: manila-privileged-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacController_privileged_bindingYamlBytes() ([]byte, error) {
	return _rbacController_privileged_bindingYaml, nil
}

func rbacController_privileged_bindingYaml() (*asset, error) {
	bytes, err := rbacController_privileged_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/controller_privileged_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacNode_privileged_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-node-privileged-binding
subjects:
  - kind: ServiceAccount
    name: manila-csi-driver-node-sa
    namespace: openshift-manila-csi-driver
roleRef:
  kind: ClusterRole
  name: manila-privileged-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacNode_privileged_bindingYamlBytes() ([]byte, error) {
	return _rbacNode_privileged_bindingYaml, nil
}

func rbacNode_privileged_bindingYaml() (*asset, error) {
	bytes, err := rbacNode_privileged_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/node_privileged_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacPrivileged_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-privileged-role
rules:
  - apiGroups: ["security.openshift.io"]
    resourceNames: ["privileged"]
    resources: ["securitycontextconstraints"]
    verbs: ["use"]
`)

func rbacPrivileged_roleYamlBytes() ([]byte, error) {
	return _rbacPrivileged_roleYaml, nil
}

func rbacPrivileged_roleYaml() (*asset, error) {
	bytes, err := rbacPrivileged_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/privileged_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacProvisioner_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-csi-provisioner-binding
subjects:
  - kind: ServiceAccount
    name: manila-csi-driver-controller-sa
    namespace: openshift-manila-csi-driver
roleRef:
  kind: ClusterRole
  name: manila-external-provisioner-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacProvisioner_bindingYamlBytes() ([]byte, error) {
	return _rbacProvisioner_bindingYaml, nil
}

func rbacProvisioner_bindingYaml() (*asset, error) {
	bytes, err := rbacProvisioner_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/provisioner_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacProvisioner_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-external-provisioner-role
rules:
  - apiGroups: [""]
    resources: ["persistentvolumes"]
    verbs: ["get", "list", "watch", "create", "delete"]
  - apiGroups: [""]
    resources: ["persistentvolumeclaims"]
    verbs: ["get", "list", "watch", "update"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["storageclasses"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["events"]
    verbs: ["get", "list", "watch", "create", "update", "patch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["csinodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: [""]
    resources: ["nodes"]
    verbs: ["get", "list", "watch"]
  - apiGroups: ["storage.k8s.io"]
    resources: ["volumeattachments"]
    verbs: ["get", "list", "watch"]
`)

func rbacProvisioner_roleYamlBytes() ([]byte, error) {
	return _rbacProvisioner_roleYaml, nil
}

func rbacProvisioner_roleYaml() (*asset, error) {
	bytes, err := rbacProvisioner_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/provisioner_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacSnapshotter_bindingYaml = []byte(`kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-csi-snapshotter-binding
subjects:
  - kind: ServiceAccount
    name: manila-csi-driver-controller-sa
    namespace: openshift-manila-csi-driver
roleRef:
  kind: ClusterRole
  name: manila-external-snapshotter-role
  apiGroup: rbac.authorization.k8s.io
`)

func rbacSnapshotter_bindingYamlBytes() ([]byte, error) {
	return _rbacSnapshotter_bindingYaml, nil
}

func rbacSnapshotter_bindingYaml() (*asset, error) {
	bytes, err := rbacSnapshotter_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/snapshotter_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _rbacSnapshotter_roleYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: manila-external-snapshotter-role
rules:
- apiGroups: [""]
  resources: ["persistentvolumes"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["persistentvolumeclaims"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["storage.k8s.io"]
  resources: ["storageclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: [""]
  resources: ["events"]
  verbs: ["list", "watch", "create", "update", "patch"]
- apiGroups: [""]
  resources: ["secrets"]
  verbs: ["get", "list"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotclasses"]
  verbs: ["get", "list", "watch"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotcontents"]
  verbs: ["create", "get", "list", "watch", "update", "delete"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshotcontents/status"]
  verbs: ["update"]
- apiGroups: ["snapshot.storage.k8s.io"]
  resources: ["volumesnapshots"]
  verbs: ["get", "list", "watch", "update"]
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["create", "list", "watch", "delete"]
- apiGroups: ["coordination.k8s.io"]
  resources: ["leases"]
  verbs: ["get", "watch", "list", "delete", "update", "create"]
`)

func rbacSnapshotter_roleYamlBytes() ([]byte, error) {
	return _rbacSnapshotter_roleYaml, nil
}

func rbacSnapshotter_roleYaml() (*asset, error) {
	bytes, err := rbacSnapshotter_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rbac/snapshotter_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"controller.yaml":    controllerYaml,
	"controller_sa.yaml": controller_saYaml,
	"csidriver.yaml":     csidriverYaml,
	"namespace.yaml":     namespaceYaml,
	"node.yaml":          nodeYaml,
	"node_nfs.yaml":      node_nfsYaml,
	"node_sa.yaml":       node_saYaml,
	"rbac/controller_privileged_binding.yaml": rbacController_privileged_bindingYaml,
	"rbac/node_privileged_binding.yaml":       rbacNode_privileged_bindingYaml,
	"rbac/privileged_role.yaml":               rbacPrivileged_roleYaml,
	"rbac/provisioner_binding.yaml":           rbacProvisioner_bindingYaml,
	"rbac/provisioner_role.yaml":              rbacProvisioner_roleYaml,
	"rbac/snapshotter_binding.yaml":           rbacSnapshotter_bindingYaml,
	"rbac/snapshotter_role.yaml":              rbacSnapshotter_roleYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"controller.yaml":    {controllerYaml, map[string]*bintree{}},
	"controller_sa.yaml": {controller_saYaml, map[string]*bintree{}},
	"csidriver.yaml":     {csidriverYaml, map[string]*bintree{}},
	"namespace.yaml":     {namespaceYaml, map[string]*bintree{}},
	"node.yaml":          {nodeYaml, map[string]*bintree{}},
	"node_nfs.yaml":      {node_nfsYaml, map[string]*bintree{}},
	"node_sa.yaml":       {node_saYaml, map[string]*bintree{}},
	"rbac": {nil, map[string]*bintree{
		"controller_privileged_binding.yaml": {rbacController_privileged_bindingYaml, map[string]*bintree{}},
		"node_privileged_binding.yaml":       {rbacNode_privileged_bindingYaml, map[string]*bintree{}},
		"privileged_role.yaml":               {rbacPrivileged_roleYaml, map[string]*bintree{}},
		"provisioner_binding.yaml":           {rbacProvisioner_bindingYaml, map[string]*bintree{}},
		"provisioner_role.yaml":              {rbacProvisioner_roleYaml, map[string]*bintree{}},
		"snapshotter_binding.yaml":           {rbacSnapshotter_bindingYaml, map[string]*bintree{}},
		"snapshotter_role.yaml":              {rbacSnapshotter_roleYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
