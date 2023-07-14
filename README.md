# ClearDB

This is a project of a SQL Database in Golang, it is a side project of mine, imagining how to code a SQL database, it is in development and I do not know when can be used.

## How to run

Just run the main.go file.

## Commands

### New database

Command: NEW DB dbName;

* The command itself ins't case sensitive, you can run like this nEw Db, but the database name it is case sensitive.
* This command will just craete the database.

### List databases

Command: LIST DBS;

* This command it ins't case sensitive as well, you can use this LiST DbS
* This command for now will just return a table with the database name and tables quantities.
  
### Using database

Command: USE db_name;

* This command serves to use the database itself, for table creation for example.
  
### Creatig table

Command: CREATE TABLE table_name (name data_type validations);

* Where it is name it is for the attribute name, data_type it is for the data_type itself, for now there is (id, int, string, boolean, json and float) for validations, there is only the required field.
  
### Listing Tables

Command: LIST TABLES;

* This command will list the table names and how much rows there is inside the table.

### Inseting data

There is 2 ways to insert data, the short syntax and the normal.

Imagine the following table user, with fields:
[id id, name string, age int]

#### Short Syntax

Example:
INSERT user: 0, 'string value', 24;

#### Normal Syntax

Example:
INSERT user: {
  id:: 0,
  name:: 'string value',
  age:: 24
};

### Finding on database

In SQL we have the SELECT, and put the columns that you want and '*' for SELECT every column, but in clear-db it is very different, here we have find in command. The command follows:

Command: find in users;

With the command, it will pick from the table all columns.

But there is a way to select the columns that you want, with:

Command: find {name, age} in users;

This will return only the columns name and age.