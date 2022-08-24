# agelimit

Quickly test a file against an age limit in your shell scripts

## Examples

Check if a file is over 12 hours old:

```shell
agelimit -v 12h ./downloaded_file.json
```

```
agelimit: Error, limit breached! ./downloaded_file.json is over 12h old.
```

Age limit errors use exit code 1.

Check if a file is over 1 day old:

```shell
agelimit -v 1d ./ETL.log
```

```
agelimit: Success! ./ETL.log is less than 1 day old
```

Age limit successes use exit code 0.


## Program Usage

```shell
Usage: ./agelimit [OPTIONS] maxage file

  maxage      The maximum age for your file and a unit specifier. For example, 100s
              for 100 seconds, 20m for 20 minutes, 3h for 3 hours or 2d for 2 days
  file        The path to the file to be tested
  -s    Silence all output
  -v    Enable verbose output
  -version
        Display the version number and quit
```

## Real World Usage

So, how would you use this in a script?

Say you had a data pipeline that needs to be run infrequently, but the data can't be more than 4 hours old for business reasons.
You could add something like the following to your script:

(In all of the following examples we're checking an age limit of 4 hours for the file, `etl.log`)

```shell
#!/usr/bin/env bash

if ./agelimit 4h ./etl.log; then 
    echo "File not old yet"
    # Proceed as normal
else
    echo "File is now old"
    # regenerate the file
fi
```

Or in a simpler form:

```shell
#!/usr/bin/env bash

if ! ./agelimit 4h ./etl.log; then 
    echo "File is now old"
    # regenerate the file
fi
```

An example that uses the explicit exit code emitted by `agelimit`:

```shell
#!/usr/bin/env bash

if ./agelimit 4h ./etl.log; then 

if [[ $? -eq 0 ]]; then 
    echo "File not old yet"
    # Proceed as normal
else
    echo "File is now old"
    # regenerate the file
fi
```


## Exit codes

The standard exit codes used by agelimit are as follows:

* Age limit **not** breached - 0
* Age limit breached - 1
* Program error - 99

This last one includes things like input typos and specifying files that don't exist and so on.
**Always** check your `agelimit` commands run as expected _before_ you push them to production!
Test with the `-v` flag to see verbose output.


## Tips and tricks

* Use in a `Makefile` or `justfile` to regenerate files if they're older than business rules dictate
* Use in build scripts to make sure you don't hammer resources
* Use in data pipelines to only regenerate files when necessary
* Use with log files for ETL, to ensure the ETL itself is as fresh as required.

## Building agelimit

This project uses a `Makefile` to build the various target architectures.
You'll also need Go installed.
Run `make` to build them all.
All builds end up in a "builds" directory.

## License

This project is released under the MIT License. See LICENSE.md file for more info.
