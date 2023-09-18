# Database and Data Center Security


🔭 Explore
---
- [Biggest Data Breaches in US History](https://www.upguard.com/blog/biggest-data-breaches-us)


The Need and Challenge for Database Security
---
- focuses on relational database management systems (RDBMS)
  - RDBMS dominates database market
- Organizational databases concentrate sensitive information
  - user credentials, financial data, etc.
- no sufficient protection on complex DBMS consist of a heterogeneous mixture of databases
  - no full-time database security personnel
- DBMS have versatile interaction protocol -  Structured Query Language (SQL)
- part or all of the corporate databases are hosted on cloud


Databases and Database Management Systems
---
- Databases
  - Structured collection of data stored for use by users and applications
  - Contain relationships between data items and groups of data items
  - Sometimes contain sensitive data that needs to be secured
- Database management system (DBMS)
  - Suite of programs for constructing and maintaining databases
  - Offers ad hoc query facilities to multiple users and applications
- Query language
  - a uniform interface to databases


DBMS Architecture
---
- ![CSPP4e Figure 5.1](../../figs/dbms.jpg)
- two protections
  - OS security mechanisms on 
    - users and processes
  - DBMS builtin security mechanisms on 
    - databases, tables, views, records and stored procedures


Relational Databases
---
- Collections of tables (relations) consisting of rows and columns
  - Each column (field, attributes) holds a particular type of data
    - Ideally has one column where all values are unique, 
      - forming an identifier/key for that row
  - Each row (record, tuple) contains a specific value for each column
- Multiple tables linked together by keys (relationship)
- Use SQL to manipulate and access the database
- Primary key
  - Uniquely identifies a row
  - Consists of one or more column names
- Foreign key
  - Links one table to attributes in another
- View/virtual table
  - Result of a query
  - often used for security purposes


Structured Query Language (SQL)
---
- Standardized language to 
  - define schema, 
  - manipulate, and query data
- Several similar versions of ANSI/ISO standard
- can be used to:
  - Create databases, tables and views
  - Insert and delete data in tables
  - Retrieve data with query statements


SQL Injection (SQLi) Attacks
---
- One of the most prevalent and dangerous network-based security threats
- Designed to exploit the nature of Web application pages
- Sends malicious SQL commands to the database server
- common attack goals
  - extract bulk of data
  - Modify or delete data
  - Execute arbitrary operating system commands
  - Launch denial-of-service (DoS) attacks



🖊️ Explore 
---
- [Typical SQLi Attacks](https://www.w3schools.com/sql/sql_injection.asp)



SQLi Attack vectors
---
- User input fields
- Server variables
  - through HTTP  headers
- website cookies
- Physical user input
  - e.g. barcodes, RFID tags, or even paper forms scanned using optical character recognition (OCR)
- Second-order injection
    - attack from within the system itself triggered by attacker


SQLi Attack types
---
- Categorized based on the relationship between 
  - the SQL command injection channel (IC) and
  - the channel for retrieving results (OC)
- inband attack 
- inferential attack
- Out-of-Band Attack


SQLi Countermeasures
Database Access Control
SQL-Based Access Definition
Cascading Authorizations
Role-Based Access Control
Inference
Database Encryption
Data Center Security
Data Center Elements
Data Center Security Considerations
TIA-492