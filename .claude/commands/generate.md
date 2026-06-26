# Regenerate Code

Regenerate Go code from Minecraft server reports and decompiled source.

## Instructions

First ensure the JSON reports exist for the target version (downloads the
vanilla server jar and runs its data generator):

```bash
cd vanilla_server_reports && ./generate.sh <version>
```

Then run the code generation:

```bash
cd pkg/data && go generate ./... && go fmt ./...
```

Verify the generated code compiles:

```bash
go build ./...
```

If generation fails, check that the JSON report files are present in
`vanilla_server_reports/generated/reports/` and that `../mcsrc/current/` holds
the decompiled source for the target version.
