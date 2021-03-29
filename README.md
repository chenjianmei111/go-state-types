# Filecoin state types
[![CircleCI](https://circleci.com/gh/chenjianmei111/go-state-types.svg?style=svg)](https://circleci.com/gh/chenjianmei111/go-state-types)
[![codecov](https://codecov.io/gh/chenjianmei111/go-state-types/branch/master/graph/badge.svg)](https://codecov.io/gh/chenjianmei111/go-state-types)

This repository contains primitive and low level types used in the Filecoin blockchain state representation.

These are primarily intended for use by [Actors](https://github.com/chenjianmei111/specs-actors) and other
modules that read chain state directly.

## Versioning

Blockchain state versioning does not naturally align with common semantic versioning conventions.

Any change in behaviour, including repairing any error that may have affected blockchain evaluation,
must be released in a major version change. We intend that to be a rare event for the contents of 
this repository.

## License
This repository is dual-licensed under Apache 2.0 and MIT terms.

Copyright 2020. Protocol Labs, Inc.
