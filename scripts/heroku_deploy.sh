# Assuming Heroku CLI is set up correcly, this script deploys your branch to Heroku
# Run with `make deploy`

heroku --version
if [ "$?" -gt "0" ]; then
  echo "MAKE: Heroku not installed! Would you like to install it? (y/n)"
  echo "MAKE: Note that you must have Homebrew installed for this to work."
  read choice
  if [ "$choice" == "y" ]; then
    echo "MAKE: Installing heroku..."
    brew install heroku/brew/heroku
    heroku --version
    heroku login
    echo "MAKE: Please enter the name of your Heroku app: "
    read app
    git remote add heroku https://git.heroku.com/"$app".git
    echo "MAKE: Process complete - please run 'make deploy' again."
    exit
  else
    echo "MAKE: Alright, bye!"
    exit
  fi
fi

branch_name=$(git symbolic-ref --short -q HEAD)
git push heroku "$branch_name":master --force