-- v6: Add ctfdpasswordsent column to students table

ALTER TABLE students ADD COLUMN ctfdpasswordsent BOOLEAN NOT NULL DEFAULT FALSE;
