# aliyun-disk

[![Build Status](https://travis-ci.org/pragkent/aliyun-disk.svg?branch=master)](https://travis-ci.org/pragkent/aliyun-disk)

Aliyun Disk Kubernetes FlexVolume Driver

## Usage
1. Copy aliyun-disk binary to kubelet volume plugin directory. (`/usr/libexec/kubernetes/kubelet-plugins/volume/exec/pragkent.me~aliyun-disk`)
2. Add three env variables to kube-controller-manager:
  - ALIYUN_ACCESS_KEY
  - ALIYUN_ACCESS_SECRET
  - ALIYUN_REGION

## Install

To install, use `go get`:

```bash
$ go get github.com/pragkent/aliyun-disk
```

## Examples
### Volume
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
  - name: myfrontend
    image: nginx
    volumeMounts:
    - mountPath: /var/www/html
      name: my-disk
  volumes:
    - name: my-disk
      flexVolume:
        driver: pragkent.me/aliyun-disk
        fsType: ext4
        options:
          diskId: "d-12345"
```

### Static PersistentVolume
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: myvol
  labels:
    type: aliyundisk
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteOnce
  flexVolume:
    driver: pragkent.me/aliyun-disk
    fsType: ext4
    options:
      diskId: "d-12345"
```


### Dynamic PersistentVolume and StorageClass
Deploy [aliyundisk-provisioner](https://github.com/pragkent/aliyundisk-provisioner) on your cluster first.

```yaml
kind: StorageClass
apiVersion: storage.k8s.io/v1
metadata:
  name: ssd
provisioner: pragkent.me/aliyun-disk
parameters:
  type: cloud_ssd

---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: my-pvd-claim
spec:
  accessModes:
    - ReadWriteOnce
  storageClassName: ssd
  resources:
    requests:
      storage: 5Gi
```
