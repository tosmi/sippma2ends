# SIPPMA export

1. copy dump to local
2. Run `make database-ephemeral` in toplevel folder to bootstrap pgsql
2. pg_restore --dbname=sippma ~/Downloads/20240719212201_sippma.dump --host=localhost --port=5432 --username=sippma -W
3. export data is currently in ~/Documents/latido