# da (да)

A personal collection of cross-platform utilities written into a single busybox-style binary.

Based off https://github.com/rwxrob/bonzai.

Composable commands for your own bonzai are available in `pkg/`.

* fidi - implements https://github.com/filipVisko/fidi
* bon - echos out a go template for bonazi commands

## Install

```bash
go install github.com/filipVisko/da@latest
```

## Tab completion

Completion is provided by the bash builtin `complete`. Add the following to your shell's rc.

```bash
complete -C da da
```

## Roadmap

- Additional git helpers
- Local https://cht.sh
- Vim macros and helper utilities
- Replacements for standard shell applications to ensure consistent behavior
