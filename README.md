# Borders

This is meant to be a small app to tally border checks in Schengen.
It takes in border crossings, and one can specify if there was a border check and, if so, if identification of the entity logging the crossing was checked.

The application was meant as a single-user operation, but was extended to support multiple user accounts somewhat.

It displays aggregates of the information it stores over a public-facing page.
Administrative tasks (Adding crossings, managing users, transports or countries) is gated behind a login.

> [!WARNING]  
> All user accounts have the same set of permissions. Be sure to only add trusted users, as they can create, edit or delete countries and means of transportation, as well as add further users.
>
> **Users cannot be deleted over the web interface (as of now).**

> [!TIP]
> One can manually delete users from the database.
> Open it using `sqlite3 borders.db` (or whatever the database file is configured to be), and run `DELETE FROM users WHERE name = 'example';`. Replace `example` with the username to be deleted.
>
> **Caution: this is irreversible.**

## Running the app

The app uses an sqlite3 database. As such, it does not depend on other software to be running.
Assuming a go toolchain >= 1.24.3, just type `go run .` in a terminal.

> [!NOTE]
> The sqlite driver uses CGO. running with `CGO_ENABLED=0` will not work.

The app will be available on `http://localhost:8080`.
The first time the application is started, it will log a link of the form `http://127.0.0.1:8080/setup/xTbJm5lrGfI` that can be used to set up an user account.
Unless it is configured otherwise through environment variables, it will also create a `borders.db` and `key.bin` in the folder in which it is invoked.
As long as this isn't done, statistics will not be available.

> [!TIP]
> The "http://127.0.0.1:8080" part of the logged message is only an example - if the application is configured to be available on another URL (behind a reverse proxy, ...), the endpoint will work there as well.

> [!NOTE]
> Once setup is done, the application should be restarted, so that the setup endpoints aren't exposed anymore.
> This isn't critical however, as accessing them once the application is set up will only show an error message.

> [!WARNING]
> The contents of the `key.bin` files are sensitive, as they can be used to forge tokens to log into the website.

A compose file is also provided for convenience.
It can be started as is, and the application will be available on `http://localhost:8080`, as with running it directly from the command line.
The `borders.db` and `key.bin` files need to be persisted, so they should be placed in a volume or similar.
The compose file does this by default.
Other settings can be changed by uncommenting the relevant lines and changing the example values.

## Adjusting configuration options

The Behaviour of the application can be changed by using environment variables:

| Environment Variable |             Default Value              | Use                                                                |
| :------------------: | :------------------------------------: | ------------------------------------------------------------------ |
|     `SERVE_PORT`     |                  8080                  | Port on which the application listens                              |
|      `DB_FILE`       |             `./borders.db`             | Where the sqlite database should be saved                          |
|      `OG_TITLE`      | `WHAT THE FUCK IS A SCHENGEN 🇩🇪🇩🇪🇩🇪🇩🇪` | What should be used as title in OpenGraph tags (for link previews) |
|    `JWT_KEY_FILE`    |              `./key.bin`               | Where the JWT key file should be saved                             |
|   `JWT_KEY_LENGTH`   |                   16                   | Length (in bytes) that a JWT key should be                         |
