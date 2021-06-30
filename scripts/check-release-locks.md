# How to Update

If you are making a proto change and "make release-lock-status" is
failing, you will need to update the status file. First keep in mind
that this check is to prevent backwards-incompatible changes against
previous releases.

1. First ensure the changes you are making are not breaking backwards
   compatibility with any of these releases.

1. Run `REFRESH=true scripts/check-release-locks.sh`.

Lock files should not be updated. These should be the `proto.lock`
result of running `protolock init` in HEAD of the corresponding
release branch.
