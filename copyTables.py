import psycopg2 as psy
import pandas as pd

print("Starting script...")

#postgres://csce315904_4user:Helpme012@csce-315-db.engr.tamu.edu/csce315904_4db

conn = psy.connect(
    database = "csce315904_4db",
    user = "csce315904_4user",
    password = "Helpme012",
    host = "csce-315-db.engr.tamu.edu",
    port = "5432"
)

conn.autocommit = True
cursor = conn.cursor()

csvFile = 'cleanedData/crew.csv'

sql = '''COPY Crew(titleID,directors,writers)
FROM STDIN
DELIMITER '\t'
CSV HEADER;'''

cursor.copy_expert(sql, open(csvFile, "r"))


print("Script ended")
