import pandas as pd
from sqlalchemy import create_engine
engine = create_engine('postgresql://postgres:1221425pg@localhost:5432/sahih')
df = pd.read_csv('data.csv', encoding='utf-8')[['nama doa asli', 'detail_doa', 'bantuan baca', 'artinya', 'panduan doa', 'keutamaan', 'sumber']]
df.rename(columns={
    'nama doa asli' : 'name', 
    'detail_doa' : 'doa', 
    'bantuan baca' : 'reading_assist', 
    'artinya': 'meaning', 
    'panduan doa': 'guidance',
    'keutamaan':'virtue',
    'sumber': 'source'
}, inplace=True)
df.insert(0, 'id', range(1, 1 + len(df)))
df.to_sql('doa_collection', engine, if_exists='fail', index=False, chunksize=1000)