+++
title = "Updates"
weight = 4
type = "article"
+++

# Updates

If you have kddp installed and a new version comes out you don't have to reinstall it completely.
The `kddp update` command does this for you.

## Standard case

If you just run `kddp update` kddp will check if there is a newer version.
If this is the case, you will be asked if you want to install it.
If you answer yes ('j'), the update will begin.

kddp will then:
* download the latest version for your system,
* update the `kddp` executable,
* update the runtime and standard libraries,
* update the DDP language server (DDPLS).
  
If everything worked without errors, you should restart the terminal and be able to see the latest version with `kddp version`.

## Options

`kddp update` takes several optional command line options:

* `--wortreich`: tells you everything that happens during the update
* `--vergleiche_version`: checks whether a new version is available without installing it
* `--pre_release`: includes Github pre-releases in the search (explained below)
* `--jetzt`: immediately downloads and installs a new version, if available, without asking

## Versions

A DDP version consists of 4 parts:

`<Major>.<Minor>.<Patch>-<Addition>`

Example: 1.2.3-alpha

* Major: this number increases whenever a major feature is released, old DDP code may no longer work
* Minor: this number increases whenever a feature that is compatible with legacy code is released
* Patch: this number always increases when bug fixes or similar are released
* Addition: one of (-pre|-alpha|-beta). Signals test versions

## How it works

`kddp update` simply looks in the [Github-Releases](https://github.com/DDP-Projekt/Kompilierer/releases) from the compiler repo to see if there is a release whose tag is newer than the currently installed version .
If this is the case, it downloads the corresponding archive file and updates itself with it.

If the `--pre_release` option is specified, releases marked as "pre-release" will also be included in the search.