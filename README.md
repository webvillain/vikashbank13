database operational commands for perfroming operations on database with user choice as per requirment


// Create Table Sql Statement
CREATE TABLE IF NOT EXISTS USERS(
    "Id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "FirstName" TEXT,
    "LastName" TEXT,
    "Email"TEXT
);

// GET ALL USERS
SELECT (Id,FirstName,LastName,Email) FROM USERS;

// GET USER BY ID
SELECT (Id,FirstName,LastName,Email) FROM USERS WHERE Id = ?;

// DELETE USER BY ID
DELETE FROM USERS WHERE Id = ?;


// UPDATE USER BY ID
UPDATE USERS
SET FirstName = ?,
    LastName = ?,
    Email = ?
WHERE
   Id = ?;


// CREATE NEW USER
INSERT INTO USERS (FirstName,LastName,Email)VALUES(?,?,?);