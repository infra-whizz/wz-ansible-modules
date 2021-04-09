# Ansible Collection - whizz.embedded

Ansible collection to run binary/embedded modules.

# Adding Binary Modules

Extending is super-simple:

1. Symlink `your_module` in `action` directory to `lib/basemodule.py`
2. Pre-compile your binaries into `library` directory with the following schema `<name>-<platform>-<arch>`.
   E.g.: `your_module-linux-x86_64`
