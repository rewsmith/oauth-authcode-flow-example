Tyk OAuth Sample
================

This is a quick project that shows the Tyk OAuth request cycle from start to finish.

To try this project out:

1. In your Tyk Gateway, create an API and call it `oauth2`
2. Set the Access Method to "Oauth 2.0"
3. Select "Allowed Access Types: Authorization codes"
4. Select "Allowed Authorize Types: Authorization Code"
5. Set the login redirect for this API to be: `http://localhost:8000/login`
6. Take note of the API ID
7. Add an oauth client to it and set the redirect to be `http://localhost:8000/final`
8. Take note of the client ID
9. Create a policy that has access to this API, take not of the Policy ID

The .env file includes some default variable values.  Update accordingly, as per your installation:

1. Set the `API_LISTENPATH` to `oauth2` (or whatever the listen path is for your OAuth API)
2. Set `ORG_ID` to be your Org ID (Go to users -> select your user, it is under RPC credentials)
3. Set `POLICY_ID` to be your policy ID
4. Set `API_ID` to be your API ID
5. Set `GATEWAY_URL` to be the host path to your gateway e.g. http://domain.com:port (note no trailing slash)
NB: If running as Docker, set `GATEWAY_URL` to `http://host.docker.internal:8080` or override when running the container (see below) 
6. Set `ADMIN_SECRET` to your the secret in your `tyk.conf`
7. Set `CLIENT_ID` to the value of your client ID (can be overriden when you run the app)
8. Set `REDIRECT_URI` to the value of your client (can be overriden when you run the app)

To run the app you can either run as go:

	go run *.go

Or, a Dockerfile is provided, so you can build a Docker image and run as a container:

    docker run -e GATEWAY_URL=http://host.docker.internal:8080 -p 8000:8000 --name <Docker-container-name> <Docker-image-name>
    
Then visit:

http://localhost:8000

1. Set the `Client ID` field to the value of your client ID
2. Set the `Redirect URI` value to the one of your client

If you've set everything up correctly, you should be taken through a full OAuth authorisation codee flow.

This app emulates two parties:

1. The requester (client)
2. The identity provider portal (your login page)

We make use of the Tyk REST API Authorization endpoint to complete the request cycle, you can see an API client in the `util.go` file.