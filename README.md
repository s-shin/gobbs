gobbs - a sample web application by golang
==========================================

```sh
### Prepare Project
make prepare

### Debug
# prepare database
make initdb
# run application for debug
make run

### Test
make test

### Build
# output "gobbs" binary.
make build

### Archive
# output gobbs.zip.
make archive

### Graceful restart & shutdown
# restart
kill -USR2 (pid)
# shutdown
kill -TERM (pid)
```
