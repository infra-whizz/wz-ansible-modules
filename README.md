# Overview

This is a collection of Ansible-compatible binary modules, used by
Whizz on those places there are no Python interpreter (or you just
want to be faster).

# Requirements (building)

Your system needs to have these installed:

- GNU Make
- GCC
- GCCGO
- File (libmagic)

# Structure

As of today, binary modules aren't strongest part of Ansible
ecosystem. Whizz is using them with the following structure:

```
  bin
  ╰─ <system>
     ╰─ <cpu>
        ╰─ <namespace>
	       ╰─ <module>
```

For example, a module `packaging.os.apt` for x86_64 Linux will appear
following:
```
  bin/linux/x86_64/packaging/os/apt
```

NOTE: Binary modules do not have `.py` file extension.

# Building

All modules are built with GCCGO instead, as it is needed for
cross-compilation, so the code has to be kept GCC compliant.

```
python3 build.py
```

This should either compile everything into `./build` directory or ask
you to install missing bits.
