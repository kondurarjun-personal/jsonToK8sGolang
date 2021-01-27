# jsonToK8sGolang

Library to deserialize json structures for k8s objects such as pods, statefulSets etc..

Example input JSON file example_pod.json:

```
{
    "apiVersion": "v1",
    "kind": "Pod",
    "metadata": {
        "annotations": {
            "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"name\":\"nginx\"},\"name\":\"nginx\",\"namespace\":\"default\"},\"spec\":{\"containers\":[{\"image\":\"nginx\",\"name\":\"nginx\",\"ports\":[{\"containerPort\":80}]}]}}\n",
            "kubernetes.io/psp": "eks.privileged"
        },
        "creationTimestamp": "2021-01-27T00:59:48Z",
        "labels": {
            "name": "nginx"
        },
        "name": "nginx",
        "namespace": "default",
        "resourceVersion": "39005029",
        "selfLink": "/api/v1/namespaces/default/pods/nginx",
        "uid": "4e8791fd-11a9-4ae8-8695-a379a9a4770f"
    },
    "spec": {
        "containers": [
            {
                "image": "nginx",
                "imagePullPolicy": "Always",
                "name": "nginx",
                "ports": [
                    {
                        "containerPort": 80,
                        "protocol": "TCP"
                    }
                ],
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "terminationMessagePolicy": "File",
                "volumeMounts": [
                    {
                        "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                        "name": "default-token-97bds",
                        "readOnly": true
                    }
                ]
            }
        ],
        "dnsPolicy": "ClusterFirst",
        "enableServiceLinks": true,
        "nodeName": "ip-192-168-80-80.us-west-2.compute.internal",
        "priority": 0,
        "restartPolicy": "Always",
        "schedulerName": "default-scheduler",
        "securityContext": {},
        "serviceAccount": "default",
        "serviceAccountName": "default",
        "terminationGracePeriodSeconds": 30,
        "tolerations": [
            {
                "effect": "NoExecute",
                "key": "node.kubernetes.io/not-ready",
                "operator": "Exists",
                "tolerationSeconds": 300
            },
            {
                "effect": "NoExecute",
                "key": "node.kubernetes.io/unreachable",
                "operator": "Exists",
                "tolerationSeconds": 300
            }
        ],
        "volumes": [
            {
                "name": "default-token-97bds",
                "secret": {
                    "defaultMode": 420,
                    "secretName": "default-token-97bds"
                }
            }
        ]
    },
    "status": {
        "conditions": [
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-27T00:59:48Z",
                "status": "True",
                "type": "Initialized"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-27T00:59:54Z",
                "status": "True",
                "type": "Ready"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-27T00:59:54Z",
                "status": "True",
                "type": "ContainersReady"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-27T00:59:48Z",
                "status": "True",
                "type": "PodScheduled"
            }
        ],
        "containerStatuses": [
            {
                "containerID": "docker://8ac0f511a15c91beafe6b4a5f3aa2126d74e194db8ba32aa65d55b5f15f1e282",
                "image": "nginx:latest",
                "imageID": "docker-pullable://nginx@sha256:10b8cc432d56da8b61b070f4c7d2543a9ed17c2b23010b43af434fd40e2ca4aa",
                "lastState": {},
                "name": "nginx",
                "ready": true,
                "restartCount": 0,
                "started": true,
                "state": {
                    "running": {
                        "startedAt": "2021-01-27T00:59:54Z"
                    }
                }
            }
        ],
        "hostIP": "192.168.80.80",
        "phase": "Running",
        "podIP": "192.168.75.69",
        "podIPs": [
            {
                "ip": "192.168.75.69"
            }
        ],
        "qosClass": "BestEffort",
        "startTime": "2021-01-27T00:59:48Z"
    }
}
```

Run the json file through the tool using "<executable> example_pod.json > output.js"

Output JS file (once beautified through a VS code extension):

