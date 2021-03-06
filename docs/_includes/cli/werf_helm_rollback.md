{% if include.header %}
{% assign header = include.header %}
{% else %}
{% assign header = "###" %}
{% endif %}
Rollback a release to the specified revision

{{ header }} Syntax

```shell
werf helm rollback RELEASE_NAME REVISION [options]
```

{{ header }} Options

```shell
      --cleanup-on-fail=false:
            Allow deletion of new resources created in this rollback when rollback failed
      --force=false:
            Force resource update through delete/recreate if needed
      --helm-release-storage-namespace='kube-system':
            Helm release storage namespace (same as --tiller-namespace for regular helm, default    
            $WERF_HELM_RELEASE_STORAGE_NAMESPACE, $TILLER_NAMESPACE or 'kube-system')
      --helm-release-storage-type='configmap':
            helm storage driver to use. One of 'configmap' or 'secret' (default                     
            $WERF_HELM_RELEASE_STORAGE_TYPE or 'configmap')
  -h, --help=false:
            help for rollback
      --home-dir='':
            Use specified dir to store werf cache files and dirs (default $WERF_HOME or ~/.werf)
      --kube-config='':
            Kubernetes config file path
      --kube-context='':
            Kubernetes config context (default $WERF_KUBE_CONTEXT)
      --no-hooks=false:
            Prevent hooks from running during rollback
      --recreate-pods=false:
            Perform pods restart for the resource if applicable
      --releases-history-max=0:
            Max releases to keep in release storage. Can be set by environment variable             
            $WERF_RELEASES_HISTORY_MAX. By default werf keeps all releases.
      --timeout=300:
            Time in seconds to wait for any individual Kubernetes operation (like Jobs for hooks)
      --tmp-dir='':
            Use specified dir to store tmp files and dirs (default $WERF_TMP_DIR or system tmp dir)
      --wait=false:
            If set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a       
            Deployment are in a ready state before marking the release as successful. It will wait  
            for as long as --timeout
```

