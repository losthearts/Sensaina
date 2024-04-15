# Sensaina 繊細な

A tool made to convert sensitivities between games.

- Use `sensanai` for the tui application.
- Use `sensanai -gui` for the gui application.

## INSTALLATION.

- Clone the git repository and cd into it:
  ```sh
  git clone git@github.com:losthearts/Sensaina.git
  cd Sensaina
  ```
- Run `go install`.
- Make sure that your `go/bin` directory is added to `$PATH`.

## BUILD.

- Clone the git repository and cd into it:
  ```sh
  git clone git@github.com:losthearts/Sensaina.git
  cd Sensaina
  ```
- Run `go build sensaina.go`.
- This will generate an executable binary.

## USAGE.

```
Sensaina 繊細な
A tool to convert sensitivities between games
Currently CS:GO, Overwatch, and Valorant is supported.

Use without any flags to run the TUI

--gui || -g
        Use this flag to run the Sensaina GUI
===

sensanai --flag value
--from
        From Game.
--to
        To Game.
--sens
        Initial Sensitivity.
--idpi
        Initial DPI.
--fdpi
        Final DPI.
```

---

> Inspired from: [convert-sens](https://github.com/succumbs/convert-sens/).
