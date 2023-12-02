+++
title = "Installation"
weight = 2
+++

# Installation

This article describes how to install the compiler (kddp).

If you have already installed kddp and want to update it to a newer version, see [Updates](/Bedienungsanleitung/EN/Einstieg/Updates)

## 1. Download

Go to the [DDP compiler Github repo](https://github.com/DDP-Projekt/Kompilierer) and download the latest [release](https://github.com/DDP-Projekt/Kompilierer/releases) for your system.
For Windows the DDP-&lt;Version&gt;-&lt;Operating System&gt;-&lt;CPU Architecture&gt;.zip and for Linux DDP-&lt;Version&gt;-&lt;Operating System&gt;-&lt;CPU Architecture&gt;.tar.gz.
For Windows users who already have gcc installed, there are also -no-mingw variants.

Example file names:
- DDP-0.0.1-windows-amd64.zip
- DDP-0.0.1-windows-amd64-no-mingw.zip
- DDP-0.0.1-linux-amd64.tar.gz

## 2. Unzip

Unzip the downloaded archive to the location where you want to install DDP.
This location, i.e. the path and folder name, should not be changed after the installation!

On Windows, simply right-click on the archive and select "Extract All".
On Linux use the `$tar` command.

After extracting, you should see the correspondingly unzipped folder. It may be that it's nested (i.e. DDP-0.0.1-windows-amd64/DDP-0.0.1-windows-amd64). If this happened, move the folder to the correct location BEFORE INSTALLING.
The name of the folder doesn't matter, so feel free to rename it from DDP-0.0.1-windows-amd64 to just DDP.

## 3. Install

Go to the DDP directory that was just created and run the file `ddp-setup.exe` (or `ddp-setup` on Linux).
A console window should open that will guide you through the installation.

At the moment the setup is only in English to make it accessible to everyone, so don't be surprised.

### Express Installation

If you don't want to agree to all the features individually, you can simply run `ddp-setup` in the console with the `--force` flag.
This means that every setup question is automatically answered with “yes”.

### Custom installation

Without the `--force` flag, the setup will continue as normal

The setup will ask you questions (blue) and report warnings (yellow) and errors (red).
The questions are yes/no questions that can be answered with y/n (yes/no).
Warnings should be read, but may be ignored in some circumstances.

If an error occurs it is very likely that the entire installation has failed.
In this case, you should first read the error to see if you can change something obvious, otherwise feel free to create an [Issue](https://github.com/DDP-Projekt/Installer/issues) with all the necessary information.

The setup is designed so that by default you should answer every question with yes (y) to ensure that everything works as well as possible.

The setup will
* on Windows, unpack Mingw64 if necessary (but won't add it to the PATH),
* compare the kddp version with the (installed) GCC version and
* if necessary, recompile the DDP runtime/stdlib,
* set the environment variable DDPPATH (this MUST be set, otherwise the compiler will not work),
* add kddp to your PATH (or the entire DDP/bin folder, which will also contain DDPLS),
* install the DDP language server (DDPLS),
* install vscode-ddp (the DDP extension for Visual Studio Code). If VSCode is installed
* delete files that are no longer needed at the end to save space

A finished DDP installation can be between 20MB and around 1.1GB in size, depending on the operating system and whether mingw64 had to be installed separately or not.

## 4. Testing

To test the installation, see [First Program](/Bedienungsanleitung/EN/Einstieg/Erstes%20Programm)