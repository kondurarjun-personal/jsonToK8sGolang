package main

import (
	"fmt"
	"io/ioutil"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

var jsonExamplePod = `
{
    "apiVersion": "v1",
    "kind": "Pod",
    "metadata": {
        "annotations": {
            "kubernetes.io/psp": "eks.privileged",
            "traffic.sidecar.istio.io/excludeOutboundPorts": "8089,8191,9997,7777,9000,17000,17500,19000",
            "traffic.sidecar.istio.io/includeInboundPorts": "8000,8088"
        },
        "creationTimestamp": "2021-01-22T23:01:31Z",
        "generateName": "splunk-test1-standalone-",
        "labels": {
            "app.kubernetes.io/component": "standalone",
            "app.kubernetes.io/instance": "splunk-test1-standalone",
            "app.kubernetes.io/managed-by": "splunk-operator",
            "app.kubernetes.io/name": "standalone",
            "app.kubernetes.io/part-of": "splunk-test1-standalone",
            "controller-revision-hash": "splunk-test1-standalone-85fbdb76b7",
            "statefulset.kubernetes.io/pod-name": "splunk-test1-standalone-0"
        },
        "name": "splunk-test1-standalone-0",
        "namespace": "default",
        "ownerReferences": [
            {
                "apiVersion": "apps/v1",
                "blockOwnerDeletion": true,
                "controller": true,
                "kind": "StatefulSet",
                "name": "splunk-test1-standalone",
                "uid": "cd689931-5036-4a40-b34a-4f0777840065"
            }
        ],
        "resourceVersion": "37980869",
        "selfLink": "/api/v1/namespaces/default/pods/splunk-test1-standalone-0",
        "uid": "71c9613f-4334-4bfb-b648-fde4cfe12dab"
    },
    "spec": {
        "affinity": {
            "podAntiAffinity": {
                "preferredDuringSchedulingIgnoredDuringExecution": [
                    {
                        "podAffinityTerm": {
                            "labelSelector": {
                                "matchExpressions": [
                                    {
                                        "key": "app.kubernetes.io/instance",
                                        "operator": "In",
                                        "values": [
                                            "splunk-test1-standalone"
                                        ]
                                    }
                                ]
                            },
                            "topologyKey": "kubernetes.io/hostname"
                        },
                        "weight": 100
                    }
                ]
            }
        },
        "containers": [
            {
                "env": [
                    {
                        "name": "SPLUNK_HOME",
                        "value": "/opt/splunk"
                    },
                    {
                        "name": "SPLUNK_START_ARGS",
                        "value": "--accept-license"
                    },
                    {
                        "name": "SPLUNK_DEFAULTS_URL",
                        "value": "/mnt/splunk-secrets/default.yml"
                    },
                    {
                        "name": "SPLUNK_HOME_OWNERSHIP_ENFORCEMENT",
                        "value": "false"
                    },
                    {
                        "name": "SPLUNK_ROLE",
                        "value": "splunk_standalone"
                    },
                    {
                        "name": "SPLUNK_DECLARATIVE_ADMIN_PASSWORD",
                        "value": "true"
                    }
                ],
                "image": "splunk/splunk:edge",
                "imagePullPolicy": "IfNotPresent",
                "livenessProbe": {
                    "exec": {
                        "command": [
                            "/sbin/checkstate.sh"
                        ]
                    },
                    "failureThreshold": 3,
                    "initialDelaySeconds": 300,
                    "periodSeconds": 30,
                    "successThreshold": 1,
                    "timeoutSeconds": 30
                },
                "name": "splunk",
                "ports": [
                    {
                        "containerPort": 8000,
                        "name": "http-splunkweb",
                        "protocol": "TCP"
                    },
                    {
                        "containerPort": 8088,
                        "name": "http-hec",
                        "protocol": "TCP"
                    },
                    {
                        "containerPort": 8089,
                        "name": "https-splunkd",
                        "protocol": "TCP"
                    },
                    {
                        "containerPort": 9000,
                        "name": "tcp-dfsmaster",
                        "protocol": "TCP"
                    },
                    {
                        "containerPort": 9997,
                        "name": "tcp-s2s",
                        "protocol": "TCP"
                    },
                    {
                        "containerPort": 17000,
                        "name": "tcp-dfccontrol",
                        "protocol": "TCP"
                    },
                    {
                        "containerPort": 19000,
                        "name": "tcp-datareceive",
                        "protocol": "TCP"
                    }
                ],
                "readinessProbe": {
                    "exec": {
                        "command": [
                            "/bin/grep",
                            "started",
                            "/opt/container_artifact/splunk-container.state"
                        ]
                    },
                    "failureThreshold": 3,
                    "initialDelaySeconds": 10,
                    "periodSeconds": 5,
                    "successThreshold": 1,
                    "timeoutSeconds": 5
                },
                "resources": {
                    "limits": {
                        "cpu": "4",
                        "memory": "8Gi"
                    },
                    "requests": {
                        "cpu": "100m",
                        "memory": "512Mi"
                    }
                },
                "terminationMessagePath": "/dev/termination-log",
                "terminationMessagePolicy": "File",
                "volumeMounts": [
                    {
                        "mountPath": "/opt/splunk/etc",
                        "name": "pvc-etc"
                    },
                    {
                        "mountPath": "/opt/splunk/var",
                        "name": "pvc-var"
                    },
                    {
                        "mountPath": "/mnt/splunk-secrets",
                        "name": "mnt-splunk-secrets"
                    },
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
        "hostname": "splunk-test1-standalone-0",
        "nodeName": "ip-192-168-39-106.us-west-2.compute.internal",
        "priority": 0,
        "restartPolicy": "Always",
        "schedulerName": "default-scheduler",
        "securityContext": {
            "fsGroup": 41812,
            "runAsUser": 41812
        },
        "serviceAccount": "default",
        "serviceAccountName": "default",
        "subdomain": "splunk-test1-standalone-headless",
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
                "name": "pvc-etc",
                "persistentVolumeClaim": {
                    "claimName": "pvc-etc-splunk-test1-standalone-0"
                }
            },
            {
                "name": "pvc-var",
                "persistentVolumeClaim": {
                    "claimName": "pvc-var-splunk-test1-standalone-0"
                }
            },
            {
                "name": "mnt-splunk-secrets",
                "secret": {
                    "defaultMode": 420,
                    "secretName": "splunk-test1-standalone-secret-v1"
                }
            },
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
                "lastTransitionTime": "2021-01-22T23:01:37Z",
                "status": "True",
                "type": "Initialized"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-22T23:02:18Z",
                "status": "True",
                "type": "Ready"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-22T23:02:18Z",
                "status": "True",
                "type": "ContainersReady"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2021-01-22T23:01:37Z",
                "status": "True",
                "type": "PodScheduled"
            }
        ],
        "containerStatuses": [
            {
                "containerID": "docker://59246ec7c2839e34fa7318466278cbd7f8f7ec700613c404a29f6d5c9112eecf",
                "image": "splunk/splunk:edge",
                "imageID": "docker-pullable://splunk/splunk@sha256:b464d33f80aaf27301d7723771d0ee4285855c0c228614da3f49a795c7f89dc2",
                "lastState": {},
                "name": "splunk",
                "ready": true,
                "restartCount": 0,
                "started": true,
                "state": {
                    "running": {
                        "startedAt": "2021-01-22T23:01:47Z"
                    }
                }
            }
        ],
        "hostIP": "192.168.39.106",
        "phase": "Running",
        "podIP": "192.168.51.188",
        "podIPs": [
            {
                "ip": "192.168.51.188"
            }
        ],
        "qosClass": "Burstable",
        "startTime": "2021-01-22T23:01:37Z"
    }
}
`

func main() {
	// Variable holding data to be deserialized
	var byteValue []byte

	// If file argument exists read file contents to deserialize
	if len(os.Args) > 1 {
		jsonFileName := os.Args[1]
		fmt.Println("\nReading file: ", jsonFileName)

		// Open our jsonFile
		jsonFile, err := os.Open(jsonFileName)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Printf("Unable to read file %s, error: %+v", jsonFileName, err)
			return
		}
		fmt.Println("Successfully read file ", jsonFileName)

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// Read from json file
		byteValue, _ = ioutil.ReadAll(jsonFile)
	} else {
		// If no file argument, read from example Pod Json
		byteValue = []byte(jsonExamplePod)
	}

	// Deserialize the file
	deser := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := deser(byteValue, nil, nil)
	if err != nil {
		fmt.Printf("Unable to deserialize %#v", err)
	}

	// Extract pod object
	switch obj.(type) {
	case *corev1.Pod:
		fmt.Printf("Deserialized pod: \n%#v\n", obj.(*corev1.Pod))
		break
	case *appsv1.StatefulSet:
		fmt.Printf("Deserialized statefulSet: \n%#v\n", obj.(*appsv1.StatefulSet))
		break
	default:
		fmt.Printf("Didn't find object type")
	}

	if len(os.Args) > 2 {
		// Write contents to a file
		_ = writeContentsToString(obj)

		fmt.Println("Writing to file ", os.Args[2])
		//writeFile(os.Args[2], string(text))
	}
}

// writeFile writes the output to a .js file & beautifies it
func writeFile(path string, text string) {
	// check if file exists
	_, err := os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var _, err = os.Create(path)
		if err != nil {
			return
		}
		fmt.Println("Output file created due to inexistence", path)
	}

	// Open file using READ & WRITE permission.
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString(text)
	if err != nil {
		return
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		return
	}
}

func writeContentsToString(obj runtime.Object) []byte {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fmt.Println(obj) // this gets captured

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	//fmt.Printf("Captured: %s", out) // prints: Captured: Hello, playground
	return out
}
