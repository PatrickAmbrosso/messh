# messh
Cross-Platform SSH Manager


messh - lists the available SSH connections
messh list - lists the available SSH connections (alias to the above)
messh list

messh add <name> - adds a new SSH connection (base command does interactively)
messh add <name> <host> <user> <port> <key | also has option to create on the fly?? needed?>

messh remove <name> - removes an SSH connection by name with a --confirm flag

messh edit <name> - edits an SSH connection by name with a --confirm flag (base command does interactively)
messh edit <name> <host> <user> <port> <key>


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

### Management of Keys

messh keys
- Base command lists the various operations that can be performed with the keys command and its subcommands.
- It has flags to specify/override the path to which the keys are managed from (as an array of paths). This can be also managed from the config file at ~/.config/messh-config.yaml.


subcommands:
- generate
  - Generates a new SSH Key Pair (interactively or with flags)
  - It has options to customize the key type, key size, output path, passphrase, comment, force, and no-confirm flags.
  - It can also generate keys with unique names for unique identification (uses a adjective-noun combination).
  - Keys should be located in the ~/.ssh directory for the current user (unless specified otherwise).

messh keys generate (both interactive & flag-based supported)
messh keys list
messh keys export
messh keys remove
messh keys prune (removes expired, invalid, and revoked keys)

messh keygen - generates a new SSH key pair (base command does interactively)
messh keygen <the necessary things for the keygen command>

messh fmt - formats the SSH connections file

messh sessions - lists the past SSH sessions
messh sessions --export - exports the past SSH sessions to a file
messh sessions prune - clears the past SSH sessions




