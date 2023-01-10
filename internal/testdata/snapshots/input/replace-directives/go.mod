module sg/replace-directives

go 1.19


require (
	github.com/dghubble/gologin v2.2.0+incompatible
)

replace (
	github.com/dghubble/gologin => github.com/sourcegraph/gologin v1.0.2-0.20181110030308-c6f1b62954d8
)
