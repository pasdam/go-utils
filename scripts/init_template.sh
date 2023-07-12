#!/bin/sh

SCRIPT_DIR="$(cd "$( dirname "$0" )" >/dev/null && pwd )"
REPO_DIR="$( dirname "$SCRIPT_DIR" )"

TEMPLATE_MODULE=$(egrep "module " "$REPO_DIR/go.mod" | sed "s/module //")
echo "Template module: $TEMPLATE_MODULE"

CURRENT_MODULE=$(git config --get remote.origin.url | sed -E 's#(^.+@)|(\.git$)|(https?://)##g' | sed 's#:#/#')
echo "Current module: $CURRENT_MODULE"

sed -i '' "s|$TEMPLATE_MODULE|$CURRENT_MODULE|g" "$REPO_DIR/README.md"
sed -i '' "s|$TEMPLATE_MODULE|$CURRENT_MODULE|" "$REPO_DIR/go.mod"

rm "$REPO_DIR/main.go"
rm -rf "$REPO_DIR/pkg/app"

while true; do
  printf "Is this a library (y|yes / n|no)? "
  read response
  case $response in
  y | yes) DELETE_DOCKER=true; break ;;
  n | no) break ;;
  *) echo 1>&2 "Invalid answer: $response" ;;
  esac
done
if [ -n "$DELETE_DOCKER" ]; then
  rm "$REPO_DIR/Dockerfile"
  echo "Please remove the docker config form .github/dependabot.yml and from .github/workflows/ci.yaml"
fi

rm "$SCRIPT_DIR/init_template.sh"
