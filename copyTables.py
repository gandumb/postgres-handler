import psycopg2 as psy
import pandas as pd
import codecs

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

#Crew
print("Doing crew operations...")

cursor.execute("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'crew')")
if (cursor.fetchone()[0]):
    cursor.execute("DROP TABLE Crew")

sql = '''CREATE TABLE Crew(
        titleID TEXT PRIMARY KEY UNIQUE,
        directors TEXT,
        writers TEXT
        );'''

cursor.execute(sql)

csvFile = 'cleanedData/crew.csv'

sql = '''COPY Crew(titleID,directors,writers)
        FROM STDIN
        DELIMITER '\t'
        CSV HEADER;'''

cursor.copy_expert(sql, open(csvFile, "r"))

#Principals
print("Doing principals operations...")

cursor.execute("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'principals')")
if (cursor.fetchone()[0]):
    cursor.execute("DROP TABLE Principals")

sql = '''CREATE TABLE Principals(
    titleID TEXT PRIMARY KEY UNIQUE,
    principalID TEXT,
    category TEXT,
    jobs TEXT,
    characters TEXT
    );'''

cursor.execute(sql)

csvFile = 'cleanedData/principals.csv'

sql = '''COPY Principals(titleID, principalID, category, jobs, characters)
        FROM STDIN
        DELIMITER '\t'
        CSV HEADER;'''

cursor.copy_expert(sql, open(csvFile, "r"))

#Users
print("Doing users operations...")

cursor.execute("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'users')")
if (cursor.fetchone()[0]):
    cursor.execute("DROP TABLE Users")

sql = '''CREATE TABLE Users(
        customerID INT PRIMARY KEY UNIQUE,
        rating DOUBLE PRECISION,
        timeDate DATE,
        titleID TEXT
        );'''

cursor.execute(sql)

csvFile = 'cleanedData/customer_ratings.csv'

sql = '''COPY Users(customerID, rating, timeDate, titleID)
        FROM STDIN
        DELIMITER '\t'
        CSV HEADER;'''

cursor.copy_expert(sql, open(csvFile, "r"))

#Media
print("Doing media operations...")

cursor.execute("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'media')")
if (cursor.fetchone()[0]):
    cursor.execute("DROP TABLE Media")

sql = '''CREATE TABLE Media(
        titleID TEXT PRIMARY KEY UNIQUE,
        mediaType TEXT,
        title TEXT,
        startYear INT,
        endYear INT,
        runtime INT,
        genre TEXT,
        releaseYear INT,
        rating DOUBLE PRECISION,
        voting INT
        );'''

cursor.execute(sql)

csvFile = 'cleanedData/titles.csv'

sql = '''COPY Media(titleID, mediaType, title, startYear, endYear, runtime, genre, releaseYear, rating, voting)
        FROM STDIN
        DELIMITER '\t'
        CSV HEADER;'''

cursor.copy_expert(sql, open(csvFile, "r"))

#Names
print("Doing names operations...")

cursor.execute("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'names')")
if (cursor.fetchone()[0]):
    cursor.execute("DROP TABLE Names")

sql = '''CREATE TABLE Names(
        principalID TEXT PRIMARY KEY UNIQUE,
        pname TEXT,
        birthYear INT,
        deathYear INT,
        profession TEXT
        );'''

cursor.execute(sql)

csvFile = 'cleanedData/names.csv'

sql = '''COPY Names(principalID, pname, birthYear, deathYear, profession)
        FROM STDIN
        DELIMITER '\t'
        CSV HEADER;'''

cursor.copy_expert(sql, open(csvFile, "r"))

print("Script ended")
