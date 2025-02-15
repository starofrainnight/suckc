# SuckC compiler

SuckC is a simple compiler of SuckC language which base of C&C++ language

## Usage

```bash
$ suckc [file]
```

If you want to see the verbose logs output:

```bash
$ GLOG_logtostderr=1 suckc [file]
```

If you want to debug the compiler:

```bash
$ suckc -d [file]
```

## Design

### Expressions

Expressions node will be evaluated to a string value
