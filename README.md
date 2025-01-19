# HSOreCTF

Source code for the [HSOreCTF.tech](http://HSOreCTF.tech) website.

![healthcheck]( https://healthchecks.io/badge/01b7201d-dab8-4530-8754-58cd26/ITZPiwk3-2/HSOreCTF.svg)
[![Dependabot Updates](https://github.com/Ublius/HSOreCTF/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/Ublius/HSOreCTF/actions/workflows/dependabot/dependabot-updates)

## Development Workflow

Install Go and [gow](https://github.com/mitranim/gow). Then run:
```
LOG_CONSOLE=1 gow -e=yaml,go,html,css run ./cmd/hsorectf/
```
which will automatically restart the app whenever you make a change.

## License

The code is licensed under AGPLv3+. All of the content of the website (besides
the Colorado School of Mines logo) is Copyright (c) Mines Oresec 2023.
