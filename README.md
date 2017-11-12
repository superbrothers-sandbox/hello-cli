## hello-cli

```
$ make
$ bin/hello --help
```

### Cannot autocomplete alias of command

Generated bash completion does not work correctly with alias of command.

```
$ bash --version
GNU bash, version 4.4.12(1)-release (x86_64-unknown-linux-gnu)
Copyright (C) 2016 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>

This is free software; you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
$ source $(brew --prefix)/share/bash-completion/bash_completion
$ source <(bin/hello completion)
$ alias h=${PWD}/bin/hello
$ complete -o default -F __start_hello h
$ h <tab>
```

**Actual output:**
```
$ h hello
```

**Expected output:**
```
$ h <tab>
completion  say
```

- https://github.com/kubernetes/kubectl/issues/120 