```
& v1.Pod {
    TypeMeta: v1.TypeMeta {
        Kind: "Pod",
        APIVersion: "v1"
    },
    ObjectMeta: v1.ObjectMeta {
        Name: "nginx",
        GenerateName: "",
        Namespace: "default",
        SelfLink: "/api/v1/namespaces/default/pods/nginx",
        UID: "4e8791fd-11a9-4ae8-8695-a379a9a4770f",
        ResourceVersion: "39005029",
        Generation: 0,
        CreationTimestamp: v1.Time {
            Time: time.Time {
                wall: 0x0,
                ext: 63747305988,
                loc: ( * time.Location)(0x2793800)
            }
        },
        DeletionTimestamp: ( * v1.Time)(nil),
        DeletionGracePeriodSeconds: ( * int64)(nil),
        Labels: map[string] string {
            "name": "nginx"
        },
        Annotations: map[string] string {
            "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"Pod\",\"metadata\":{\"annotations\":{},\"labels\":{\"name\":\"nginx\"},\"name\":\"nginx\",\"namespace\":\"default\"},\"spec\":{\"containers\":[{\"image\":\"nginx\",\"name\":\"nginx\",\"ports\":[{\"containerPort\":80}]}]}}\n",
            "kubernetes.io/psp": "eks.privileged"
        },
        OwnerReferences: [] v1.OwnerReference(nil),
        Finalizers: [] string(nil),
        ClusterName: "",
        ManagedFields: [] v1.ManagedFieldsEntry(nil)
    },
    Spec: v1.PodSpec {
        Volumes: [] v1.Volume {
            v1.Volume {
                Name: "default-token-97bds",
                VolumeSource: v1.VolumeSource {
                    HostPath: ( * v1.HostPathVolumeSource)(nil),
                    EmptyDir: ( * v1.EmptyDirVolumeSource)(nil),
                    GCEPersistentDisk: ( * v1.GCEPersistentDiskVolumeSource)(nil),
                    AWSElasticBlockStore: ( * v1.AWSElasticBlockStoreVolumeSource)(nil),
                    GitRepo: ( * v1.GitRepoVolumeSource)(nil),
                    Secret: ( * v1.SecretVolumeSource)(0xc00047d400),
                    NFS: ( * v1.NFSVolumeSource)(nil),
                    ISCSI: ( * v1.ISCSIVolumeSource)(nil),
                    Glusterfs: ( * v1.GlusterfsVolumeSource)(nil),
                    PersistentVolumeClaim: ( * v1.PersistentVolumeClaimVolumeSource)(nil),
                    RBD: ( * v1.RBDVolumeSource)(nil),
                    FlexVolume: ( * v1.FlexVolumeSource)(nil),
                    Cinder: ( * v1.CinderVolumeSource)(nil),
                    CephFS: ( * v1.CephFSVolumeSource)(nil),
                    Flocker: ( * v1.FlockerVolumeSource)(nil),
                    DownwardAPI: ( * v1.DownwardAPIVolumeSource)(nil),
                    FC: ( * v1.FCVolumeSource)(nil),
                    AzureFile: ( * v1.AzureFileVolumeSource)(nil),
                    ConfigMap: ( * v1.ConfigMapVolumeSource)(nil),
                    VsphereVolume: ( * v1.VsphereVirtualDiskVolumeSource)(nil),
                    Quobyte: ( * v1.QuobyteVolumeSource)(nil),
                    AzureDisk: ( * v1.AzureDiskVolumeSource)(nil),
                    PhotonPersistentDisk: ( * v1.PhotonPersistentDiskVolumeSource)(nil),
                    Projected: ( * v1.ProjectedVolumeSource)(nil),
                    PortworxVolume: ( * v1.PortworxVolumeSource)(nil),
                    ScaleIO: ( * v1.ScaleIOVolumeSource)(nil),
                    StorageOS: ( * v1.StorageOSVolumeSource)(nil),
                    CSI: ( * v1.CSIVolumeSource)(nil),
                    Ephemeral: ( * v1.EphemeralVolumeSource)(nil)
                }
            }
        },
        InitContainers: [] v1.Container(nil),
        Containers: [] v1.Container {
            v1.Container {
                Name: "nginx",
                Image: "nginx",
                Command: [] string(nil),
                Args: [] string(nil),
                WorkingDir: "",
                Ports: [] v1.ContainerPort {
                    v1.ContainerPort {
                        Name: "",
                        HostPort: 0,
                        ContainerPort: 80,
                        Protocol: "TCP",
                        HostIP: ""
                    }
                },
                EnvFrom: [] v1.EnvFromSource(nil),
                Env: [] v1.EnvVar(nil),
                Resources: v1.ResourceRequirements {
                    Limits: v1.ResourceList(nil),
                    Requests: v1.ResourceList(nil)
                },
                VolumeMounts: [] v1.VolumeMount {
                    v1.VolumeMount {
                        Name: "default-token-97bds",
                        ReadOnly: true,
                        MountPath: "/var/run/secrets/kubernetes.io/serviceaccount",
                        SubPath: "",
                        MountPropagation: ( * v1.MountPropagationMode)(nil),
                        SubPathExpr: ""
                    }
                },
                VolumeDevices: [] v1.VolumeDevice(nil),
                LivenessProbe: ( * v1.Probe)(nil),
                ReadinessProbe: ( * v1.Probe)(nil),
                StartupProbe: ( * v1.Probe)(nil),
                Lifecycle: ( * v1.Lifecycle)(nil),
                TerminationMessagePath: "/dev/termination-log",
                TerminationMessagePolicy: "File",
                ImagePullPolicy: "Always",
                SecurityContext: ( * v1.SecurityContext)(nil),
                Stdin: false,
                StdinOnce: false,
                TTY: false
            }
        },
        EphemeralContainers: [] v1.EphemeralContainer(nil),
        RestartPolicy: "Always",
        TerminationGracePeriodSeconds: ( * int64)(0xc000469ad0),
        ActiveDeadlineSeconds: ( * int64)(nil),
        DNSPolicy: "ClusterFirst",
        NodeSelector: map[string] string(nil),
        ServiceAccountName: "default",
        DeprecatedServiceAccount: "default",
        AutomountServiceAccountToken: ( * bool)(nil),
        NodeName: "ip-192-168-80-80.us-west-2.compute.internal",
        HostNetwork: false,
        HostPID: false,
        HostIPC: false,
        ShareProcessNamespace: ( * bool)(nil),
        SecurityContext: ( * v1.PodSecurityContext)(0xc00031d570),
        ImagePullSecrets: [] v1.LocalObjectReference(nil),
        Hostname: "",
        Subdomain: "",
        Affinity: ( * v1.Affinity)(nil),
        SchedulerName: "default-scheduler",
        Tolerations: [] v1.Toleration {
            v1.Toleration {
                Key: "node.kubernetes.io/not-ready",
                Operator: "Exists",
                Value: "",
                Effect: "NoExecute",
                TolerationSeconds: ( * int64)(0xc000469b00)
            }, v1.Toleration {
                Key: "node.kubernetes.io/unreachable",
                Operator: "Exists",
                Value: "",
                Effect: "NoExecute",
                TolerationSeconds: ( * int64)(0xc000469b20)
            }
        },
        HostAliases: [] v1.HostAlias(nil),
        PriorityClassName: "",
        Priority: ( * int32)(0xc000469a70),
        DNSConfig: ( * v1.PodDNSConfig)(nil),
        ReadinessGates: [] v1.PodReadinessGate(nil),
        RuntimeClassName: ( * string)(nil),
        EnableServiceLinks: ( * bool)(0xc000469a49),
        PreemptionPolicy: ( * v1.PreemptionPolicy)(nil),
        Overhead: v1.ResourceList(nil),
        TopologySpreadConstraints: [] v1.TopologySpreadConstraint(nil),
        SetHostnameAsFQDN: ( * bool)(nil)
    },
    Status: v1.PodStatus {
        Phase: "Running",
        Conditions: [] v1.PodCondition {
            v1.PodCondition {
                Type: "Initialized",
                Status: "True",
                LastProbeTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 0,
                        loc: ( * time.Location)(nil)
                    }
                },
                LastTransitionTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 63747305988,
                        loc: ( * time.Location)(0x2793800)
                    }
                },
                Reason: "",
                Message: ""
            }, v1.PodCondition {
                Type: "Ready",
                Status: "True",
                LastProbeTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 0,
                        loc: ( * time.Location)(nil)
                    }
                },
                LastTransitionTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 63747305994,
                        loc: ( * time.Location)(0x2793800)
                    }
                },
                Reason: "",
                Message: ""
            }, v1.PodCondition {
                Type: "ContainersReady",
                Status: "True",
                LastProbeTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 0,
                        loc: ( * time.Location)(nil)
                    }
                },
                LastTransitionTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 63747305994,
                        loc: ( * time.Location)(0x2793800)
                    }
                },
                Reason: "",
                Message: ""
            }, v1.PodCondition {
                Type: "PodScheduled",
                Status: "True",
                LastProbeTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 0,
                        loc: ( * time.Location)(nil)
                    }
                },
                LastTransitionTime: v1.Time {
                    Time: time.Time {
                        wall: 0x0,
                        ext: 63747305988,
                        loc: ( * time.Location)(0x2793800)
                    }
                },
                Reason: "",
                Message: ""
            }
        },
        Message: "",
        Reason: "",
        NominatedNodeName: "",
        HostIP: "192.168.80.80",
        PodIP: "192.168.75.69",
        PodIPs: [] v1.PodIP {
            v1.PodIP {
                IP: "192.168.75.69"
            }
        },
        StartTime: ( * v1.Time)(0xc0004bed00),
        InitContainerStatuses: [] v1.ContainerStatus(nil),
        ContainerStatuses: [] v1.ContainerStatus {
            v1.ContainerStatus {
                Name: "nginx",
                State: v1.ContainerState {
                    Waiting: ( * v1.ContainerStateWaiting)(nil),
                    Running: ( * v1.ContainerStateRunning)(0xc0004bece0),
                    Terminated: ( * v1.ContainerStateTerminated)(nil)
                },
                LastTerminationState: v1.ContainerState {
                    Waiting: ( * v1.ContainerStateWaiting)(nil),
                    Running: ( * v1.ContainerStateRunning)(nil),
                    Terminated: ( * v1.ContainerStateTerminated)(nil)
                },
                Ready: true,
                RestartCount: 0,
                Image: "nginx:latest",
                ImageID: "docker-pullable://nginx@sha256:10b8cc432d56da8b61b070f4c7d2543a9ed17c2b23010b43af434fd40e2ca4aa",
                ContainerID: "docker://8ac0f511a15c91beafe6b4a5f3aa2126d74e194db8ba32aa65d55b5f15f1e282",
                Started: ( * bool)(0xc000469ba5)
            }
        },
        QOSClass: "BestEffort",
        EphemeralContainerStatuses: [] v1.ContainerStatus(nil)
    }
}
```
