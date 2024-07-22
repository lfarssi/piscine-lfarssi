find . -maxdepth 1 -not -name '.*' -not -name '.' -not -name '..' -printf '%A@ %f\n' |
sort -n |
cut -d' ' -f2- |
awk '{if (system("[ -d ""$0"" ]") == 0) print $0"/"; else print $0}' |
tr '\n' ',' |
sed 's/,$/\n/'
