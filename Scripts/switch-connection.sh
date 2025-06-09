#!/bin/bash

PROFILE=$1
if [[ "$PROFILE" != "local" && "$PROFILE" != "test" ]]; then
  echo "Usage: $0 [local|test]"
  exit 1
fi

SECRETS_FILES=(
  "$HOME/.microsoft/usersecrets/ed4421dd-2538-4137-b9d0-eb19bf67dd36/secrets.json"
  "$HOME/.microsoft/usersecrets/6a496182-b158-4603-b355-de91cb9c41e1/secrets.json"
  "$HOME/.microsoft/usersecrets/F4C2A280-B7A3-4D05-A76A-244E13325587/secrets.json"
)

for SECRETS_PATH in "${SECRETS_FILES[@]}"; do
  if [[ -f "$SECRETS_PATH" ]]; then
    TMP_PATH="${SECRETS_PATH}.tmp"
    jq --arg profile "$PROFILE" \
      '.["Database:ConnectionString"] = .["Database:ConnectionString:" + ($profile | ascii_upcase[0:1] + $profile[1:])]' \
      "$SECRETS_PATH" > "$TMP_PATH" && mv "$TMP_PATH" "$SECRETS_PATH"
    echo "✔ Switched $SECRETS_PATH to '$PROFILE'"
  else
    echo "⚠ File not found: $SECRETS_PATH"
  fi
done

SECRETS_PATH="$HOME/.microsoft/usersecrets/F4C2A280-B7A3-4D05-A76A-244E13325587/secrets.json"
TMP_PATH="${SECRETS_PATH}.tmp"

jq --arg profile "$PROFILE" '
  .["ConnectionStrings:Database:Admin:ConnectionString"] =
    .["ConnectionStrings:Database:Admin:ConnectionString:" + ($profile | ascii_upcase[0:1] + $profile[1:])] |
  .["ConnectionStrings:Database:Platform:ConnectionString"] =
    .["ConnectionStrings:Database:Platform:ConnectionString:" + ($profile | ascii_upcase[0:1] + $profile[1:])]
' "$SECRETS_PATH" > "$TMP_PATH" && mv "$TMP_PATH" "$SECRETS_PATH"

echo "✔ Switched $SECRETS_PATH to '$PROFILE'"