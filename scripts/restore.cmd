@echo off
DATE=`date +%m.%_e.%Y`
psql -c "CREATE DATABASE drugrecord;" postgres://postgres:Zoo123@localhost/?sslmode=disable
psql postgres://postgres:Zoo123@localhost/drugrecord?sslmode=disable drugrecord < %1backup-$DATE.sql
