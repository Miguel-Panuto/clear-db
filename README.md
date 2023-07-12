# ClearDB

This is a project of a SQL Database in Golang, it is a side project of mine, imagining how to code a SQL database, it is in development and I do not know when can be used.

## How to run

Just run the main.go file.

## Commands

  * NEW DB dbName;
    * The command itself ins't case sensitive, you can run like this nEw Db, but the database name it is case sensitive.
    * This command will just craete the database.
  * LIST DBS;
    * This command it ins't case sensitive as well, you can use this LiST DbS
    * This command for now will just return a table with the database name and tables quantities.
  * USE db_name;
    * This command serves to use the database itself, for table creation for example.
  * CREATE TABLE table_name (name data_type validations);
    * Where it is name it is for the attribute name, data_type it is for the data_type itself, for now there is (id, int, string, boolean, json and float) for validations, there is only the required field.
  * LIST TABLES;
    * This command will list the table names and how much rows there is inside the table.