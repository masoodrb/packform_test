# Setup Instructions


## For Backend

In terminal swith to `backend` folder and modify `.env` file with your postgresql connection settings.

make sure the database that you specify here exits in DB.

after that run command `go run bin/import_data/main.go` in terminal.

to run web server run command `go run bin/server/main.go`.

to create table and import data from csv files.


## For Frontend

make sure you have node `v19` installed. This is developed with `node v19.4.0`.

In terminal swith to `frontend` folder and run command `npm install`.

after all packages are installed runn command `npm run build`.

once the build process is completed run command `npm run preview`.

open any Browser and open link ` http://localhost:4173/orders` to see the following page.

![screenshot](/screenshot.png?raw=true "Screen Shot")