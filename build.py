import platform
import shutil
import os
import sys


def check_requirements() -> None:
    """
    Check main requirements
    """
    required = {
        "gcc": "GNU C Compiler",
        "gccgo": "GCC backend for Go compiler",
        "make": "GNU Make utility for directing complilation",
        "file": "libmagic based utility 'file'",
    }
    for binexe in required:
        if not shutil.which(binexe):
            sys.stderr.write("Error: {} is missing. Check your distribution which package to install.\n".format(required[binexe]))
            sys.exit(1)

def find_binaries(directory: str) -> list:
    """
    Get binaries from the current directory.
    """
    out = []
    for fname in os.listdir(directory):
        with os.popen("file {}".format(os.path.join(directory, fname))) as fh:
            ftype = list(filter(None, fh.read().strip().replace(",", "").split(" ")[1:]))
            if "ELF" in ftype:
                out.append(os.path.join(directory, fname))
    return out

def compile_modules(current_dir: str) -> None:
    """
    Compile modules via GNU Make
    """
    for root, dirs, files in os.walk(os.path.join(current_dir, "modules"), topdown=False):
        for name in files:
            if name == "Makefile":
                os.chdir(root)
                print("-" * 80)
                print("Compiling module at %s", root)
                for opt in ["gcc", "strip"]:
                    out = os.system("make {}".format(opt))
                    if out:
                        print("Error: module {} compilation failed. Check the output above.".format(root))
                        sys.exit(1)
    os.chdir(current_dir)

def remove_build(current_dir: str) -> None:
    """
    Remove any build that was done before, if exists.
    """
    build_path = os.path.join(current_dir, "build")
    if os.path.exists(build_path):
        shutil.rmtree(build_path, ignore_errors=True)

def distribute(current_dir: str) -> None:
    """
    Distribute compiled modules.
    """
    build_root = os.path.join(current_dir, "build", "bin", platform.system().lower(), platform.processor())
    os.makedirs(build_root, exist_ok=True)
    for root, dirs, files in os.walk(os.path.join(current_dir, "modules"), topdown=False):
        for name in files:
            if name == "Makefile":
                for bin_file in find_binaries(root):
                    dst = os.path.dirname(os.path.join(build_root, bin_file[len(os.path.join(current_dir, "modules")) + 1:]))
                    os.makedirs(dst, exist_ok=True)
                    shutil.move(bin_file, dst)
                    print("Moved", bin_file)

def main():
    """
    Build all binary modules
    """
    current_dir = os.path.dirname(os.path.abspath(sys.argv[0]))
    check_requirements()
    compile_modules(current_dir)
    remove_build(current_dir=current_dir)
    distribute(current_dir=current_dir)

    print("Done")

if __name__ == "__main__":
    main()
