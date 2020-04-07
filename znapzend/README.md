znapzend
========
Znapzend Helm Chart for automated ZFS snapshot & replication

Current chart version is `0.3.2`





## Chart Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` |  |
| env | object | `{}` | A dict with KEY: VALUE pairs |
| fullnameOverride | string | `""` |  |
| host.zfsDevice | string | `"/dev/zfs"` | The device on the host which is used by the 'zfs' binary within the container |
| image.pullPolicy | string | `"IfNotPresent"` |  |
| image.repository | string | `"docker.io/oetiker/znapzend"` | Znapzend image repository |
| image.tag | string | `"master"` | Znapzend image tag (version) |
| imagePullSecrets | list | `[]` | List of image pull secrets if you use a privately hosted image |
| metrics.enabled | bool | `true` | Enable the znapzend metrics exporter for Prometheus |
| metrics.env | object | `{}` | A dict with KEY: VALUE pairs as environment variables for the exporter |
| metrics.image.pullPolicy | string | `"IfNotPresent"` |  |
| metrics.image.repository | string | `"docker.io/braindoctor/znapzend-exporter"` | Exporter image repository |
| metrics.image.tag | string | `"v0.1.0"` | Exporter image tag |
| metrics.ingress.annotations | object | `{}` |  |
| metrics.ingress.enabled | bool | `false` | Useful if your Prometheus is outside of the cluster |
| metrics.ingress.hosts | list | `[{"host":null,"paths":[]}]` | See Kubernetes Docs for a guide to setup TLS on Ingress |
| metrics.ingress.tls | list | `[]` |  |
| metrics.jobs.register | list | `[]` | String list of datasets that should be registered right at startup |
| metrics.resources | object | `{}` |  |
| metrics.service.enabled | bool | `true` | Whether to enable a Service object for metrics endpoint if metrics.enabled is true |
| metrics.service.nodePort | int | `0` | NodePort if service type is not ClusterIP |
| metrics.service.port | int | `8080` | Port on which the service is reachable |
| metrics.service.type | string | `"ClusterIP"` |  |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` |  |
| podSecurityContext | object | `{}` |  |
| replicaCount | int | `1` | Only increase if you have multiple Nodes with AntiAffinity, otherwise it does not make sense to run > 1 replicas on the same node |
| resources | object | `{}` |  |
| securityContext | object | `{"allowPrivilegeEscalation":true,"privileged":true}` | The current image requires to run privileged in order to access ZFS |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.create | bool | `true` | Specifies whether a service account should be created |
| serviceAccount.name | string | `nil` | If not set and create is true, a name is generated using the fullname template |
| ssh.config | string | `nil` | ssh_config(5)-compatible file content to configure SSH options when connecting |
| ssh.externalSecretName | string | `nil` | Set this value if you provide your own secret with SSH config |
| ssh.identities | object | `{}` | Provide a private key for each SSH identity, see values.yaml for an example |
| ssh.knownHosts | string | `nil` | List of {host, pubKey} dicts where the public key of each host is configured |
| ssh.path | string | `"/root/.ssh"` | Path where your SSH config and identities get mounted in the container |
| tolerations | list | `[]` |  |