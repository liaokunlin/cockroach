// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 921
	`ALTER`: {
		//line sql.y: 922
		Category: hGroup,
		//line sql.y: 923
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER DATABASE
`,
	},
	//line sql.y: 931
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 932
		Category: hDDL,
		//line sql.y: 933
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 955
		SeeAlso: `https://www.cockroachlabs.com/docs/alter-table.html
`,
	},
	//line sql.y: 966
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 967
		Category: hDDL,
		//line sql.y: 968
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 970
		SeeAlso: `https://www.cockroachlabs.com/docs/alter-view.html
`,
	},
	//line sql.y: 977
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 978
		Category: hDDL,
		//line sql.y: 979
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 981
		SeeAlso: `https://www.cockroachlabs.com/docs/alter-database.html
`,
	},
	//line sql.y: 988
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 989
		Category: hDDL,
		//line sql.y: 990
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 998
		SeeAlso: `https://www.cockroachlabs.com/docs/alter-index.html
`,
	},
	//line sql.y: 1208
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1209
		Category: hCCL,
		//line sql.y: 1210
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1227
		SeeAlso: `RESTORE, https://www.cockroachlabs.com/docs/backup.html
`,
	},
	//line sql.y: 1235
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1236
		Category: hCCL,
		//line sql.y: 1237
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1253
		SeeAlso: `BACKUP, https://www.cockroachlabs.com/docs/restore.html
`,
	},
	//line sql.y: 1267
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1268
		Category: hCCL,
		//line sql.y: 1269
		Text: `
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   comma = '...'          [CSV-specific]
   comment = '...'        [CSV-specific]
   nullif = '...'         [CSV-specific]

`,
		//line sql.y: 1287
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1382
	`CANCEL`: {
		//line sql.y: 1383
		Category: hGroup,
		//line sql.y: 1384
		Text: `CANCEL JOB, CANCEL QUERY
`,
	},
	//line sql.y: 1390
	`CANCEL JOB`: {
		ShortDescription: `cancel a background job`,
		//line sql.y: 1391
		Category: hMisc,
		//line sql.y: 1392
		Text: `CANCEL JOB <jobid>
`,
		//line sql.y: 1393
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOB
`,
	},
	//line sql.y: 1401
	`CANCEL QUERY`: {
		ShortDescription: `cancel a running query`,
		//line sql.y: 1402
		Category: hMisc,
		//line sql.y: 1403
		Text: `CANCEL QUERY <queryid>
`,
		//line sql.y: 1404
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1412
	`CREATE`: {
		//line sql.y: 1413
		Category: hGroup,
		//line sql.y: 1414
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW
`,
	},
	//line sql.y: 1428
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1429
		Category: hDML,
		//line sql.y: 1430
		Text: `DELETE FROM <tablename> [WHERE <expr>] [RETURNING <exprs...>]
`,
		//line sql.y: 1431
		SeeAlso: `https://www.cockroachlabs.com/docs/delete.html
`,
	},
	//line sql.y: 1439
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 1440
		Category: hCfg,
		//line sql.y: 1441
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 1453
	`DROP`: {
		//line sql.y: 1454
		Category: hGroup,
		//line sql.y: 1455
		Text: `DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP USER
`,
	},
	//line sql.y: 1464
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 1465
		Category: hDDL,
		//line sql.y: 1466
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1467
		SeeAlso: `https://www.cockroachlabs.com/docs/drop-index.html
`,
	},
	//line sql.y: 1479
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 1480
		Category: hDDL,
		//line sql.y: 1481
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1482
		SeeAlso: `https://www.cockroachlabs.com/docs/drop-table.html
`,
	},
	//line sql.y: 1494
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 1495
		Category: hDDL,
		//line sql.y: 1496
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 1497
		SeeAlso: `https://www.cockroachlabs.com/docs/drop-index.html
`,
	},
	//line sql.y: 1517
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 1518
		Category: hDDL,
		//line sql.y: 1519
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 1520
		SeeAlso: `https://www.cockroachlabs.com/docs/drop-database.html
`,
	},
	//line sql.y: 1540
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 1541
		Category: hPriv,
		//line sql.y: 1542
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 1543
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 1585
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 1586
		Category: hMisc,
		//line sql.y: 1587
		Text: `
EXPLAIN <statement>
EXPLAIN [( [PLAN ,] <planoptions...> )] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, HELP, EXPLAIN, EXECUTE

Plan options:
    TYPES, EXPRS, METADATA, QUALIFY, INDENT, VERBOSE, DIST_SQL

`,
		//line sql.y: 1598
		SeeAlso: `https://www.cockroachlabs.com/docs/explain.html
`,
	},
	//line sql.y: 1654
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 1655
		Category: hMisc,
		//line sql.y: 1656
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 1657
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1679
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 1680
		Category: hMisc,
		//line sql.y: 1681
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 1682
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 1705
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 1706
		Category: hMisc,
		//line sql.y: 1707
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 1708
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 1728
	`GRANT`: {
		ShortDescription: `define access privileges`,
		//line sql.y: 1729
		Category: hPriv,
		//line sql.y: 1730
		Text: `
GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1740
		SeeAlso: `REVOKE, https://www.cockroachlabs.com/docs/grant.html
`,
	},
	//line sql.y: 1748
	`REVOKE`: {
		ShortDescription: `remove access privileges`,
		//line sql.y: 1749
		Category: hPriv,
		//line sql.y: 1750
		Text: `
REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 1760
		SeeAlso: `GRANT, https://www.cockroachlabs.com/docs/revoke.html
`,
	},
	//line sql.y: 1843
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 1844
		Category: hCfg,
		//line sql.y: 1845
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 1846
		SeeAlso: `https://www.cockroachlabs.com/docs/set-vars.html
`,
	},
	//line sql.y: 1876
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 1877
		Category: hCfg,
		//line sql.y: 1878
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 1879
		SeeAlso: `SHOW CLUSTER SETTING, SET SESSION,
https://www.cockroachlabs.com/docs/cluster-settings.html
`,
	},
	//line sql.y: 1897
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 1898
		Category: hCfg,
		//line sql.y: 1899
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }

`,
		//line sql.y: 1904
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
https://www.cockroachlabs.com/docs/set-vars.html
`,
	},
	//line sql.y: 1921
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 1922
		Category: hTxn,
		//line sql.y: 1923
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 1930
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
https://www.cockroachlabs.com/docs/set-transaction.html
`,
	},
	//line sql.y: 2105
	`SHOW`: {
		//line sql.y: 2106
		Category: hGroup,
		//line sql.y: 2107
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE TABLE, SHOW CREATE VIEW, SHOW USERS, SHOW TRANSACTION, SHOW BACKUP,
SHOW JOBS, SHOW QUERIES, SHOW SESSIONS, SHOW TRACE
`,
	},
	//line sql.y: 2132
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2133
		Category: hCfg,
		//line sql.y: 2134
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2135
		SeeAlso: `https://www.cockroachlabs.com/docs/show-vars.html
`,
	},
	//line sql.y: 2156
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2157
		Category: hCCL,
		//line sql.y: 2158
		Text: `SHOW BACKUP <location>
`,
		//line sql.y: 2159
		SeeAlso: `https://www.cockroachlabs.com/docs/show-backup.html
`,
	},
	//line sql.y: 2167
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2168
		Category: hCfg,
		//line sql.y: 2169
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2172
		SeeAlso: `https://www.cockroachlabs.com/docs/cluster-settings.html
`,
	},
	//line sql.y: 2189
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2190
		Category: hDDL,
		//line sql.y: 2191
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2192
		SeeAlso: `https://www.cockroachlabs.com/docs/show-columns.html
`,
	},
	//line sql.y: 2200
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2201
		Category: hDDL,
		//line sql.y: 2202
		Text: `SHOW DATABASES
`,
		//line sql.y: 2203
		SeeAlso: `https://www.cockroachlabs.com/docs/show-databases.html
`,
	},
	//line sql.y: 2211
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2212
		Category: hPriv,
		//line sql.y: 2213
		Text: `SHOW GRANTS [ON <targets...>] [FOR <users...>]
`,
		//line sql.y: 2214
		SeeAlso: `https://www.cockroachlabs.com/docs/show-grants.html
`,
	},
	//line sql.y: 2222
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2223
		Category: hDDL,
		//line sql.y: 2224
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 2225
		SeeAlso: `https://www.cockroachlabs.com/docs/show-indexes.html
`,
	},
	//line sql.y: 2243
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 2244
		Category: hDDL,
		//line sql.y: 2245
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 2246
		SeeAlso: `https://www.cockroachlabs.com/docs/show-constraints.html
`,
	},
	//line sql.y: 2259
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 2260
		Category: hMisc,
		//line sql.y: 2261
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 2262
		SeeAlso: `CANCEL QUERY
`,
	},
	//line sql.y: 2278
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 2279
		Category: hMisc,
		//line sql.y: 2280
		Text: `SHOW JOBS
`,
		//line sql.y: 2281
		SeeAlso: `CANCEL JOB, PAUSE JOB, RESUME JOB
`,
	},
	//line sql.y: 2289
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 2290
		Category: hMisc,
		//line sql.y: 2291
		Text: `
SHOW [KV] TRACE FOR SESSION
SHOW [KV] TRACE FOR <statement>
`,
		//line sql.y: 2294
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 2315
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 2316
		Category: hMisc,
		//line sql.y: 2317
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
	},
	//line sql.y: 2333
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 2334
		Category: hDDL,
		//line sql.y: 2335
		Text: `SHOW TABLES [FROM <databasename>]
`,
		//line sql.y: 2336
		SeeAlso: `https://www.cockroachlabs.com/docs/show-tables.html
`,
	},
	//line sql.y: 2348
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 2349
		Category: hCfg,
		//line sql.y: 2350
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 2351
		SeeAlso: `https://www.cockroachlabs.com/docs/show-transaction.html
`,
	},
	//line sql.y: 2370
	`SHOW CREATE TABLE`: {
		ShortDescription: `display the CREATE TABLE statement for a table`,
		//line sql.y: 2371
		Category: hDDL,
		//line sql.y: 2372
		Text: `SHOW CREATE TABLE <tablename>
`,
		//line sql.y: 2373
		SeeAlso: `https://www.cockroachlabs.com/docs/show-create-table.html
`,
	},
	//line sql.y: 2381
	`SHOW CREATE VIEW`: {
		ShortDescription: `display the CREATE VIEW statement for a view`,
		//line sql.y: 2382
		Category: hDDL,
		//line sql.y: 2383
		Text: `SHOW CREATE VIEW <viewname>
`,
		//line sql.y: 2384
		SeeAlso: `https://www.cockroachlabs.com/docs/show-create-view.html
`,
	},
	//line sql.y: 2392
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 2393
		Category: hPriv,
		//line sql.y: 2394
		Text: `SHOW USERS
`,
		//line sql.y: 2395
		SeeAlso: `CREATE USER, DROP USER, https://www.cockroachlabs.com/docs/show-users.html
`,
	},
	//line sql.y: 2447
	`PAUSE JOB`: {
		ShortDescription: `pause a background job`,
		//line sql.y: 2448
		Category: hMisc,
		//line sql.y: 2449
		Text: `PAUSE JOB <jobid>
`,
		//line sql.y: 2450
		SeeAlso: `SHOW JOBS, CANCEL JOB, RESUME JOB
`,
	},
	//line sql.y: 2458
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 2459
		Category: hDDL,
		//line sql.y: 2460
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 2486
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE TABLE,
https://www.cockroachlabs.com/docs/create-table.html
https://www.cockroachlabs.com/docs/create-table-as.html
`,
	},
	//line sql.y: 2820
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 2821
		Category: hDML,
		//line sql.y: 2822
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2823
		SeeAlso: `https://www.cockroachlabs.com/docs/truncate.html
`,
	},
	//line sql.y: 2831
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 2832
		Category: hPriv,
		//line sql.y: 2833
		Text: `CREATE USER <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 2834
		SeeAlso: `DROP USER, SHOW USERS, https://www.cockroachlabs.com/docs/create-user.html
