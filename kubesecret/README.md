### kubesecret

command line tool for managing Kubernetes secrets in clusters. The underlying pinnings
are the standard kubernetes client libraries and will honor any and all permissions
within the cluster(s). The assumption when using this tool is that you already have
access to the cluster(s) and a valid `~/.kube/config` file.

### build

`go install .` should install a `kubesecret` binary built for your platform into your
`$GOPATH/bin` directory.

### usage

The Command actions are `create`, `get`, `update` and `delete`.

* create - create a secret in the specified cluster and namespace. If no values are
given for the secret an empty one will be created.
* get - get the secret from the server and write the contents out in `cwd` to a file named
`${secret}.json`
* delete - delete the secret from the server
* update - update the secret to either add or replace a value in the secret.

An action must be supplied. All actions must have a secret specified for which to act upon.

`kube-secret -cluster staging create my-nifty-secret`

`kube-secret -cluster production -namespace cluster get prometheus`

`kube-secret -cluster test -name token -value slartibartfast update lemmy`

Options :

`-value` can be in the form of `@foo.json` where a file named foo.json will be read and used
as the contents for the secret. If the value is not preceded with an `@` that value will be
used for the secret contents.

Secret contents will base base64 encoded at runtime. There is no need to pre-encode them.

```
Usage:
  kubesecret [flags]
  kubesecret [command]

Available Commands:
  create      create a secret in the cluster
  delete      delete secret from cluster
  get         get secret from cluster
  help        Help about any command
  update      update an item in a secret

Flags:
  -c, --cluster string     cluster for secret manipulation
  -h, --help               help for kubesecret
      --kube-conf string   kubectl config file to use (default "/Users/brunton/.kube/config")
      --loglevel string    log level (default "info")
  -n, --namespace string   namespace for secret manipulation (default "default")

Use "kubesecret [command] --help" for more information about a command.
```