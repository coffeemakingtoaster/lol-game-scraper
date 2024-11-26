# This could also be part of the bash script...but for my (and your) sanity we will do it like this
import sqlite3
import csv
import json

con = sqlite3.connect("./tmp.db")

cur = con.cursor()

cur.execute("SELECT * FROM matches")
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
        match_data = json.loads(match_data)
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

print(f"Exported {cnt} matches!")
