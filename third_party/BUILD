package(default_visibility = ["PUBLIC"])

go_get(
    name = "require",
    get = "github.com/stretchr/testify/require",
    deps = [
        ":assert",
        ":difflib",
        ":objx",
        ":spew",
        ":yaml",
    ],
)

go_get(
    name = "assert",
    get = "github.com/stretchr/testify/assert",
    deps = [
        ":difflib",
        ":objx",
        ":spew",
        ":yaml",
    ],
)

go_get(
    name = "difflib",
    get = "github.com/pmezard/go-difflib/difflib",
)

go_get(
    name = "objx",
    get = "github.com/stretchr/objx",
)

go_get(
    name = "spew",
    get = "github.com/davecgh/go-spew/spew",
)

go_get(
    name = "yaml",
    get = "gopkg.in/yaml.v3",
)