`,
	},
	//line sql.y: 2852
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 2853
		Category: hDDL,
		//line sql.y: 2854
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 2855
		SeeAlso: `CREATE TABLE, SHOW CREATE VIEW, https://www.cockroachlabs.com/docs/create-view.html
`,
	},
	//line sql.y: 2869
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 2870
		Category: hDDL,
		//line sql.y: 2871
		Text: `
CREATE [UNIQUE] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 2879
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE INDEX,
https://www.cockroachlabs.com/docs/create-index.html
`,
	},
	//line sql.y: 3018
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 3019
		Category: hTxn,
		//line sql.y: 3020
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 3021
		SeeAlso: `SAVEPOINT, https://www.cockroachlabs.com/docs/savepoint.html
`,
	},
	//line sql.y: 3029
	`RESUME JOB`: {
		ShortDescription: `resume a background job`,
		//line sql.y: 3030
		Category: hMisc,
		//line sql.y: 3031
		Text: `RESUME JOB <jobid>
`,
		//line sql.y: 3032
		SeeAlso: `SHOW JOBS, CANCEL JOB, PAUSE JOB
`,
	},
	//line sql.y: 3040
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 3041
		Category: hTxn,
		//line sql.y: 3042
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 3043
		SeeAlso: `RELEASE, https://www.cockroachlabs.com/docs/savepoint.html
