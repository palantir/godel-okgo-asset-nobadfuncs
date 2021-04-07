<p align="right">
<a href="https://autorelease.general.dmz.palantir.tech/palantir/godel-okgo-asset-nobadfuncs"><img src="https://img.shields.io/badge/Perform%20an-Autorelease-success.svg" alt="Autorelease"></a>
</p>

godel-okgo-asset-nobadfuncs
===========================
godel-okgo-asset-nobadfuncs is an asset for the gödel [okgo plugin](https://github.com/palantir/okgo). It provides the functionality of the [go-nobadfuncs](https://github.com/palantir/go-nobadfuncs) check.

This check verifies that a set of blacklisted functions are not referenced.

You can delegate out to the underlying check config flags with: `./godelw run-check nobadfuncs -- --print-all $pkgPath`
