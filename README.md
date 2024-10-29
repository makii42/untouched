untouched
=========

Stupid simple tool to run in a CI pipeline or similar to detect changed
files from code generators.

WHY
---

This might come in handy if you check in your generated code from `protoc`,
`openapi` and the like, and you want to make sure the generated code in your
repo did not get changed during your CI build.

What exactly
------------

`untouched` executes a git command and exits with a non-zero return code if it
finds files in the working tree were modified.  It will ignore a few [git 
status indicators][untouched-ignores].

As of now, it takes one flag `-diff`, which triggers an additional `git diff`
once it discoveres files that were touched, so it's obvious in your pipeline's
output what changes were made.

What's next
-----------

This tools is not perfect, but works for a few use cases for now. If it stays
useful can think of a number of flags to be added like:

* Ignore certain folders or file(s)
* Ignore certain file extensions
* Adjust the git status indicators ignored
* ... you get the whiff.

PRs are welcome.

[untouched-ignores]: main.go#L15-21
