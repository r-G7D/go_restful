import sqlite3

with sqlite3.connect("go_restful.db") as db : 
    cursor = db.cursor