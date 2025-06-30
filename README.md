# messh
Cross-Platform SSH Manager

messh BASE COMMAND

messh config <action> <flags>
- base command just displays current active configuration & source
- actions:
  - init <flags>: initializes the config file
  - show <flags>: shows the current configuration (option to export, use template)

messh keys <action> <flags>
- base command just displays the stats of keys, like keys count, expiry status, default params, etc.
- actions:
  - generate <flags>: generates a new SSH key pair
  - list <flags>: lists the SSH keys
  - export <flags>: exports the SSH keys to a file
  - remove <flags>: removes the SSH keys
  - prune <flags>: prunes the dangling SSH keys

messh conn <action> <flags>
- base command is used to connect to an SSH host
- actions:
  - show <flags>: lists the available SSH connections
  - add <name> <flags>: adds a new SSH connection
  - edit <name> <flags>: edits an SSH connection
  - remove <name> <flags>: removes an SSH connection

messh history <flags>
- no subcommands, just flags that controls the output of the history
- flags:
  - quiet: suppresses the history output
  - file: specifies the path to the config file (dir or file - file should end with .log, other extens are overridden to .log)
  - ids (comma-separated): specifies the ids of the SSH connections to be included in the history
  - names (comma-separated): specifies the names of the SSH connections to be included in the history
  - match (regex): specifies the match pattern to be used for the SSH connections to be included in the history

messh doctor <flags>
  - no subcommands, just flags that controls the output of the watcher
  - flags:
    - quiet: suppresses the watcher output
    - file: specifies the path to the config file (dir or file - file should end with .log, other extens are overridden to .log)
  - notes:
    - complains if both file and quiet flags are set
    - warns if extn is not .log

### Management of themes
- messh theme is managed as a subcommand of the main messh command (or maybe a persistent global flag?)
- Invalid theme name will lead to fallback to default theme
- Theme are configured in the ~/.config/messh/messh-config.yaml file
- The theme files are stored in the ~/.config/messh/themes directory
- Other default themes are embedded in the binary

messh theme [subcommand] [flags]

subcommands:
- list
- add <name>
  - allows option to supply a theme.yaml file or url to theme.yaml file
  - themes are stored in the 


<!-- messh watcher <action> <flags>
- base command returns the status of the watcher service (uses karadinos/service)
- actions:
  - status <flags>: returns the status of the watcher service
  - start <flags>: starts the watcher service
  - stop <flags>: stops the watcher service
  - restart <flags>: restarts the watcher service
  - install <flags>: installs the watcher service
  - uninstall <flags>: uninstalls the watcher service -->