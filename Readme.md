# OTool
a simple tool for onmyoji

# Structure

```
.
├── Rakefile
├── build.sh
├── go.mod
├── go.sum
├── main.go
├── makefile
├── output
│        └── bin
│        ├── otool
│        ├── otool-arm64-0.0.1.tar.gz
│        ├── otool-linux64-0.0.1.tar.gz
│        ├── otool-mac64-0.0.1.tar.gz
│        ├── otool-win64-0.0.1.tar.gz
│        └── otool.exe
└── pkg
    ├── cmd
    │    ├── onmyoji_command.go
    │    ├── onmyoji_command_test.go
    │    └── root_command.go
    └── onmyoji
        ├── onmyoji.go
        └── onmyoji_test.go
```

# How to use
1. Find the right file which can run on your os.
2. If your os is windows, unzip the `tar.gz` file and just run otool.exe directly.
3. If your os is not windows, unzip the `tar.gz` file and `./otool`
4. You can add the flag `-d 'your path'` to set the download path. (default path is `./OnmyojiPictures`)