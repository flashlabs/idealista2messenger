# Idealista to Messenger

Idealista2Messenger (I2M) is an app that provides the ability to forward notification emails from idealista.com 
to Meta's Messenger accounts.

## Business Driver

As an idealista.com user I want to share its offer notifications with group of people via different channel than email.

## How the application works

1. I2M authorizes to Google API with the OAuth 2.0 credentials.
2. Using Gmail API I2M fetches emails specified by the custom query.
3. Using Facebook's Open Graph API the I2M tells the Meta application to send a message on behalf of the FB Page to its user

## Before you start

You will need:

1. Google Cloud project.
2. Google account with Gmail enabled.
3. Facebook Page.
4. Meta App.


## Configuration

### Google

#### Create Project
1. [Create](https://console.cloud.google.com/projectcreate) the Google Project

#### Enable the Gmail API
1. In the Google Cloud console, [enable the Gmail API](https://console.cloud.google.com/flows/enableapi?apiid=gmail.googleapis.com).

#### Add required privileges for the project
1. In the OAuth screen project's section you need to add privileges to Gmail API with `.../auth/gmail.modify` scope.

#### Obtain the `credentials.json` file 
1. In the Google Cloud console, go to **Menu > APIs & Services > Credentials**.
2. Click **Create Credentials > OAuth client ID**.
3. Click **Application type > Desktop app**.
4. In the **Name** field, type a name for the credential. This name is only shown in the Google Cloud console.
5. Click **Create**. The OAuth client created screen appears, showing your new Client ID and Client secret.
6. Click **OK**. The newly created credential appears under **OAuth 2.0 Client IDs**.
7. Save the downloaded JSON file as `credentials.json`, and move the file to `config` directory.

#### Obtain the `token.json` file
1. The first time you run the I2M, it prompts you to authorize access:
   - If you're not already signed in to your Google Account, you're prompted to sign in. If you're signed in to multiple accounts, select one account to use for authorization.
   - Click Accept.
   - Copy the `code` from the browser, paste it into the command-line prompt, and press Enter.
2. `token.json` file will be created in `config` directory.
3. Authorization information is stored in the file system, so the next time you run the I2M, you aren't prompted for authorization.

### Meta

1. Create Facebook Page
2. Create Meta App
   - App might be in dev mode, you will just need to invite users to participate f.e. as a testers
   - If you invite someone, their not getting any notification, invitations might be accepted [here](https://developers.facebook.com/settings/developer/requests/)
3. Get a Page Access Token

#### Obtain Long Lived Paged Access Token
Long-lived Page Access Token procedure:
1. Get User Access Token:
   - App: {your meta app}
   - User or Page token: User token
   - Permissions: pages_show_list, pages_messaging, pages_read_engagement
   - Tool: https://developers.facebook.com/tools/explorer/
2. Go with token to Access Token Debugger:
   - Debug token
   - Click “Extend Access Token”
   - Tool: https://developers.facebook.com/tools/debug/accesstoken/
3. With long-lived User Access Token get long-lived Page Access Token:
   - Graph API version: v15.0
   - App Scoped User ID: {your app scoped user id} 
   - CURL: `curl -i -X GET "https://graph.facebook.com/{graph-api-version}/{app-scoped-user-id}/accounts?access_token={long-lived-user-access-token}"`
   - Docs: https://developers.facebook.com/docs/facebook-login/guides/access-tokens/get-long-lived/#long-lived-page-token
4. Long-lived Page Access Token save in `config/page_access_token.json` file.

#### Obtain Facebook Page Scoped User IDs for sending messages
1. CURL: `curl -i -X GET "https://graph.facebook.com/{latest-api-version}/{page-id}/conversations?fields=participants&access_token={page-access-token}"`
2. Docs: https://developers.facebook.com/docs/messenger-platform/get-started

#### Putting all together in the `.env` file

1. `FB_PAGE_ID` - Facebook Page ID
2. `FB_PAGE_RECIPIENTS` - Comma separated recipients lists (no whitespace allowed) 

### Run the I2M application

If everything is set up correctly, you should be able run application and see similar output:

```
go run main.go
Idealista2Messenger 2022.12.10 17:44:37
Main process
Sending 1 message(s) to [<PAGE_SCOPED_USER_ID> <PAGE_SCOPED_USER_ID>]
Message <MESSAGE_ID> to <PAGE_SCOPED_USER_ID> sent
Message <MESSAGE_ID> to <PAGE_SCOPED_USER_ID> sent
Application "Idealista 2 Messenger" has finished processing
```

### Troubleshooting

#### Gmail token expired or revoked

If your Gmail token is expired or revoked and you can see similar output:
```
Main process
Unable to retrieve messages: Get "https://gmail.googleapis.com/gmail/v1/users/me/messages?alt=json&prettyPrint=false&q=is%3Aunread+from%3Aidealista.com": oauth2: cannot fetch token: 400 Bad Request
Response: {
  "error": "invalid_grant",
  "error_description": "Token has been expired or revoked."
}
Application "Idealista 2 Messenger" has finished processing
```

1. Delete the `config/token.json` file: `rm -rf config/token.json`
2. Run the application to regenerate it (See _Obtain the `token.json` file_ section)

#### Other issues
If you have any questions, just submit an Issue with all the details. 