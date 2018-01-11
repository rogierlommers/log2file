## log2file
Sometimes you just want your (CLI) application to log certain information to a logfile. No more, no less. This package provides an easy way to achieve that.

## usage
Somewhere in your init() function, do this:

```
	if err := log2file.Activate(); err != nil {
		log.Fatal(err)
	}
```

from now on, you can do this:

```
  log2file.Write("line")
```

And this will be written to a file on disk, called `2018-01-11-T19-57-43.log`. If you want more control of the filename, you should initialize by using `log2file.ActivateForFile("/tmp/logfile.txt")`
