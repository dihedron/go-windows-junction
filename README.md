# junctions

Manipulate junctions on Windows and symlinks on UNIX.

## Create

```bash
$> dist/linux/amd64/junction create --target=_tests/target/test.txt --link=_tests/mylink
```

where `target` is the filesystem object the link will point to, and `link` is the name of the link/junction. At the end of the day, it will look as follows:

```bash
$> ls -la _tests
total 12
drwxrwsr-x 3 dihedron users 4096 ago  1 15:28 .
drwxrwsr-x 9 dihedron users 4096 ago  1 15:26 ..
lrwxrwxrwx 1 dihedron users   55 ago  1 15:28 mylink -> /workspaces/junction/_tests/target/test.txt
drwxrwsr-x 2 dihedron users 4096 ago  1 15:02 target
```

## Remove

In order to remove a link:

```bash
$> dist/linux/amd64/junction remove --link=_tests/mylink
```

