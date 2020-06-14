@echo off
pg_dump postgres://postgres:Zoo123@localhost/drugrecord?sslmode=disable > %1\backup.sql