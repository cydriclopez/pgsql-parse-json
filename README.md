# pgsql-parse-json

## Parse JSON in PostgreSQL to save records

> ***This tutorial requires some knowledge in Linux, Docker, Git, Angular, PostgreSQL, and Go Programming Language.***

### Table of Contents
1. Introduction
2. Goal
3. Prerequisites
4. Clone this repo
5. Client-side Angular code
6. PostgreSQL database code
7. Compile Go server-side code
8. Server-side Go code
9. Conclusion

### 1. Introduction

We are continuing where we left-off in the previous tutorial [Go POST JSON passthru controller](https://github.com/cydriclopez/go-post-json-passthru).

The added code are: 1.) Postgresql database code, 2.) Go code package ***common*** to connect to Postgresql, and 3.) alter Go code package ***treedata*** to call the Postgresql stored-function code and pass the client Tree component JSON data.

In this tutorial we do not have new Angular sourcecode in the ***src/client*** folder. Instead what we have is the compiled version, from the previous tutorial [Go POST JSON passthru controller](https://github.com/cydriclopez/go-post-json-passthru), in the ***src/client/dist-static*** folder. I have a ***.gitignore*** rule to block the ***dist*** folder so I had to rename to ***dist-static***.


### 2. Goal

Our goal is to save our tree component JSON data from our app screen:
<br/>
<kbd><img src="images/primeng-tree-demo2.png" width="650"/></kbd>
<br/>

Into our Go pass-thru controller server-side screen console:
<br/>
```bash
:webserv .
2022/09/05 13:07:59 PostgreSQL 14.2 (Debian 14.2-1.pgdg110+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
2022/09/05 13:07:59
Serving static folder: .
Listening on port: :3000
Press Ctrl-C to stop server
```
<code>
2022/09/05 13:09:42 jsonData: [{"label":"Documents","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Documents Folder","children":[{"label":"Work","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Work Folder","children":[{"label":"Expenses.doc","icon":"pi pi-file","data":"Expenses Document"},{"label":"Resume.doc","icon":"pi pi-file","data":"Resume Document"}],"toexpand":false},{"label":"Home","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Home Folder","children":[{"label":"Invoices.txt","icon":"pi pi-file","data":"Invoices for this month"}],"toexpand":false}],"toexpand":true},{"label":"Pictures","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Pictures Folder","children":[{"label":"barcelona.jpg","icon":"pi pi-image","data":"Barcelona Photo"},{"label":"logo.jpg","icon":"pi pi-image","data":"PrimeFaces Logo"},{"label":"primeui.png","icon":"pi pi-image","data":"PrimeUI Logo"}],"toexpand":true},{"label":"Movies","expandedIcon":"pi pi-folder-open","collapsedIcon":"pi pi-folder","data":"Movies Folder","children":[{"label":"Al Pacino","data":"Pacino Movies","children":[{"label":"Scarface","icon":"pi pi-video","data":"Scarface Movie"},{"label":"Serpico","icon":"pi pi-video","data":"Serpico Movie"}]},{"label":"Robert De Niro","data":"De Niro Movies","children":[{"label":"Goodfellas","icon":"pi pi-video","data":"Goodfellas Movie"},{"label":"Untouchables","icon":"pi pi-video","data":"Untouchables Movie"}]}]}]
</code>
<br/>
<br/>

And finally, as records in a table in our Postgresql database:

```
postgres=# select * from tree_data;

 key | parent |     label      |    icon     |   expandedicon    | collapsedicon |          data           | leaf | toexpand
-----+--------+----------------+-------------+-------------------+---------------+-------------------------+------+----------
   1 |      0 | data           |             |                   |               | data                    | f    | t
   2 |      1 | Documents      |             | pi pi-folder-open | pi pi-folder  | Documents Folder        | f    | t
   3 |      2 | Work           |             | pi pi-folder-open | pi pi-folder  | Work Folder             | f    | f
   4 |      3 | Expenses.doc   | pi pi-file  |                   |               | Expenses Document       | t    | f
   5 |      3 | Resume.doc     | pi pi-file  |                   |               | Resume Document         | t    | f
   6 |      2 | Home           |             | pi pi-folder-open | pi pi-folder  | Home Folder             | f    | f
   7 |      6 | Invoices.txt   | pi pi-file  |                   |               | Invoices for this month | t    | f
   8 |      1 | Pictures       |             | pi pi-folder-open | pi pi-folder  | Pictures Folder         | f    | t
   9 |      8 | barcelona.jpg  | pi pi-image |                   |               | Barcelona Photo         | t    | f
  10 |      8 | logo.jpg       | pi pi-image |                   |               | PrimeFaces Logo         | t    | f
  11 |      8 | primeui.png    | pi pi-image |                   |               | PrimeUI Logo            | t    | f
  12 |      1 | Movies         |             | pi pi-folder-open | pi pi-folder  | Movies Folder           | f    | f
  13 |     12 | Al Pacino      |             |                   |               | Pacino Movies           | f    | f
  14 |     13 | Scarface       | pi pi-video |                   |               | Scarface Movie          | t    | f
  15 |     13 | Serpico        | pi pi-video |                   |               | Serpico Movie           | t    | f
  16 |     12 | Robert De Niro |             |                   |               | De Niro Movies          | f    | f
  17 |     16 | Goodfellas     | pi pi-video |                   |               | Goodfellas Movie        | t    | f
  18 |     16 | Untouchables   | pi pi-video |                   |               | Untouchables Movie      | t    | f
(18 rows)

postgres=#
```

### 3. Prerequisites

As mentioned before, this tutorial builds on the previous tutorial [Go POST JSON passthru controller](https://github.com/cydriclopez/go-post-json-passthru). I suggest you go thru them in sequence especially if you are new to Angular and the Go programming language.

I assume that you have a working [Angular](https://github.com/cydriclopez/docker-ng-dev), [PostgreSQL](https://github.com/cydriclopez/docker-pg-dev) and [Go](https://github.com/cydriclopez/go-static-server#3-install-go) installations. Please checkout the previous tutorials that cover these topics.


### 4. Clone this repo



### 5. Client-side Angular code

As already mentioned, in this tutorial we do not have new Angular sourcecode in the ***src/client*** folder. Instead what we have is the compiled version, from the previous tutorial [Go POST JSON passthru controller](https://github.com/cydriclopez/go-post-json-passthru), in the ***src/client/dist-static*** folder. I have a ***.gitignore*** rule to block the ***dist*** folder so I had to rename to ***dist-static***.

Further down below we will run our Go server-side app ***webserv*** and pass this Angular compiled folder:

```bash
user1@penguin:~/Projects/pgsql-parse-json$
:cd src/client/dist-static/primeng-quickstart-cli
:webserv .
```


### 6. PostgreSQL database code


### 7. Compile Go server-side code


### 8. Server-side Go code


### 9. Conclusion


Under construction!
But you can peek into the completed committed source-code.
Happy coding! 😊

---
