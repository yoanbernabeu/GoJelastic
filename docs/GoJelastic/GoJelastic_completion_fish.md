## GoJelastic completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	GoJelastic completion fish | source

To load completions for every new session, execute once:

	GoJelastic completion fish > ~/.config/fish/completions/GoJelastic.fish

You will need to start a new shell for this setup to take effect.


```
GoJelastic completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --token string   A token is required
      --url string     A url is required
```

### SEE ALSO

* [GoJelastic completion](GoJelastic_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 22-Dec-2022