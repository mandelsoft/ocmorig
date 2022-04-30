## ocm ocm versions show

show dedicated versions (semver compliant)

### Synopsis

```
ocm ocm versions show [<options>] <component> {<version pattern>}
```

### Options

```
  -h, --help          help for show
  -l, --latest        show only latest version
  -r, --repo string   repository name or spec
  -s, --semantic      show semantic version
```

### Description


Match versions of a component against some patterns.

If the <code>--repo</code> option is specified, the given names are interpreted
relative to the specified repository using the syntax

<center><code>&lt;component>[:&lt;version>]</code></center>

If no <code>--repo</code> option is specified the given names are interpreted 
as located OCM component version references:

<center><code>[&lt;repo type>::]&lt;host>[:&lt;port>][/&lt;base path>]//&lt;component>[:&lt;version>]</code></center>

Additionally there is a variant to denote common transport archives
and general repository specifications

<center><code>[&lt;repo type>::]&lt;filepath>|&lt;spec json>[//&lt;component>[:&lt;version>]]</code></center>

The <code>--repo</code> option takes an OCM repository specification:

<center><code>[&lt;repo type>::]&lt;configured name>|&lt;file path>|&lt;spec json></code></center>

For the *Common Transport Format* the types <code>directory</code>,
<code>tar</code> or <code>tgz</code> is possible.

Using the JSON variant any repository type supported by the 
linked library can be used:

Dedicated OCM repository types:
- `ComponentArchive`

OCI Repository types (using standard component repository to OCI mapping):
- `ArtefactSet`
- `CommonTransportFormat`
- `DockerDaemon`
- `Empty`
- `OCIRegistry`
- `oci`


### Examples

```

$ ocm transfer ctf ctf.tgz ghcr.io/mandelsoft/components

```

### SEE ALSO

##### Parents

* [ocm ocm versions](ocm_ocm_versions.md)	 - Commands acting on component version names
* [ocm ocm](ocm_ocm.md)	 - Dedicated command flavors for the Open Component Model
* [ocm](ocm.md)	 - ocm command line client
