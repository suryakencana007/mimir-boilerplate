#!/bin/sh
set -eu

touch go.mod

CONTENT=$(cat <<-EOD
module {{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}

go 1.15

require (
github.com/golang/mock v1.3.1
)
EOD
)

echo "$CONTENT" > go.mod
