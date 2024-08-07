# AL
Al stands for account language.

It's basically a subset of yaml that disallows escape characters etc.
This means that usernames, passwords, etc can easily be stored/read without having to consider strange edge cases.
It only supports one layer of nesting.

A `.al` file looks like:
```
foo:
    bar: baz
    splat: rat
x:
    y: z
```

In sum, we only really have 2 "special" characters.
1. ":" the delimeter
2. "    ": the nester(4 spaces)
