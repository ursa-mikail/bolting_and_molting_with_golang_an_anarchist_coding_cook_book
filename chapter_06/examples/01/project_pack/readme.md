
# Adding an External Library for Email Validation
<pre>
# run go get to import the external library:

% go mod init project_pack
% cd project_pack

project_pack % go get github.com/go-playground/validator/v10
project_pack % go mod tidy

project_pack % go run main.go
[ERROR] Authentication failed: invalid credentials

</pre>

By using packages and modules, you're delegating specific tasks to specialized code:

The `auth` package is responsible for authentication.
The `logging` package handles logging errors.
The `db` package connects to a database.
The `email` package validates email addresses.

You're essentially "passing the buck" by reusing and delegating responsibilities to external libraries (like the validator library) and maintaining modularity. 
