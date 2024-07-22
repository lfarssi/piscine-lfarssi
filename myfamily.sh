URL="https://learn.zone01oujda.ma/assets/superhero/all.json"

if [ -z "$HERO_ID" ]; then
  echo "HERO_ID is not set. Please export HERO_ID with the appropriate value."
  exit 1
fi

relatives=$(curl -s "$URL" | jq -r --arg id "$HERO_ID" '
  .[] | select(.id == ($id | tonumber)) | .connections.relatives
')

if [ -z "$relatives" ]; then
  echo "No relatives found for HERO_ID $HERO_ID"
else
  echo "$relatives"
fi