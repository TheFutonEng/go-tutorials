# Leanred from This Section

I still do not understand how modules work.  Modules were disabled on my system after trying to use Go workspaces.

This will happen anywhere in the directory structure when trying to create a module. 

```
[rmengert@Robs-MBP:~]
$ go mod init got-test
go: modules disabled by GO111MODULE=off; see 'go help modules'
[rmengert@Robs-MBP:~]
```

This doesn't impact being able to proceed through the material and run `go test` where needed.  

#TODO: Understand modules and fix my laptop to be able to use them properly.