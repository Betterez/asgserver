Simple web server test
===========================
This "server" is a test http server that enables you to change the output version.
it was made for one purpose only – testing load balancers consistency.

building
------------
`make` is required. just run `make` in the root folder.

Running
-------------
`make run` will run with version =1.
You can set up `"VERSION=2"` for version 2 and so on.
version must be a number.
