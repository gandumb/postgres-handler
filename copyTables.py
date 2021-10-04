print("Starting script...")

from sqlalchemy import create_engine
import pandas as pd


engine = create_engine('postgres://csce315904_4user:Helpme012@csce-315-db.engr.tamu.edu/csce315904_4db')
df = pd.read_csv('rawData/crew.csv')
df.to_sql('Crew', engine, if_exists='replace')


print("Script ended")
