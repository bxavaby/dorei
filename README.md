<div align="center">

<img src="assets/help.png" width="746" alt="dorei help" />

───────────────────

![Deps](https://img.shields.io/badge/0-deps?style=plastic&label=deps&labelColor=000000&color=00ADD8)
[![Last Commit](https://img.shields.io/github/last-commit/bxavaby/rnm?style=plastic&labelColor=000000&color=00ADD8)](https://github.com/bxavaby/rnm/commits/main)

<br>

### dorei (奴隷) is a command-line tool that runs a **standalone daemon**, which uses a simplified interval format combined with absolute granularity (second), thence operating as a *lesser-cron* of sorts
#

<br>

<div align="left">

## Roadmap

- [x] Some basic, but COOL flags (needs expansion)
- [x] Config at ~/.config/dorei/dorei.conf
- [ ] System shell execution (no sandboxing)
- [ ] Simple error handling (log + continue)
- [ ] Concurrent timers (1 goroutine per task)
- [ ] Log notifications via matrix (notify)
- [ ] ...
- [ ] Nix-shell command wrapper ([nix] prefix in config)
- [ ] Hot-reload dorei.conf (avoid daemon restart)
- [ ] More notification services w/ notify
- [ ] Docker instead or in addition to nix-shell
- [ ] Prometheus metrics for monitoring
- [ ] Frontend??

<br>

## Safety Mechanisms

- [ ] Blacklist dangerous commands like ```rm -rf``` and ```:(){:|:&};:```
- [ ] Dry-run option, to print the scheduled commands
- [ ] Editor list matching, to prevent stupid errors
- [ ] Logging w/ timestamp via cli or matrix (if enabled)
- [ ] More to come ([Suggest Feature](../../issues))

</div>

<br>

───────────────────

*Hope this will serve you somehow.*

**[Report Bug](../../issues)** | **[Suggest Feature](../../issues)**

**MIT License © 2025 bxavaby**

</div>
