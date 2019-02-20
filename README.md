# dfm

dfm is a cli for managing your dotfiles. It's heavily influence by the spectacular `rcm` tool, but is os agnostic and written in Go. It's also a single binary.

Commands:

* `dfm init` - intitialise a new dotfiles directory
* `dfm ls` - list files managed by dfm
* `dfm add` - convert a file into a dotfile managed by dfm
* `dfm sync` - synchronise a dofile folder
* `dfm rm` - remove a symlink
* `dfm help` - list the help instructions


## Contributing

Compile

* `make compile`