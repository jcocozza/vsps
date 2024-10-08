# VSPS
vsps is a Very Simple Password Service.

It is a CLI tool for managing passwords.

Everything is stored locally.

Under the hood, it is just a simple file which can be edited directly.
Otherwise, it's just a bit of fluff on top.

## Account Language
vsps is just a file. A basic account looks like this:
```
<account_name>:
    username: <username>
    password: <password>
```

I wrote a simple parser for a "custom" file format which I call `al` (for account language).

Account Name, Username and Password are the only fields that are done by default.

Arbitrary data can also be associated with an account.
This is done in the CLI with the `-i`.
```
<account_name>:
    username: <username>
    password: <password>
    account_number: <account_number>
    foo: bar
    What_is_your_favorite_color: blue
```

This allows you to easily store other information with your account.
For example, you can add security questions, credit card information, and others.

Just be sure to keep to proper al format if you opt to edit it directly.
Things will break with improper formats/unexpected layouts.

vsps does not support any special nesting in the al file.

## Build
You can build the CLI tool with `go build` from the main directory.

(I still need to figure out a proper way to do this)

## Autocompletion
Using the `completion` command, you can generate an autocompletion script to make CLI usage easier.
You can also download the completion script from the releases where you got the binary from.

Disclaimer: I've only really tested this on zsh.

### Generate Completion
Run the following command:
```
vsps completion [shell-name]
```
Where shell name can be one of: `bash`, `fish`, `powershell`, or `zsh`

To be able to use the completions:
```
vsps completion [shell-name] > /path/to/completion/vsps.[shell-name]
```
Then see below for how to source the completion.
(Again, you can also get the completion script from the releases with the binary)

### Bash
One time:
```bash
source /path/to/completion/vsps.bash
```

To persist, add the following line to your `.bashrc` file.
```bash
source /path/to/completion/vsps.bash
```
### Zsh
To persist add the following 2 lines to your `.zshrc`
```zsh
autoload -U compinit; compinit
source /path/to/completion/vsps.zsh
```

### Fish
To persist, add the following line to your `config.fish`
```fish
source /path/to/completion/vsps.fish
```

### Powershell
Run the following command:
```ps1
$vspsCompletion = "$HOME\Documents\WindowsPowerShell\Microsoft.PowerShell_profile.ps1"
./vsps completion powershell > $vspsCompletion
```

## Encryption
vsps also offers basic encryption.
If you like, you can create encrypted accounts.
In the CLI, this is via the `-e` flag.

To encrypt accounts, you will need to provide a master password. This will NOT be persisted ANYWHERE.
Without it, you cannot access your encrypted passwords.
*DO NOT LOSE YOUR MASTER PASSWORD*

Note that encrypted accounts are kept in a separate file from regular accounts.

If you lose your master password, you can reset by deleting the encrypted file.
You will lose any accounts that are encrypted.
*DO NOT LOSE YOUR MASTER PASSWORD*

## To Do (not in any particular order)
- Implement Unit tests for testable stuff (maybe?)
- Clear clipboard after copying after 30-45 sec

## Versioning and Releases
For updating purposes all releases are of the form `vX.Y.Z`.

Each binary will be named in the following format: `vsps_*_<OS>_<ARCH>_<version>`.
For example: `vsps_cli_darwin_amd64_v0.0.1`.

Proper naming in the releases is important because the updater will check the github api for the latest release and will use these formats to update.

As of `v0.0.3`, the CLI has an updater.

## Caveats
Account names with spaces in them need to be surrounded with '' when accessing them in the terminal.
For example, if I want to get my account named foo bar then I need to do: `vsps -s 'foo bar'`.
This is because when reading in from args the cli thinks the entry ends when it encounters the space character.

'' should NOT need to be used when creating and updating accounts. Whenever a user is prompted for input, the use of a bufio.Reader means that it can read after it encounters a space character.

## Apple Users
Be aware that you will probably need allow vsps to open in your settings.
You will probably want to put the binary in `/usr/local/bin` or add it to your path.
