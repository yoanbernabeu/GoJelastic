## GoJelastic completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	GoJelastic completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
GoJelastic completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --token string   A token is required
      --url string     A url is required
```

### SEE ALSO

* [GoJelastic completion](GoJelastic_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 30-Dec-2022