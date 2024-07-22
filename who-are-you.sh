URL="https://learn.zone01oujda.ma/assets/superhero/all.json"
NAME=$(curl -s $URL | jq -r ' .[] | select(.id==70) | .name')
echo "\"$NAME\""