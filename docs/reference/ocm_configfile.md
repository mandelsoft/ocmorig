## ocm configfile &mdash; Configuration File

### Description


The command line client supports configuring by a given configuration file.
If existent by default the file <code>$HOME/.ocmconfig</code> will be read.
Using the option <code>--config</code> an alternative file can be specified.

The file format is yaml. It uses the same type mechanism used for all
kinds of typed specification in the ocm area. The file must have the type of
a configuration specification. Instead, the command line client supports
a generic configuration specification able to host a list of arbitrary configuration
specifications. The type for this spec is <code>generic.config.ocm.gardener.cloud/v1</code>.

The following configuration types are supported:

- <code>attributes.config.ocm.gardener.cloud</code>
  The config type <code>attributes.config.ocm.gardener.cloud</code> can be used to define a list
  of arbitrary attribute specifications:
  
  <pre>
      type: attributes.config.ocm.gardener.cloud
      attributes:
         &lt;name>: &lt;yaml defining the attribute>
         ...
  </pre>

- <code>credentials.config.ocm.gardener.cloud</code>
  The config type <code>credentials.config.ocm.gardener.cloud</code> can be used to define a list
  of arbitrary configuration specifications:
  
  <pre>
      type: credentials.config.ocm.gardener.cloud
      consumers:
        - identity:
            &lt;name>: &lt;value>
            ...
          credentials:
            - &lt;credential specification>
            ... credential chain
      repositories:
         - repository: &lt;repository specification>
           credentials:
            - &lt;credential specification>
            ... credential chain
      aliases:
         &lt;name>: 
           repository: &lt;repository specification>
           credentials:
            - &lt;credential specification>
            ... credential chain
  </pre>

- <code>generic.config.ocm.gardener.cloud</code>
  The config type <code>generic.config.ocm.gardener.cloud</code> can be used to define a list
  of arbitrary configuration specifications:
  
  <pre>
      type: generic.config.ocm.gardener.cloud
      configurations:
        - type: &lt;any config type>
          ...
        ...
  </pre>

- <code>keys.config.ocm.gardener.cloud</code>
  The config type <code>keys.config.ocm.gardener.cloud</code> can be used to define
  public and private keys. A key value might be given by one of the fields:
  - <code>path</code>: path of file with key data
  - <code>data</code>: base64 encoded binary data
  - <code>stringdata</code>: data a string parsed by key handler
  
  <pre>
      type: keys.config.ocm.gardener.cloud
      privateKeys:
         &lt;name>:
           path: &lt;file path>
         ...
      publicKeys:
         &lt;name>:
           data: &lt;base64 encoded key representation>
         ...
  </pre>

- <code>memory.credentials.config.ocm.gardener.cloud</code>
  The config type <code>memory.credentials.config.ocm.gardener.cloud</code> can be used to define a list
  of arbitrary credentials stored in a memory based credentials repository:
  
  <pre>
      type: memory.credentials.config.ocm.gardener.cloud
      repoName: default
      credentials:
        - credentialsName: ref
          reference:  # refer to a credential set stored in some other credential repository
            type: Credentials # this is a repo providing just one explicit credential set
            properties:
              username: mandelsoft
              password: specialsecret
        - credentialsName: direct
          credentials: # direct credential specification
              username: mandelsoft2
              password: specialsecret2
  </pre>

- <code>oci.config.ocm.gardener.cloud</code>
  The config type <code>oci.config.ocm.gardener.cloud</code> can be used to define
  OCI registry aliases:
  
  <pre>
      type: oci.config.ocm.gardener.cloud
      aliases:
         &lt;name>: &lt;OCI registry specification>
         ...
  </pre>

- <code>ocm.cmd.config.ocm.gardener.cloud</code>
  The config type <code>ocm.cmd.config.ocm.gardener.cloud</code> can be used to 
  configure predefined aliases for dedicated OCM repositories and 
  OCI registries.
  
  <pre>
     type: ocm.cmd.config.ocm.gardener.cloud
     ocmRepositories:
     &lt;name>: &lt;specification of OCM repository>
     ...
     ociRepositories:
     &lt;name>: &lt;specification of OCI registry>
     ...
  </pre>

- <code>scripts.ocm.config.ocm.gardener.cloud</code>
  The config type <code>scripts.ocm.config.ocm.gardener.cloud</code> can be used to define transfer scripts:
  
  <pre>
      type: scripts.ocm.config.ocm.gardener.cloud
      scripts:
        &lt;name>:
          path: &lt;>file path>
        &lt;other name>:
          script: &lt;>nested script as yaml>
  </pre>



### Examples

```

type: generic.config.ocm.gardener.cloud/v1
configurations:
  - type: credentials.config.ocm.gardener.cloud
    repositories:
      - repository:
          type: DockerConfig/v1
          dockerConfigFile: "~/.docker/config.json"
          propagateConsumerIdentity: true
   - type: attributes.config.ocm.gardener.cloud
     attributes:  # map of attribute settings
       compat: true
#  - type: scripts.ocm.config.ocm.gardener.cloud
#    scripts:
#      "default":
#         script:
#           process: true

```

### SEE ALSO

##### Parents

* [ocm](ocm.md)	 &mdash; Open Component Model command line client

