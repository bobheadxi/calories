# This script handles local database creation to make testing easier.
# Run with `make db`

postgres -V
if [ "$?" -gt "0" ]; then
  echo "MAKE: Postgres not installed! Would you like to install it? (y/n)"
  echo "MAKE: Note that you must have Homebrew installed for this to work."
  read choice
  if [ "$choice" == "y" ]; then
    echo "MAKE: Installing postgres..."
    brew install postgresql
    initdb /usr/local/var/postgres
    echo "MAKE: Process complete - please run 'make db' again."
    exit
  else
    echo "MAKE: Alright, bye!"
    exit
  fi
fi

echo "MAKE: Killing existing postgres processes..."
pg_ctl -D /usr/local/var/postgres stop -s -m fast
pg_ctl -D /usr/local/var/postgres start
createuser -s postgres
createdb calories_test_db
psql -d calories_test_db -a -f ./scripts/test_db_setup.sql
echo "MAKE: Local database ready to go!"