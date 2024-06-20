# VSPS

vsps is a Very Simple Password Service.

It's a package for managing passwords via CLI or GUI.

Everything is stored locally.

Under the hood, it is just a yaml file which can be edited directly.
Otherwise, it's just a bit of fluff on top.

Just be sure to keep to proper yaml format if you opt to edit it directly.
Things will break with improper formats/unexpected layouts.

vsps does not support any special nesting in the yaml file.

## YAML

vsps is just a yaml file. A basic account looks like this: 
```yaml
<account_name>:
    username: <username>
    password: <password>
```

Account Name, Username and Password are the only fields that are done by default.

Arbitrary data can also be associated with an account.
This is done in the CLI with the `-i`.
```yaml
<account_name>:
    username: <username>
    password: <password>
    account_number: <account_number>
    foo: bar
    What_is_your_favorite_color: blue
```

This allows you to easily store other information with your account.
For example, you can add security questions, credit card information, and others.

## Build
You can build the CLI tool with `go build` from the main directory.

The GUI can be built with `go build` in the gui directory.

Alternatively, use the `go_compile` script to generate several binaries.

(I still need to figure out a proper way to do this)

## Encryption
vsps also offers basic encryption. 
If you like, you can create encrypted accounts.
In the CLI, this is via the `-e` flag. In the GUI, simply select the 'encrypted accounts' tab to manage your encrypted accounts.

To encrypt accounts, you will need to provide a master password. This will NOT be persisted ANYWHERE.
Without it, you cannot access your encrypted passwords. 
*DO NOT LOSE YOUR MASTER PASSWORD*

Note that encrypted accounts are kept in a separate file from regular accounts.

## To Do
- Implement Unit tests for testable stuff (maybe?)
- password length
- provide option to move accounts between encrypted and unencrypted.
- reset option in case of loss of master password -- simply remove the encrypted file
- Copy password on account creation
- Clear clipboard after copying after 30-45 sec
- Open button in GUI
