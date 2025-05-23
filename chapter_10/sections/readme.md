# Chapter 10: File I/O — Reading and Writing Your Way to Freedom
## "File I/O Follies: Reading, Writing, and Weeping"

File I/O in Go looks deceptively clean—`os.Open`, `ioutil.ReadFile`, `os.Create`, `defer f.Close()`—but like most things that seem simple, it's an illusion propped up by runtime panic and missing nuance.

You will encounter files that refuse to open, permissions that don’t match your user, encoding issues no one told you about, and directories that vanish because the context was a temp folder in a Docker build. Welcome to the real world of File I/O.

Buffered vs unbuffered? You’ll Google that 30 times. Binary vs text? Welcome to encoding hell. Some files are unknown, and they may flood the memory you cannot afford. Some files are ports or gateways to the unknown. Some files are not meant to be open. 

Append vs overwrite? Check your flags, because one mistake can nuke a production log.
Concurrency? You better lock that file or prepare to debug phantom writes.

And error handling? Most just slap on `if err != nil { panic(err) }` and pretend they're "handling" it. Spoiler: they aren’t.

The real I/O problems:
- Error Swallowing: Ignored `defer` or missing flush kills your output silently.
- MIME and Encodings: Everyone assumes UTF-8. The universe does not.
- Filesystem Diversity: Works fine on your Mac, crashes on Linux CI.
- Permission Mismatch: File opened, yes—but not writable? Enjoy debugging that blindfolded.

In Golang, working with files is less about elegance and more about pain management. You’re not just reading and writing—you’re praying it won’t break again for reasons outside your codebase. Welcome to I/O: where everything is your fault, even when it’s not.



<hr>

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