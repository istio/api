## How to update

1. First ensure the changes you are making are not breaking backwards
   compatibility with any of these releases.

1. Edit `scripts/release-locks.sh` and comment out the line `rm status`.

1. Run `make release-lock-status`.

1. Copy the file `status` over the file
    `scripts/locks/release-<ver>/proto.lock.status`, corresponding to
   the release version.

1. `scripts/relese-locks.sh` can be reverted.

1. Include `scripts/locks/release-<ver>/proto.lock.status` in your PR
   and justification for the change.

Lock files should not be updated. These should be the `proto.lock`
result of running `protolock init` in HEAD of the corresponding
release branch.
