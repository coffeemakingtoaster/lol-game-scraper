#!/usr/bin/env bash
echo "Bash version: ${BASH_VERSION}"
# Output merged database
MERGED_DB="tmp.db"

# Remove old merged database if it exists
if [[ -f "$MERGED_DB" ]]; then
  rm "$MERGED_DB"
fi

# Get a list of all .db files in the current directory
DB_FILES=(./in/*.db)

# Check if there are at least two .db files
if [[ ${#DB_FILES[@]} -lt 2 ]]; then
  echo "There must be at least two .db files in the current directory to merge."
  exit 1
fi

# Initialize the merged database with the schema from the first file
cp "${DB_FILES[0]}" "$MERGED_DB"

# Loop through all remaining .db files and merge them into the merged database
for DB in "${DB_FILES[@]:1}"; do
  echo "Merging $DB into $MERGED_DB..."
  sqlite3 "$MERGED_DB" <<EOF
ATTACH DATABASE '$DB' AS db_to_merge;
BEGIN TRANSACTION;
-- Dynamically copy all data from db_to_merge to main database
SELECT 'INSERT INTO ' || name || ' SELECT * FROM db_to_merge.' || name || ';'
FROM sqlite_master 
WHERE type = 'table' AND name NOT LIKE 'sqlite_%'
ORDER BY name;

COMMIT;
DETACH DATABASE db_to_merge;
EOF
done

echo "All databases merged into $MERGED_DB"

sqlite3 "$MERGED_DB" "DELETE FROM matches WHERE is_relevant IS FALSE;"

python3 ./csv_export.py

rm "$MERGED_DB"