`,
	},
	//line sql.y: 3057
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 3058
		Category: hTxn,
		//line sql.y: 3059
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3067
		SeeAlso: `COMMIT, ROLLBACK, https://www.cockroachlabs.com/docs/begin-transaction.html
`,
	},
	//line sql.y: 3080
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 3081
		Category: hTxn,
		//line sql.y: 3082
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 3085
		SeeAlso: `BEGIN, ROLLBACK, https://www.cockroachlabs.com/docs/commit-transaction.html
`,
	},
	//line sql.y: 3098
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 3099
		Category: hTxn,
		//line sql.y: 3100
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 3101
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, https://www.cockroachlabs.com/docs/rollback-transaction.html
`,
	},
	//line sql.y: 3215
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 3216
		Category: hDDL,
		//line sql.y: 3217
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 3218
		SeeAlso: `https://www.cockroachlabs.com/docs/create-database.html
`,
	},
	//line sql.y: 3287
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 3288
		Category: hDML,
		//line sql.y: 3289
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 3294
		SeeAlso: `UPSERT, UPDATE, DELETE, https://www.cockroachlabs.com/docs/insert.html
`,
	},
	//line sql.y: 3311
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 3312
		Category: hDML,
		//line sql.y: 3313
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 3317
		SeeAlso: `INSERT, UPDATE, DELETE, https://www.cockroachlabs.com/docs/upsert.html
`,
	},
	//line sql.y: 3393
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 3394
		Category: hDML,
		//line sql.y: 3395
		Text: `UPDATE <tablename> [[AS] <name>] SET ... [WHERE <expr>] [RETURNING <exprs...>]
`,
		//line sql.y: 3396
		SeeAlso: `INSERT, UPSERT, DELETE, https://www.cockroachlabs.com/docs/update.html
`,
	},
	//line sql.y: 3564
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 3565
		Category: hDML,
		//line sql.y: 3566
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 3577
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 3578
		Category: hDML,
		//line sql.y: 3579
		Text: `
SELECT [DISTINCT]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 3591
		SeeAlso: `https://www.cockroachlabs.com/docs/select.html
`,
	},
	//line sql.y: 3651
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 3652
		Category: hDML,
		//line sql.y: 3653
		Text: `TABLE <tablename>
`,
		//line sql.y: 3654
		SeeAlso: `SELECT, VALUES, https://www.cockroachlabs.com/docs/table-expressions.html
`,
	},
	//line sql.y: 3893
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 3894
		Category: hDML,
		//line sql.y: 3895
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 3896
		SeeAlso: `SELECT, TABLE, https://www.cockroachlabs.com/docs/table-expressions.html
`,
	},
	//line sql.y: 4001
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 4002
		Category: hDML,
		//line sql.y: 4003
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index hints:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 4021
		SeeAlso: `https://www.cockroachlabs.com/docs/table-expressions.html
`,
	},
}
