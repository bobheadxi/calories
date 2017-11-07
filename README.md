# calories
WIP

### Setup
1. Make sure you have [Go installed](https://golang.org/doc/install)
2. Clone the repository into `$GOPATH/src/bobheadxi`
3. If you aren't a repository collaborator, you will have to create your own Heroku instance:
```bash
heroku create
heroku config:set HEROKU_URI==your-heroku-uri
```
4. Now you have to set up your Facebook app page. Create a Facebook page for your app at https://developers.facebook.com/apps/
5. Click on "Messenger" under "Products". Generate a token under "Token Generation", and set up Webhook under "Webhooks". Use the URL of your Heroku server as your webhook URL (append "/webhook/" to the end)
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