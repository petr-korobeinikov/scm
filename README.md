# scm 💪

---

## It's time 🕒 to clone ⤵️ interesting 🧐 repo faster ⏩ and cleaner 🧹

![Usage example](demo.svg)

Just type:

```shell
scm https://github.com/pkorobeinikov/scm
```

It will clone `https://github.com/pkorobeinikov/scm` into `~/Workspace/github.com/pkorobeinikov/scm`.

It's also possible to clone `hg`-repo. So command:

```shell
scm hg http://hg.nginx.org/nginx
```

will clone `scm hg http://hg.nginx.org/nginx/` into `~/Workspace/hg.nginx.org/nginx`.

## Configuration

Put this into your `.rc`-file:

```shell
export SCM_WORKSPACE_DIR="~/Projects"         # defaults to ~/Workspace
export SCM_WORKSPACE_DIR_DEFAULT_PERM="0755"  # defaults to 0755
```

## Building from source

```shell
go build -o ~/Bin/scm main.go
```

## Running tests

```shell
 go test  -cover -v ./internal
```
