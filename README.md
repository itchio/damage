# damage

[![Build Status](https://travis-ci.org/itchio/damage.svg?branch=master)](https://travis-ci.org/itchio/damage)
[![GoDoc](https://godoc.org/github.com/itchio/damage?status.svg)](https://godoc.org/github.com/itchio/damage)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/itchio/damage/blob/master/LICENSE)
[![No Maintenance Intended](http://unmaintained.tech/badge.svg)](http://unmaintained.tech/)

damage deals with DMG files, which are compressed little bundles of evil.

In particular, it should be able to:

  * Give you information about a DMG file
    * Whether or not it has a SLA (Software License Agreement), for example
    * And extracting the text of said SLA
  * Attach/detach a DMG file
    * Bypassing said SLA if you tell it to

## Prerequisites

hdiutil must be in your `$PATH`.

That probably means being on macOS.
