## Structure

main package: `oidc_copycat`
inside `oidc_copycat`:
    - directory.go => responsible for user management
        - defines User model
        - defines DirectoryReaderWriter interface
        - defines DirectoryService{} interface
    - applications.go / appregistry.go => responsible for managing OIDC applications
        - defines Application model
        - defines ApplicationReaderWriter interface
        - contains ApplicationRegistry{} service struct
    - sessions.go
        - defines Session model (for use by ory/oauth2 package )
        - defines SessionReaderWriter
        - defines SessionService{} struct
    - storage.go => implements the various reader/writer interfaces using BadgerDB

### Main ideas

* The `storage` implementation is injected into the different `*Service` structs
* When testing, the service structs get a mock impl of the `storage` interface
* HTTP handlers are defined in a `http` package
    * Handlers receive service(s) that were built on server startup
