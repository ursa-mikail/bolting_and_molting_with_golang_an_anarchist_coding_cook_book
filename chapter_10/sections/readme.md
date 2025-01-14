
We automate the scaffolding of the <a href="https://github.com/ursa-mikail/golang-gaia-basic-structure/tree/main"> golang-gaia-basic-structure</a>.

<pre>
chmod +x make_go.sh
# Run the script with your desired module name:
# ./make_go.sh example.com/demo
./make_go.sh test-app

# Resulting Structure
After running the script, the structure will look like this:

└── test-app
    ├── go.mod
    ├── go.sum
    ├── images.db
    ├── images_export.csv
    ├── libs
    │ ├── db
    │ │ └── db.go
    │ └── p0
    │     └── p0.go
    ├── main.go
    └── utils
        ├── crypto.go
        └── util_00.go

# modify the generated main.go, db.go, crypto.go

# test run:
% cd test-app 
% go get github.com/mattn/go-sqlite3
% go run main.go

out:
</pre>
```
Image Record: map[contents:example_data hmac:42558374b30a8eec6e7e5220a5a3bf4ee6921ed19bb2304dd4ce1604fd16ebbf id:1 name:example_image sha256:d7f2db9e66297f3ac43a9ddcad1c9ec43c1becbba3b87dd1689ace47b9afed7c status:active status_signature:b9c69ada4404dba178f64ce0bae66602567df9e826790a41abf2d7454e069542 status_url: team:team_a team_owner:owner_a]
Data exported to images_export.csv
```