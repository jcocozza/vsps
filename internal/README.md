# AL
Al stands for account language.

It's basically a subset of yaml that doesn't care about anything except newlines, and the delimeter ":".
This means that usernames, passwords, etc can easily be stored/read without having to consider strange edge cases or worrying about escape character matching.
At this time, empty newlines will cause an error and are not supported.

A well structured `.al` file looks like:
```
my_acct1:
    username: user
    password: pass
my_acct2:
    username: foo@foo.com
    extra-info: foo-bar
```

Note that only only one layer of nesting is supported. Nesing is supported for either 2-spaces or 4-spaces, but not tabs.
The following is _invalid_.
```
a:
    b:
        c: 123
```

In sum, we only really have a few "special" characters.
1. ": " the delimeter
2. "  ", "    ": the nester(2 or 4 spaces)
3. "\n" newline

From a technical perspective, al also treats "username" and "password" as special string types which are identified by the tokenizer.
While accounts are not required to have either of these, they are common enough that I think it makes sense to make them special.

4. "username"
5. "password"
