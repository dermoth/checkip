# Remote IP Address Check API

## Description

This simple application returns the remote address of the calling user. It
supports both IPv4 and IPv6.

## Setup instructions

1. Create a new Google App Engine project and open the shell

2. Check out the code

        CHECKIPDIR=src/$DEVSHELL_PROJECT_ID/checkip
        mkdir -p $CHECKIPDIR
        git clone https://github.com/dermoth/checkip.git $CHECKIPDIR
        cd $CHECKIPDIR/src

3. Test the app using the web preview - run the app on the shell with:

        goapp serve app.yaml

    NB: The app will likely show an IP address of `127.0.0.1` as the request
    is being proxied at multiple levels...

4. Deploy with (increase version as needed):

        goapp deploy -application $DEVSHELL_PROJECT_ID -version 0

5. Create Custom domains for checkip to your liking. To have IPv4-only and
   IPv6-only url's, create specific domain names for them and only add the
   A (IPv4) or AAAA (IPv6) records from Google. If you add both sets, the
   IP returned will depend on the IP version supported.

    NB: Custom domains are required to set up version-specific endpoints.
    Also keep in mind browsers usually favor IPv4 over 6to4 and Terredo
    tunnels.
 
## Bugs and missing features

X-Forwarded-For and other "client ip" headers are not looked at. Although some
services out there returns those and some setups migth actually require it
(ex. runnning behind a load-balancer) it is always hard to know for sure which
one is valid when multiple addresses are encountered. At the very least a
check for RFC1918 addresses should be made, and possibly a blacklist of local
gateways. When runnign straight up on Google App Engine, the remote address
is that of the client.

For now this app has been written in a minimalistic fashion to reduce load on
my free usage quota to a strict minimum. Additional features could be
implemented such as request parameters to return data in other formats (HTML,
JSON, etc.). Please let me know if you have such interest.

## Contact

Thomas Guyot-Sionnest <Thomas@Guyot-Sionnest.net>

