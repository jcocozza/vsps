# VSPS

vsps is a Very Simple Password Service.

It's a package for managing passwords via CLI or GUI.

Everything is stored locally.

Under the hood, it is just a yaml file which can be edited directly.
Otherwise, it's just a bit of fluff on top.

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
- provide option to move accounts between encrypted and decrypted.
- reset option in case of loss of master password
- GUI: copy on double click
- Copy password on account creation
