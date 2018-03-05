# calories
[![Build Status](https://travis-ci.org/bobheadxi/calories.svg?branch=dev)](https://travis-ci.org/bobheadxi/calories) [![Coverage Status](https://coveralls.io/repos/github/bobheadxi/calories/badge.svg?branch=31-unit-tests)](https://coveralls.io/github/bobheadxi/calories?branch=31-unit-tests)

WIP

### Setup
1. Make sure you have [Go](https://golang.org/doc/install) and the [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli#download-and-install) installed
2. Clone the repository into `$GOPATH/src/github.com/bobheadxi`
3. If you aren't a repository collaborator, you will have to create your own Heroku and [Postgres](https://devcenter.heroku.com/categories/heroku-postgres) instance:
```bash
heroku create
heroku config:set HEROKU_URI=your-heroku-uri
heroku config:set DATABASE_URL=your-database-url
```
4. Now you have to set up your Facebook app page. Create a Facebook page for your app at https://developers.facebook.com/apps/
5. Click on "Messenger" under "Products". Generate a token under "Token Generation", and set up Webhook under "Webhooks". Use the URL of your Heroku server as your webhook URL (append `/webhook` to the end). Make sure you at least subscribe to `messages` and `messaging_postbacks`.
6. Save your tokens as Config Variables on Heroku. These are required for the app to run.
```bash
heroku config:set FB_PAGE_ID=your-page-id
heroku config:set FB_TOKEN=your-fb-token
```
7. With all this set up, you should be good to go! Deploy the bot to your Heroku instance: 
```bash
git push heroku dev:master
```
Now you can message your Facebook page to try out the bot. Note that until your application gets submitted for review and accepted, other users will have to be added as testers before the bot will respond to them.

### Development

- `make` -> builds the project

- `make deps` -> installs project dependencies

- `make db` -> sets up local database shenanigans for testing

- `make test` -> runs all tests with coverage reporting

- `make clean` -> cleans up project (eg remove binaries)

- `make deploy` -> deploys your branch to Heroku (and sets up the Heroku CLI if you haven't)
