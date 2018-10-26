# damage

damage deals with DMG files, which are compressed little bundles of evil.

In particular, it should be able to:

  * Give you information about a DMG file
    * Whether or not it has a EULA, for example
    * And extracting the text of said EULA
  * Attach/detach a DMG file
    * Bypassing said EULA if you tell it to

## Prerequisites

hdiutil must be in your `$PATH`.

That probably means being on macOS.
