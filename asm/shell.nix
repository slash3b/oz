with (import <nixpkgs> {});

mkShell {
    buildInputs = [
        nasm
    ];
}
