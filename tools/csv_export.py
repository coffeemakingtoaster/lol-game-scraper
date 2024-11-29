import sqlite3
import csv
import json
import os
import shutil

TMP_DB = "./tmp.db"

# merge dbs
try:
    os.remove(TMP_DB)
except:
    pass

db_files = []

for file in os.listdir("./in"):
    if file.endswith(".db"):
        db_files.append(os.path.join("./in", file))

if len(db_files) == 0:
    print("no dbs")
    exit(1)

shutil.copyfile(db_files[0], TMP_DB)
print(f"{db_files[0]} as base")

con = sqlite3.connect(TMP_DB)

cur = con.cursor()

for i in range(1, len(db_files)):
    file = db_files[i]
    print(f"Merging {file}")
    copy_conn = sqlite3.connect(file)
    copy_cursor = copy_conn.cursor()
    copy_cursor.execute("SELECT * FROM matches WHERE is_relevant IS TRUE")
    for row in copy_cursor:
        # Construct an insert query based on the number of columns in the row
        placeholders = ", ".join("?" for _ in row)
        insert_query = f"INSERT OR IGNORE INTO matches VALUES ({placeholders})"
        cur.execute(insert_query, row)

    # Commit changes to Database B
    con.commit()

cur.execute("SELECT * FROM matches WHERE is_relevant IS TRUE")
match_data_dict = {
    "matchId": "",
    "winner_BOTTOM": "",
    "winner_JUNGLE": "",
    "winner_MIDDLE": "",
    "winner_TOP": "",
    "winner_UTILITY": "",
    "loser_BOTTOM": "",
    "loser_JUNGLE": "",
    "loser_MIDDLE": "",
    "loser_TOP": "",
    "loser_UTILITY": "",
}
wanted_keys = list(match_data_dict.keys())
cnt = 0
with open("./export/export.csv", "w", newline="") as csvfile:
    w = csv.DictWriter(csvfile, match_data_dict.keys())
    w.writeheader()
    for row in cur:
        match_id, match_data, _ = row
        try:
            match_data = json.loads(match_data)
        except:
            print("Could not parse match data for one game.")
            continue
        match_data_dict["matchId"] = match_id
        for player in match_data["info"]["participants"]:
            status = "winner" if player["win"] else "loser"
            champion = player["championName"]
            role = player["individualPosition"]
            key = f"{status}_{role}"
            match_data_dict[key] = champion
        if len(match_data_dict.keys()) > len(wanted_keys):
            for key in list(match_data_dict.keys()):
                if key not in wanted_keys:
                    del match_data_dict[key]
        w.writerow(match_data_dict)
        cnt += 1

os.remove(TMP_DB)
print(f"Exported {cnt} matches!")
