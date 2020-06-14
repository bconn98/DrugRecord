psql -c "CREATE DATABASE drugrecord;" postgres://postgres:Zoo123@localhost/?sslmode=disable
psql postgres://postgres:Zoo123@localhost/drugrecord?sslmode=disable

rem Setup db from moms backup
rem delete all rows
rem create new backup to pass with installer for setting up schema