/* Welcome to dorei/main.go

dorei (also 奴隷, i.e. 'slave', in japanese) is
a command-line interval based scheduler in Go.
I like to think of it as a lesser-cron, which
pairs a simplified interval format with absolute
granularity (second), resulting in minimal syntax
complexity, but less flexibility. Cron and quartz
are still recommended, if one needs more control.

FEATURES:
---------
- minimal dependencies

- a dorei.conf at ~/.config/dorei/dorei.conf

- concurrent timers (1 lightweight goroutine per
task + ticker)

- system shell command exec, supporting any script
or pipeline allowed by the OS

- simple error handling (log and continue, do not
attempt to run again if it first fails)

- log notifications via matrix w/ notify, fully
configureable in dorei.conf (only matrix for now)

- some basic, but COOL flags:

> -a to add an entry to the config interactively
> -d/--dry-run to print all scheduled commands
without running
> -m/--matrix to toggle notifications on/off
using matrix as the default service (for now)
> -v to print the version and build date
> -h to print the help message

To add an interactive entry, dorei would have to
read from conf and check whether the field "EDITOR"
is populated or not. By requesting the editor's exec
command on first run, dorei is able to store it in
the .env file. Of course, the user input would be
checked against a list, consisting of:

nano; vi; vim; neovim; nvim; micro; etc

This list contains the exec commands most people
use. If dorei does not find a match, it assumes
that the user entered the wrong command.
This is a safety mechanism, however it can be
bypassed by editing the file directly at
~/.config/dorei/dorei.conf

dorei is a standalone daemon: it runs until you
kill it. Be careful with the commands you add to
dorei.conf, as some could have unintended effects.

SAFETY MECHANISMS:
------------------
- blacklist dangerous commands like 'rm -rf',
'shutdown', ':(){:|:&};:', etc, by scanning them
on config load and warning/rejecting (can still
be bypassed with more complex commands)
- dry-run flag
- editor list matching
- every command is logged w/ timestamp via cli
or matrix, if the service is enabled
- more to come (check roadmap below)

>0.1.0 ROADMAP (to be only):
----------------------------
+ implement nix-shell internally, to wrap every
scheduled command using a general shell.nix

A NIX_SHELL boolean flag would be added to .env,
toggleable using the '-n/--nix' flag. This adds
an optional safety mechanism, since nix-shell
provides isolated and hermetic environments,
thus preventing unintended changes to the user's
system. Caveats: performance overhead, nix must
be installed, environment limitations. Could
instead be embedded in each entry in dorei.conf
(e.g. 60s: [nix] /home/user/scripts/backup.sh),
although that would defeat the broader security
purpose established above.

+ hot-reload dorei.conf (avoids restarting the
daemon)

+ add more notify optional services, apart from
matrix,

+ containerize runtimes with docker instead or
in addition to nix-shell

+ expose prometheus metrics for usage monitoring

+ frontend??
----------------------------

Hope this will serve you somehow.
Don't whip it, though :(

Licensed under the MIT License.
Copyright (c) 2025 bxavaby

Repo: https://github.com/bxavaby/dorei
*/

package main

import (
	"github.com/bxavaby/dorei/cli"
	"os"
)

func main() {
	os.Exit(cli.Run())
}
