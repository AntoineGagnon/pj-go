# What is this?

This is a simple tool that returns the first folder that matches the String entered as the first argument.

The script expects all your code projects to be in the ~/Code folder.

To use it in your terminal, add the following alias:

`alias pj='f() { cd `pj-go $1` };f'`